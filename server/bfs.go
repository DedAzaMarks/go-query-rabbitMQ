package main

import (
	"fmt"
	"hw05-query/config"
	"log"
)

func reverse(s []string) []string {
	first := 0
	last := len(s) - 1
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
	return s
}

func bfs(request *config.Request) (int, []string, error) {
	from := request.From
	to := request.To

	if from == to {
		return 0, []string{from}, nil
	}

	previous := make(map[string]string)
	queue := []string{from}
	pathFound := false

queue:
	for len(queue) != 0 {
		cur := queue[0]
		log.Printf("current page: %s", cur)
		queue = queue[1:]
		neighbours, err := getNeighbours(cur)
		if err != nil {
			log.Printf("ERROR: %v", err)
		}
		for _, neighbour := range neighbours {
			if _, ok := previous[neighbour]; !ok {
				previous[neighbour] = cur
				queue = append(queue, neighbour)
			}
			if neighbour == to {
				pathFound = true
				break queue
			}
		}
	}

	if !pathFound {
		return 0, nil, fmt.Errorf("no path from: %s to: %s", from, to)
	}

	var res []string
	cur := to
	for cur != from {
		res = append(res, cur)
		cur = previous[cur]
	}
	res = append(res, from)
	log.Printf("FOUND!")
	return len(res), reverse(res), nil
}
