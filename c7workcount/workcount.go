package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type kv struct { /* to sort the map*/
	key   string
	value int
}

func textFormat(text string) []string {
	text = strings.ToLower(text)    /*to lower case*/
	text = strings.TrimSpace(text)  /*remove space at the beginen and at the end*/
	return strings.Split(text, " ") /*create slice per space*/
}

func regExpText(text string) string {
	r, _ := regexp.Compile("[^a-zA-Z0-9 ]") /*letterNumbers :=ra-zA-Z0-9]*/
	return r.ReplaceAllString(text, "")
}

/*this function count  words and return each word and how many time they appear*/
func WordCount(text string) {
	//
	wordTotal := make(map[string]int) /*creating a map to return each word and how many times there are*/
	kvSlice := []kv{}                 /* createing slice of key value to sort the map*/

	ohneSpace := []string{} /* create slice with all word separate with space*/

	text = regExpText(text)      /*remove all but letter and number using regexp*/
	ohneSpace = textFormat(text) /* func to format text and return []string*/
	//

	for _, v := range ohneSpace {
		wordTotal[v]++
	}

	kvSlice = sortMap(wordTotal) /*sorting map*/
	printSortedMap(kvSlice)

}
func printSortedMap(s []kv) {
	for _, v := range s {
		fmt.Printf("%v:%v\n", v.key, v.value)
	}
}
func sortMap(map1 map[string]int) []kv {

	kvSlice := []kv{}
	for k, v := range map1 {

		kvSlice = append(kvSlice, kv{k, v})
		// fmt.Printf("Key %v Value %v \n", k, v)
	}
	sort.Slice(kvSlice, func(x, y int) bool {
		return kvSlice[x].value > kvSlice[y].value
	})
	return kvSlice

}
func main() {
	WordCount(`The code for a column-span cell looks like a regular <td> cell, except for the code
	colspan=”[number of columns to span]”. The closing tag is the same as for
	a regular <td> cell. Notice that a <td> with the colspan feature replaces the same
	number of regular <td>s as the number of columns that are spanned. In the first row,
	there are three regular <td>s. In the second row, where two columns are spanned, there’s
	one regular <td> plus the span. In the third row, where three columns are spanned, there’s
	no regular <td>. You can make table headings span columns, too. The code is...`)
	fmt.Println("=====================================")
	fmt.Println("=============================================================")

	WordCount(`For anyone wanting things to work like they do for Omar, I could not get
	 it to work unless i first did $ git config credential.helper store. Note: The credentials
	  will be saved unencrypted on a file inside your home directory, therefore use it with discretion.`)

}
