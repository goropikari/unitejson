package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
)

var re = regexp.MustCompile(`//.*\n`)

func removeComment(file []byte) []byte {
	return re.ReplaceAll(file, []byte{})
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal(fmt.Errorf("usage: unitejson {json file1} {json file2} ..."))
		return
	}

	files := [][]byte{}
	for i, filepath := range os.Args {
		if i == 0 {
			continue
		}

		data, err := os.ReadFile(filepath)
		if err != nil {
			log.Fatal(err)
			return
		}
		files = append(files, removeComment(data))
	}

	j, err := UniteJSON(files)
	if err != nil {
		log.Fatal(err)
		return
	}
	unitefile, err := json.Marshal(j)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(unitefile))
}
