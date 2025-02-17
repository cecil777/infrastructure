package mongodb

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestObjectID2String(t *testing.T) {
	testStr := "测试ObjectID"
	testStruct := NewObjectID2String()
	testStruct.objectId, _ = primitive.ObjectIDFromHex(testStr)

	fmt.Println(testStruct.Generate())
}
