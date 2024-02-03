package repository

import "errors"

const Limit = 10

var ErrNoData = errors.New("no data found")

func Page(page int64) int64 {
	return int64(Limit * (page - 1))
}

type Result[T any] struct {
	data []T
	err  error
}

func (r Result[T]) ExpectOne() (T, error) {
	var defaultData T
	if r.err != nil {
		return defaultData, r.err
	}

	if len(r.data) == 0 {
		return defaultData, ErrNoData
	}

	if len(r.data) > 1 {
		return defaultData, errors.New("too many results")
	}

	return r.data[0], nil
}

func (r Result[T]) ExpectMany() ([]T, error) {
	if r.err != nil {
		return nil, r.err
	}

	if len(r.data) == 0 {
		return nil, ErrNoData
	}

	return r.data, nil
}

func Error[T any](err error) Result[T] {
	return Result[T]{err: err}
}

func Data[T any](data []T) Result[T] {
	return Result[T]{data: data}
}
