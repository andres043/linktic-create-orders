package output

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
	"linktic-create-orders/internal/order/adapter/output/dto"
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
	orderDto := dto.ToDto(order)
	orderDto.OrderID = uuid.New().String()

	item, err := attributevalue.MarshalMap(orderDto)
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
