package main

import (
	"github.com/valyala/fasthttp"

	"github.com/alexpts/go-next/next"
)

func main() {
	apiV1 := next.NewApp()
	apiV1.Get(`/users/`, next.Config{}, func(ctx *next.HandlerCxt) {
		ctx.Response.AppendBodyString(`v1 - users`)
	})
	//...

	apiV2 := next.NewApp()
	apiV2.Get(`/users/`, next.Config{}, func(ctx *next.HandlerCxt) {
		ctx.Response.AppendBodyString(`v2 - users`)
	})
	// ...

	reuseApp := next.NewApp()
	reuseApp.Get(`/users/`, next.Config{}, func(ctx *next.HandlerCxt) {
		ctx.Response.AppendBodyString(`reuse - users`)
	})
	//...

	app := next.NewApp()
	app.Mount(apiV1, `/api/v1`).
		Mount(apiV2, `/api/v2`).
		Mount(reuseApp, ``)

	server := &fasthttp.Server{
		Handler: app.FasthttpHandler,
	}

	_ = server.ListenAndServe(":3000")
}
