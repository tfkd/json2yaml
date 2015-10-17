package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	indentString = "  "
	arrayString  = "- "
	keyString    = ": "
)

func printHeader() {
	fmt.Print("---\n\n")
}

func walkJson(o interface{}) {
	walkJsonInternal(o, 0, false, "")
}

func walkJsonInternal(o interface{}, indent int, isArray bool, key string) {
	for i := 0; i < indent; i++ {
		fmt.Print(indentString)
	}
	if isArray {
		fmt.Print(arrayString)
	}
	if key != "" {
		fmt.Print(key + keyString)
	}
	switch v := o.(type) {
	case []interface{}:
		if len(v) == 0 {
			fmt.Println("[]")
			return
		}
		if isArray || key != "" {
			fmt.Println("")
		}
		for _, vv := range v {
			walkJsonInternal(vv, indent+1, true, "")
		}
	case map[string]interface{}:
		if len(v) == 0 {
			fmt.Println("{}")
			return
		}
		if isArray || key != "" {
			fmt.Println("")
		}
		for k, vv := range v {
			walkJsonInternal(vv, indent+1, isArray, k)
		}
	case bool:
		fmt.Println(v)
	case float64:
		fmt.Println(v)
	case string:
		fmt.Printf("\"%s\"\n", v)
	default:
		fmt.Println("null")
	}
}

func readJson(f string) ([]byte, error) {
	if f == "-" {
		return ioutil.ReadAll(os.Stdin)
	}
	return ioutil.ReadFile(f)
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("Usage: json2yaml JSONFILE")
		return
	}

	b, err := readJson(flag.Arg(0))
	if err != nil {
		fmt.Println(err)
	}

	printHeader()
	var f interface{}
	err = json.Unmarshal(b, &f)
	if err != nil {
		fmt.Println(err)
	}
	walkJson(f)
}
