package db

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/irisco88/tracking-api/db/clickhouse"
	"github.com/irisco88/tracking-api/db/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

//go:generate mockgen -source=$GOFILE -destination=mock_db/conn.go -package=$GOPACKAG
type TrackingDB interface {
	GetPgConn() *pgxpool.Pool
	GetChConn() driver.Conn
	clickhouse.AVLDBConn
	postgres.FMSDataBaseConn
}

var _ TrackingDB = &TrackingDataBase{}

type TrackingDataBase struct {
	chConn driver.Conn
	pgConn *pgxpool.Pool
	*clickhouse.AVLDataBase
	*postgres.FMSDataBase
}

func (tdb *TrackingDataBase) GetPgConn() *pgxpool.Pool {
	return tdb.pgConn
}

func (tdb *TrackingDataBase) GetChConn() driver.Conn {
	return tdb.chConn
}

func NewTrackingDB(chURL, pgURL string) (*TrackingDataBase, error) {
	avlDbConn, err := clickhouse.ConnectToAvlDB(chURL)
	if err != nil {
		return nil, err
	}
	fmsConn, err := postgres.ConnectToFmsDB(pgURL)
	if err != nil {
		return nil, err
	}
	return &TrackingDataBase{
		FMSDataBase: postgres.NewFmsDataBase(fmsConn),
		AVLDataBase: clickhouse.NewAvlDataBase(avlDbConn),
	}, nil
}
