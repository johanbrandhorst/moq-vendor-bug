package main

import "github.com/lib/pq"

//go:generate moq -out generated.go -pkg main ./ MyInterface

type MyInterface interface {
	DoSomething(pq.ByteaArray) pq.BoolArray
}

func main() {
}
