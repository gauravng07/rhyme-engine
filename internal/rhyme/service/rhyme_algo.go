package service

import (
	"math"
	"sort"
	"strings"
	"unicode"
)

type PhonemeModifier struct {
	T string
	Mod map[string]bool
}

func FindPhoneticWithModifiers() map[string]PhonemeModifier{
	phoneticWithModifier := make(map[string]PhonemeModifier)
	for ph := range PHONETIC_ALPHABET {
		modifier := make(map[string]bool)
		for i := range ph {
			char := ph[i]
			if unicode.IsUpper(rune(char)) {
				if i == 0 {
					modifier[string(char)] = true
				} else {
					modifier[string(char)] = false
				}
			}
		}
		if len(modifier) != 0 {
			var trimmed string
			for char := range modifier {
				trimmed = strings.Replace(ph, char, "", 1)
			}
			phoneticWithModifier[ph]= PhonemeModifier{
				T: trimmed,
				Mod: modifier,
			}
		}
	}
	return phoneticWithModifier
}

func SplitWordIntoPhonetics(words []string, phoneticWithModifier map[string]PhonemeModifier) map[string]string {
	phoneticWordsList := make(map[string]string)
	phoneticLargest := make([]string,0)
	phoneticLargest = []string{"Chere", "Cath", "easE", "Cear", "Cour", "eigh", "ourE", "thin", "Cey", "all", "eye", "onV", "omV", "air", "ear", "ere", "ure", "ing", "oth", "eas", "ar", "Sa", "ur", "ie", "yC", "ee", "ea", "ut", "pu", "ue", "oo", "ow", "ou", "ay", "oE", "oy", "oi", "ck", "ss", "sh", "tt", "ch", "th", "is", "ys", "ge", "u", "a", "e", "i", "o", "b", "d", "f", "g", "h", "y", "c", "k", "l", "m", "n", "p", "r", "s", "t", "v", "w", "z", "j"}

	for _, w := range words {
		original := w
		w = strings.ToLower(w)
		phoneticWord := w
		wordAdded := false
		for _, phoneme := range phoneticLargest {
			modifiers := make(map[string]bool)
			if !hasLetter(w) {
				phoneticWordsList[original] = phoneticWord
				wordAdded = true
				break
			}

			if _, ok := phoneticWithModifier[phoneme]; ok {
				m := phoneticWithModifier[phoneme]
				phoneme, modifiers = m.T, m.Mod
			}

			if strings.Contains(w, phoneme) {
				index := strings.Index(w, phoneme)
				if len(modifiers) > 0 {
					for mod := range modifiers {
						if mod == "S" {
							replace := strings.HasPrefix(w, phoneme)
							if replace {
								w = strings.Replace(w, phoneme, "*", 1)
								phoneticWord = strings.Replace(phoneticWord, phoneme, PHONETIC_ALPHABET[mod+phoneme], 1)
							}
						} else if mod == "E" {
							replace := strings.HasSuffix(w, phoneme)
							if replace {
								w = strings.Replace(w, phoneme, "*", 1)
								phoneticWord = strings.Replace(phoneticWord, phoneme, PHONETIC_ALPHABET[phoneme + mod ], 1)
							}
						} else {
							var charInModPos uint8
							if v, ok := modifiers[mod]; ok && v {
								charInModPos = w[index - 1]
							} else {
								charInModPos = w[index + len(mod)]
							}

							if checkCharModValidity(charInModPos, mod) {
								w = strings.Replace(w, string(charInModPos) + phoneme, "*",1)
								if _, ok := modifiers[mod]; ok {
									phoneticWord = strings.Replace(phoneticWord, string(charInModPos) + phoneme, PHONETIC_ALPHABET[mod + phoneme ],1)
								} else {
									phoneticWord = strings.Replace(phoneticWord, string(charInModPos) + phoneme, PHONETIC_ALPHABET[phoneme + mod ],1)
								}
							}
						}
					}
				} else {
					w = strings.Replace(w, phoneme, "*",1)
					phoneticWord = strings.Replace(phoneticWord, phoneme, PHONETIC_ALPHABET[phoneme],1)
				}
			}
		}
		if !wordAdded {
			phoneticWordsList[original] = phoneticWord
		}
	}
	return phoneticWordsList
}

