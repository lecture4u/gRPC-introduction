package main

import (
	"context"
	pb "github.com/lecture4u/gRPC-introduction/calc/calc"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

const port = ":50051"

var decreaseUnit = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

type calculator struct {
	data   int64
	record []*pb.CalcRequest
}

// server is used to implement helloworld.GreeterServer
type server struct {
	pb.UnimplementedCalculatorServer
	c *calculator
	m *sync.RWMutex
}

func (s *server) CalcOperation(ctx context.Context, in *pb.CalcRequest) (*pb.CalcResponse, error) {
	log.Printf("received: %s%s%v", in.GetOp().String(), " , ", in.GetVal())
	processCalculation(s.c, in)

	return &pb.CalcResponse{
		Error:  true,
		Result: s.c.data,
	}, nil
}

func (s *server) CalcOperations(stream pb.Calculator_CalcOperationsServer) error {
	for {
		o, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.CalcResponse{
				Error:  true,
				Result: s.c.data,
			})
		}
		if err != nil {
			return err
		}
		processCalculation(s.c, o)
	}
}

//GetCurrentOperationResults(*CalcRequest, Calculator_GetCurrentOperationResultsServer) error
func (s *server) GetCurrentOperationResults(in *pb.CalcRequest, stream pb.Calculator_GetCurrentOperationResultsServer) error {
	for _, r := range s.c.record {
		if err := stream.Send(&pb.CalcResponse{
			Error:  true,
			Result: r.Val,
		}); err != nil {
			return err
		}
	}
	return nil
}

func (s *server) CalcJointOperation(in *pb.CalcRequest, stream pb.Calculator_CalcJointOperationServer) error {
	log.Printf("received: %s%s%v", in.GetOp().String(), " , ", in.GetVal())
	processCalculation(s.c, in)

	for _, d := range decreaseUnit {
		<-time.After(time.Millisecond * 500)
		processCalculation(s.c, &pb.CalcRequest{
			Op:  pb.Operation_sub,
			Val: int64(d),
		})

		stream.Send(&pb.CalcResponse{
			Error:  true,
			Result: s.c.data,
		})
	}
	return nil
}

func (s *server) CalcJointOperations(stream pb.Calculator_CalcJointOperationsServer) error {
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		log.Printf("received: %s%s%v", r.GetOp().String(), " , ", r.GetVal())

		processCalculation(s.c, r)
		go func() {
			for _, d := range decreaseUnit {
				<-time.After(time.Millisecond)
				s.m.Lock()
				processCalculation(s.c, &pb.CalcRequest{
					Op:  pb.Operation_sub,
					Val: int64(d),
				})
				result := s.c.data
				s.m.Unlock()
				stream.Send(&pb.CalcResponse{
					Error:  true,
					Result: result,
				})
			}
		}()
	}
	return nil
}

func processCalculation(c *calculator, in *pb.CalcRequest) {
	switch in.GetOp() {
	case pb.Operation_add:
		c.data += in.GetVal()
	case pb.Operation_sub:
		c.data -= in.GetVal()
	case pb.Operation_mul:
		c.data *= in.GetVal()
	case pb.Operation_div:
		c.data /= in.GetVal()
	case pb.Operation_clear:
		c.data = 0
	default:
		log.Panic("invalid operation")
	}
	c.record = append(c.record, in)
}

func main() {
	server := server{
		c: &calculator{},
		//m: new(sync.RWMutex),
		m: &sync.RWMutex{},
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterCalculatorServer(s, &server)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
