package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"strings"
	"time"

	"github.com/lecture4u/gRPC-introduction/plum/plum"
)

var operation = flag.String("o", "getPeerState", "")

func main() {
	flag.Parse()
	conn, cErr := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if cErr != nil {
		log.Fatalf("could not make connection with the node: %v", cErr)
	}
	defer conn.Close()

	farmerClient := plum.NewFarmerClient(conn)

	switch *operation {
	case "getPeerState":
		handleGetPeerState(farmerClient)
	case "getPeerStateStream":
		handleGetPeerStateStream(farmerClient)
	}
}

func handleGetPeerStateStream(client plum.FarmerClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	stream, err := client.GetPeerStateStream(ctx, &plum.Empty{})
	if err != nil {
		log.Fatalf("could not get peer state by streaming: %v", err)
	}

	for {
		ps, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("colud not get peer state by streaming: %v", err)
		}
		log.Println(formatPeerState(ps))
	}
}

func handleGetPeerState(client plum.FarmerClient) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	r, err := client.GetPeerState(ctx, &plum.Empty{})
	if err != nil {
		log.Printf("could not get the state of peer: %v", err)
	}
	log.Println(formatPeerState(r))
}

func formatPeerState(p *plum.PeerState) string {
	var s string

	t, err := makeTarget(p.GetIpv4(), p.GetPort())
	if err != nil {
		log.Println("could not make target, ", err)
	}

	s += fmt.Sprintf("\n| %s |\n", "================ PEER ================")
	s += fmt.Sprintf("|%-17s| %-20d |\n", "ID", p.GetId())
	s += fmt.Sprintf("|%-17s| %-20s |\n", "Address", t)
	s += fmt.Sprintf("|%-17s| %-20s |\n", "Role", p.Role.String())
	s += fmt.Sprintf("|%-17s| %-20d |\n", "Round", p.GetConsensusRound())
	s += fmt.Sprintf("|%-17s| %-20d |\n", "Current Primary", p.GetCurrentPrimary())
	s += fmt.Sprintf("|%-17s| %-20s |\n", "Consensus Phase", p.GetConsensusPhase().String())
	s += fmt.Sprintf("|%-17s| %-20d |\n", "Vote[Prepare]", p.Vote[int32(plum.PBFTPhase_PBFTPrepare)])
	s += fmt.Sprintf("|%-17s| %-20d |\n", "Vote[Commit]", p.Vote[int32(plum.PBFTPhase_PBFTPrepare)])
	s += fmt.Sprintf("|%-17s| %-20d |\n", "Vote[RoundChange]", p.Vote[int32(plum.PBFTPhase_PBFTRoundChange)])
	s += fmt.Sprintf("|%-17s| %-20s |\n", "Consensus State", p.GetConsensusState().String())
	s += fmt.Sprintf("|%-17s| %-20d |\n", "Block Height", p.GetBlockHeight())
	s += fmt.Sprintf("|%-17s| %-20d |\n", "Queue Length", p.GetQueueLength())
	s += fmt.Sprintf("|%-17s| %-20d |\n", "Heap Length", p.GetHeapLength())
	s += fmt.Sprintf("|%-17s| %-20f |\n", "Reputation", p.Reputation)
	s += fmt.Sprintf("|%-17s| %-20d |\n", "Selected Count", p.SelectedCount)
	s += fmt.Sprintf("|%-17s| %-20d |\n", "Tentative SC", p.TentativeSelectedCount)

	return s
}

func makeTarget(ipv4 string, port string) (string, error) {
	if len(port) == 0 {
		return "", fmt.Errorf("failed to make full target - port is missing")
	}

	if port[0] == byte(':') {
		port = strings.Trim(port, ":")
	}

	return fmt.Sprintf("%s:%s", ipv4, port), nil
}
