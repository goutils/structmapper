package structmapper

import "testing"

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

type E struct {
	ArrayField []int
}

type F struct {
	B []B
}

type DestinationB struct {
	Field1 string
	Field2 int
}

type DestinationF struct {
	B []DestinationB
}

func TestAutoMapWorksForSameFieldSet(t *testing.T) {
	from := A{"a", "1"}
	to := A{}

	err := AutoMap(from, &to)

	if from != to || err != nil {
		t.Fail()
	}
}

//TODO: Intentinally ignored different types
func TestAutoMapWorksForDifferentFieldSet(t *testing.T) {
	from := A{"a", "1"}
	to := B{}

	err := AutoMap(from, &to)

	if from.Field1 != to.Field1 || err != nil {
		t.Fail()
	}
}

func TestAutoMapWorksForEmbeddedFieldSet(t *testing.T) {
	from := C{B{"a", 1}}
	to := D{}

	err := AutoMap(from, &to)

	if from.B.Field1 != to.B.Field1 || err != nil {
		t.Fail()
	}

}

func TestAutoMapCanCopyArrayFields(t *testing.T) {
	from := E{[]int{1, 2, 3}}
	to := E{[]int{}}

	err := AutoMap(from, &to)

	if err != nil || len(from.ArrayField) != len(to.ArrayField) {
		t.Fail()
	}
}

func TestAutoMapCanCopyArrayStructFields(t *testing.T) {
	from := F{[]B{B{Field1: "1", Field2: 2}}}
	to := DestinationF{[]DestinationB{}}

	err := AutoMap(from, &to)

	if err != nil || len(from.B) != len(to.B) {
		t.Fail()
	}
}

func TestAutoMapCanHandleNil(t *testing.T) {
	to := A{}

	err := AutoMap(nil, &to)
	if err.Error() != "Cannot map nil" {
		t.Fail()
	}
}
