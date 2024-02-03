package grpc

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type HeaderResult struct {
	res []string
	err error
}

func (r HeaderResult) ExpectOne() (string, error) {
	if r.err != nil {
		return "", r.err
	}

	if len(r.res) != 1 {
		return "", status.Errorf(codes.InvalidArgument, "more than one header")
	}

	return r.res[0], nil
}

func (r HeaderResult) ExpectMany() ([]string, error) {
	if r.err != nil {
		return nil, r.err
	}

	return r.res, nil
}

func Header(ctx context.Context, key string) HeaderResult {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return HeaderResult{nil, status.Errorf(codes.DataLoss, "failed to get metadata")}
	}

	val := md.Get(key)
	if len(val) == 0 {
		return HeaderResult{nil, status.Errorf(codes.Unauthenticated, fmt.Sprintf("missing %v header", key))}
	}

	return HeaderResult{val, nil}
}
