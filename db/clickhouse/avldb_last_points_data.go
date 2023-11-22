package clickhouse

import (
	"context"
	devicepb "github.com/irisco88/protos/gen/device/v1"
	"go.uber.org/zap"
	"strings"
)

func (adb *AVLDataBase) GetLastPointsData(ctx context.Context, dataFilter string) ([]*devicepb.AVLData, error) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	dataFilterQ := "  "
	if strings.Contains(dataFilter, "from:") || strings.Contains(dataFilter, "to:") {
		logger.Info("dataFilterQ: ",
			zap.Any("5:", dataFilterQ),
			zap.Any("6:", dataFilter),
		)
		if strings.Contains(dataFilter, "from:") {
			froms := strings.Split(dataFilter, "from:")[1]
			from := strings.Split(froms, "&")[0]
			dataFilterQ = " timestamp >= '" + from + "' "
		}
		if strings.Contains(dataFilter, "to:") {
			tos := strings.Split(dataFilter, "to:")[1]
			to := strings.Split(tos, "&")[0]
			if strings.Contains(dataFilter, "from:") {
				dataFilterQ = dataFilterQ + " and "
			}
			dataFilterQ = dataFilterQ + " timestamp <= '" + to + "' "
		}
		dataFilterQ = " where " + dataFilterQ + " order by timestamp desc"
		logger.Info("dataFilterQ: ",
			zap.Any("7:", dataFilterQ),
			zap.Any("8:", dataFilter),
		)
	}

	devicesLastPointsDatasQuery := "SELECT imei,timestamp AS ts,toUInt8(priority) AS priority,longitude,latitude,toInt32(altitude),toInt32(angle),toInt32(satellites),toInt32(speed), toUInt32(event_id),io_elements FROM avlpoints " + dataFilterQ

	rows, err := adb.GetChConn().Query(ctx, devicesLastPointsDatasQuery, &dataFilterQ)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	lastPointsData := make([]*devicepb.AVLData, 0)
	for rows.Next() {
		var (
			lastPointData = &devicepb.AVLData{}
			gps           = &devicepb.GPS{}
			priority      uint8
			elements      = make(map[string]float64)
		)
		err := rows.Scan(
			&lastPointData.Imei,
			&lastPointData.Timestamp,
			&priority,
			&gps.Longitude,
			&gps.Latitude,
			&gps.Altitude,
			&gps.Angle,
			&gps.Satellites,
			&gps.Speed,
			&lastPointData.EventId,
			&elements,
		)
		if err != nil {
			return nil, err
		}
		lastPointData.Gps = gps
		lastPointData.Priority = devicepb.PacketPriority(priority)
		for Name, Value := range elements {
			lastPointData.IoElements = append(lastPointData.IoElements, &devicepb.IOElement{
				ElementName:  Name,
				ElementValue: Value,
			})
		}
		lastPointsData = append(lastPointsData, lastPointData)
	}
	return lastPointsData, nil
}
