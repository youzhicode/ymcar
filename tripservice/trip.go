package tripservice

import (
	"context"

	trippb "github.com/youzhicode/ymcar/proto/gen/go"
)

var TripService = &Service{}

type Service struct {
	trippb.UnimplementedTripServiceServer
}

func (*Service) GetTrip(c context.Context, req *trippb.GetTripRequest) (*trippb.GetTripResponse, error) {
	return &trippb.GetTripResponse{
		Id: req.Id,
		Trip: &trippb.Trip{
			Start:       "abc",
			End:         "def",
			DurationSec: 3600,
			FeeCent:     10000,
			StartPos: &trippb.Location{
				Latitude:  30,
				Longitude: 50,
			},
			EndPos: &trippb.Location{
				Latitude:  40,
				Longitude: 70,
			},
			PathLocations: []*trippb.Location{
				{
					Latitude:  31,
					Longitude: 51,
				},
				{
					Latitude:  35,
					Longitude: 58,
				},
			},
			Status: trippb.TripStatus_IN_PROCESS,
		},
	}, nil
}
