package c1anagram

import (
	"fmt"
	"strings"
)

func palindrom(str1 string, str2 string) bool {
	lowerStr1 := []byte(strings.ToLower(str1))
	lowerStr2 := []byte(strings.ToLower(str2))
	if len(str1) != len(str2) {
		return false
	}
	temStr := make([]byte, len(lowerStr1))
	for i := 0; i < len(lowerStr1); i++ {

		temStr[i] = lowerStr1[len(lowerStr1)-i-1]
		//fmt.Println(string(lowerStr1))
		//fmt.Println(string(temStr))
	}
	//fmt.Println(string(temStr), string(lowerStr2))
	return string(temStr) == string(lowerStr2)
}

func anagram(str1 string, str2 string) bool {
	lowerStr1 := []byte(strings.ToLower(str1))
	lowerStr2 := []byte(strings.ToLower(str2))
	if len(str1) != len(str2) {
		return false
	}
	anagram := false
	for _, v := range lowerStr1 {
		for _, h := range lowerStr2 {
			if v == h {
				anagram = true

			} else {
				anagram = false
			}
		}
	}
	return anagram
}

func Palindrom(str1, str2 string) {

	pal := palindrom(str1, str2)
	if pal {
		fmt.Println("Are Palindrom")
	} else {
		fmt.Println("Not Palindrom")
	}

}

func Anagram(str1, str2 string) {

	ana := anagram(str1, str2)
	if ana {
		fmt.Println("Are Anagram")
	} else {
		fmt.Println("Not Anagram")
	}

}
