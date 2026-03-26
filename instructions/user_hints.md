## User Hints

With hints user can specify the serialization options for the type.
 
### User-provided Type

These hints could be applied to the user-provided type.

#### Serializer Name
 
With the `mus:name` hint, the user can specify a custom name for the 
generated MUS serializer.
 
```go
// mus:name = CustomSerializer
type Foo struct {
  ...
}
```
 
Example:
 
```go
// mus:name = CustomFooSer
type Foo struct {
 	num int
}
 
var CustomFooSer = customFooSer{}
 
type customFooSer struct{}
 
func (s customFooSer) Marshal( ... ) (n int) { ... }
...
```

#### Validator

```go
// mus:vl = ...
T
```

Where:

- vl: com.Validator[T] interface implementation

Example:

```go
// mus:vl = ValidateFoo
type Foo int

func (s fooSer) Unmarshal( ... ) (v Foo, n int, err error) {
	var v1 int
	v1, n, err = raw.Int.Unmarshal(...)
	if err != nil {
		return
	}
	v = Foo(v1)
	if err = ValidateFoo(v); err != nil {
		return
	}
	return
}
...
```

#### Interface Implementations

With the `mus:impls` hint, the user can specify the types that implement an 
interface.

```go
// mus:impls = Foo, Bar, Baz
```

Example:

```go
// mus:impls = Foo, Bar
type MyInterface interface {
  Method()
}
```

The generator will use this hint to create a serializer for the interface type.

#### Generated Code Location

With the `mus:path` hint, the user can specify the location of the serializer 
for the type. This hint is used for both:
1. Specifying the location (package) where the code should be generated.
2. Specifying the location where the generator should look for a previously 
   generated serializer for a defined type.

```go
// mus:path = package_name
T
```

Example:

```go
// mus:path = github.com/user/project/pkg
type Foo struct {
	...
}
```

<!-- *********************************************************************** -->

### Struct Field Type or Underlying Type

These hints could be applied to the struct field type or underlying type.

#### Number Encoding

For all number types the user can specify the encoding.

```go
// mus:enc = raw
int
```

Where:

- enc: could be "varint" or "raw".

Example:

```go
// mus:enc = raw.Int
type Foo int

type fooSer struct{}

func (s fooSer) Marshal( ... ) (n int) {
	return raw.Int.Marshal( ... )
}

func (s fooSer) Unmarshal( ... ) (v Foo, n int, err error) {
	var v1 int
	v1, n, err = raw.Int.Unmarshal( ... )
	if err != nil {
		return
	}
	v = Foo(v1)
	return
}

func (s fooSer) Size(v Foo) (size int) {
	return raw.Int.Size(int(v))
}

func (s fooSer) Skip( ... ) (n int, err error) {
	return raw.Int.Skip( ... )
}
```

Nested encoding hints are also supported. For example, to specify the 
encoding for a slice element:

```go
// mus:elem:enc = raw
type Foo []int

// This slice serializer will use raw.Int for the int type.
var sliceQ65A = ord.NewSliceSer(raw.Int)
```

#### Time

```go
// mus:time_unit = ...
time.Time
```

Where:

- type_unit: sec -> TimeUnix, 
  milli -> TimeUnixMilli, 
  macro -> TimeUnixMicro, 
	nano -> TimeUnixNano, 
	sec_utc -> TimeUnixUTC, 
	milli_utc -> TimeUnixMilliUTC, 
	micro_utc -> TimeUnixMicroUTC, 
	nano_utc -> TimeUnixNanoUTC 

Example:

```go
// mus:time_unit = milli
type Foo time.Time

type fooSer struct{}

func (s fooSer) Marshal(v Foo, ... ) (n int) {
	return raw.TimeUnixMilli.Marshal(time.Time(v), ... )
}
...
```

#### Field Validator

```go
// mus:vl = ...
T
```

Where:

- vl: com.Validator[T] interface implementation

Example:

