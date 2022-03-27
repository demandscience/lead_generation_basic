package nosql

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"lead_generation_basic/config"
)

var connectionString = "mongodb://%s:%s@%s:%d/?directConnection=%s&serverSelectionTimeoutMS=%d&retryWrites=%s"
var Db = connect()

// connect sets up mongo db client
func connect() *mongo.Client {
	connectionString := fmt.Sprintf(connectionString,
		config.GetString("nosqldb.user"),
		config.GetString("nosqldb.password"),
		config.GetString("nosqldb.host"),
		config.GetInt("nosqldb.port"),
		config.GetString("nosqldb.direct_conn"),
		config.GetInt("nosqldb.server_timeout_ms"),
		config.GetString("nosqldb.retry_writes"))
	clientOptions := options.Client().ApplyURI(connectionString)

	db, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic("Failed to initialize mongodb client: " + err.Error())
	}
	err = db.Ping(context.TODO(), nil)
	if err != nil {
		panic("Failed to ping mongodb: " + err.Error())
	}

	return db
}
