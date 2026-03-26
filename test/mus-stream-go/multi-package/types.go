package multipackage

import bar "github.com/mus-format/mus-skill-go/test/mus-stream-go/multi-package/sub"

// Should use a serializer from another package.

type Foo struct {
	bar.Bar
}

type Zoo bar.Bar
