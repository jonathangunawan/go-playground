# gRPC Example

This project is to show how to make Go apps that use gRPC as communication protocol with another apps.

## Step by step make gRPC Server

1. Make proto folder, so we will place every proto file in this folder
2. Make your proto file
3. Generate Go file based on proto file using proto command (see genGRPC.sh)
4. Make the function based on .proto that you make, in this case I put that function in usecase folder (maybe I should named it handler folder instead of usecase?)
5. I recomend to make client package (checkout README inside client folder, for the detail

## How to run this project

1. Run server.go inside cmd folder to run server app
2. Run user.go inside user folder to run user app
3. You can test through endpoint. for example 127.0.0.1:8080/add/1/2



## Tips for Unit Test in the user app

It's easy to make unit test for user app, you just need to make interface based on function that you only need instead of use generated interface in proto folder. So even if the gRPC server have 100.000 function under Math interface, you don't need to make a whole 100.000 mock. So NewMath() will implementor of New interface in the user app.