func FindBestMatches(phoneticList map[string]string, inputWordList map[string]string, word string) []map[string][]string {
	scores, inputLen := findPossibleMatches(phoneticList, inputWordList, word)
	bestMatchedWords := make([]map[string][]string, 0)
	bestScores := sortScores(scores)
	limitedScores := make([]map[string][]string, 0)
	for _, v := range bestScores {
		if v < float64(inputLen * 10) {
			limitedScores = append(limitedScores, scores[v])
		}
	}

	if len(limitedScores) == 0 {
		return nil
	} else if len(limitedScores[0]) > 1 {
		bestMatchedWords = append(bestMatchedWords, limitedScores[0])
	} else {
		difference := bestScores[len(bestScores) - 1] - bestScores[0]
		best := bestScores[0] + 5
		compare := math.Max(float64(difference/2), float64(best))
		bestWithinRange := make([]map[string][]string, 0)
		for k := range scores {
			if k < compare {
				bestWithinRange = append(bestWithinRange, scores[k])
			}
		}
		bestMatchedWords = append(bestMatchedWords, bestWithinRange...)
	}
	return bestMatchedWords
}

func sortScores(scores map[float64]map[string][]string) []float64 {
	keys := make([]float64, 0)
	for k := range scores {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}

func findPossibleMatches(phoneticList map[string]string, inputWordList map[string]string, word string) (map[float64]map[string][]string, float64) {
	scores := make(map[float64]map[string][]string, 0)
	formattedInputWord := filterSlice(inputWordList[word])
	for i, j := 0, len(formattedInputWord)-1; i < j; i, j = i+1, j-1 {
		formattedInputWord[i], formattedInputWord[j] = formattedInputWord[j], formattedInputWord[i]
	}
	inputLen := float64(len(formattedInputWord))

	for k, word := range phoneticList {
		total := -1.0
		wordS := filterSlice(word)
		for i, j := 0, len(wordS)-1; i < j; i, j = i+1, j-1 {
			wordS[i], wordS[j] = wordS[j], wordS[i]
		}
		numberIncorrectAtBeginning := 0
		eligible := true
		successiveInputs := 0
		for _, phoneticType := range wordS {
			var score float64
			indexInInput, ok := Find(formattedInputWord, phoneticType)
			indexWordS, _ := Find(wordS, phoneticType)
			if ok {
				numberIncorrectAtBeginning = -1
				distance := math.Abs(float64(indexInInput) - float64(indexWordS))
				if successiveInputs > 0 {
					score = distance
				} else {
					if distance > 2 {
						wordLen := float64(len(wordS))
						score = (inputLen / wordLen ) * 10
					} else {
						score = distance + float64((indexInInput) * 6)
					}
				}
				successiveInputs = successiveInputs + 1
				if total == -1 {
					total = score
				} else {
					total = total + score
				}
			} else {
				if numberIncorrectAtBeginning != -1 {
					numberIncorrectAtBeginning = numberIncorrectAtBeginning + 1
					if numberIncorrectAtBeginning == int(math.Round(float64(inputLen)/3)) {
						eligible = false
						break
					}
				}
				successiveInputs = 0
				wordLen := float64(len(wordS))
				score = (inputLen / wordLen) * 10
				if total == -1 {
					total = score
				} else {
					total = total + score
				}
			}
		}

		if total > -1 && eligible {
			if _, ok := scores[total];ok {
				scores[total][k] = wordS
			} else {
				scores[total] = map[string][]string{
					k: wordS,
				}
			}
		}
	}
	return scores, inputLen
}

func hasLetter(str string) bool {
	for _, letter := range str {
		if unicode.IsLetter(letter) {
			return true
		}
	}
	return false
}

func checkCharModValidity(char uint8, modifier string) bool {
	if modifier == "C" {
		if strings.Contains(string(char), CONSONANTS) {
			return true
		}
	} else if modifier == "V" {
		if strings.Contains(string(char), VOWELS) {
			return true
		}
	}
	return false
}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func filterSlice(t string) []string {
	k := strings.Split(t, "/")
	temp := make([]string, 0)
	for _, d := range k {
		if len(d) > 0 {
			temp = append(temp, d)
		}
	}
	return temp
}

