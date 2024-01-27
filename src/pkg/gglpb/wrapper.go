package gglpb

import (
	pb "google.golang.org/protobuf/types/known/wrapperspb"
)

func String(val string) *pb.StringValue {
	if val == "" {
		return nil
	}

	return &pb.StringValue{Value: val}
}

func Bool(val bool) *pb.BoolValue {
	return &pb.BoolValue{Value: val}
}

func Int32(val int32) *pb.Int32Value {
	return &pb.Int32Value{Value: val}
}

func Int64(val int64) *pb.Int64Value {
	return &pb.Int64Value{Value: val}
}

func Float(val float32) *pb.FloatValue {
	return &pb.FloatValue{Value: val}
}

func Double(val float64) *pb.DoubleValue {
	return &pb.DoubleValue{Value: val}
}
