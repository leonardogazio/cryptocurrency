package server

import (
	"context"
	"errors"
	"fmt"

	"github.com/leonardogazio/cryptocurrency/currencyserver/model"
	"github.com/leonardogazio/cryptocurrency/currencyserver/repo"
	"github.com/leonardogazio/cryptocurrency/proto/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2/bson"
)

//CreateCurrency creates a currency document
func (s *CurrencyServiceServer) CreateCurrency(ctx context.Context, req *pb.Currency) (*pb.Currency, error) {
	data := model.Currency{
		Code: req.GetCode(),
		Name: req.GetName(),
	}

	if err := data.Validate(); err != nil {
		return nil, err
	}

	res, err := repo.Database.Collection("currencies").InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	id := res.InsertedID.(primitive.ObjectID)
	req.Id = id.Hex()

	return req, nil
}

//UpdateCurrency updates a currency document
func (s *CurrencyServiceServer) UpdateCurrency(ctx context.Context, req *pb.Currency) (*pb.Currency, error) {
	id, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Could not convert supplied request id to a MongoDB ObjectId: %s", err.Error()),
		)
	}

	data := model.Currency{
		Code: req.GetCode(),
		Name: req.GetName(),
	}
	if err := data.Validate(); err != nil {
		return nil, err
	}

	result := repo.Database.Collection("currencies").FindOneAndUpdate(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": data},
		options.FindOneAndUpdate().SetReturnDocument(1),
	)

	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Could not find currency with supplied code: %s", err.Error()),
		)
	}
	return &pb.Currency{
		Id:   data.ID.Hex(),
		Code: data.Code,
		Name: data.Name,
	}, nil
}

//DeleteCurrency deletes a currency document
func (s *CurrencyServiceServer) DeleteCurrency(ctx context.Context, req *pb.DeleteCurrencyReq) (res *pb.DeleteCurrencyRes, err error) {
	id, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	if r, _ := repo.Database.Collection("currencies").DeleteOne(ctx, bson.M{"_id": id}); r.DeletedCount == 0 {
		return nil, errors.New("no currency record found by given code")
	}
	return &pb.DeleteCurrencyRes{
		Success: true,
	}, err
}

//ListCurrencies returns a stream with currencies list
func (s *CurrencyServiceServer) ListCurrencies(ctx context.Context, req *pb.ListCurrenciesReq) (*pb.CurrencyList, error) {
	data := &model.Currency{}

	filter := bson.M{
		"$or": []interface{}{
			bson.M{"code": bson.M{"$regex": primitive.Regex{Pattern: req.GetFilter(), Options: "i"}}},
			bson.M{"name": bson.M{"$regex": primitive.Regex{Pattern: req.GetFilter(), Options: "i"}}},
		},
	}

	cursor, err := repo.Database.Collection("currencies").Find(context.Background(), filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	defer cursor.Close(context.Background())

	items := []*pb.Currency{}
	for cursor.Next(context.Background()) {
		if err := cursor.Decode(data); err != nil {
			return nil, status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %s", err.Error()))
		}

		items = append(items, &pb.Currency{
			Id:   data.ID.Hex(),
			Code: data.Code,
			Name: data.Name,
		})
	}
	if err := cursor.Err(); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Unkown cursor error: %s", err.Error()))
	}
	return &pb.CurrencyList{Items: items}, nil
}
