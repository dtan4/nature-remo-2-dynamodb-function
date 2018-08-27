package main

import (
	"context"
	"fmt"
	"os"

	"github.com/dtan4/nature-remo-2-dynamodb-function/natureremo"
	"github.com/pkg/errors"
)

func run(args []string) error {
	if len(os.Args) != 2 {
		return errors.New("Nature Remo API access token must be provided")
	}

	accessToken := os.Args[1]
	nc := natureremo.NewClient(accessToken)

	ctx := context.Background()

	metrics, err := nc.GetRoomMetrics(ctx)
	if err != nil {
		return errors.Wrap(err, "cannot get room metrics")
	}

	for id, m := range metrics {
		fmt.Printf("id: %s, temperature: %f, humidity: %f, illumination: %f, createdAt: %s\n", id, m.Temperature, m.Humidity, m.Illumination, m.CreatedAt.Local().String())
	}

	return nil
}

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
