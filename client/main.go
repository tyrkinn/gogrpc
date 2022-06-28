package main

import (
	"context"
	"flag"
	"gogrpc/proto"
	"gogrpc/utils"
	"log"
	"strconv"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	serverAddr = "localhost:50051"
)

func checkAndGetArgs() (int32, int32) {
	if flag.NArg() < 2 {
		log.Fatal("You should provide two numbers")
	}
	firstTwoArgs := flag.Args()[:2]
	var intArgsSlice []int32
	for _, arg := range firstTwoArgs {
		v, err := strconv.Atoi(strings.Trim(arg, " "))
		utils.CheckErr(err, &utils.CheckErrArgs{Message: "Args must be int: " + arg})
		intArgsSlice = append(intArgsSlice, int32(v))
	}
	return intArgsSlice[0], intArgsSlice[1]
}

func main() {
	flag.Parse()
	x, y := checkAndGetArgs()
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	utils.CheckErr(err, nil)
	adderClient := proto.NewAdderClient(conn)
	res, err := adderClient.Add(context.Background(), &proto.AddRequest{X: x, Y: y})
	utils.CheckErr(err, nil)
	log.Println(res.GetResult())
}
