package benchmarks

import (
	"sync"
	"testing"
	"time"

	grpcprotobuf "github.com/plutov/benchmark-grpc-protobuf-vs-http-json/grpc-protobuf"
	"github.com/plutov/benchmark-grpc-protobuf-vs-http-json/grpc-protobuf/proto"
	"golang.org/x/net/context"
	g "google.golang.org/grpc"
)

func init() {
	go grpcprotobuf.Start()
	time.Sleep(time.Second)
}

func BenchmarkGRPCProtobuf(b *testing.B) {
	b.Log(b.N)
	// b.Log()
	/*
		conn, err := g.Dial("127.0.0.1:60000", g.WithInsecure())
		if err != nil {
			b.Fatalf("grpc connection failed: %v", err)
		}

		client := proto.NewAPIClient(conn)

		for n := 0; n < b.N; n++ {
			doGRPC(client, b)
		}

		goRouting := 100
		aCount := b.N
		//aCount := b.N / goRouting
	*/
	goRouting := 100
	aCount := b.N
	conn, err := g.Dial("127.0.0.1:60000", g.WithInsecure())
	// conn, err := g.Dial("192.168.1.81:30600", g.WithInsecure())
	if err != nil {
		b.Fatalf("grpc connection failed: %v", err)
	}
	var n sync.WaitGroup
	for i := 1; i <= goRouting; i++ {
		n.Add(1)
		go func(amount int) {

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
