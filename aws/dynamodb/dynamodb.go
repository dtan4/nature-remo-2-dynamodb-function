package dynamodb

import (
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/pkg/errors"

	"github.com/dtan4/nature-remo-2-dynamodb-function/types"
)

// Client represents a wrapper of DynamoDB API client
type Client struct {
	api dynamodbiface.DynamoDBAPI
}

// NewClient creates DynamoDB API client wrapper
func NewClient(api dynamodbiface.DynamoDBAPI) *Client {
	return &Client{
		api: api,
	}
}

// InsertRoomMetrics inserts room metrics to DynamoDB
func (c *Client) InsertRoomMetrics(table string, metrics map[string]*types.RoomMetrics) error {
	// do nothing when the given metrics is empty
	if len(metrics) == 0 {
		return nil
	}

	// TODO(dtan4): use batchWriteItem API to reduce API calls

	for id, m := range metrics {
		input := &dynamodb.PutItemInput{
			Item: map[string]*dynamodb.AttributeValue{
				"ID": {
					S: aws.String(id),
				},
				"Temperature": {
					N: aws.String(strconv.FormatFloat(m.Temperature, 'f', -1, 64)),
				},
				"Humidity": {
					N: aws.String(strconv.FormatFloat(m.Humidity, 'f', -1, 64)),
				},
				"Illumination": {
					N: aws.String(strconv.FormatFloat(m.Illumination, 'f', -1, 64)),
				},
				"CreatedAt": {
					N: aws.String(strconv.FormatInt(m.CreatedAt.Unix(), 10)),
				},
			},
			TableName: aws.String(table),
		}

		if _, err := c.api.PutItem(input); err != nil {
			return errors.Wrapf(err, "cannot insert metrics: %#v", m)
		}
	}

	return nil
}
