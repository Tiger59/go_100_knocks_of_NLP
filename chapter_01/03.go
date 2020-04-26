package main

import (
	"fmt"
	"strings"
)

func main(){
	s := []int{}
	sentence:= "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	sentence = strings.Replace(sentence, ",", "", -1)
	sentence = strings.Replace(sentence, ".", "", -1)
	words 	:= strings.Split(sentence," ")
	 for i:=0 ;i<len(words);i++ {
		 s =append(s,len(words[i]))
	 }
	 fmt.Print(s)
}
