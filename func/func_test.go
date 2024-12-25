package func_

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

type postReq func(ctx context.Context, req *req, resp *resp) (status int)

type req struct {
	path string
}

type resp struct {
	status int
	data   string
}

func handleReq(ctx context.Context, req *req, callback postReq) any {
	// handle req ...
	resp := &resp{
		data: "data",
	}
	resp.status = callback(ctx, req, resp)
	return resp
}

type PostReqCallback struct {
	req  *req
	resp *resp
}

func (cb *PostReqCallback) postReq(ctx context.Context, req *req, resp *resp) (status int) {
	cb.req = req
	cb.resp = resp
	return 200
}

func TestPassObjectMethodToFuncTypeParam(t *testing.T) {
	cb := &PostReqCallback{}
	req := &req{
		path: "/shopping_cart",
	}
	r := handleReq(context.Background(), req, cb.postReq)
	assert.Equal(t, req, cb.req)
	expectedResp := &resp{
		status: 200,
		data:   "data",
	}
	assert.Equal(t, expectedResp, cb.resp)
	assert.Equal(t, expectedResp, r)
}
