package config

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func UrlToTitle(link string) (string, error) {
	u, err := url.Parse(link)
	if err != nil {
		return "", err
	}
	if strings.Contains(u.Host, "wikipedia.org") == false {
		return "", fmt.Errorf("link \"from\" is not wikipedia.org page")
	}
	if strings.HasPrefix(u.Path, "/wiki/") == false {
		return "", fmt.Errorf("link has no \"/wiki/\" prefix")

	}
	return strings.TrimPrefix(u.Path, "/wiki/"), nil
}

func TitleToUrl(title string) string {
	return WikipediaWiki + title
}
