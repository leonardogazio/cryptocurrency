package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
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
	fmt.Print("Enter operation type [create|read|update|delete|rate|ratestream]: ")
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
		fmt.Print("Enter currency ID: ")
		id, _ := reader.ReadString('\n')
		res, err = cli.DeleteCurrency(context.Background(), &pb.DeleteCurrencyReq{Id: strings.TrimSpace(id)})
	case "rate":
		data := &pb.RateCurrencyReq{}

		fmt.Print("Enter currency ID: ")
		id, _ := reader.ReadString('\n')
		fmt.Print("Enter vote [0 = downvote / 1 = upvote]: ")
		vote, _ := reader.ReadString('\n')
		data.CurrencyId = strings.TrimSpace(id)
		data.Vote = strings.TrimSpace(vote)

		res, err = cli.RateCurrency(context.Background(), data)
	case "ratestream":
		fmt.Print("Enter currency ID: ")
		id, _ := reader.ReadString('\n')

		data := &pb.RateCurrencyReq{
			CurrencyId: strings.TrimSpace(id),
		}

		stm, _err := cli.RatingSumStream(context.Background(), data)
		if err != nil {
			err = _err
		}

		for {
			res, _err = stm.Recv()
			if _err == io.EOF {
				err = errors.New("server streaming stopped")
				break
			}
			if _err != nil {
				err = _err
				break
			}
			fmt.Println(res)
		}
	default:
		err = fmt.Errorf("invalid operation type %q", opr)
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