```go
type Foo struct {
  // mus:vl = ValidateNum
	num int
	// mus:vl = ValidateStr
	str string
}

func (s fooSer) Unmarshal( ... ) (v Foo, n int, err error) {
	var v1 int
	v.num, n, err = raw.Int.Unmarshal( ... )
	if err != nil {
		return
	}
	if err = ValidateNum(v.num); err != nil {
		return
	}
	var n1 int
	v.str, n1, err = ord.String.Unmarshal( ... )
	n += n1
	if err != nil {
		return
	}
	err = ValidateStr(v.str)
	return
}
```

#### String

String type hints require creating and using an anonymous serializer for that 
type. 

```go
// mus:lenSer = ...
// mus:lenVl = ...
string
```

Where:

- lenSer: mus.Serializer[int] interface implementation
- lenVl: com.Validator[int] interface implementation

Example:

```go
// mus:lenSer = raw.Int
type Foo string

// Marshal, Unmarshal, Size, Skip methods use this serializer.
var str`generated hash value` = ord.NewStringSer(strops.WithLenSer(raw.Int))
```

Example valid:

```go
// mus:lenSer = raw.Int
// mus:lenVl: ValidateLength
type Foo string

// Marshal, Unmarshal, Size, Skip methods use this serializer.
var str`generated hash value` = ord.NewValidStringSer(
	stropts.WithLenSer(raw.Int), 
  stropts.WithLenValidator(ValidateLength),
)
```

#### Array

```go
// mus:lenSer = ...
// mus:elemVl = ...
[N]T
```

Where:

- lenSer: mus.Serializer[int] interface implementation
- elemVl: com.Validator[T] interface implementation

Example:

```go
// mus:lenSer = raw.Int
// mus:elem:enc = raw
type ArrayInt [3]int

// Marshal, Unmarshal, Size, Skip methods use this serializer.
var array`generated hash value` = unsafe.NewArraySer(raw.Int, 
  arropts.WithLenSer[int](raw.Int), 
)
```

Example valid:

```go
// mus:lenSer = raw.Int
// mus:elemVl: ValidateElement
// mus:elem:enc = raw
type ArrayInt [3]int

// Marshal, Unmarshal, Size, Skip methods use this serializer.
var array`generated hash value` = unsafe.NewValidArraySer(raw.Int, 
  arropts.WithLenSer[int](raw.Int), 
  arropts.WithElemValidator[int](ValidateElement),
)
```

#### Slice

```go
// mus:lenSer = ...
// mus:lenVl = ...
// mus:elemVl = ...
[]T
```

Where:

- lenSer: mus.Serializer[int] interface implementation
- lenVl: com.Validator[int] interface implementation
- elemVl: com.Validator[T] interface implementation

Example:

```go
// mus:lenSer = raw.Int
// mus:elem:enc = raw
type SliceInt []int

// Marshal, Unmarshal, Size, Skip methods use this serializer.
var slice`generated hash value` = ord.NewSliceSer(raw.Int, 
  slopts.WithLenSer[int](raw.Int), 
)
```

Example valid:

```go
// mus:lenSer = raw.Int
// mus:lenVl: ValidateLength
// mus:elemVl: ValidateElement
// mus:elem:enc = raw
type SliceInt []int

// Marshal, Unmarshal, Size, Skip methods use this serializer.
var slice`generated hash value` = ord.NewValidSliceSer(raw.Int, 
  slopts.WithLenSer[int](raw.Int), 
	slopts.WithLenValidator[int](ValidateLength), 
  slopts.WithElemValidator[int](ValidateElement),
)
```

#### Byte Slice

```go
// mus:lenSer = ...
// mus:lenVl = ...
[]byte
```
Where:

- lenSer: mus.Serializer[int] interface implementation
- lenVl: com.Validator[int] interface implementation

Example:

```go
// mus:lenSer = raw.Int
type ByteSlice []byte 

// Marshal, Unmarshal, Size, Skip methods use this serializer.
var bs`generated hash value` = ord.NewByteSliceSer(bslops.WithLenSer(raw.Int))
```

Example valid:

```go
// mus:lenSer = raw.Int
// mus:lenVl: ValidateLength
type ByteSlice []byte 

// Marshal, Unmarshal, Size, Skip methods use this serializer.
var bs`generated hash value` = ord.NewValidByteSliceSer(
	bslopts.WithLenSer(raw.Int), 
  bslopts.WithLenValidator(ValidateLength),
)
```

#### Map

