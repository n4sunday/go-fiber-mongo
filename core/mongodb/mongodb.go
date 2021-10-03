package mongodb

import (
	"sync"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	Collection *mongo.Collection
	Mux        sync.Mutex
}

func (r *Repo) GetLookup(collection string, localField string, foreignID string, as string) primitive.M {
	return primitive.M{
		"$lookup": primitive.M{
			"from":         collection,
			"localField":   localField,
			"foreignField": foreignID,
			"as":           as,
		},
	}
}
