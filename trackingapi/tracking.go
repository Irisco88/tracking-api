package trackingapi

import (
	"github.com/nats-io/nats.go"
	trackingpb "github.com/openfms/protos/gen/tracking/v1"
	"github.com/openfms/tracking-api/db"
	"go.uber.org/zap"
)

type TrackingService struct {
	trackingpb.UnimplementedTrackingServiceServer
	nats       *nats.Conn
	logger     *zap.Logger
	trackingDB db.TrackingDB
}

func NewTrackingService(logger *zap.Logger, natsConn *nats.Conn, dbConn db.TrackingDB) *TrackingService {
	return &TrackingService{
		logger:     logger,
		nats:       natsConn,
		trackingDB: dbConn,
	}
}
