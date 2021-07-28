package main

import (
	"github.com/coolryze/crawler/engine"
	"github.com/coolryze/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
