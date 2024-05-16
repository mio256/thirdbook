package handler_test

import (
	"context"
	"testing"

	"github.com/mio256/thirdbook/usecase/handler"
	"github.com/stretchr/testify/assert"
)

func TestHandler_PingGet(t *testing.T) {
	h := handler.NewHandler()
	ctx := context.Background()
	res, err := h.PingGet(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "pong", res.Message.Value)
}
