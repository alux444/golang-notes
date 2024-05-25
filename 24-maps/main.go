package main

import (
	"fmt"
	"sort"
)

func maps() {
	// each map is a pointer to a map header
	// nil map - is readonly
	var dict map[string]string

	// literal maps
	map1 := map[string]string{
		"hello": "hi",
		"bye":   "goodbye",
		"test":  "testing",
	}
	map1["test"] = "t"

	map2 := map[string]string{}
	map2["hi"] = "wow"

	fmt.Printf("Number of keys: %d\n", len(dict))
	fmt.Println(map1)
	fmt.Println(map2)
}

func mapFuncs() {
	map1 := map[string]string{
		"hello": "hi",
		"bye":   "goodbye",
		"test":  "testing",
	}

	//checking existence
	value, ok := map1["random"]
	if !ok {
		fmt.Println(ok)
		fmt.Println(value)
	}

	//better syntax
	query := "bye"
	if value, ok := map1[query]; ok {
		fmt.Printf("%q means %#v\n", query, value)
	}

	delete(map1, "hello")
	delete(map1, "notkey")
	fmt.Println(map1)
}

func cloning() {
	dict := map[string]string{
		"good":    "iyi",
		"great":   "harika",
		"perfect": "mükemmel",
		"awesome": "mükemmel",
	}

	clone := make(map[string]string, len(dict))
	for key, value := range dict {
		clone[key] = value
	}
	fmt.Println(clone)
}

func exercise() {
	houses := map[string][]string{
		"gryffindor": {"weasley", "hagrid", "dumbledore", "lupin"},
		"hufflepuf":  {"wenlock", "scamander", "helga", "diggory", "bobo"},
		"ravenclaw":  {"flitwick", "bagnold", "wildsmith", "montmorency"},
		"slytherin":  {"horace", "nigellus", "higgs", "bobo", "scorpius"},
		"bobo":       {"wizardry", "unwanted"},
	}

	house, students := "ravenclaw", houses["ravenclaw"]
	if students == nil {
		fmt.Printf("Sorry. I don't know anything about %q.\n", house)
		return
	}

	clone := append([]string(nil), students...)
	sort.Strings(clone)

	fmt.Printf("~~~ %s students ~~~\n\n", house)
	for _, student := range clone {
		fmt.Printf("\t+ %s\n", student)
	}
}

func main() {
	maps()
	mapFuncs()
	cloning()
	exercise()
}
