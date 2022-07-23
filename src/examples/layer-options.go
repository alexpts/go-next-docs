package main

import (
	"github.com/valyala/fasthttp"

	"github.com/alexpts/go-next/next"
)

func main() {
	app := next.NewApp()

	handler := func(ctx *next.HandlerCxt) {
		name, _ := ctx.UriParams()["name"]
		ctx.Response.AppendBodyString(`Hello ` + name)
	}

	app.Use(next.Config{
		`Path`:     `/hello/{name}/`,
		`Name`:     `HelloAction`,
		`Methods`:  `GET|POST`,
		`Priority`: 100,
		`Restrictions`: next.Restrictions{
			`name`: `[a-z]+`,
		},
	}, handler)

	server := &fasthttp.Server{
		Handler: app.FasthttpHandler,
	}

	_ = server.ListenAndServe(":3000")
}
