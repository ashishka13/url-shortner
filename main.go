package main

import "url-shortner/handler"

var Url = "https://youtube.com/@SpaceX"

var UrlMap = make(map[string]string)

const TinyUrlPrefix = "tu.ai/"

func main() {
	UrlMap["gu"] = "https://google.com"
	handler.StartController(UrlMap)
}
