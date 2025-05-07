package guest

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func insertGuestData(guest Guest, client *mongo.Client) {
	collection := client.Database("wildleap").Collection("guests")

	_, err := collection.InsertOne(context.Background(), guest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Guest data saved successfully")
}
