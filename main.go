package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var Name string
	var Color string
	var Attitude string
	var Word string
	words := []string{"so boring",
		"fuck you",
		"kill then all",
		"boo ~~~~~~~~~~",
		"no ~~~~~",
		"I am the Lord of Darkness",
		"Akira is best!!!!!",
		"I am the master of the world",
		"no !",
		"shifffffffffft!!!!!!!!",
		"by jim",
		"test,test",
		"php have the world",
		"you can call me , my number is : fuck you ",
		"go away",
		"golang 16",
		"i like it"}

	flag.StringVar(&Name, "name", "boring", "cat name")
	flag.StringVar(&Color, "color", "black", "cat color")
	flag.StringVar(&Attitude, "attitude", "", "cat attitude")
	flag.StringVar(&Word, "word", "boring!", "cat Word")

	flag.Parse()

	if Word != "" {
		words = append(words, Word)
	}
	rand.Seed(time.Now().Unix()) //Seed生成的随机数
	i := random(1, len(words))
	if i > len(words) {
		i = len(words) - 1
	}
	fmt.Println(Color + " " + Name + " say :" + words[i])
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
