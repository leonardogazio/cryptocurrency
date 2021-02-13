package server

import (
	"context"
	"log"
	"net"
	"os"
	"testing"

	"github.com/icrowley/fake"
	"github.com/leonardogazio/cryptocurrency/currencyserver/repo"
	"github.com/leonardogazio/cryptocurrency/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

var lis *bufconn.Listener
var ctx context.Context
var updateRequests []*pb.Currency

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestMain(m *testing.M) {
	if err := repo.Open(); err != nil {
		log.Fatalf("Server exited with error: %v", err)
	}

	bufSize := 1024 * 1024
	ctx = context.Background()
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterCurrencyServiceServer(s, &CurrencyServiceServer{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	os.Exit(m.Run())
}

func TestCreateCurrency(t *testing.T) {
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	cli := pb.NewCurrencyServiceClient(conn)

	requests := make([]*pb.Currency, 10)
	for i := 0; i <= 6; i++ {
		requests[i] = &pb.Currency{Code: fake.CurrencyCode(), Name: fake.Word()}
	}

	//user input error cases
	requests[7] = &pb.Currency{Code: fake.CurrencyCode()}
	requests[8] = &pb.Currency{Name: fake.Word()}
	requests[9] = &pb.Currency{}

	for _, req := range requests {
		res, err := cli.CreateCurrency(ctx, req)
		if err != nil {
			t.Errorf("CreateCurrency test failed: %s", err.Error())
			continue
		}
		updateRequests = append(updateRequests, res)
		log.Printf("CreateCurrency Response: %+v", res)
	}
}

func TestUpdateCurrency(t *testing.T) {
	if len(updateRequests) == 0 {
		t.Fatal("No Currency ID for this testing, execute TestCreateCurrency first.")
	}

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	cli := pb.NewCurrencyServiceClient(conn)

	for _, req := range updateRequests {
		req.Code = fake.CurrencyCode()
		req.Name = fake.Word()
		res, err := cli.UpdateCurrency(ctx, req)
		if err != nil {
			t.Errorf("UpdateCurrency test failed: %s", err.Error())
			continue
		}
		log.Printf("UpdateCurrency Response: %+v", res)
	}
}

func TestListCurrencies(t *testing.T) {
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	cli := pb.NewCurrencyServiceClient(conn)

	res, err := cli.ListCurrencies(ctx, &pb.ListCurrenciesReq{})
	if err != nil {
		t.Errorf("ListCurrencies test failed: %s", err.Error())
		return
	}
	for _, currency := range res.GetItems() {
		log.Printf("ListCurrencies Response: %+v", currency)
	}
}

func TestDeleteCurrency(t *testing.T) {
	if len(updateRequests) == 0 {
		t.Fatal("No Currency ID for this testing, execute TestCreateCurrency first.")
	}

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	cli := pb.NewCurrencyServiceClient(conn)

	for _, req := range updateRequests {
		res, err := cli.DeleteCurrency(ctx, &pb.DeleteCurrencyReq{Id: req.GetId()})
		if err != nil {
			t.Errorf("DeleteCurrency test failed: %s", err.Error())
			continue
		}
		log.Printf("DeleteCurrency Response: %+v", res)
	}
}
