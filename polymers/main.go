package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"unicode"
)

func main() {
	data, err := ioutil.ReadFile("../polymer.txt")
	if err != nil {
		log.Fatal(err)
	}
	result := day5Result(data)
	fmt.Println(result)
}

func day5Result(list []byte) map[string]int {
	letters := []string{"aA", "bB", "cC", "dD", "eE", "fF", "gG", "hH", "iI", "jJ", "kK", "lL", "mM", "nN", "oO", "pP", "qQ", "rR", "sS", "tT", "uU", "vV", "wW", "xX", "yY", "zZ"}
	res := make(map[string]int, 0)

	s := make([]rune, 0)
	for _, val := range string(list) {
		s = append(s, val)
	}
	result := loopOver(s)
	res["all"] = len(runeString(result))

	for _, val := range letters {
		input := strings.Replace(string(list), val[0:1], "", -1)
		input2 := strings.Replace(input, val[1:], "", -1)
		s := make([]rune, 0)
		for _, val := range input2 {
			s = append(s, val)
		}
		result := loopOver(s)
		res[val] = len(runeString(result))
	}
	return res
}

func runeString(x []rune) string {
	result := ""
	for _, val := range x {
		result += string(val)
	}
	return result
}

func loopOver(list []rune) []rune {
	l := list
	for index, char := range l {
		if index < len(l)-1 {
			if findInteraction(char, l[index+1]) {
				l = append(l[:index], l[index+1:]...)
				l = append(l[:index], l[index+1:]...)
				l = loopOver(l)
			}
		}
	}
	return l
}

func findInteraction(a, b rune) bool {
	if unicode.ToLower(a) == unicode.ToLower(b) {
		if unicode.IsUpper(b) && unicode.IsLower(a) { //|| (unicode.IsUpper(b) && unicode.IsLower(a)) {
			return true
		}
		if unicode.IsLower(b) && unicode.IsUpper(a) { //|| (unicode.IsUpper(b) && unicode.IsLower(a)) {
			return true
		}
	}
	return false
}
