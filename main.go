package main

import "fmt"

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
