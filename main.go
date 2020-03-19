package main

import (
	"context"
	"fmt"

	"github.com/tikv/client-go/config"
	"github.com/tikv/client-go/rawkv"
)

func main()  {

	cli, err := rawkv.NewClient(context.TODO(), []string{"47.110.155.53:32814", "47.110.155.53:32816", "47.110.155.53:32818"}, config.Default())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	errPut := cli.Put(context.TODO(), []byte("foo"), []byte("bar"))
	if errPut != nil {
		panic(errPut)
	}

	val, errGet := cli.Get(context.TODO(), []byte("foo"))
	if errGet != nil {
		panic(errPut)
	}

	fmt.Println(string(val))

	errDel := cli.Delete(context.TODO(), []byte("foo"))
	if errDel != nil {
		panic(errDel)
	}
}
