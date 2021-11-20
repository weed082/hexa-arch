package main

import "log"

type child struct {
	userIdx int
}
type parent struct {
	child *child
}

func main() {
	child := &child{1}
	parent := &parent{child: child}
	log.Printf("%p", &parent.child)
	slice(parent)
}

func slice(parent *parent) {
	log.Printf("%p", &parent.child)
}
