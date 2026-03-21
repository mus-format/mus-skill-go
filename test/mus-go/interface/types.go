package intr

import com "github.com/mus-format/common-go"

// Should generate an interface serializer for the following cases:
// 1. All DTMs already defined.
// 2. Partial DTMs defs.
// 3. No DTMs.

const (
	FooDTM com.DTM = iota + 1
	BarDTM
)

// mus:impls = Foo, Bar
type Interface1 any

// mus:impls = Car, Foo
type Interface2 any

// mus:impls = Zoo
type Interface3 any

// -----------------------------------------------------------------------------

type Foo struct {
	Num int
}

type Bar string

type Zoo bool

type Car struct {
	float float32
}
