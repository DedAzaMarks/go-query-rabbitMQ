package main

import (
	"flag"
	"hw05-query/config"
	"log"
	"net/rpc"
)

func main() {
	from := flag.String("from", "https://en.wikipedia.org/wiki/Kernick", "start page of wikipedia")
	to := flag.String("to", "https://en.wikipedia.org/wiki/French_language", "end page of wikipedia")
	flag.Parse()

	fromTitle, err := config.UrlToTitle(*from)
	config.FailOnError(err, "can't get from page title")
	toTitle, err := config.UrlToTitle(*to)
	config.FailOnError(err, "can't get to page title")

	request := config.Request{From: fromTitle, To: toTitle}
	log.Printf("request:%+v\n", request)

	client, err := rpc.DialHTTP("tcp", config.ADDRESS)
	config.FailOnError(err, "client DialHTTP tcp error")

	var respond config.Respond
	config.FailOnError(client.Call("Pathfinder.GetPath", request, &respond), "client call error")
	log.Printf("%+v", respond)
}
