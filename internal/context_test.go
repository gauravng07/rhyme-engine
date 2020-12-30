package internal_test

import (
	"context"
	"rhyme-engine/internal"
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_ShouldGetCorrelationIdIfContextHaveCorrelationId(t *testing.T) {
	ctx := context.WithValue(context.Background(), internal.ContextKeyCorrelationID, "123.123")
	id := internal.GetContextValue(ctx, internal.ContextKeyCorrelationID)
	assert.Equal(t, id, "123.123")
}

func Test_ShouldGetEmptyCorrelationIdIfContextDoesNotHaveCorrelationId(t *testing.T) {
	ctx := context.Background()
	id := internal.GetContextValue(ctx, internal.ContextKeyCorrelationID)
	assert.Equal(t, id, "")
}

func Test_ShouldSetCorrelationIdInContext(t *testing.T) {
	ctx := internal.SetContextWithValue(context.Background(), internal.ContextKeyCorrelationID, "123.123")
	assert.Equal(t, ctx.Value(internal.ContextKeyCorrelationID), "123.123")
}
