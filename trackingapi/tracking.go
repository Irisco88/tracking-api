package trackingapi

import (
	"github.com/nats-io/nats.go"
	trackingpb "github.com/openfms/protos/gen/tracking/v1"
	"go.uber.org/zap"
)

type TrackingService struct {
	trackingpb.UnimplementedTrackingServiceServer
	nats   *nats.Conn
	logger *zap.Logger
}

func NewTrackingService(logger *zap.Logger, natsConn *nats.Conn) *TrackingService {
	return &TrackingService{
		logger: logger,
		nats:   natsConn,
	}
}
