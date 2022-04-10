package config

import "fmt"

const HOST = "localhost"
const PORT = "42069"

var ADDRESS = fmt.Sprintf("%s:%s", HOST, PORT)

const QueueRequest = "request"
const QueueRespond = "respond"

const WikipediaWiki = "https://en.wikipedia.org/wiki/"
