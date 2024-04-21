package clickhouse

import (
	"context"
	"fmt"
	devicepb "github.com/irisco88/protos/gen/device/v1"
	"strconv"
	"strings"
)

const devicesLastPointsQuery = `
SELECT
    imei,
    timestamp AS ts,
    toUInt8(priority) AS priority,
    longitude,
    latitude,
    toInt32(altitude),
    toInt32(angle),
    toInt32(satellites),
    toInt32(speed),
    toUInt32(event_id),
    io_elements
FROM
    lastpoints
WHERE
    imei IN (?);
`

// GetLastPoints returns last point of devices filtered by imei
func (adb *AVLDataBase) GetLastPoints(ctx context.Context, imeiList []string) ([]*devicepb.AVLData, error) {
	rows, err := adb.GetChConn().Query(ctx, devicesLastPointsQuery, imeiList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	lastPoints := make([]*devicepb.AVLData, 0)
	for rows.Next() {
		var (
			lastPoint = &devicepb.AVLData{}
			gps       = &devicepb.GPS{}
			priority  uint8
			elements  = make(map[string]string)
		)
		err := rows.Scan(
			&lastPoint.Imei,
			&lastPoint.Timestamp,
			&priority,
			&gps.Longitude,
			&gps.Latitude,
			&gps.Altitude,
			&gps.Angle,
			&gps.Satellites,
			&gps.Speed,
			&lastPoint.EventId,
			&elements,
		)
		if err != nil {
			return nil, err
		}
		lastPoint.Gps = gps
		lastPoint.Priority = devicepb.PacketPriority(priority)

		for Name, Value := range elements {
			firstFloat, secondFloat, err := splitsString(Value)
			if err != nil {
				lastPoint.IoElements = append(lastPoint.IoElements, &devicepb.IOElement{
					ElementName:  Name,
					ElementValue: 0,
					NormalValue:  0,
				})
			} else {
				lastPoint.IoElements = append(lastPoint.IoElements, &devicepb.IOElement{
					ElementName:  Name,
					ElementValue: firstFloat,
					NormalValue:  secondFloat,
				})
			}
		}

		lastPoints = append(lastPoints, lastPoint)
	}
	return lastPoints, nil
}
func splitsString(elValue string) (val float64, normalVal float64, err error) {
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
