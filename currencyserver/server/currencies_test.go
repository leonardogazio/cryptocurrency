package server

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/leonardogazio/cryptocurrency/currencyserver/repo"
	"github.com/leonardogazio/cryptocurrency/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	if err := repo.Open(); err != nil {
		log.Fatalf("Server exited with error: %v", err)
	}

	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterCurrencyServiceServer(s, &CurrencyServiceServer{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestCreateCurrency(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	cli := pb.NewCurrencyServiceClient(conn)

	requests := []*pb.Currency{
		{Code: "BTC", Name: "Bitcoin"},
		{Code: "ADA", Name: "Cardano"},
		{Code: "ETH", Name: "Ethereum"},
		{Code: "DOT"},
		{Name: "Litecoin"},
		{},
	}

	for _, req := range requests {
		res, err := cli.CreateCurrency(ctx, req)
		if err != nil {
			t.Errorf("CreateCurrency test failed: %s", err.Error())
			continue
		}
		log.Printf("Response: %+v", res)
	}
}
