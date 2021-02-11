package model

import (
	"errors"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Currency data
type Currency struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Code string             `bson:"code"`
	Name string             `bson:"name"`
}

//Validate validates data input
func (c Currency) Validate() (err error) {
	if strings.TrimSpace(c.Code) == "" || strings.TrimSpace(c.Name) == "" {
		err = errors.New("code/name should have values")
	}
	return
}
