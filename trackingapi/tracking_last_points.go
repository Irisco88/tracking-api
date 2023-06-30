package trackingapi

import (
	"context"
	trackingv1 "github.com/openfms/protos/gen/tracking/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (ts *TrackingService) LastPoints(ctx context.Context, req *trackingv1.LastPointsRequest) (
	*trackingv1.LastPointsResponse, error) {
	if len(req.ImeiList) == 0 {
		return nil, status.Error(codes.InvalidArgument, "imei list required")
	}
	lastPoints, err := ts.trackingDB.GetLastPoints(ctx, req.GetImeiList())
	if err != nil {
		ts.logger.Error("failed to get last points", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &trackingv1.LastPointsResponse{
		Points: lastPoints,
	}, nil
}
