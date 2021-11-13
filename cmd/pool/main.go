package main

import "log"

type type1 struct {
	Field1 string
	Field2 int
}
type type2 struct {
	Field1 string
	Field2 int
}

func main() {
	c := []int{1, 2, 3, 4, 5, 6, 7}
	c = append(c[:2], c[3:]...)
	log.Println(c)
}
