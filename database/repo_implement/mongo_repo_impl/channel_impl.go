package mongo_repo_impl

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type ChannelRepoImplMongo struct {
	Db *mongo.Database
}

//func NewChannelRepo(db *mongo.Database) repository.ChannelRepo {
//return &ChannelRepoImplMongo{
//Db: db,
//}
//}
