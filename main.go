/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

//go:generate protoc -I ./protos --go_out=plugins=grpc:./tikv ./protos/roykv-for-tikv.proto

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"github.com/tikv/client-go/config"
	"github.com/tikv/client-go/rawkv"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "roykv-for-tikv/tikv"
)

const (
	port = ":50055"
)

var rawKvClient *rawkv.Client

func getRawKvClient() *rawkv.Client {
	cli, err := rawkv.NewClient(context.TODO(), []string{"47.110.155.53:32814", "47.110.155.53:32816", "47.110.155.53:32818"}, config.Default())
	if err != nil {
		panic(err)
	}
	return cli
}

// server is used to implement roykvtikv.tikvServer
type tikvServer struct{}

// Del implements roykvtikv.tikvServer
func (s *tikvServer) Del(ctx context.Context, in *pb.DelRequest) (*pb.DelReply, error) {
	//todo
	return &pb.DelReply{Deleted: 1}, nil
}

// Set implements roykvtikv.tikvServer
func (s *tikvServer) Set(ctx context.Context, in *pb.SetRequest) (*pb.SetReply, error) {
	//todo
	return &pb.SetReply{Result: true}, nil
}

// Get implements roykvtikv.tikvServer
func (s *tikvServer) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetReply, error) {
	//todo
	return &pb.GetReply{Value: "bar"}, nil
}

// Exist implements roykvtikv.tikvServer
func (s *tikvServer) Exist(ctx context.Context, in *pb.ExistRequest) (*pb.ExistReply, error) {
	//todo
	return &pb.ExistReply{Existed: true}, nil
}

// Scan implements roykvtikv.tikvServer
func (s *tikvServer) Scan(ctx context.Context, in *pb.ScanRequest) (*pb.ScanReply, error) {
	//todo
	return &pb.ScanReply{Data: []*pb.KVEntry{&pb.KVEntry{Key:"foo", Value:"bar"}}}, nil
}

// MGet implements roykvtikv.tikvServer
func (s *tikvServer) MGet(ctx context.Context, in *pb.MGetRequest) (*pb.MGetReply, error) {
	//todo
	var data map[string]string
	data = make(map[string]string)
	data["foo"] = "bar"
	return &pb.MGetReply{Data: data}, nil
}

// GetAll implements roykvtikv.tikvServer
func (s *tikvServer) GetAll(ctx context.Context, in *pb.GetAllRequest) (*pb.GetAllReply, error) {
	//todo
	var data map[string]string
	data = make(map[string]string)
	data["foo"] = "bar"
	return &pb.GetAllReply{Data: data}, nil
}

// Count implements roykvtikv.tikvServer
func (s *tikvServer) Count(ctx context.Context, in *pb.CountRequest) (*pb.CountReply, error) {
	//todo
	return &pb.CountReply{Count: 1}, nil
}

func main() {
	rawKvClient = getRawKvClient()
	defer rawKvClient.Close()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTiKVServer(s, &tikvServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
