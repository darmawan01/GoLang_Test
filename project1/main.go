package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

func main() {
	fmt.Print("Enter the text: ")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")
	// remove the normal text symbol
	var re = regexp.MustCompile("[.,?]")
	input = re.ReplaceAllString(input, "")

	words := strings.Split(input, " ")

	tmp := make(map[string]int)
	for _, str := range words {
		key := strings.ToLower(str)
		tmp[key] = 0

		for _, str2 := range words {
			if key == strings.ToLower(str2) {
				tmp[key] += 1
			}
		}
	}

	p := make(PairList, len(tmp))
	i := 0
	for k, v := range tmp {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)

	out(p)
}

func out(p PairList) {
	f, err := os.OpenFile("output.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := make(map[string]int)
	for i, item := range p {
		if i+1 <= 10 {
			data[item.Key] = item.Value
		}

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
