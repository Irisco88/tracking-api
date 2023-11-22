package clickhouse

import (
	"context"
	devicepb "github.com/irisco88/protos/gen/device/v1"
	"strings"
)

func (adb *AVLDataBase) GetLastPointsData(ctx context.Context, dataFilter string) ([]*devicepb.AVLData, error) {
	//logger, _ := zap.NewDevelopment()
	//defer logger.Sync()
	dataFilterQ := "  "
	if strings.Contains(dataFilter, "from=") || strings.Contains(dataFilter, "to=") || strings.Contains(dataFilter, "imei=") || strings.Contains(dataFilter, "priority=") || strings.Contains(dataFilter, "eventid=") {
		dataFrom := ""
		dataTo := ""
		dataImei := ""
		dataPriority := ""
		dataEventId := ""
		if strings.Contains(dataFilter, "from=") {
			froms := strings.Split(dataFilter, "from=")[1]
			from := strings.Split(froms, "&")[0]
			dataFrom = " and  timestamp >= '" + from + "' "
		}
		if strings.Contains(dataFilter, "to=") {
			tos := strings.Split(dataFilter, "to=")[1]
			to := strings.Split(tos, "&")[0]
			dataTo = " and timestamp <= '" + to + "' "
		}
		if strings.Contains(dataFilter, "imei=") {
			imeis := strings.Split(dataFilter, "imei=")[1]
			imei := strings.Split(imeis, "&")[0]
			dataImei = " and imei = '" + imei + "' "
		}
		if strings.Contains(dataFilter, "priority=") {
			prioritys := strings.Split(dataFilter, "priority=")[1]
			priority := strings.Split(prioritys, "&")[0]
			dataPriority = " and priority = '" + priority + "' "
		}
		if strings.Contains(dataFilter, "eventid=") {
			eventids := strings.Split(dataFilter, "eventid=")[1]
			eventid := strings.Split(eventids, "&")[0]
			dataEventId = " and event_id = " + eventid + " "
		}
		dataFilterQ = " where imei!='0' " + dataFrom + dataTo + dataImei + dataPriority + dataEventId + " order by timestamp desc"
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
