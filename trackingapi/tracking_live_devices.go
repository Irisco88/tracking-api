package trackingapi

import (
	"errors"
	"github.com/golang/protobuf/proto"
	devicepb "github.com/irisco88/protos/gen/device/v1"
	trkpb "github.com/irisco88/protos/gen/tracking/v1"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

func (ts *TrackingService) LiveDevices(stream trkpb.TrackingService_LiveDevicesServer) error {
	ctx := stream.Context()
	req, err := stream.Recv()
	if err != nil {
		ts.logger.Error("error receiving request", zap.Error(err))
		return status.Error(codes.Internal, "internal error")
	}
	if len(req.ImeiList) == 0 {
		return status.Error(codes.InvalidArgument, "imei list is required")
	}
	subChan, err := ts.nats.SubscribeSync("device.lastpoint.*")
	if err != nil {
		ts.logger.Error("error subscribing to last point channel", zap.Error(err))
		return status.Error(codes.Internal, "internal error")
	}
	defer func() {
		if err := subChan.Unsubscribe(); err != nil {
			ts.logger.Error("error unsubscribing from subscription", zap.Error(err))
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			msg, err := subChan.NextMsgWithContext(ctx)
			if err != nil {
				ts.logger.Error("error getting new last point", zap.Error(err))
				return status.Error(codes.Internal, "internal error")
			}
			if !slices.ContainsFunc(req.ImeiList, func(imei string) bool {
				return strings.Contains(msg.Subject, imei)
			}) {
				continue
			}
			point := &devicepb.AVLData{}
			if e := proto.Unmarshal(msg.Data, point); e != nil {
				ts.logger.Error("error unmarshal last point", zap.Error(errors.Unwrap(e)))
				continue
			}
			resp := &trkpb.LiveDevicesResponse{
				Point: point,
			}
			if err := stream.Send(resp); err != nil {
				ts.logger.Error("error sending data", zap.Error(err))
				return status.Error(codes.Internal, "internal error")
			}
		}
	}
}
