package value

type Type string

const (
	BoolValueType   Type = "BOOL"
	StringValueType Type = "STRING"
	FloatValueType  Type = "FLOAT"
	IntValueType    Type = "INT"
	JsonValueType   Type = "JSON"
)

var (
	values = []Type{
		BoolValueType,
		StringValueType,
		FloatValueType,
		IntValueType,
		JsonValueType,
	}
)

func ParseType(s string) *Type {
	for _, valueType := range values {
		if valueType.String() == s {
			return &valueType
		}
	}
	return nil
}

func (v Type) String() string {
	return string(v)
}
