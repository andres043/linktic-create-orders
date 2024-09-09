package output

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"linktic-create-orders/internal/order/model"
)

type DynamoDB struct {
	Client *dynamodb.Client
}

func NewDynamoDb() *DynamoDB {
	return &DynamoDB{
		Client: dynamodb.New(dynamodb.Options{
			Region:       "us-west-2",
			BaseEndpoint: aws.String("http://localhost:4566"),
		}),
	}
}

func (c *DynamoDB) CreateOrder(ctx context.Context, order model.Order) error {
	item, err := attributevalue.MarshalMap(order)
	orderID := uuid.New().String()
	item["order_id"] = &types.AttributeValueMemberS{Value: orderID}
	if err != nil {
		panic(err)
	}

	_, err = c.Client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String("orders"), Item: item,
	})
	if err != nil {
		return err
	}
	return nil
}
