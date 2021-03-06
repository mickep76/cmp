[![GoDoc](https://godoc.org/github.com/ake-persson/typecomp?status.svg)](https://godoc.org/github.com/ake-persson/typecomp)
[![codecov](https://codecov.io/gh/ake-persson/typecomp/branch/master/graph/badge.svg)](https://codecov.io/gh/ake-persson/typecomp)
[![Build Status](https://travis-ci.org/ake-persson/typecomp.svg?branch=master)](https://travis-ci.org/ake-persson/typecomp)
[![Go Report Card](https://goreportcard.com/badge/github.com/ake-persson/typecomp)](https://goreportcard.com/report/github.com/ake-persson/typecomp)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/ake-persson/mlfmt/blob/master/LICENSE)

# cmp

Package provides comparison for different Go types and a Comparer interface when using structs.

```go
package main

import (
	"fmt"
	"log"

	cmp "github.com/ake-persson/kvstore/cmp"
)

type Car struct {
	Manufacturer string
	Model        string
}

type Cars []*Car

func (c Car) String() string {
	return c.Manufacturer + " " + c.Model
}

func (c Car) Eq(b interface{}) (bool, error) {
	return c.String() == b.(Car).String(), nil
}

func (c Car) Lt(b interface{}) (bool, error) {
	return c.String() < b.(Car).String(), nil
}

func main() {
	cars := Cars{
		&Car{
			Manufacturer: "Audi",
			Model:        "Q3",
		},
		&Car{
			Manufacturer: "Audi",
			Model:        "Q5",
		},
	}

	ok, err := cmp.Eq(cars[0], cars[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ok: %v\n", ok)

	ok, err = cmp.Eq(cars[0], cars[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ok: %v\n", ok)
}
```
