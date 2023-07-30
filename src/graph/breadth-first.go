package main

import (
	"container/list"
	"fmt"
)

func search(graph map[string][]string, name string) bool {
	queue := list.New()
	for _, elem := range graph[name] {
		queue.PushBack(elem)
	}
	searched := []string{}
	for queue.Len() > 0 {
		person := queue.Remove(queue.Back()).(string)
		if !is_found(searched, person) {
			if person_is_seller(person) {
				fmt.Printf("%s is a seller", person)
				return true
			} else {
				for _, elem := range graph[person] {
					queue.PushBack(elem)
				}
				searched = append(searched, person)
			}
		}
	}
	return false
}

func person_is_seller(person string) bool {
	flag := false
	if person == "thom" {
		flag = true
	}
	return flag
}

func is_found(searched []string, name string) bool {
	flag := false
	for _, elem := range searched {
		if name == elem {
			flag = true
		}
	}
	return flag
}

func main() {
	graph := make(map[string][]string)
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}
	fmt.Printf(" - %t", search(graph, "you"))
}
