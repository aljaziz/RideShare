package domain

import "go.mongodb.org/mongo-driver/v2/bson"

type RideFareModel struct {
	ID                bson.ObjectID
	UserID            string
	PackageSlug       string
	TotalPriceInCents float64
}
