package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sync"
	"testing"

	httpjson "github.com/plutov/benchmark-grpc-protobuf-vs-http-json/http-json"
)

func init() {
	// go httpjson.Start()
	// time.Sleep(time.Second)
}

func BenchmarkHTTPJSON(b *testing.B) {
	goRouting := 100
	//aCount := b.N / goRouting
	aCount := b.N

	var n sync.WaitGroup
	for i := 1; i <= goRouting; i++ {
		n.Add(1)
		go func(amount int) {

	client := &http.Client{}

			for n := 0; n < aCount; n++ {
				doPost(client, b)
			}
			n.Done()
		}(i)
	}
	n.Wait()

}

func doPost(client *http.Client, b *testing.B) {
	u := &httpjson.User{
		Email:    "foo@bar.com",
		Name:     "Bench",
		Password: "bench",
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(u)

	resp, err := client.Post("http://192.168.1.81:30901/", "application/json", buf)
	if err != nil {
		b.Fatalf("http request failed: %v", err)
	}

	defer resp.Body.Close()

	// We need to parse response to have a fair comparison as gRPC does it
	var target httpjson.Response
	decodeErr := json.NewDecoder(resp.Body).Decode(&target)
	if decodeErr != nil {
		b.Fatalf("unable to decode json: %v", decodeErr)
	}

	if target.Code != 200 || target.User == nil || target.User.ID != "1000000" {
		b.Fatalf("http response is wrong: %v", resp)
	}
}