```go
// mus:lenSer = ...
// mus:lenVl = ...
// mus:keyVl = ...
// mus:valVl = ...
map[K]V
```

Where:

- lenSer: mus.Serializer[int] interface implementation
- lenVl: com.Validator[int] interface implementation
- keyVl: com.Validator[K] interface implementation
- valVl: com.Validator[V] interface implementation

Example:

```go
// mus:lenSer = raw.Int
// mus:value:enc = raw
type Map map[string]int 

// Marshal, Unmarshal, Size, Skip methods use this serializer.
var map`generated hash value` = ord.NewMapSer(
	mapopts.WithLenSer[string, int](raw.Int),
)
```

Example valid:

```go
// mus:lenSer = raw.Int
// mus:lenVl: ValidateLength
// mus:keyVl: ValidateKey
// mus:valVl: ValidateValue
type Map map[string]int 

// Marshal, Unmarshal, Size, Skip methods use this serializer.
var map`generated hash value` = ord.NewValidMapSer(
	mapopts.WithLenSer[string, int](raw.Int),
	mapopts.WithLenValidator[string, int](ValidateLength),
	mapopts.WithKeyValidator[string, int](ValidateKey),
	mapopts.WithValueValidator[string, int](ValidateValue),
)
```

#### Skip Field

With the `mus:skip` hint, the user can skip a struct field from the serialization 
process.

```go
// mus:skip = true
```

Example:

```go
type Foo struct {
  num int
  // mus:skip = true
  secret string
}

func (s fooSer) Marshal(v Foo, ... ) (n int) {
	return varint.Int.Marshal(v.num, ...)
}

func (s fooSer) Unmarshal( ... ) (v Foo, n int, err error) {
	v.num, n, err = varint.Int.Unmarshal( ... )
	return
}

func (s fooSer) Size(v Foo) (size int) {
	return varint.Int.Size(v.num)
}

func (s fooSer) Skip( ... ) (n int, err error) {
	return varint.Int.Skip( ... )
}
```

#### Nested Hints

Hints may be nested. Example for slice:

```go
// mus:lenVl = 
// mus:elem:lenVl = 
// mus:elem:elem:lenVl = 
[][]string
```

Where:
- first lenVl corresponds to the outer slice
- second lenVl corresponds to the inner slice
- third lenVl corresponds to the string


Example for map (hash values are fake):

```go
// mus:lenVl = ValidateMapLength 
// mus:key:lenVl = ValidateStringLength 
// mus:value:lenVl = ValidateSliceLength
// mus:value:elem:enc = raw
type Map map[string][]int

var strMUTorjOEDFYj3vrvkKR6ZwΞΞ = ord.NewValidStringSer(
	stropts.WithLenValidator(ValidateStringLength),
)

var sliceΔNqGN8oSY9KKXDASz9gfZwΞΞ = ord.NewValidSliceSer(raw.Int, 
	slopts.WithLenValidator[int](ValidateSliceLength),
)

// This map serializer uses the generated serializers for the key and value types.
var mapFΣQi8gWhMtB5OLKuuT2gOgΞΞ = ord.NewValidMapSer(
	strMUTorjOEDFYj3vrvkKR6ZwΞΞ, sliceΔNqGN8oSY9KKXDASz9gfZwΞΞ,
	mapopts.WithLenValidator[string, []int](ValidateMapLength),
)
```

Where:
- first lenVl corresponds to the map
- second lenVl corresponds to the key string
- third lenVl corresponds to the value slice
- fourth enc corresponds to the int

#### Validator Functions

If the validator hint refers to a function (e.g., `func ValidateXYZ(T) error`,
see [Validation](#validation)), the generator MUST automatically wrap it in
`com.ValidatorFn[T]` to be used with the `mus` options. If the hint refers to a
struct that already implements the `com.Validator` interface, it should be used 
as is.

Example:

```go
// mus:lenVl = ValidateLength
type Foo string

func ValidateLength(l int) error {
	return nil
}

// Marshal, Unmarshal, Size, Skip methods use this serializer.
var str`generated hash value` = ord.NewValidStringSer(
	stropts.WithLenValidator(com.ValidatorFn[int](ValidateLength)),
)
```
