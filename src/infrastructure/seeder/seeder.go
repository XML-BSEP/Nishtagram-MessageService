package seeder

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"message-service/domain"
	"time"
)

func DropDatabase(db string, mongoCli *mongo.Client, ctx context.Context){
	err := mongoCli.Database(db).Drop(ctx)
	if err != nil {
		return
	}
}

func SeedData(db string, mongoCli *mongo.Client, ctx context.Context) {
	DropDatabase(db, mongoCli, ctx)

	if cnt,_ := mongoCli.Database(db).Collection("messages").EstimatedDocumentCount(ctx, nil); cnt == 0 {
		personCollection := mongoCli.Database(db).Collection("messages")
		if err := seedMessages(personCollection, ctx); err != nil {
			log.Fatal("Error seeding data")
		}
	}

	if cnt,_ := mongoCli.Database(db).Collection("blocked_messages").EstimatedDocumentCount(ctx, nil); cnt == 0 {
		personCollection := mongoCli.Database(db).Collection("blocked_messages")
		if err := seedBlockMessages(personCollection, ctx); err != nil {
			log.Fatal("Error seeding data")
		}
	}

	if cnt,_ := mongoCli.Database(db).Collection("message_requests").EstimatedDocumentCount(ctx, nil); cnt == 0 {
		personCollection := mongoCli.Database(db).Collection("message_requests")
		if err := seedMessageRequests(personCollection, ctx); err != nil {
			log.Fatal("Error seeding data")
		}
	}


}

func seedMessages(collection *mongo.Collection, ctx context.Context) error{
	message1 := domain.Message{
		ID: "958c32e2-14b8-4c4d-8ef9-982932c34819",
 		Type: 0,
 		Timestamp: time.Now(),
 		Content: "Pozdraaav",
 		Path: "",
 		MessageFrom: domain.Profile{ID: "424935b1-766c-4f99-b306-9263731518bc"},
 		MessageTo: domain.Profile{ID: "e2b5f92e-c31b-11eb-8529-0242ac130003"},
	}

	time.Sleep(2*time.Second)

	message2 := domain.Message{
		ID: "be39e5c2-30a0-4b40-86a6-2e183f373b26",
		Type: 0,
		Timestamp: time.Now(),
		Content: "Poy",
		Path: "",
		MessageFrom: domain.Profile{ID: "424935b1-766c-4f99-b306-9263731518bc"},
		MessageTo: domain.Profile{ID: "e2b5f92e-c31b-11eb-8529-0242ac130003"},
	}

	time.Sleep(2*time.Second)

	message3 := domain.Message{
		ID: "124f443b-6a4c-4cf6-a999-6306bc518676",
		Type: 0,
		Timestamp: time.Now(),
		Content: "Kako steeeeeeee?",
		Path: "",
		MessageFrom: domain.Profile{ID: "e2b5f92e-c31b-11eb-8529-0242ac130003"},
		MessageTo: domain.Profile{ID: "424935b1-766c-4f99-b306-9263731518bc"},
	}

	message4 := domain.Message{
		ID: "1acf3f3c-fb0e-491a-b5e0-e0c2fae3b700",
		Type: 0,
		Timestamp: time.Now(),
		Content: "Poy",
		Path: "",
		MessageFrom: domain.Profile{ID: "a2c2f993-dc32-4a82-82ed-a5f6866f7d03"},
		MessageTo: domain.Profile{ID: "e2b5f92e-c31b-11eb-8529-0242ac130003"},
	}

	if _, err := collection.InsertOne(ctx, message1); err != nil {
		return err
	}

	if _, err := collection.InsertOne(ctx, message3); err != nil {
		return err
	}

	if _, err := collection.InsertOne(ctx, message2); err != nil {
		return err
	}

	if _, err := collection.InsertOne(ctx, message4); err != nil {
		return err
	}

	return nil

}

func seedBlockMessages(collection *mongo.Collection, ctx context.Context) error {
	blockMessage1 := domain.BlockMessage{
		ID : "061d564a-3f67-423e-b201-2ac99a167985",
		BlockedBy: domain.Profile{ ID: "e2b5f92e-c31b-11eb-8529-0242ac130003"},
		BlockedFor: domain.Profile{ID :"23ddb1dd-4303-428b-b506-ff313071d5d7a"},
		Timestamp: time.Now(),
	}

	if _, err := collection.InsertOne(ctx, blockMessage1); err != nil {
		return err
	}
	return nil

}

func seedMessageRequests(collection *mongo.Collection, ctx context.Context) error {
	messageRequest1 := domain.Message{
		ID: "108b9b70-a27c-4724-b396-16e0115f03f4",
		MessageFrom: domain.Profile{ID: "ead67925-e71c-43f4-8739-c3b823fe21bb"}, //user5
		MessageTo: domain.Profile{ID: "e2b5f92e-c31b-11eb-8529-0242ac130003"}, //user1
		Timestamp: time.Now(),
		Content: "Kako si?",
		Path: "",
		Type: 0,
	}

	if _, err := collection.InsertOne(ctx, messageRequest1); err != nil {
		return err
	}
	return nil

}


