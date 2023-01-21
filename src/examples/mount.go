package main

import (
	"github.com/valyala/fasthttp"

	"github.com/alexpts/go-next/next"
	"github.com/alexpts/go-next/next/layer"
)

func main() {
	apiV1 := next.ProvideMicroApp(nil, nil)
	apiV1.Get(`/users/`, layer.Layer{}, func(ctx *layer.HandlerCtx) {
		ctx.Response.AppendBodyString(`v1 - users`)
	})
	//...

	apiV2 := next.ProvideMicroApp(nil, nil)
	apiV2.Get(`/users/`, layer.Layer{}, func(ctx *layer.HandlerCtx) {
		ctx.Response.AppendBodyString(`v2 - users`)
	})
	// ...

	reuseApp := next.ProvideMicroApp(nil, nil)
	reuseApp.Get(`/users/`, layer.Layer{}, func(ctx *layer.HandlerCtx) {
		ctx.Response.AppendBodyString(`reuse - users`)
	})
	//...

	app := next.ProvideMicroApp(nil, nil)
	app.Mount(apiV1, `/api/v1`).
		Mount(apiV2, `/api/v2`).
		Mount(reuseApp, ``)

	server := &fasthttp.Server{
		Handler: app.FastHttpHandler,
	}

	_ = server.ListenAndServe(":3000")
}
