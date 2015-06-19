package main

import (
	"fmt"
	"math/rand"
	"time"
)

// return random metasyntactic variable
func main() {
	m := []string{
		"foo",
		"bar",
		"baz",
		"foobar",
		"hoge",
		"piyo",
		"fuga",
		"kuke",
		"hogehoge",
		"hage",
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(m[r.Intn(len(m))])
}
