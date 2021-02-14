package model

import (
	"errors"
	"time"

	"github.com/leonardogazio/cryptocurrency/common"
)

//Rating data
type Rating struct {
	Currency Currency  `bson:"currency"`
	Datetime time.Time `bson:"datetime"`
	Vote     int       `bson:"vote"`
}

//Validate validates data input
func (r Rating) Validate() error {
	if (r.Currency == Currency{}) {
		return errors.New("Rating should have a Currency")
	}
	if (r.Datetime == time.Time{}) {
		return errors.New("Rating should have a Datetime")
	}
	acceptableVoteValues := common.ValueSlice{0, 1}
	if !acceptableVoteValues.Has(r.Vote) {
		return errors.New("Vote value is invalid. Acceptable values: 0 = Downvote / 1 = Upvote")
	}
	return nil
}
