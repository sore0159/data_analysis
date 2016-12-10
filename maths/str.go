package maths

import (
	"strings"
)

func (vs Vars) String() string {
	parts := make([]string, len(vs))
	for i, v := range vs {
		parts[i] = v.Name
	}
	return strings.Join(parts, ", ")
}

func (v Var) String() string {
	return v.Name
}
