package tests

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/strikesecurity/strikememongo"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/mgo.v2/bson"
)

const (
	// collection constants
	usersCollectionName = "users"
)

var (
	// collection variables
	usersCollection	*mongo.Collection

	databaseName 	= ""
	mongoURI 		= ""
	database		*mongo.Database
)

func TestMain(m *testing.M) {
	mongoServer, err := strikememongo.StartWithOptions(&strikememongo.Options{MongoVersion: "4.2.0", ShouldUseReplica: true})
	if err != nil {
		log.Fatal(err)
	}

	mongoURI = mongoServer.URIWithRandomDB()
	splitDatabaseName := strings.Split(mongoURI, "/")
	databaseName = splitDatabaseName[len(splitDatabaseName) - 1]

	defer mongoServer.Stop()

	setup()

	m.Run()
}

func setup() {
	startApplication()
	createCollections()
	teardown()
}

func createCollections() {
	err := database.CreateCollection(context.Background(), usersCollectionName)
	if err != nil {
		log.Fatalf("error creating collection: %s", err.Error())
	}

	usersCollection = database.Collection(usersCollectionName)
}

func startApplication() {
	dbClient, ctx, err := initDB()
	if err != nil {
		log.Fatal("error connecting to database", err)
	}

	err = dbClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error connecting to database", err)
	}

	database = dbClient.Database(databaseName)
}

func initDB() (client *mongo.Client, ctx context.Context, err error) {
	uri := fmt.Sprintf("%s%s", mongoURI, "?retryWrites=false")
	client, err = mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return
}


func teardown() {
	usersCollection.DeleteMany(context.Background(), bson.M{})
}