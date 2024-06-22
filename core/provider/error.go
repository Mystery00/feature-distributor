package provider

import (
	"feature-distributor/common/alert"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`^([A-Z_]+): (\d+)$`)

func DealError(err error) error {
	s := err.Error()
	if strings.HasPrefix(s, "error code: ") {
		ss := strings.TrimPrefix(s, "error code: ")
		group := re.FindStringSubmatch(ss)
		i, _ := strconv.Atoi(group[2])
		code := alert.Code(i)
		return alert.Error(code)
	}
	return err
}
