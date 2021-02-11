package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/leonardogazio/cryptocurrency/proto/pb"
	"google.golang.org/grpc"
)

func persistCurrency(cli pb.CurrencyServiceClient, cur *pb.Currency, createnew bool) (res *pb.Currency, err error) {
	if createnew {
		res, err = cli.CreateCurrency(context.Background(), cur)
	} else {
		res, err = cli.UpdateCurrency(context.Background(), cur)
	}
	return
}

func commandExec(cli pb.CurrencyServiceClient, reader *bufio.Reader) {
	fmt.Print("Enter operation type [create|read|update|delete]: ")
	opr, _ := reader.ReadString('\n')
	opr = strings.ToLower(strings.TrimSpace(opr))

	var res interface{}
	var err error

	switch opr {
	case "create", "update":
		data := &pb.Currency{}

		if opr == "update" {
			fmt.Print("Enter currency ID: ")
			id, _ := reader.ReadString('\n')
			data.Id = strings.TrimSpace(id)
		}

		fmt.Print("Enter new currency code [Eg. BTC]: ")
		code, _ := reader.ReadString('\n')
		fmt.Print("Enter new currency name [Eg. Bitcoin]: ")
		name, _ := reader.ReadString('\n')

		data.Code = strings.TrimSpace(code)
		data.Name = strings.TrimSpace(name)

		res, err = persistCurrency(cli, data, opr == "create")
	case "read":
		fmt.Print("Enter filter [currency CODE or NAME]: ")
		filter, _ := reader.ReadString('\n')
		res, err = cli.ListCurrencies(context.Background(), &pb.ListCurrenciesReq{Filter: strings.TrimSpace(filter)})
	case "delete":
		fmt.Print("Enter currency code: ")
		id, _ := reader.ReadString('\n')
		res, err = cli.DeleteCurrency(context.Background(), &pb.DeleteCurrencyReq{Id: strings.TrimSpace(id)})
	default:
		err = fmt.Errorf("invalid crud operation type %q", opr)
	}

	if err != nil {
		log.Printf("Error: %s\r\n", err.Error())
	} else {
		log.Printf("Response: %+v\r\n", res)
	}

	commandExec(cli, reader)
}

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to server: %s\r\n", err.Error())
	}
	defer conn.Close()

	cli := pb.NewCurrencyServiceClient(conn)
	commandExec(cli, bufio.NewReader(os.Stdin))
}
