package attribute

type Type string

const (
	StringAttributeType Type = "STRING"
	NumberAttributeType Type = "NUMBER"
	TimeAttributeType   Type = "TIME"
)

var (
	values = []Type{
		StringAttributeType,
		NumberAttributeType,
		TimeAttributeType,
	}
)

func ParseType(s string) *Type {
	for _, t := range values {
		if t.String() == s {
			return &t
		}
	}
	return nil
}

func (v Type) String() string {
	return string(v)
}
