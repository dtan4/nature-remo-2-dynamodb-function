package main

import (
	"context"
	"fmt"
	"os"

	"github.com/pkg/errors"

	"github.com/dtan4/nature-remo-2-dynamodb-function/aws"
	"github.com/dtan4/nature-remo-2-dynamodb-function/natureremo"
)

func run(args []string) error {
	if len(os.Args) != 2 {
		return errors.New("DynamoDB table name must be provided")
	}

	table := os.Args[1]

	accessToken := os.Getenv("NATURE_REMO_ACCESS_TOKEN")

	nc := natureremo.NewClient(accessToken)

	ctx := context.Background()

	metrics, err := nc.GetRoomMetrics(ctx)
	if err != nil {
		return errors.Wrap(err, "cannot get room metrics")
	}

	for id, m := range metrics {
		fmt.Printf("id: %s, temperature: %f, humidity: %f, illumination: %f, createdAt: %s\n", id, m.Temperature, m.Humidity, m.Illumination, m.CreatedAt.Local().String())
	}

	if err := aws.Initialize(""); err != nil {
		return errors.Wrap(err, "cannot initialize AWS API clients")
	}

	if err := aws.DynamoDB.InsertRoomMetrics(table, metrics); err != nil {
		return errors.Wrap(err, "cannot insert metrics to DynamoDB")
	}

	fmt.Println("inserted to DynamoDB successfully")

	return nil
}

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
