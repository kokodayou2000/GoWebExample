package main

import (
	"fmt"
)

type Middleware func(string) string

func A(out string) Middleware {

	return func(in string) string {
		fmt.Println("A->> " + "in- " + in + "out " + out)
		return " A "
	}
}

func B(out string) Middleware {

	return func(in string) string {
		fmt.Println("B->> " + "in- " + in + "out " + out)
		return " B "
	}
}

func Chain(str string, middleware ...Middleware) string {
	for _, m := range middleware {
		str += " " + m(str)
	}
	return str
}

func main() {
	str := Chain("chain", A("< A >"), B("< B >"))
	fmt.Println("end " + str)
}
