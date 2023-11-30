package trackingapi

import (
	"context"
	trackingv1 "github.com/irisco88/protos/gen/tracking/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (ts *TrackingService) LastPointsData(ctx context.Context, req *trackingv1.LastPointsDataRequest) (
	*trackingv1.LastPointsDataResponse, error) {
	lastPoints, err := ts.trackingDB.GetLastPointsData(ctx, req.DataFilter)
	if err != nil {
		ts.logger.Error("failed to get last points", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &trackingv1.LastPointsDataResponse{
		Points: lastPoints,
	}, nil
}
