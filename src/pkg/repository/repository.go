package repository

import "errors"

const Limit = 10

func Page(page int64) int64 {
	return int64(Limit * (page - 1))
}

type Result[T any] struct {
	data []T
	err  error
}

// One returns the first item in the data collection.
// if there are more than one item, it will return an error.
// if an error occurred during the query, it will return the error.
func (r Result[T]) One() (T, error) {
	var defaultData T
	if r.err != nil {
		return defaultData, r.err
	}

	if len(r.data) == 0 {
		return defaultData, nil
	}

	if len(r.data) > 1 {
		return defaultData, errors.New("too many results")
	}

	return r.data[0], nil
}

// Many returns all items in the data collection.
// if an error occurred during the query, it will return the error.
func (r Result[T]) Many() ([]T, error) {
	if r.err != nil {
		return nil, r.err
	}

	return r.data, nil
}

func Error[T any](err error) Result[T] {
	return Result[T]{err: err}
}

func Data[T any](data []T) Result[T] {
	return Result[T]{data: data}
}
