## Anonymous Serializers

For a `string` (with a hint), pointer, array, slice or a map type generate an
anonymous serializer.

Anonymous serializer should be created with a constructor from 
[Predefined Constructors](./predefined_constructors) and declared as a variable.

Anonymous serializer variable name is build like:

```
// here "+" sign denotes a concatenation

// for a string serializer
str + Base64 (Hash (lenSer + lenVl))

// for an array serializer
array + Base64 (Hash (element type + lenSer + elemVl))

// for a byte slice serializer
byteSlice + Base64 (Hash (lenSer + lenVl))

// for a slice serializer
slice + Base64 (Hash (element type + lenSer + lenVl + elemVl))

// for a map serializer
map + Base64 (Hash (key type + value type + lenSer + lenVl + keyVl + valVl))

// for a pointer serializer
ptr + Base64 (Hash (element type + baseSer))
```

To create valid Go identifiers, the standard Base64 characters MUST be replaced 
as follows:
- `+` -> `Δ`
- `/` -> `Σ`
- `=` -> `Ξ`

Use hash (MD5) in a name to ensure unique names across the package. For
example:

```go
var (
	mapAUEFΣr5wWΣsQKmdS5ns98QΞΞ   = ord.NewMapSer(ord.String, varint.Int64)
	ptry4kBΔSqZ1hMnnFP3Y5bQ4wΞΞ   = ord.NewPtrSer(ord.Bool)
	sliceoVLU4RNiPduaSOBCraZgCwΞΞ = ord.NewSliceSer(varint.Uint16)
)
```

### Reference Implementation (Python)

To ensure consistency in hash generation, you can use the following Python script:

```python
import hashlib
import base64

def get_mus_anon_var_name(prefix, data):
    # 1. MD5 Hash of the data
    md5_hash = hashlib.md5(data.encode('utf-8')).digest()
    
    # 2. Base64 Encode
    b64_str = base64.b64encode(md5_hash).decode('utf-8')
    
    # 3. Project-specific substitutions
    # + -> Δ (U+0394 Delta)
    # / -> Σ (U+03A3 Sigma)
    # = -> Ξ (U+039E Xi)
    substitutions = {
        '+': 'Δ',
        '/': 'Σ',
        '=': 'Ξ',
    }
    
    for char, replacement in substitutions.items():
        b64_str = b64_str.replace(char, replacement)
        
    return f"{prefix}{b64_str}"

# Example usage:
# print(get_mus_anon_var_name("map", "stringint64"))
# print(get_mus_anon_var_name("str", "stringValidateStringLength"))
```
