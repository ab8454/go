package main

import "fmt"

type part struct {
	description string
	count       int
}

func showInfo(p part) {
	fmt.Println("Description : ", p.description)
	fmt.Println("Count :", p.count)
}

func minimunOrder (description string) part{
	var p part
	p.description = description
	p.count = 1000
	return p
}

func main() {

	p:=minimunOrder("Hex bolts")
	fmt.Println(p.description,p.count)
/*
	var bolts part
	bolts.description = "Hex Bolts"
	bolts.count = 100

	showInfo(bolts)
	*/
}
