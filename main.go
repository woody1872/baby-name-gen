package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strings"
	"time"
)

//go:embed boys.txt
var bf string

//go:embed girls.txt
var gf string

func main() {
	var gender string
	flag.StringVar(&gender, "gender", "", "gender of name to generate")
	var numNames int
	flag.IntVar(&numNames, "results", 1, "number of results to return")
	flag.Parse()

	if gender == "" {
		fmt.Print("Enter gender (b/boy or g/girl): ")
		fmt.Scanln(&gender)
	}

	bNames := strings.Split(bf, " ")
	bNamesLen := len(bNames)
	gNames := strings.Split(gf, " ")
	gNamesLen := len(gNames)

	fmt.Println(bNamesLen, bNames, reflect.TypeOf(bNames))
	os.Exit(0)

	min := 1
	var max int
	var namesList []string
	gender = strings.ToLower(gender)
	if gender == "b" || gender == "boy" {
		namesList = bNames
		max = bNamesLen
	} else if gender == "g" || gender == "girl" {
		namesList = gNames
		max = gNamesLen
	} else {
		log.Fatalln("Invalid gender")
	}

	rand.Seed(time.Now().UnixNano())
	randNums := make([]int, 0, max)
	for i := 0; i < numNames; i++ {
		randNums = append(randNums, rand.Intn(max-min)+min)
	}

	for i, v := range namesList {
		for _, rn := range randNums {
			if i == rn {
				fmt.Println(v[i])
			}
		}
	}
}
