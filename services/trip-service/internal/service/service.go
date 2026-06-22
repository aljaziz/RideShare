package service

import (
	"aljaziz/RideShare/services/trip-service/internal/domain"
	tripTypes "aljaziz/RideShare/services/trip-service/pkg/types"
	"aljaziz/RideShare/shared/env"
	"aljaziz/RideShare/shared/types"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type service struct {
	repo domain.TripRepository
}

func NewService(repo domain.TripRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateTrip(ctx context.Context, fare *domain.RideFareModel) (*domain.TripModel, error) {
	t := &domain.TripModel{
		ID:       bson.NewObjectID(),
		UserID:   fare.UserID,
		Status:   "pending",
		RideFare: fare,
	}

	return s.repo.CreateTrip(ctx, t)
}

func (s *service) GetRoute(ctx context.Context, pickup, destination *types.Coordinate, useOSRMAPI bool) (*tripTypes.OsrmApiResponse, error) {
	if !useOSRMAPI {
		// Return a simple mock response in case we don't want to rely on an external API
		return &tripTypes.OsrmApiResponse{
			Routes: []struct {
				Distance float64 `json:"distance"`
				Duration float64 `json:"duration"`
				Geometry struct {
					Coordinates [][]float64 `json:"coordinates"`
				} `json:"geometry"`
			}{
				{
					Distance: 5.0, // 5km
					Duration: 600, // 10 minutes
					Geometry: struct {
						Coordinates [][]float64 `json:"coordinates"`
					}{
						Coordinates: [][]float64{
							{pickup.Latitude, pickup.Longitude},
							{destination.Latitude, destination.Longitude},
						},
					},
				},
			},
		}, nil
	}

	// or use our self hosted API (check the course lesson: "Preparing for External API Failures")
	baseURL := env.GetString("OSRM_API", "http://router.project-osrm.org")

	url := fmt.Sprintf("%s/route/v1/driving/%f,%f;%f,%f?overview:full&geometries=geojson", baseURL, pickup.Longitude, pickup.Latitude, destination.Longitude, destination.Latitude)

	log.Printf("Started Fetching from OSRM API: URL: %s", url)

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch route from OSRM API: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read the response: %v", err)
	}

	log.Printf("Got response from OSRM API %s", string(body))

	var routeResponse tripTypes.OsrmApiResponse
	if err := json.Unmarshal(body, &routeResponse); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	return &routeResponse, nil
}
