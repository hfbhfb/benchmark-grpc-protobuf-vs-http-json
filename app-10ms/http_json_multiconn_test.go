package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"testing"
	"time"

	httpjson "github.com/plutov/benchmark-grpc-protobuf-vs-http-json/http-json"
)

func init() {

}

func BenchmarkHTTPJSONMultiConn(b *testing.B) {
	b.Log(b.N)
	for i := 0; i < 100; i++ {
		conn, err := net.Dial("tcp", "192.168.1.81:30901")
		if err != nil {
			fmt.Println("Error connecting:", err)
			return
		}
		defer conn.Close()

		client := &http.Client{
			Transport: &http.Transport{
				Dial: func(network, addr string) (net.Conn, error) {
					return conn, nil
				},
			},
			Timeout: 10 * time.Second,
		}

		for n := 0; n < b.N; n++ {
			doPostMultiConn(client, b)
		}
	}

}

func doPostMultiConn(client *http.Client, b *testing.B) {
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
