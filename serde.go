package main

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

func readItem(i int64) item {
	file, err := ioutil.ReadFile(strconv.FormatInt(i, 10) + ".json")
	check(err)

	data := item{}

	err = json.Unmarshal([]byte(file), &data)
	check(err)

	return data
}

func writeItem(i item) {
	b, err := json.Marshal(i)
	check(err)

	err = ioutil.WriteFile("1.json", b, 0644)
	check(err)
}
