package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/jonathangunawan/go-playground/grpc/proto"

	"github.com/jonathangunawan/go-playground/grpc/client"

	"github.com/julienschmidt/httprouter"
)

var (
	serverTarget = "127.0.0.1:4040"
)

func main() {
	pm, err := client.NewMath(serverTarget)
	if err != nil {
		panic(err)
	}

	g := httprouter.New()
	g.GET("/add/:a/:b", func(resp http.ResponseWriter, req *http.Request, param httprouter.Params) {
		a, err := strconv.ParseInt(param.ByName("a"), 10, 64)
		if err != nil {
			resp.WriteHeader(http.StatusBadRequest)
		}

		b, err := strconv.ParseInt(param.ByName("b"), 10, 64)
		if err != nil {
			resp.WriteHeader(http.StatusBadRequest)
		}

		// Using normal Math Client
		ctx := context.Background()
		protoReq := &proto.MathRequest{FirstParam: a, SecondParam: b}
		if response, err := pm.Client.Add(ctx, protoReq); err == nil {
			resp.WriteHeader(http.StatusOK)
			resp.Header().Set("Content-Type", "application/json")
			data, _ := json.Marshal(response)
			resp.Write(data)
		} else {
			resp.WriteHeader(http.StatusInternalServerError)
		}
	})

	g.GET("/mult/:a/:b", func(resp http.ResponseWriter, req *http.Request, param httprouter.Params) {
		a, err := strconv.ParseInt(param.ByName("a"), 10, 64)
		if err != nil {
			resp.WriteHeader(http.StatusBadRequest)
		}

		b, err := strconv.ParseInt(param.ByName("b"), 10, 64)
		if err != nil {
			resp.WriteHeader(http.StatusBadRequest)
		}

		// Using normal Math Client
		ctx := context.Background()
		protoReq := &proto.MathRequest{FirstParam: a, SecondParam: b}
		if response, err := pm.Client.Multiply(ctx, protoReq); err == nil {
			resp.WriteHeader(http.StatusOK)
			resp.Header().Set("Content-Type", "application/json")
			data, _ := json.Marshal(response)
			resp.Write(data)
		} else {
			resp.WriteHeader(http.StatusInternalServerError)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", g))
}
