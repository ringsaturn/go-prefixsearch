package prefixsearch_test

import (
	"fmt"
	"sort"

	"github.com/ringsaturn/go-prefixsearch"
)

// ExampleAutoComplete just creates an object and does simple test
func ExampleSearchTree_AutoComplete() {
	st := prefixsearch.New[int]()
	st.Add("Hello world!", 1)
	st.Add("New impressions", 2)
	st.Add("Hello golang!", 3)
	st.Add("Привет, мир!", 4)
	st.Add("Надо же :)", 5)

	// converting []interface{} to []int to use sort.Ints function :)
	intResult := []int{}
	for _, x := range st.AutoComplete("HE") {
		intResult = append(intResult, x)
	}

	sort.Ints(intResult)
	fmt.Println(intResult)
	// Output: [1 3]
}

type Item struct {
	ID      int
	Name    string
	Comment string
}

func (i *Item) LessThan(j Item) bool { return i.ID < j.ID }
func (i *Item) EqualTo(j Item) bool  { return i.ID == j.ID }

// Support of unicode symbols and using struct as value
func ExampleSearchTree_AutoComplete_unicode() {

	// Item define as:
	// 	type Item struct {
	// 		ID      int
	// 		Name    string
	// 		Comment string
	// 	}

	// 	func (i *Item) LessThan(j Item) bool { return i.ID < j.ID }
	// 	func (i *Item) EqualTo(j Item) bool  { return i.ID == j.ID }

	data := []Item{
		{1, "Hello world!", "First example"},
		{2, "New impressions", "Second example"},
		{3, "Hello golang!", "Some other important info"},
		{4, "Привет, мир!", "Unicode symbols also work"},
		{5, "こんにちは世界", "Even this one may work"},
	}

	st := prefixsearch.New[Item]()
	for _, x := range data {
		st.Add(x.Name, x)
	}

	fmt.Println(st.AutoComplete("こん"))
	// Output: [{5 こんにちは世界 Even this one may work}]
}

// ExampleSearch shows another possible usage of this package
func ExampleSearchTree_Search() {
	st := prefixsearch.New[int]()
	st.Add("Hello world!", 1)
	st.Add("New impressions", 2)
	st.Add("Hello golang!", 3)
	st.Add("Привет, мир!", 4)
	st.Add("Надо же :)", 5)

	fmt.Println(st.Search("HE"))
	fmt.Println(st.Search("HELLO WORLD"))
	fmt.Println(st.Search("HELLO WORLD!"))
	fmt.Println(st.Search("HELLO WORLD!!"))

	// Output:
	// 0
	// 0
	// 1
	// 0
}
