package main

import (
	"github.com/valyala/fasthttp"
)

// MyHandler : for Structured
type MyHandler struct {
}

// HandleFastHTTP : Handler for HTTP
func (h *MyHandler) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("content-type", "application/json")
	ctx.Response.SetBody([]byte("{\"message\":\"Hello\""))
}

func main() {
	myHandler := &MyHandler{}
	fasthttp.ListenAndServe(":8080", myHandler.HandleFastHTTP)
}
