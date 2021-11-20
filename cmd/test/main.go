package main

import "log"

func main() {
	s := &[]int{}
	slice(s)
	log.Printf("slice address %p, slice value %v", s, s)
}

func slice(s *[]int) {
	*s = append(*s, 1)
	log.Println(s)
}
