package variables

import (
	"strings"
)

func ParsePartialDepth(stream string) string {
	sli := strings.Split(stream, "@")
	switch len(sli) {
	case 2, 3:
		return sli[0]
	default:
		return ""
	}
}

func Ptr[T any](v T) *T { return &v }
