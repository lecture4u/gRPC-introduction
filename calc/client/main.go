package main

import (
	"context"
	"flag"
	pb "github.com/lecture4u/gRPC-introduction/calc/calc"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

var (
	operation = flag.String("o", "calcOperation", "")
)

func main() {
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCalculatorClient(conn)

	switch *operation {
	case "calcOperation":
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		printResult(calcOperation(ctx, c))
	case "calcOperations":
		printResult(calcOperations(c))
	case "getCurrentOperationResults":
		getCurrentOperationResults(c)
	case "calcJointOperation":
		calcJointOperation(c)
	case "calcJointOperations":
		calcJointOperations(c)
	}
}

func printResult(r *pb.CalcResponse) {
	if r.GetError() != false {
		log.Printf("Calculation Result: %v", r.GetResult())
	}
}

func calcJointOperations(c pb.CalculatorClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := c.CalcJointOperations(ctx)
	if err != nil {
		log.Fatalf("could not start streaming: %v", err)
	}

	var ops []*pb.CalcRequest
	for i := 0; i < 5; i++ {
		ops = append(ops, &pb.CalcRequest{
			Op:  pb.Operation_add,
			Val: 500,
		})
	}

	waitCh := make(chan struct{})
	go func() {
		for {
			r, err := stream.Recv()
			if err == io.EOF {
				log.Printf("receive done")
				close(waitCh)
				return
			}
			if err != nil {
				log.Fatalf("got invalid value: %v", err)
			}
			log.Printf("Result: %v | %v", r.GetResult(), r.GetError())
		}
	}()

	for _, op := range ops {
		<-time.After(time.Millisecond * 100)
		if err := stream.Send(op); err != nil {
			log.Fatalf("could not send: %v", err)
		}
	}
	stream.CloseSend()
	<-waitCh
}

func calcJointOperation(c pb.CalculatorClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := c.CalcJointOperation(ctx, &pb.CalcRequest{
		Op:  pb.Operation_add,
		Val: 100,
	})
	if err != nil {
		log.Fatalf("could not start streaming: %v", err)
	}
	log.Printf("Joint operation results")
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("got invalid value: %v", err)
		}
		log.Printf("result: %v | %v", r.GetResult(), r.GetError())
	}
}

func getCurrentOperationResults(c pb.CalculatorClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := c.GetCurrentOperationResults(ctx, &pb.CalcRequest{})
	if err != nil {
		log.Fatalf("could not start streaming: %v", err)
	}
	log.Printf("Current operation results")
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("got invalid value: %v", err)
		}
		log.Printf("result: %v | %v", r.GetResult(), r.GetError())
	}
}

func calcOperations(c pb.CalculatorClient) *pb.CalcResponse {
	stream, ctxErr := c.CalcOperations(context.Background())
	if ctxErr != nil {
		log.Fatalf("could not stream: %v", ctxErr)
	}

	var ops []*pb.CalcRequest
	for i := 1; i < 11; i++ {
		ops = append(ops, &pb.CalcRequest{
			Op:  pb.Operation_add,
			Val: int64(i),
		})
	}

	//send
	for _, op := range ops {
		if err := stream.Send(op); err != nil {
			log.Fatalf("could not send: %v", err)
		}
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("could not get response from the server: %v", err)
	}
	return response
}

func calcOperation(ctx context.Context, c pb.CalculatorClient) *pb.CalcResponse {
	r, err := c.CalcOperation(ctx, &pb.CalcRequest{
		Op:  pb.Operation_add,
		Val: 10,
	})

	if err != nil {
		log.Fatalf("could not calc: %v", err)
	}
	return r
}
