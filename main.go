package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

var (
	Json   []byte
	Output map[string]interface{}

	FlattenArg bool
	KeyArg     bool
	SearchArg  string
)

func Flatten() (out map[string]interface{}) {
	out = make(map[string]interface{})

	var flattener func(prefix string, m interface{})

	flattener = func(prefix string, m interface{}) {
		switch m.(type) {
		case map[string]interface{}:
			for k, v := range m.(map[string]interface{}) {
				p := ""
				if prefix != "" {
					p = prefix + "/"

				}
				flattener(p+k, v)
			}
		// case []string:
		// 	for i, v := range m.([]string) {
		// 		p := ""
		// 		if prefix != "" {
		// 			p = prefix + "/"
		// 		}

		// 		flattener(p+string(i), v)
		// 	}
		default:
			out[prefix] = m
		}

	}

	flattener("", Output)

	return out
}

func Search() {

}

func init() {
	flag.BoolVar(&FlattenArg, "f", false, "Flatten the JSON tree")
	flag.BoolVar(&KeyArg, "k", true, "Show the Keys only")
	flag.StringVar(&SearchArg, "s", "", "Search the JSON tree")
}

func main() {
	var (
		err error
	)

	flag.Parse()

	args := flag.Args()

	// Parse Json
	Json, err = ioutil.ReadFile(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read file %s", args[0])
		os.Exit(-1)
	}

	err = json.Unmarshal(Json, &Output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to unmarshal %s", err)
		os.Exit(-1)
	}

	if FlattenArg {
		out := Flatten()
		var keys []string
		for k, _ := range out {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			if SearchArg != "" {
				if strings.Contains(k, SearchArg) {
					fmt.Println(k)
				}
			} else {
				fmt.Println(k)
			}
		}
		return
	}

}
