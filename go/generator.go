package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func getDictFromJSON(filename string) map[string]interface{} {
	f, err := ioutil.ReadFile(filename)
	checkErr(err)

	var dict map[string]interface{}
	json.Unmarshal([]byte(f), &dict)

	return dict
}

func generateNamesFromDict(dict map[string]interface{}, count int) []string {
	result := make([]string, count)
	firsts := dict["first"].([]interface{})
	lasts := dict["last"].([]interface{})
	countFirstNames := len(firsts)
	countLastNames := len(lasts)

	// get rid of deterministic "random" sequence by new one
	randSeed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randSeed)

	for i := 0; i < count; i++ {
		randFirst := random.Intn(countFirstNames)
		randLast := random.Intn(countLastNames)
		result[i] = strings.Join([]string{firsts[randFirst].(string), lasts[randLast].(string)}, " ")
	}

	return result
}

func writeJSONToFile(arr []string, filename string) string {
	jsonOutput, err := json.Marshal(arr)
	checkErr(err)

	err = ioutil.WriteFile(filename, jsonOutput, 0644)
	checkErr(err)

	return filename
}

func simpleOutput (arr []string) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	outputPtr := flag.String("output", "print", "'print' -> simple output, 'json' -> to file")
	countPtr := flag.Int("count", 10, "count names")
	flag.Parse()

	dict := getDictFromJSON("../data/dict.json")

	names := generateNamesFromDict(dict, *countPtr)
	
	switch *outputPtr {
	case "json":
		filename := writeJSONToFile(names, "names.json")
		fmt.Println(filename)
	case "print":
		simpleOutput(names)
	default:
		simpleOutput(names)
	}
	
}
