package trackingapi

import (
	"context"
	trackingv1 "github.com/irisco88/protos/gen/tracking/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (ts *TrackingService) AllPointsData(ctx context.Context, req *trackingv1.AllPointsDataRequest) (
	*trackingv1.AllPointsDataResponse, error) {
	AllPoints, err := ts.trackingDB.GetAllPointsData(ctx, req.DataFilter)
	if err != nil {
		ts.logger.Error("failed to get All points", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &trackingv1.AllPointsDataResponse{
		Points: AllPoints,
	}, nil
}
