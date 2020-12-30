#!/usr/bin/env bash

main() {
  install_go_libraries
  build
  echo "========   DONE   ========"
}

install_go_libraries() {
  echo "   Installing GO dependencies    "
  install_dep github.com/golang/mock/mockgen
  install_dep golang.org/x/tools/cmd/goimports
  install_dep github.com/mcubik/goverreport
  go mod tidy
}

build() {
  echo "         Building App            "
  make build
}

main