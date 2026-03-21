package multipackage

import bar "github.com/mus-format/mus-gen-skill-go/test/mus-go/multi-package/sub"

// Should use a serializer from another package.

type Foo struct {
	bar.Bar
}

type Zoo bar.Bar
