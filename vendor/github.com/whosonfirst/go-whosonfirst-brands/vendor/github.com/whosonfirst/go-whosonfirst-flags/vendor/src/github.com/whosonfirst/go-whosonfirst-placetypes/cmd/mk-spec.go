package main

// As in: go run cmd/mk-spec.go > placetypes/spec.go

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {

	latest_spec := "https://raw.githubusercontent.com/whosonfirst/whosonfirst-placetypes/master/data/placetypes-spec-latest.json"

	spec := flag.String("spec", latest_spec, "...")

	flag.Parse()

	rsp, err := http.Get(*spec)
	defer rsp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(rsp.Body)

	if err != nil {
		log.Fatal(err)
	}

	ts := time.Now()

	fmt.Printf("%s\n\n", "package placetypes")

	fmt.Printf("/* %s */\n", *spec)
	fmt.Printf("/* This file was generated by robots (%s) at %s */\n\n", "cmd/mk-spec.go", ts.UTC())
	fmt.Printf("const Specification string = `%s`", strings.Trim(string(body), "\n"))
}
