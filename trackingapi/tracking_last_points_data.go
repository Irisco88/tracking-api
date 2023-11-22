package trackingapi

import (
	"context"
	trackingv1 "github.com/irisco88/protos/gen/tracking/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (ts *TrackingService) LastPointsData(ctx context.Context, req *trackingv1.LastPointsDataRequest) (
	*trackingv1.LastPointsResponse, error) {
	lastPoints, err := ts.trackingDB.GetLastPointsData(ctx, req.DataFilter)

	ts.logger.Info("checkLastPoints",
		zap.Any("1:", req),
		zap.Any("3:", req.DataFilter),
	)
	if err != nil {
		ts.logger.Error("failed to get last points", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal error")
	}
	ts.logger.Info("checkLastPoints",
		zap.Any("4:", lastPoints),
		zap.Any("5:", err),
	)
	return &trackingv1.LastPointsResponse{
		Points: lastPoints,
	}, nil
}
