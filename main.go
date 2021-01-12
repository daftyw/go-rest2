package main

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

// MyHandler : for Structured
type MyHandler struct {
}

// HandleFastHTTP : Handler for HTTP
func (h *MyHandler) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("content-type", "application/json")
	switch string(ctx.Path()) {
	case "/ping":
		h.handlerPing(ctx)
	default:
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	}
}

func (h *MyHandler) handlerPing(ctx *fasthttp.RequestCtx) {
	jsonOut, _ := json.Marshal(struct {
		message string
	}{
		message: "pong",
	})
	ctx.Response.SetBody(jsonOut)
}

func main() {
	myHandler := &MyHandler{}
	fasthttp.ListenAndServe(":8080", myHandler.HandleFastHTTP)
}
