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
	"strconv"
	"strings"

	"google.golang.org/grpc"
	pb "roykv-for-tikv/tikv"
)

const (
	port = ":50055" //fetch grpc port from options
)

var rawKvClient *rawkv.Client

func getRawKvClient() *rawkv.Client {
	//todo fetch pd address from options
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
	var deleted uint64

	for _, key := range in.Keys {
		errDelete := rawKvClient.Delete(context.TODO(), []byte(key))
		if errDelete != nil {
			log.Println(errDelete)
		} else {
			deleted++
		}
	}

	return &pb.DelReply{Deleted: deleted}, nil
}

// Set implements roykvtikv.tikvServer
func (s *tikvServer) Set(ctx context.Context, in *pb.SetRequest) (*pb.SetReply, error) {
	var result bool

	errPut := rawKvClient.Put(context.TODO(), []byte(in.GetKey()), []byte(in.GetValue()))
	if errPut != nil {
		log.Println(errPut)
		result = false
	} else {
		result = true
	}
	return &pb.SetReply{Result: result}, nil
}

// Get implements roykvtikv.tikvServer
func (s *tikvServer) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetReply, error) {
	var value string
	value = ""

	byteVal, errGet := rawKvClient.Get(context.TODO(), []byte(in.GetKey()))
	if errGet != nil {
		log.Println(errGet)
	} else {
		if byteVal != nil {
			value = string(byteVal)
		}
	}

	return &pb.GetReply{Value: value}, nil
}

// Exist implements roykvtikv.tikvServer
func (s *tikvServer) Exist(ctx context.Context, in *pb.ExistRequest) (*pb.ExistReply, error) {
	var existed bool

	byteVal, errGet := rawKvClient.Get(context.TODO(), []byte(in.GetKey()))

	if errGet != nil {
		log.Println(errGet)
		existed = false
	} else {
		existed = byteVal != nil
	}

	return &pb.ExistReply{Existed: existed}, nil
}

// Scan implements roykvtikv.tikvServer
func (s *tikvServer) Scan(ctx context.Context, in *pb.ScanRequest) (*pb.ScanReply, error) {
	var data []*pb.KVEntry
	data = []*pb.KVEntry{}

	var startKey string
	startKey = in.GetStartKey()
	var startKeyType string
	startKeyType = in.GetStartKeyType()
	var endKey string
	endKey = in.GetEndKey()
	var endKeyType string
	endKeyType = in.GetEndKeyType()
	var keyPrefix string
	keyPrefix = in.GetKeyPrefix()

	var limit uint64
	limit = in.GetLimit()

	var count uint64
	count = 0

	var lastKey string
	lastKey = ""
	var skipFirst bool
	skipFirst = false

	for ;count < limit; {
		var listKey [][]byte
		var listVal [][]byte
		var errScan error

		if !skipFirst {
			listKey, listVal, errScan = rawKvClient.Scan(context.TODO(), []byte(startKey), []byte(endKey), int(limit))
			if errScan != nil {
				log.Println(errScan)
			}
		} else {
			listKey, listVal, errScan = rawKvClient.Scan(context.TODO(), []byte(lastKey), []byte(endKey), int(limit))
			if errScan != nil {
				log.Println(errScan)
			}
		}

		if skipFirst {
			if len(listKey) <= 0 {
				break
			}
		} else {
			if len(listKey) <= 1 {
				break
			}
		}

		for i, byteKey := range listKey {
			if skipFirst {
				if i == 0 {
					continue
				}
			}

			var key string
			key = string(byteKey)

			lastKey = key
			skipFirst = true

			if strings.HasPrefix(key, keyPrefix) {
				var matched bool
				matched = true

				var realKey string
				realKey = key[len(keyPrefix) :]

				if len(startKey) > 0 {
					var realStartKey string
					realStartKey = startKey[len(keyPrefix) :]

					if startKeyType == "integer" {
						var errAtoi error
						var realKeyInt int
						var realStartKeyInt int
						realKeyInt, errAtoi = strconv.Atoi(realKey)
						if errAtoi != nil {
							matched = false
						}

						realStartKeyInt, errAtoi = strconv.Atoi(realStartKey)
						if errAtoi != nil {
							matched = false
						}

						if matched {
							if realKeyInt < realStartKeyInt {
								matched = false
							}
						}
					} else if startKeyType == "double" {
						var errParseDouble error
						var realKeyDouble float64
						var realStartKeyDouble float64
						realKeyDouble, errParseDouble = strconv.ParseFloat(realKey, 64)
						if errParseDouble != nil {
							matched = false
						}

						realStartKeyDouble, errParseDouble = strconv.ParseFloat(realStartKey, 64)
						if errParseDouble != nil {
							matched = false
						}

						if matched {
							if realKeyDouble < realStartKeyDouble {
								matched = false
							}
						}
					}
				}
				if len(endKey) > 0 {
					var realEndKey string
					realEndKey = endKey[len(keyPrefix) :]

					if endKeyType == "integer" {
						var errAtoi error
						var realKeyInt int
						var realEndKeyInt int
						realKeyInt, errAtoi = strconv.Atoi(realKey)
						if errAtoi != nil {
							matched = false
						}

						realEndKeyInt, errAtoi = strconv.Atoi(realEndKey)
						if errAtoi != nil {
							matched = false
						}

						if matched {
							if realKeyInt < realEndKeyInt {
								matched = false
							}
						}
					} else if endKeyType == "double" {
						var errParseDouble error
						var realKeyDouble float64
						var realEndKeyDouble float64
						realKeyDouble, errParseDouble = strconv.ParseFloat(realKey, 64)
						if errParseDouble != nil {
							matched = false
						}

						realEndKeyDouble, errParseDouble = strconv.ParseFloat(realEndKey, 64)
						if errParseDouble != nil {
							matched = false
						}

						if matched {
							if realKeyDouble < realEndKeyDouble {
								matched = false
							}
						}
					}
				}

				if matched {
					data = append(data, &pb.KVEntry{Key:key, Value:string(listVal[i])})
					count++
					if count >= limit {
						break
					}
				}
			}
		}
	}

	if (count < limit) && (len(endKey) > 0) {
		lastValue, errGet := rawKvClient.Get(context.TODO(), []byte(endKey))
		if errGet != nil {
			log.Println(errGet)
		} else {
			if lastValue != nil {
				data = append(data, &pb.KVEntry{Key:endKey, Value:string(lastValue)})
			}
		}
	}

	return &pb.ScanReply{Data: data}, nil
}

