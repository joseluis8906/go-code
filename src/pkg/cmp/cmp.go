package cmp

// Criteria represents a criteria.
type Criteria struct {
	operator string
	field    string
	value    any
	page     int64
}

// Operator returns the operator of the criteria.
func (c Criteria) Operator() string {
	return c.operator
}

// Field returns the field of the criteria.
func (c Criteria) Field() string {
	return c.field
}

// Value returns the value of the criteria.
func (c Criteria) Value() any {
	return c.value
}

// Page returns the page of the criteria.
func (c Criteria) Page() int64 {
	return c.page
}

// WithPage returns a criteria with the given page.
func (c Criteria) WithPage(page int64) Criteria {
	c.page = page

	return c
}

// Equals returns a criteria for the given field and value.
func Equals(field string, value any) Criteria {
	return Criteria{
		operator: "$eq",
		field:    field,
		value:    value,
		page:     1,
	}
}

// GreaterThan returns a criteria for the given field and value.
func GreaterThan(field string, value any) Criteria {
	return Criteria{
		operator: "$gt",
		field:    field,
		value:    value,
		page:     1,
	}
}

// LessThan returns a criteria for the given field and value.
func LessThan(field string, value any) Criteria {
	return Criteria{
		operator: "$lt",
		field:    field,
		value:    value,
		page:     1,
	}
}

// Contains returns a criteria for the given field and value.
func Contains(field string, value any) Criteria {
	return Criteria{
		operator: "$regex",
		field:    field,
		value:    value,
		page:     1,
	}
}
