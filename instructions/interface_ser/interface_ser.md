# Interface Serializer

To generate a serializer for an interface type:
1. If there is no `// mus:impls` hint, skip this type.
2. For each implementation type listed in the `// mus:impls` hint:
   - Generate a common serializer (e.g., `FooMUS`).
   - Generate a DTM constant (e.g., `FooDTM`) if it doesn't already exist.
     Missing DTMs are generated in the `mus.ai.gen.go` file as a block, starting 
     from 1:
     ```go
     const (
       FooDTM com.DTM = iota + 1
       BarDTM
     )
     ```
   - Use the DTM constant to create a [Typed Serializer](./typed_ser.md) (e.g., 
    `FooTypedMUS`).
3. Generate the interface serializer that uses these typed serializers.


> [!IMPORTANT]
> **DTM Management**: Managing DTM conflicts is the **user's responsibility**. 
> There are two supported modes:
> 1. **User-defined DTMs**: The user provides all DTM constants manually.
> 2. **Generated DTMs**: The generator handles all DTM constants starting from 1.

> [!NOTE]
> Either **all** implementation DTMs must be provided by the user, or **all** must 
> be missing. If some are missing and some are present, output an error: "Partial 
> DTMs for type_name.", where `type_name` is the interface name.

Example:

```go
// User-provided source code:

const (
  FooDTM com.DTM = 1
  BarDTM com.DTM = 2
)

type Foo struct {...}
type Bar struct {...}

// mus:impls = Foo, Bar
type MyInterface interface {
  Method()
}

// -----------------------------------------------------------------------------
// Generated code:

var MyInterfaceMUS = myInterfaceMUS{}

type myInterfaceMUS struct{}

func (s myInterfaceMUS) Marshal(v MyInterface, ... ) (n int) {
  switch val := v.(type) {
  case Foo:
    return FooTypedMUS.Marshal(val, ... )
  case Bar:
    return BarTypedMUS.Marshal(val, ... )
  default:
		panic(fmt.Sprintf(com.ErrorPrefix+"unexpected %v type", t))
  }
}

func (s myInterfaceMUS) Unmarshal( ... ) (v MyInterface, n int, err error) {
  dtm, n, err := typed.DTMSer.Unmarshal( ... )
  if err != nil {
    return
  }
  switch dtm {
  case FooDTM:
    var v1 Foo
    v1, n1, err1 := FooTypedMUS.UnmarshalData( ... )
    return v1, n + n1, err1
  case BarDTM:
    var v1 Bar
    v1, n1, err1 := BarTypedMUS.UnmarshalData( ... )
    return v1, n + n1, err1
  default:
    err = common.ErrUnknownDTM
    return
  }
}

func (s myInterfaceMUS) Size(v MyInterface) (size int) {
	switch t := v.(type) {
	case Foo:
		return FooTypedMUS.Size(t)
	case Bar:
		return BarTypedMUS.Size(t)
	default:
		panic(fmt.Sprintf(com.ErrorPrefix+"unexpected %v type", t))
	}
}

func (s myInterfaceMUS) Skip( ... ) (n int, err error) {
	dtm, n, err := typed.DTMSer.Unmarshal( ... )
	if err != nil {
		return
	}
	var n1 int
	switch dtm {
	case FooDTM:
		n1, err = FooTypedMUS.SkipData( ... )
	case BarDTM:
		n1, err = BarTypedMUS.SkipData( ... )
	default:
    err = common.ErrUnknownDTM
		return
	}
	n += n1
	return
} 
```
