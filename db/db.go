package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CheckPassword(password string) int {
	// 设置mongoDB连接信息
	clientOptions := options.Client().ApplyURI("mongodb+srv://enter19961119:ZNWuC0JW1EH5a9RD@cluster0.zud9vhh.mongodb.net/test")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("Failed to connect to MongoDB:", err)
		return -1
	}
	defer client.Disconnect(context.Background())

	// 获取mongoDB集合
	collection := client.Database("vercel_chatgpt_user").Collection("users")
	// 构造查询条件
	filter := bson.M{"password": password}
	// 执行查询
	var result bson.M
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		fmt.Println("Failed to find document:", err)
		return -1
	}

	// 获取username的值
	auth, ok := result["auth"].(bool)
	if !ok {
		fmt.Println("Failed to get username")
		return -1
	}

	if auth {
		return 0
	} else {
		return -1
	}
}
