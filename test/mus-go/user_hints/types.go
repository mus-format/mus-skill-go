package userhints

import "time"

// Should apply the following hints:
// - name
// - skip
// - path
// - vl
// - time_unit
// - lenSer
// - lenVl
// - keyVl
// - valVl
// - elemVl
// - elem:lenSer
// - elem:lenVl
// - elem:elemVl
// - elem:elem:enc

// mus:name = CustomName
type Name int

type Skip struct {
	num int
	// mus:skip = true
	secret string
}

// mus:path = github.com/mus-format/mus-go/.ai/cookbook/test/user_hints/sub
type CrossPackage int

// mus:vl = ValidateStruct
type Valid struct {
	num int
}

type ValidField struct {
	// mus:vl = ValidateNum
	num int
	// mus:vl = ValidateStr
	str string
}

// mus:time_unit = nano_utc
type Time time.Time

// mus:lenSer = raw.Int
// mus:lenVl = ValidateLength
// mus:keyVl = ValidateKey
// mus:valVl = ValidateValue
type DefMap map[string]int

// mus:lenSer = raw.Int
// mus:lenVl = ValidateLength
// mus:elemVl = ValidateElement
type DefSlice []int

// mus:lenSer = raw.Int
// mus:elemVl = ValidateElement
type DefArray [3]int

// mus:lenSer = raw.Int
// mus:lenVl = ValidateLength
// mus:keyVl = ValidateKey
//
// mus:key:lenSer = varint.Int
// mus:key:lenVl = ValidateLength
//
// mus:val:lenSer = raw.Int
// mus:val:lenVl = ValidateLength
//
// mus:val:elem:lenSer = varint.Int
// mus:val:elem:elemVl = ValidateElement
// mus:val:elem:elem:enc = raw
type Nested map[string][][3]int
