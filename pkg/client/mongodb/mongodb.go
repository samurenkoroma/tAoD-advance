package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, host, port, username, password, database, authDb string) (db *mongo.Database, err error) {
	var mongoDbUrl string
	var isAuth bool
	if username == "" && password == "" {
		mongoDbUrl = fmt.Sprintf("mongodb://%s:%s", host, port)
	} else {
		isAuth = true
		mongoDbUrl = fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)
	}

	clientOpts := options.Client().ApplyURI(mongoDbUrl)

	if isAuth {
		if authDb == "" {
			authDb = database
		}
		clientOpts.SetAuth(options.Credential{AuthSource: authDb, Username: username, Password: password})
	}
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect mongodb to due error: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping mongodb to due error: %v", err)
	}
	return client.Database(database), err
}
