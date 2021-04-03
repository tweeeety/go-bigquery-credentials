package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

const (
	key       = "./key.json"
	projectID = "your-project"
	query     = "select * from your-project.datasetName.tableName"
)

func main() {

	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID, option.WithCredentialsFile(key))
	if err != nil {
		fmt.Printf("Failed to create client: %v", err)
	}
	defer client.Close()

	it, err := client.Query(query).Read(ctx)
	if err != nil {
		log.Printf("Failed to Read Query: %v", err)
	}

	for {
		var values []bigquery.Value
		err := it.Next(&values)
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println("Failed to Iterate Query:%v", err)
		}
		fmt.Println(values)
	}
}
