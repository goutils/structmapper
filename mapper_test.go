package main

import (
	"fmt"
	"testing"
)


type A struct {
	Field1 string
	Field2 string
}

type B struct {
	Field1 string
	Field2 int
}

type C struct {
	B B
}

type D struct {
	B B
}

func TestAutoMapWorksForSameFieldSet(t *testing.T) {
	from := A{"a","1"}

	result, err := AutoMap(from, A{})
	to := result.(A)

	if from != to || err != nil {
		t.Fail()
	}
	fmt.Println(from,to)
}

func TestAutoMapWorksForDifferentFieldSet(t *testing.T) {
	from := A{"a","1"}

	result, err := AutoMap(from, B{})
	to := result.(B)

	if from.Field1 != to.Field1 || err != nil {
		t.Fail()
	}
	fmt.Println(from,to)
}

func TestAutoMapWorksForEmbeddedFieldSet(t *testing.T) {
	from := C{B{"a",1}}
	to := D{}

	result, err := AutoMap(from,to)
	to = result.(D)

	if from.B != to.B || err != nil {
		t.Fail()
	}
	fmt.Println(from,to)
}
