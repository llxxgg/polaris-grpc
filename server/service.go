package main

import (
	"context"
	"errors"
	"fmt"
	"polaris-grpc/pb"
	"strings"
)

const (
	StrMaxSize = 1024
)

// Service errors
var (
	ErrMaxSize = errors.New("maximum size of 1024 bytes exceeded")

	ErrStrValue = errors.New("maximum size of 1024 bytes exceeded")
)

type StringService struct {
	Port int
}

func (s *StringService) Concat(ctx context.Context, req *pb.StringRequest) (*pb.StringResponse, error) {
	if len(req.A)+len(req.B) > StrMaxSize {
		response := pb.StringResponse{Ret: ""}
		return &response, nil
	}
	response := pb.StringResponse{Ret: fmt.Sprintf("%s, from port:%d", req.A+req.B, s.Port)}
	return &response, nil
}

func (s *StringService) Diff(ctx context.Context, req *pb.StringRequest) (*pb.StringResponse, error) {
	if len(req.A) < 1 || len(req.B) < 1 {
		response := pb.StringResponse{Ret: ""}
		return &response, nil
	}
	res := ""
	if len(req.A) >= len(req.B) {
		for _, char := range req.B {
			if strings.Contains(req.A, string(char)) {
				res = res + string(char)
			}
		}
	} else {
		for _, char := range req.A {
			if strings.Contains(req.B, string(char)) {
				res = res + string(char)
			}
		}
	}
	response := pb.StringResponse{Ret: res}
	return &response, nil
}
