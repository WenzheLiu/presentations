package main

import (
	/*
		explain why I'm importing fmt
	*/
	"fmt"
	/*
		explain why I'm importing pprof
	*/
	_ "net/http/pprof"
)

func main() {
	fmt.Println("Seattle, WA ☔")
}
