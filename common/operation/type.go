package operation

import (
	"feature-distributor/common/attribute"
	"strings"
)

type Type string

const (
	StringInType         Type = "STRING_IN"
	StringOutType        Type = "STRING_OUT"
	StringStartType      Type = "STRING_START"
	StringNotStartType   Type = "STRING_NOT_START"
	StringEndType        Type = "STRING_END"
	StringNotEndType     Type = "STRING_NOT_END"
	StringContainType    Type = "STRING_CONTAIN"
	StringNotContainType Type = "STRING_NOT_CONTAIN"
	StringRegexType      Type = "STRING_REGEX"
	StringNotRegexType   Type = "STRING_NOT_REGEX"

	NumberEqType      Type = "NUMBER_EQ"
	NumberNotEqType   Type = "NUMBER_NOT_EQ"
	NumberGtType      Type = "NUMBER_GT"
	NumberGteType     Type = "NUMBER_GTE"
	NumberLtType      Type = "NUMBER_LT"
	NumberLteType     Type = "NUMBER_LTE"
	NumberBetweenType Type = "NUMBER_BETWEEN"

	TimeBeforeType Type = "TIME_BEFORE"
	TimeAfterType  Type = "TIME_AFTER"
)

var (
	values = []Type{
		StringInType,
		StringOutType,
		StringStartType,
		StringNotStartType,
		StringEndType,
		StringNotEndType,
		StringContainType,
		StringNotContainType,
		StringRegexType,
		StringNotRegexType,
		NumberEqType,
		NumberNotEqType,
		NumberGtType,
		NumberGteType,
		NumberLtType,
		NumberLteType,
		NumberBetweenType,
		TimeBeforeType,
		TimeAfterType,
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

func (v Type) ForAttributeType(attrType attribute.Type) bool {
	attributeTypeStr := strings.Split(v.String(), "_")[0]
	return attributeTypeStr == attrType.String()
}
