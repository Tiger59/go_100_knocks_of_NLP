package main

import "fmt"

func main(){
	word1 := []rune("パトカー")
	word2 := []rune("タクシー")
	for i:=0; i<len(word1); i++{
		fmt.Print(string(word1[i])+string(word2[i]))
	}
}