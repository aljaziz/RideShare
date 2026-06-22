package domain

import (
	tripTypes "aljaziz/RideShare/services/trip-service/pkg/types"
	"aljaziz/RideShare/shared/types"
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type TripModel struct {
	ID       bson.ObjectID
	UserID   string
	Status   string
	RideFare *RideFareModel
}

type TripRepository interface {
	CreateTrip(ctx context.Context, trip *TripModel) (*TripModel, error)
}

type TripService interface {
	CreateTrip(ctx context.Context, fare *RideFareModel) (*TripModel, error)
	GetRoute(ctx context.Context, pickup, destination *types.Coordinate, useOSRMAPI bool) (*tripTypes.OsrmApiResponse, error)
}
