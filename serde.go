package main

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

func readItem(n int64) item {
	file, err := ioutil.ReadFile("items/" + strconv.FormatInt(n, 10) + ".json")
	check(err)

	data := item{}

	err = json.Unmarshal([]byte(file), &data)
	check(err)

	return data
}

func writeItem(i item, n int64) {
	b, err := json.Marshal(i)
	check(err)

	err = ioutil.WriteFile("items/" + strconv.FormatInt(n, 10)+".json", b, 0644)
	check(err)
}
