package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	path := flag.String("path", "tensorJK.log", "Enter path string")                      // go run . <path>
	level := flag.String("level", "Error", " dilter Error or Warning or Info or Missing") //go run . <path> <filter>
	flag.Parse()                                                                          //go run . -help
	f, err := os.Open(*path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() //closing file after main func exits
	res := bufio.NewReader(f)
	for {
		s, err := res.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(s, *level) {
			fmt.Println(s)
		}
	}
}