// MGet implements roykvtikv.tikvServer
func (s *tikvServer) MGet(ctx context.Context, in *pb.MGetRequest) (*pb.MGetReply, error) {
	var data map[string]string
	data = make(map[string]string)

	var byteKeys [][]byte
	byteKeys = [][]byte{}
	for _, key := range in.Keys {
		byteKeys = append(byteKeys, []byte(key))
	}

	byteVals, errBGet := rawKvClient.BatchGet(context.TODO(), byteKeys)
	if errBGet != nil {
		log.Println(errBGet)
	} else {
		for i, byteVal := range byteVals {
			if byteVal != nil {
				data[string(byteKeys[i])] = string(byteVal)
			}
		}
	}

	return &pb.MGetReply{Data: data}, nil
}

// GetAll implements roykvtikv.tikvServer
func (s *tikvServer) GetAll(ctx context.Context, in *pb.GetAllRequest) (*pb.GetAllReply, error) {
	var data map[string]string
	data = make(map[string]string)

	var keyPrefix string
	keyPrefix = ""

	var lastKey string
	lastKey = ""

	var skipFirst bool
	skipFirst = false

	for ;true; {
		var listKey [][]byte
		var listVal [][]byte
		var errScan error

		if !skipFirst {
			listKey, listVal, errScan = rawKvClient.Scan(context.TODO(), []byte(""), []byte(""), 10000)
			if errScan != nil {
				log.Println(errScan)
			}
		} else {
			listKey, listVal, errScan = rawKvClient.Scan(context.TODO(), []byte(lastKey), []byte(""), 10000)
			if errScan != nil {
				log.Println(errScan)
			}
		}

		if skipFirst {
			if len(listKey) <= 0 {
				break
			}
		} else {
			if len(listKey) <= 1 {
				break
			}
		}

		for i, byteKey := range listKey {
			if skipFirst {
				if i == 0 {
					continue
				}
			}

			var key string
			key = string(byteKey)

			lastKey = key
			skipFirst = true

			if strings.HasPrefix(key, keyPrefix) {
				data[key] = string(listVal[i])
			}
		}
	}

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
