package main

import "fmt"

func search(states_needed map[string]bool, stations map[string]map[string]bool) []string {
	final_stations := []string{}
	for len(states_needed) > 0 {
		best_station := ""
		states_covered := []string{}
		for key2, value2 := range stations {
			covered := findIntersection(states_needed, value2)
			if len(covered) > len(states_covered) {
				best_station = key2
				states_covered = covered
			}
		}
		for _, value3 := range states_covered {
			delete(states_needed, value3)
		}
		final_stations = append(final_stations, best_station)
		delete(stations, best_station)
	}

	return final_stations
}

func findIntersection(arr1, arr2 map[string]bool) []string {
	set := make(map[string]bool)
	var intersection []string

	for key, _ := range arr1 {
		set[key] = true
	}

	for key, _ := range arr2 {
		if set[key] {
			intersection = append(intersection, key)
		}
	}

	return intersection
}

func main() {
	states_needed := map[string]bool{
		"mt": true,
		"wa": true,
		"or": true,
		"id": true,
		"nv": true,
		"ut": true,
		"ca": true,
		"az": true,
	}
	stations := make(map[string]map[string]bool)
	stations["kone"] = map[string]bool{"id": true, "nv": true, "ut": true}
	stations["ktwo"] = map[string]bool{"wa": true, "id": true, "mt": true}
	stations["kthree"] = map[string]bool{"or": true, "nv": true, "ca": true}
	stations["kfour"] = map[string]bool{"nv": true, "ut": true}
	stations["kfive"] = map[string]bool{"ca": true, "az": true}
	result := search(states_needed, stations)
	fmt.Println(result)
}
