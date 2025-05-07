package guest

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func insertGuestData(guest Guest, client *mongo.Client) error {
	collection := client.Database("wildleap").Collection("guests")
	_, err := collection.InsertOne(context.Background(), guest)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
