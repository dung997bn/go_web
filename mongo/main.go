package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Product type
type Product struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Name        string             `json:"product_name" bson:"product_name"`
	Price       int                `json:"price" bson:"price"`
	Currency    string             `json:"currency" bson:"currency"`
	Quantity    int                `json:"quantity" bson:"quantity"`
	Discount    int                `json:"discount,omitempty" bson:"discount,omitempty"`
	Vendor      string             `json:"vendor" bson:"vendor"`
	Accessories []string           `json:"accessories,omitempty" bson:"accessories,omitempty"`
	SkuID       string             `json:"sku_id" bson:"sku_id"`
}

var iphone10 = Product{
	ID:          primitive.NewObjectID(),
	Name:        "Iphone 10",
	Price:       980,
	Currency:    "USD",
	Quantity:    40,
	Vendor:      "Apple",
	Accessories: []string{"charger", "headset", "slotopener"},
	SkuID:       "1234",
}

var trimmer = Product{
	ID:          primitive.NewObjectID(),
	Name:        "easy trimmer",
	Price:       120,
	Currency:    "USD",
	Quantity:    30,
	Vendor:      "Philips",
	Discount:    7,
	Accessories: []string{"charger", "comb", "cleaning oil"},
	SkuID:       "2345",
}

var speaker = Product{
	ID:          primitive.NewObjectID(),
	Name:        "speaker",
	Price:       300,
	Currency:    "USD",
	Quantity:    400,
	Vendor:      "Bosch",
	Discount:    4,
	Accessories: []string{"charger", "comb", "cleaning oil"},
	SkuID:       "4567",
}

var car = Product{
	ID:       primitive.NewObjectID(),
	Name:     "car",
	Price:    30000,
	Currency: "USD",
	Quantity: 400,
	Vendor:   "Bosch",
	Discount: 400,
	SkuID:    "456789",
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		fmt.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}
	db := client.Database("tronics")
	collection := db.Collection("products")
	//**************Insert one*******
	// res, err := collection.InsertOne(context.Background(), car)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(res.InsertedID)

	//*************Insert many***********
	// resMany, err := collection.InsertMany(context.Background(), []interface{}{iphone10, trimmer, speaker})
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(resMany.InsertedIDs)

	//************************************************************************
	//*******READ***************************************************
	//*************************************************

	//**************Equality Operator using FindOne***************
	// var findOne Product
	// err = collection.FindOne(context.Background(), bson.M{"price": 300}).Decode(&findOne)
	// fmt.Println("--------Equality Operator using FindOne--------")
	// fmt.Println(findOne)

	// {ObjectID("5f50dd2addaea439eaf1a863") speaker 300 USD 400 4 Bosch [charger comb cleaning oil] 4567}

	//**************Comparison operator using Find***************
	// var find Product
	// fmt.Println("--------Comparison operator using Find--------")
	// findCursor, err := collection.Find(context.Background(), bson.M{"price": bson.M{"$gt": 100}})
	// for findCursor.Next(context.Background()) {
	// 	err := findCursor.Decode(&find)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(find.Name)
	// }

	//****************Logical operator using Find****************
	// var findLogic Product
	// logicFilter := bson.M{
	// 	"$and": bson.A{
	// 		bson.M{"price": bson.M{"$gt": 100}},
	// 		bson.M{"quantity": bson.M{"$gt": 300}},
	// 	},
	// }
	// findLogicRes, err := collection.Find(context.Background(), logicFilter)

	// for findLogicRes.Next(context.Background()) {
	// 	err := findLogicRes.Decode(&findLogic)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(findLogic.Name)
	// }

	//**************Element operator using Find******************
	// var findElement Product
	// elementFilter := bson.M{
	// 	"accessories": bson.M{"$exists": false},
	// }

	// findElementRes, err := collection.Find(context.Background(), elementFilter)
	// for findElementRes.Next(context.Background()) {
	// 	err := findElementRes.Decode(&findElement)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(findElement.Name)
	// }

	//************ARRAY operator using Find***********
	// var findArray Product
	// arrayFilter := bson.M{
	// 	"accessories": bson.M{"$all": bson.A{"comb"}},//array accessories has comb
	// }

	// findArrayRes, err := collection.Find(context.Background(), arrayFilter)
	// for findArrayRes.Next(context.Background()) {
	// 	err := findArrayRes.Decode(&findArray)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(findArray.Name)
	// }

	//************************************************************************
	//*******UPDATE***************************************************
	//*************************************************

	//**********************Update operator for Field
	// updateFieldCondi := bson.M{"$set": bson.M{"IsEssential": "false"}}
	// updateFieldRes, err := collection.UpdateMany(context.Background(), bson.M{}, updateFieldCondi)
	// fmt.Println(updateFieldRes.ModifiedCount)

	//***********************Update operator for Array
	// updateArrayCondi := bson.M{"$addToSet": bson.M{"accessories": "manual"}}//add manual to accessories array
	// updateArrayRes, err := collection.UpdateMany(context.Background(), bson.M{}, updateArrayCondi)
	// fmt.Println(updateArrayRes.ModifiedCount)

	//******************Update operator for fiel -multiple operators
	// incCondi := bson.M{
	// 	"$mul": bson.M{
	// 		"price": 5,
	// 	},
	// }
	// incRes, err := collection.UpdateMany(context.Background(), bson.M{}, incCondi)
	// fmt.Println(incRes.ModifiedCount)

	//*******************Delete operator
	delCondi := bson.M{
		"price": bson.M{
			"$gt": 10000,
		},
	}
	delRes, err := collection.DeleteMany(context.Background(), delCondi)
	fmt.Println(delRes.DeletedCount)

}
