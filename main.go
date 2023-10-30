package main

import (
	"fmt"
	"strings"
)

func c() {
	f := fullName{firstName: "Sasha"}
	fmt.Println(f.getFirstName())
}

type fullName struct {
	firstName string
}

func (fn *fullName) getFirstName() string {
	return fn.firstName
}

func Title(s string) string {
	return strings.Title(s)
}
