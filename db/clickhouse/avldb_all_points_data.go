package clickhouse

import (
	"context"
	"fmt"
	devicepb "github.com/irisco88/protos/gen/device/v1"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

func (adb *AVLDataBase) GetAllPointsData(ctx context.Context, dataFilter string) ([]*devicepb.AVLData, error) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

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
			dataImei = " and imei in ( " + imei + " ) "
		}
		if strings.Contains(dataFilter, "priority=") {
			prioritys := strings.Split(dataFilter, "priority=")[1]
			priority := strings.Split(prioritys, "&")[0]
			dataPriority = " and priority in ( " + priority + " ) "
		}
		if strings.Contains(dataFilter, "eventid=") {
			eventids := strings.Split(dataFilter, "eventid=")[1]
			eventid := strings.Split(eventids, "&")[0]
			dataEventId = " and event_id = " + eventid + " "
		}
		dataFilterQ = " where imei!='0' " + dataFrom + dataTo + dataImei + dataPriority + dataEventId
	}
	devicesLastPointsDatasQuery := "SELECT imei,timestamp AS ts,toUInt8(priority) AS priority,longitude,latitude,toInt32(altitude),toInt32(angle),toInt32(satellites),toInt32(speed), toUInt32(event_id),io_elements FROM avlpoints " + dataFilterQ + " order by timestamp "

	rows, err := adb.GetChConn().Query(ctx, devicesLastPointsDatasQuery, &dataFilterQ)
	if err != nil {
		logger.Info("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&1",
			zap.Any("2:", err),
		)
		return nil, err
	}
	//logger.Info("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&2",
	//	zap.Any("3:", rows),
	//)
	defer rows.Close()
	lastPointsData := make([]*devicepb.AVLData, 0)
	for rows.Next() {
		var (
			lastPointData = &devicepb.AVLData{}
			gps           = &devicepb.GPS{}
			priority      uint8
			elements      = make(map[string]string)
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
			logger.Info("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&1",
				zap.Any("2:", err),
			)
			return nil, err
		}
		lastPointData.Gps = gps
		lastPointData.Priority = devicepb.PacketPriority(priority)

		for Name, Value := range elements {
			firstFloat, secondFloat, err := splitString(Value)
			if err != nil {
				lastPointData.IoElements = append(lastPointData.IoElements, &devicepb.IOElement{
					ElementName:  Name,
					ElementValue: 0,
					NormalValue:  0,
				})
			} else {
				lastPointData.IoElements = append(lastPointData.IoElements, &devicepb.IOElement{
					ElementName:  Name,
					ElementValue: firstFloat,
					NormalValue:  secondFloat,
				})
			}
		}
		lastPointsData = append(lastPointsData, lastPointData)
	}
	return lastPointsData, nil
}

func splitString(elValue string) (val float64, normalVal float64, err error) {
	underscoreIndex := strings.Index(elValue, "_")
	if underscoreIndex != -1 {
		firstPart := elValue[:underscoreIndex]
		secondPart := elValue[underscoreIndex+1:]

		firstFloat, err1 := strconv.ParseFloat(firstPart, 64)
		secondFloat, err2 := strconv.ParseFloat(secondPart, 64)

		if err1 != nil || err2 != nil {
			return 0, 0, fmt.Errorf("error converting substrings to float64")
		} else {
			return firstFloat, secondFloat, nil
		}
	} else {
		// Return 0 in both parameters if there is no underscore
		ev, err := strconv.ParseFloat(elValue, 64)
		if err != nil {
			return 0, 0, fmt.Errorf("error converting substrings to float64")
		} else {
			return ev, 0, nil
		}

	}
}
