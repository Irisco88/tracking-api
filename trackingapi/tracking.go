package trackingapi

import (
	"github.com/nats-io/nats.go"
	"github.com/openfms/authutil"
	commonpb "github.com/openfms/protos/gen/common/v1"
	trackingpb "github.com/openfms/protos/gen/tracking/v1"
	"github.com/openfms/tracking-api/db"
	"go.uber.org/zap"
)

type TrackingService struct {
	trackingpb.UnimplementedTrackingServiceServer
	nats        *nats.Conn
	logger      *zap.Logger
	trackingDB  db.TrackingDB
	authManager *authutil.AuthManager
}

func NewTrackingService(logger *zap.Logger, natsConn *nats.Conn, dbConn db.TrackingDB, auth *authutil.AuthManager) *TrackingService {
	return &TrackingService{
		logger:      logger,
		nats:        natsConn,
		trackingDB:  dbConn,
		authManager: auth,
	}
}

func (ts *TrackingService) GetAuthManager() *authutil.AuthManager {
	return ts.authManager
}

func (ts *TrackingService) GetRoleAccess(fullMethod string) []commonpb.UserRole {
	methodPerms, ok := trackingpb.TrackingServicePermission.MethodStreams[fullMethod]
	if ok {
		return methodPerms.Roles
	}
	return nil
}
