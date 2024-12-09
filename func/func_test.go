package _func

import (
	"context"
	"fmt"
	"testing"
)

type handlerFunc func(c context.Context, ctx *requestContext) any

type requestContext struct{}

type result struct {
	reqPath string
	resp    any
}

func dispatchReq(ctx context.Context, path string, handler handlerFunc) any {
	return &result{
		reqPath: path,
		resp:    handler(ctx, &requestContext{}),
	}
}

type someHandler struct {
	name string
}

func (h *someHandler) handle(c context.Context, ctx *requestContext) any {
	return "ok"
}

func TestPassObjectMethodToArgOfTypeFunc(t *testing.T) {
	h := &someHandler{name: "add to shopping-cart"}
	r := dispatchReq(context.Background(), "/shopping_cart", h.handle)
	fmt.Printf("result: %+v\n", r)
}
