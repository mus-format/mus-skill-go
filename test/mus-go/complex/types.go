package complex

import (
	"image"

	"github.com/mus-format/mus-gen-skill-go/test/mus-go/complex/sub1"
	"github.com/mus-format/mus-gen-skill-go/test/mus-go/complex/sub2"
)

// Should be able to handle a complex case.

// mus:name = ComplexFooMUS
type Foo struct {
	// mus:vl = validators.ValidateInt
	Num int
	// mus:enc = raw
	RawNum int
	Str    ValidString
	Sl     CustomSlice
	Map    CustomMap
	// mus:skip = true
	Skipped string
	Ptr     *int
	Bar1    sub1.Bar
	Bar2    sub2.Bar
	// mus:path = github.com/mus-format/mus-gen-skill-go/test/mus-go/complex/sers
	Point image.Point
	Embedded
	// mus:lenSer = raw.Int
	// mus:elem:lenSer = raw.Int
	// mus:elem:elem:vl = ValidateString
	sl [][]string
}

// -----------------------------------------------------------------------------

// mus:vl = validators.ValidateString
type ValidString string

// mus:lenSer = raw.Int
// mus:key:vl = validators.ValidateMapKey
// mus:value:elem:enc = raw
type CustomMap map[string][]int

// mus:lenSer = raw.Int
// mus:elem:vl = validators.ValidateSliceElem
type CustomSlice []int

type Embedded struct {
	Bool bool
}

// mus:impls = Impl1, Impl2
type MyInterface interface {
	Method()
}

type Impl1 struct {
	A int
}

func (i Impl1) Method() {}

type Impl2 struct {
	B string
}

func (i Impl2) Method() {}
