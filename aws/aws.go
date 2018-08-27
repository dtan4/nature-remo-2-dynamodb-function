package aws

import (
	"github.com/dtan4/nature-remo-2-dynamodb-function/aws/dynamodb"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	dynamodbapi "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/pkg/errors"
)

var (
	// DynamoDB represents DynamoDB API client
	DynamoDB *dynamodb.Client
)

// Initialize initializes AWS API clients
func Initialize(region string) error {
	var (
		sess *session.Session
		err  error
	)

	if region == "" {
		sess, err = session.NewSession()
		if err != nil {
			return errors.Wrap(err, "cannot create new AWS session.")
		}
	} else {
		sess, err = session.NewSession(&aws.Config{Region: aws.String(region)})
		if err != nil {
			return errors.Wrap(err, "cannot create new AWS session.")
		}
	}

	DynamoDB = dynamodb.NewClient(dynamodbapi.New(sess))

	return nil
}
