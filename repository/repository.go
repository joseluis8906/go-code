package repository

import (
	"strings"
)

const Limit = 10

func Page(page int64) int64 {
	return int64(Limit * (page - 1))
}

// Criteria represents a criteria.
type Criteria struct {
	Operator string
	Field    string
	Value    any
}

func Query(format string, args ...any) Criteria {
	q := format
	c := Criteria{
		Value: args[0],
	}

	if strings.Contains(q, "=") {
		c.Operator = "$eq"
		f := field(q, "=")
		c.Field = f
	} else if strings.Contains(q, "<>") {
		c.Operator = "$ne"
		f := field(q, "<>")
		c.Field = f
	} else if strings.Contains(q, ">") {
		c.Operator = "$gt"
		f := field(q, ">")
		c.Field = f
	} else if strings.Contains(q, "<") {
		c.Operator = "$lt"
		f := field(q, "<")
		c.Field = f
	} else if strings.Contains(q, ">=") {
		c.Operator = "$gte"
		f := field(q, ">=")
		c.Field = f
	} else if strings.Contains(q, "<=") {
		c.Operator = "$lte"
		f := field(q, "<=")
		c.Field = f
	} else if strings.Contains(q, " LIKE ") {
		c.Operator = "$regex"
		f := field(q, " LIKE ")
		c.Field = f
	}

	return c
}

func field(q, op string) string {
	d := strings.Split(q, op)
	return strings.Trim(d[0], " ")
}
