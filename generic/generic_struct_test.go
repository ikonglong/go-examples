package generic

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

type RequestContext struct {
	Path string
	Body string // json
}

func (r *RequestContext) JSON(httpStatus int, jsonObj any) {
	fmt.Printf("http status: %d, json body: %v\n", httpStatus, jsonObj)
}

type iReqHandler[Req any, Result any] interface {
	handle(reqCtx *RequestContext)
	buildReq(reqCtx *RequestContext) *Req
	doHandle(reqCtx *RequestContext, req *Req) Result
}

type iOKResponder[Result any] interface {
	respondOK(reqCtx *RequestContext, r Result)
}

func newReqHandlerBase[Req any, Result any](sub iReqHandler[Req, Result]) *reqHandlerBase[Req, Result] {
	return &reqHandlerBase[Req, Result]{iReqHandler: sub}
}

type reqHandlerBase[Req any, Result any] struct {
	iReqHandler[Req, Result]
}

func (h *reqHandlerBase[Req, Result]) handle(reqCtx *RequestContext) {
	req := h.buildReq(reqCtx)
	r := h.doHandle(reqCtx, req)
	h.respondOK2(reqCtx, r)
}

func (h *reqHandlerBase[Req, Result]) buildReq(reqCtx *RequestContext) *Req {
	reqPtr := new(Req)
	err := mapHTTPReqToReq(reqCtx, reqPtr)
	if err != nil {
		panic(fmt.Errorf("illegal req. Failed to build req from http req: %w", err))
	}
	return reqPtr
}

func (h *reqHandlerBase[Req, Result]) respondOK2(reqCtx *RequestContext, rs Result) {
	v, ok := h.iReqHandler.(iOKResponder[Result])
	if ok {
		fmt.Println("call respondOK on sub object")
		v.respondOK(reqCtx, rs)
		return
	}
	reqCtx.JSON(http.StatusOK, rs)
	fmt.Println("call respondOK on base")
}

func mapHTTPReqToReq(reqCtx *RequestContext, req interface{}) error {
	return json.Unmarshal([]byte(reqCtx.Body), req)
}

func newCallAPIHandler() iReqHandler[CallAPIReq, any] {
	o := &CallAPIHandler{}
	var sub iReqHandler[CallAPIReq, any] = o
	o.reqHandlerBase = newReqHandlerBase(sub)
	return o
}

type CallAPIHandler struct {
	*reqHandlerBase[CallAPIReq, any]
}

func (h *CallAPIHandler) doHandle(reqCtx *RequestContext, req *CallAPIReq) any {
	fmt.Printf("CallAPIHandler: do handle req: %v\n", req)
	return "CallAPIHandler: ok"
}

func (h *CallAPIHandler) respondOK(reqCtx *RequestContext, r any) {
	fmt.Println("call respondOK on CallAPIHandler")
}

type CallAPIReq struct {
	URL  string         `json:"url"`
	Body map[string]any `json:"body"`
}

func TestCallAPIHandler(t *testing.T) {
	h := newCallAPIHandler()
	bodyJSON, err := json.Marshal(map[string]any{
		"url": "https://openapi.com/greet",
		"body": map[string]any{
			"who":      "Bob",
			"greeting": "hello",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	reqCtx := &RequestContext{Path: "/", Body: string(bodyJSON)}
	h.handle(reqCtx)
}

func newPlaceOrderHandler() *PlaceOrderHandler {
	o := &PlaceOrderHandler{}
	var sub iReqHandler[PlaceOrderReq, *PlaceOrderResult] = o
	o.reqHandlerBase = newReqHandlerBase(sub)
	return o
}

type PlaceOrderHandler struct {
	*reqHandlerBase[PlaceOrderReq, *PlaceOrderResult]
}

func (h *PlaceOrderHandler) doHandle(reqCtx *RequestContext, req *PlaceOrderReq) *PlaceOrderResult {
	fmt.Printf("PlaceOrderHandler: do handle req: %v\n", req)
	return &PlaceOrderResult{
		Status: "ok",
	}
}

type PlaceOrderReq struct {
	URL  string         `json:"url"`
	Body map[string]any `json:"body"`
}

type PlaceOrderResult struct {
	Status string
}

func TestPlaceOrderHandler(t *testing.T) {
	h := newPlaceOrderHandler()
	bodyJSON, err := json.Marshal(map[string]any{
		"url": "https://openapi.com/place_order",
		"body": map[string]any{
			"orderItems": "this is a order item list",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	reqCtx := &RequestContext{Path: "/", Body: string(bodyJSON)}
	h.handle(reqCtx)
}
