package main

import (
	"fmt"
	"math"
)

func search(weights map[string]map[string]float64, costs map[string]float64, parents map[string]string) {
	processed := []string{}
	node := find_lowest_cost_node(costs, processed)
	for node != "" {
		cost := costs[node]
		neighbors := weights[node]
		for key, _ := range neighbors {
			new_cost := cost + neighbors[key]
			if costs[key] > new_cost {
				costs[key] = new_cost
				parents[key] = node
			}
		}
		processed = append(processed, node)
		node = find_lowest_cost_node(costs, processed)
	}
	fmt.Println(costs["fin"])
	key := "fin"
	fmt.Printf("fin - ")
	for i := len(parents) - 1; i >= 0; i-- {
		if key != "fin" && key != "start" {
			fmt.Printf("%s - ", key)
		}
		key = parents[key]

	}
	fmt.Printf("start")
}

func find_lowest_cost_node(costs map[string]float64, processed []string) string {
	lowest_cost := math.Inf(1)
	lowest_cost_node := ""
	for keyCost, _ := range costs {
		cost := costs[keyCost]
		flag := false
		if cost < lowest_cost {
			for keyProc, _ := range processed {
				if processed[keyProc] == keyCost {
					flag = true
					break
				}
			}
			if !flag {
				lowest_cost = cost
				lowest_cost_node = keyCost
			}
		}
	}
	return lowest_cost_node
}

func main() {
	weights := make(map[string]map[string]float64)
	weights["start"] = make(map[string]float64)
	weights["start"]["a"] = 6
	weights["start"]["b"] = 2
	weights["a"] = make(map[string]float64)
	weights["a"]["fin"] = 1
	weights["b"] = make(map[string]float64)
	weights["b"]["a"] = 3
	weights["b"]["fin"] = 5
	weights["fin"] = make(map[string]float64)

	costs := make(map[string]float64)
	costs["a"] = 6
	costs["b"] = 2
	costs["fin"] = math.Inf(1)

	parents := make(map[string]string)
	parents["a"] = "start"
	parents["b"] = "start"
	parents["fin"] = ""

	search(weights, costs, parents)
}
