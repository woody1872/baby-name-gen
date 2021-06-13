package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

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

	var nameFile string // don't remove
	var err error       // don't remove
	gender = strings.ToLower(gender)
	if gender == "b" || gender == "boy" {
		nameFile, err = filepath.Abs("boys.txt")
		if err != nil {
			log.Fatalln(err)
		}
	} else if gender == "g" || gender == "girl" {
		nameFile, err = filepath.Abs("boys.txt")
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Fatalln("Invalid gender")
	}

	f, err := os.Open(nameFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	min := 1
	max, err := lineCounter(f)
	if err != nil {
		log.Fatalln(err)
	}

	rand.Seed(time.Now().UnixNano())
	randNums := make([]int, 0, numNames)
	for i := 0; i < numNames; i++ {
		randNums = append(randNums, rand.Intn(max-min)+min)
	}

	f.Seek(0, 0)
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		i += 1
		for _, n := range randNums {
			if i == n {
				fmt.Println(scanner.Text())
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
