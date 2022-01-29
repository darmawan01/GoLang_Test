package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"time"

	"github.com/zhexuany/wordGenerator"

	"log"
	"os"
)

func main() {
	arg := os.Args[1]

	if arg == "gen" {
		generate()
	} else {
		process(arg)
	}
}

func process(path string) {
	start := time.Now()

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var txts []string
	for scanner.Scan() {
		txts = append(txts, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	out(txts)

	fmt.Printf("\nTotal data %d", len(txts))
	fmt.Printf("\nEnded process %s ", time.Since(start).String())
}

func generate() {
	f, err := os.OpenFile("GoLang_Test.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 1000000; i++ {
		str := wordGenerator.GetWord(5)
		_, err = fmt.Fprintln(f, str)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func out(strs []string) {
	f, err := os.OpenFile("output.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := map[string]interface{}{
		"data": strs,
	}

	out, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	_, err = fmt.Fprintln(f, string(out))
	if err != nil {
		fmt.Println(err)
	}
}
