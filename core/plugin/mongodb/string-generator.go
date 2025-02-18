package mongodb

import (
	"github.com/cecil777/infrastructure/core/object"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type objectID2String struct{}

func (o *objectID2String) Generate() string {
	return primitive.NewObjectID().Hex()
}

func NewStringGenerator() object.IStringGenerator {
	return &objectID2String{}
}
