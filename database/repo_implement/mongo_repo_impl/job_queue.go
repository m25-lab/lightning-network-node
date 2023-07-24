package mongo_repo_impl

import (
	"context"
	"fmt"
	"time"

	"github.com/m25-lab/lightning-network-node/database/models"
	repo "github.com/m25-lab/lightning-network-node/database/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type JobQueueRepoImplMongo struct {
	Db *mongo.Database
}

func (mongo *JobQueueRepoImplMongo) Publish(data *models.JobQueueData) error {
	if data == nil {
		return fmt.Errorf("Nil job data")
	}

	if data.ReadyTime.Unix() < time.Now().Unix() {
		return fmt.Errorf("Job has expired")
	}

	if _, err := mongo.Db.Collection(JobQueue).InsertOne(context.Background(), data); err != nil {
		return err
	}
	return nil
}

func (mongo *JobQueueRepoImplMongo) Consume(ctx context.Context, topic string) (*models.JobQueueData, error) {
	jobData := models.JobQueueData{}
	now := time.Now()
	response := mongo.Db.Collection(JobQueue).FindOne(ctx, bson.M{
		"topic": topic,
		"ready_time": bson.M{
			"$gte": now,
		},
	})
	if err := response.Decode(&jobData); err != nil {
		return nil, err
	}

	return &jobData, nil
}

func (mongo *JobQueueRepoImplMongo) MarkConsumed(ctx context.Context, id string, topic string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = mongo.Db.Collection(JobQueue).DeleteMany(ctx, bson.M{
		"_id": objectID,
	})
	if err != nil {
		return err
	}
	return nil
}

func NewJobQueueRepo(db *mongo.Database) repo.JobQueueRepo {
	return &JobQueueRepoImplMongo{
		Db: db,
	}
}
