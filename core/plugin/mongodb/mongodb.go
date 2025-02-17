package mongodb

import "go.mongodb.org/mongo-driver/bson/primitive"

type objectID2String struct {
	objectId primitive.ObjectID
}

func (o *objectID2String) Generate() string {
	return o.objectId.Hex()
}

func NewObjectID2String() *objectID2String {
	return &objectID2String{}
}
