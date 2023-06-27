package trackingapi

import (
	"context"
	"errors"
	"github.com/golang/protobuf/proto"
	devicepb "github.com/openfms/protos/gen/device/v1"
	trkpb "github.com/openfms/protos/gen/tracking/v1"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

func (ts *TrackingService) LiveDevices(req *trkpb.LiveDevicesRequest, stream trkpb.TrackingService_LiveDevicesServer) error {
	ctx := stream.Context()
	subChan, err := ts.nats.SubscribeSync("device.lastpoint.*")
	if err != nil {
		ts.logger.Error("error subscribe on last point channel", zap.Error(err))
		return status.Error(codes.Internal, "internal error")
	}
	defer func() {
		if err := subChan.Unsubscribe(); err != nil {
			ts.logger.Error("error unsub from subscription", zap.Error(err))
		}
	}()
	for {
		msg, err := subChan.NextMsgWithContext(ctx)
		if err != nil {
			if errors.Is(err, context.Canceled) ||
				errors.Is(err, context.DeadlineExceeded) {
				return nil
			}
			ts.logger.Error("can't get new last point", zap.Error(err))
			return status.Error(codes.Internal, "internal error")
		}
		if !slices.ContainsFunc(req.ImeiList, func(imei string) bool {
			return strings.Contains(msg.Subject, imei)
		}) {
			continue
		}
		go ts.SendLastPoint(stream, msg.Data)
	}
}

func (ts *TrackingService) SendLastPoint(out trkpb.TrackingService_LiveDevicesServer, msgData []byte) {
	point := &devicepb.AVLData{}
	if e := proto.Unmarshal(msgData, point); e != nil {
		ts.logger.Error("error unmarshal last point",
			zap.String("data", string(msgData)),
			zap.Error(errors.Unwrap(e)))
		return
	}
	if sendErr := out.Send(&trkpb.LiveDevicesResponse{Point: point}); sendErr != nil {
		ts.logger.Error("can not send data",
			zap.Error(sendErr))
		return
	}
}
