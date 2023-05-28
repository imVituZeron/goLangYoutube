package main

import (
	"fmt"
	"sort"
)

type user struct {
	Name    string
	Last    string
	Age     int
	Sayings []string
}

type byAge []user

func (a byAge) Len() int           { return len(a) }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type byLast []user

func (l byLast) Len() int           { return len(l) }
func (l byLast) Less(i, j int) bool { return l[i].Last < l[j].Last }
func (l byLast) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

func main() {
	u1 := user{
		Name: "Jame",
		Last: "Jond",
		Age:  32,
		Sayings: []string{
			"ddd",
			"zsdfsadfsf",
			"ldasdfsf",
		},
	}

	u2 := user{
		Name: "Bayek",
		Last: "Egpty",
		Age:  22,
		Sayings: []string{
			"Aya",
			"Snake",
			"Bristo",
		},
	}

	users := []user{u1, u2}

	for _, v := range users {
		sort.Strings(v.Sayings)
	}

	sort.Sort(byAge(users))
	showDetails(users)

	sort.Sort(byLast(users))
	fmt.Println(users)
}

func showDetails(A []user) {
	for i, v := range A {
		fmt.Println(i, "| Full name:", v.Name, v.Last, "| Age:", v.Age)
		for _, value := range v.Sayings {
			fmt.Println("\t", "Sayings:", value)
		}
		fmt.Println("\n")
	}
}
