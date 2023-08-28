package main

import (
	"sync"
	"testing"

	"github.com/plutov/benchmark-grpc-protobuf-vs-http-json/grpc-protobuf/proto"
	"golang.org/x/net/context"
	g "google.golang.org/grpc"
)

func init() {
	// go grpcprotobuf.Start()
	// time.Sleep(time.Second)
}

func BenchmarkGRPCProtobuf(b *testing.B) {
	goRouting := 8
	aCount := b.N / goRouting

	var n sync.WaitGroup
	for i := 1; i <= goRouting; i++ {
		n.Add(1)
		go func(amount int) {

			conn, err := g.Dial("192.168.1.81:30800", g.WithInsecure())
			// conn, err := g.Dial("192.168.1.81:30600", g.WithInsecure())
			if err != nil {
				b.Fatalf("grpc connection failed: %v", err)
			}

			client := proto.NewAPIClient(conn)

			for n := 0; n < aCount; n++ {
				doGRPC(client, b)
			}
			n.Done()
		}(i)
	}
	n.Wait()

}

func doGRPC(client proto.APIClient, b *testing.B) {
	resp, err := client.CreateUser(context.Background(), &proto.User{
		Email:    "foo@bar.com",
		Name:     "Bench",
		Password: "bench",
	})

	if err != nil {
		b.Fatalf("grpc request failed: %v", err)
	}

	if resp == nil || resp.Code != 200 || resp.User == nil || resp.User.Id != "1000000" {
		b.Fatalf("grpc response is wrong: %v", resp)
	}
}
