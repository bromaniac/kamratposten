package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func readItem(n int64) item {
	file, err := ioutil.ReadFile("items/" + strconv.FormatInt(n, 10))
	check(err)

	data := item{}

	err = json.Unmarshal([]byte(file), &data)
	check(err)

	return data
}

func writeItem(i item, n int64) {
	i.ID = n

	if i.Parent == 0 {
		f, err := os.OpenFile("threads.dat", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		check(err)
		defer f.Close()
		f.WriteString(fmt.Sprintf("%d\n", n))
		i.Kind = "thread"
	}

	b, err := json.Marshal(i)
	check(err)

	err = ioutil.WriteFile("items/"+strconv.FormatInt(n, 10), b, 0644)
	check(err)
}
