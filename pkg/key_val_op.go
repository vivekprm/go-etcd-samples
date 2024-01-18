package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		switch err {
		case context.Canceled:
			log.Fatalf("ctx is canceled by another routine: %v", err)
		case context.DeadlineExceeded:
			log.Fatalf("ctx is attached with a deadline is exceeded: %v", err)
		case rpctypes.ErrEmptyKey:
			log.Fatalf("client-side error: %v", err)
		default:
			log.Fatalf("bad cluster endpoints, which are not etcd servers: %v", err)
		}
	}

	// Creating Key
	pctx, pcancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	presp, err := cli.Put(pctx, "sample_key", "sample_value")
	pcancel()
	if err != nil {
		fmt.Printf("Error in putting key %v\n", err)
	}

	fmt.Printf("Response: %v\n", presp)

	// Getting the key
	gctx, gcancel := context.WithTimeout(context.Background(), 40*time.Millisecond)
	gresp, err := cli.Get(gctx, "sample_key", clientv3.WithRev(presp.Header.Revision))
	gcancel()

	if err != nil {
		log.Fatal(err)
	}
	for _, ev := range gresp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}

	// delete the keys
	// dctx, dcancel := context.WithTimeout(context.Background(), 40*time.Millisecond)
	// dresp, err := cli.Delete(dctx, "sample_key", clientv3.WithPrefix())
	// dcancel()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Deleted all keys:", int64(len(gresp.Kvs)) == dresp.Deleted)

	defer cli.Close()
}
