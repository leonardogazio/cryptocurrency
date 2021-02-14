package server

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/leonardogazio/cryptocurrency/currencyserver/model"
	"github.com/leonardogazio/cryptocurrency/currencyserver/repo"
	"github.com/leonardogazio/cryptocurrency/proto/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2/bson"
)

func voteCount(ctx context.Context, currencyID primitive.ObjectID, vote int32) (int64, error) {
	return repo.Database.Collection("ratings").CountDocuments(ctx, bson.M{"currency._id": currencyID, "vote": vote})
}

//RateCurrency upvotes/downvotes a currency and returns currency ratings sum
func (s *CurrencyServiceServer) RateCurrency(ctx context.Context, req *pb.RateCurrencyReq) (res *pb.RatingSumRes, err error) {
	currencyID, err := primitive.ObjectIDFromHex(req.GetCurrencyId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "no valid Currency id provided")
	}

	currency := model.Currency{}
	result := repo.Database.Collection("currencies").FindOne(ctx, bson.M{"_id": currencyID})
	if err := result.Decode(&currency); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("No currency found by given object ID %q", req.GetCurrencyId()))
	}

	vote, err := strconv.Atoi(req.GetVote())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Vote argument is not a valid number.")
	}

	data := model.Rating{
		Currency: currency,
		Datetime: time.Now(),
		Vote:     vote,
	}
	if err := data.Validate(); err != nil {
		return nil, err
	}

	_, err = repo.Database.Collection("ratings").InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	upvoteCount, err := voteCount(ctx, currencyID, 1)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not count upvotes")
	}
	downvoteCount, err := voteCount(ctx, currencyID, 0)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not count downvotes")
	}

	return &pb.RatingSumRes{
		Currency: &pb.Currency{
			Id:   data.Currency.ID.Hex(),
			Code: data.Currency.Code,
			Name: data.Currency.Name,
		},
		Upvotes:   int32(upvoteCount),
		Downvotes: int32(downvoteCount),
	}, err
}

//RatingSumStream streams periodically to the client given Currency rating sum
func (s *CurrencyServiceServer) RatingSumStream(req *pb.RateCurrencyReq, stream pb.CurrencyService_RatingSumStreamServer) error {
	currencyID, err := primitive.ObjectIDFromHex(req.GetCurrencyId())
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "no valid Currency id provided")
	}

	ctx := context.Background()

	currency := model.Currency{}
	result := repo.Database.Collection("currencies").FindOne(ctx, bson.M{"_id": currencyID})
	if err := result.Decode(&currency); err != nil {
		return status.Errorf(codes.NotFound, fmt.Sprintf("No currency found by given object ID %q", req.GetCurrencyId()))
	}

	lastCount := struct {
		Upvotes, Downvotes int32
	}{}

	timer := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-stream.Context().Done():
			return nil
		case <-timer.C:
			upvoteCount, err := voteCount(ctx, currencyID, 1)
			if err != nil {
				return status.Errorf(codes.Internal, "could not count upvotes")
			}
			downvoteCount, err := voteCount(ctx, currencyID, 0)
			if err != nil {
				return status.Errorf(codes.Internal, "could not count downvotes")
			}

			if lastCount.Upvotes != int32(upvoteCount) || lastCount.Downvotes != int32(downvoteCount) {
				stream.Send(&pb.RatingSumRes{
					Currency: &pb.Currency{
						Id:   currency.ID.Hex(),
						Code: currency.Code,
						Name: currency.Name,
					},
					Upvotes:   int32(upvoteCount),
					Downvotes: int32(downvoteCount),
				})
				lastCount.Upvotes = int32(upvoteCount)
				lastCount.Downvotes = int32(downvoteCount)
			}
		}
	}
}
