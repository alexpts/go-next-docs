package examples

import (
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/alexpts/go-next/next"
)

type RouterConfig struct {
	Path         string
	Methods      string
	Controller   string
	Name         string
	Priority     int
	Handler      next.Handler
	Restrictions next.Restrictions
}

var HandlerMap = map[string]next.Handler{
	"otherwise": MainPageAppHandler2,
	"mainPage":  MainPageAppHandler2,
	"mainPage2": MainPageAppHandler2,
	"hello":     MainPageAppHandler2,
}

func CreateLayers(projectDir string) []*next.Layer {
	rawData, err := os.ReadFile(projectDir + "/config/router.yml")

	var routes map[string]RouterConfig
	err = yaml.Unmarshal(rawData, &routes)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return factoryLayers(routes)
}

func factoryLayers(routes map[string]RouterConfig) []*next.Layer {
	var layers []*next.Layer
	var methods []string

	for name, route := range routes {
		handler := HandlerMap[route.Controller]
		if route.Name == `` {
			route.Name = name
		}

		if route.Methods != `` {
			methods = strings.Split(route.Methods, `|`) // GET|POST|PUT
		}

		layer := next.Layer{
			Handlers:     []next.Handler{handler},
			Priority:     route.Priority,
			Name:         route.Name,
			Path:         route.Path,
			Restrictions: route.Restrictions,
			Methods:      methods,
		}

		layers = append(layers, &layer)
	}

	return layers
}

func MainPageAppHandler2(ctx *next.HandlerCxt) {
	ctx.Response.AppendBodyString(`MainPageAppHandler2`)
}
