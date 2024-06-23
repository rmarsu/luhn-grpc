package adder

import (
	"context"
	"fmt"
	"luhn/algo"
	api "luhn/api/proto"
	"strconv"
)

type GrpcServer struct {
	api.UnimplementedValidServer
}

func (s *GrpcServer) Valid(ctx context.Context, req *api.ValidRequest) (*api.ValidResponse, error) {
	text := req.Number
	number, _ := strconv.Atoi(text)
	result := algo.Valid(number)
	fmt.Println(number)
	return &api.ValidResponse{
		Valid: result ,
	}, nil

}
