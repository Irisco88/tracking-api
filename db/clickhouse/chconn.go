package clickhouse

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	devicepb "github.com/irisco88/protos/gen/device/v1"
	"net"
	"time"
)

type AVLDBConn interface {
	GetChConn() driver.Conn
	GetLastPoints(ctx context.Context, imeiList []string) ([]*devicepb.AVLData, error)
	GetLastPointsData(ctx context.Context, dataFilter string) ([]*devicepb.AVLData, error)
}

var _ AVLDBConn = &AVLDataBase{}

type AVLDataBase struct {
	clickhouseConn driver.Conn
}

func (adb *AVLDataBase) GetChConn() driver.Conn {
	return adb.clickhouseConn
}

func NewAvlDataBase(rawConn driver.Conn) *AVLDataBase {
	return &AVLDataBase{
		clickhouseConn: rawConn,
	}
}

func ConnectToAvlDB(databaseURL string) (driver.Conn, error) {
	opts, err := clickhouse.ParseDSN(databaseURL)
	if err != nil {
		return nil, err
	}
	opts.DialContext = func(ctx context.Context, addr string) (net.Conn, error) {
		//dialCount++
		var d net.Dialer
		return d.DialContext(ctx, "tcp", addr)
	}
	opts.Compression = &clickhouse.Compression{
		Method: clickhouse.CompressionLZ4,
	}
	opts.DialTimeout = time.Second * 30
	opts.MaxOpenConns = 5
	opts.MaxIdleConns = 5
	opts.ConnMaxLifetime = time.Duration(10) * time.Minute
	opts.ConnOpenStrategy = clickhouse.ConnOpenInOrder

	conn, err := clickhouse.Open(opts)
	if err != nil {
		return nil, err
	}
	if e := conn.Ping(context.Background()); e != nil {
		return nil, e
	}
	return conn, nil
}
