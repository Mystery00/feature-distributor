package enum

import "feature-distributor/common/attribute"

type AttributeType int8

const (
	StringAttributeType AttributeType = 1 + iota
	NumberAttributeType
	TimeAttributeType
)

var (
	attributeValues = []AttributeType{
		StringAttributeType,
		NumberAttributeType,
		TimeAttributeType,
	}
	attributeTypeMap = map[AttributeType]attribute.Type{
		StringAttributeType: attribute.StringAttributeType,
		NumberAttributeType: attribute.NumberAttributeType,
		TimeAttributeType:   attribute.TimeAttributeType,
	}
)

func ParseAttributeType(s string) AttributeType {
	vType := attribute.ParseType(s)
	if vType == nil {
		panic("unknown attribute type")
	}
	for k, v := range attributeTypeMap {
		if v == *vType {
			return k
		}
	}
	panic("unknown attribute type")
}

func (v AttributeType) String() string {
	return attributeTypeMap[v].String()
}

func AttributeTypeEnum(v int8) AttributeType {
	for _, t := range attributeValues {
		if int8(t) == v {
			return t
		}
	}
	panic("unknown attribute type")
}
