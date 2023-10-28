package clickhouse

import (
	"context"
	devicepb "github.com/irisco88/protos/gen/device/v1"
)

const devicesLastPointsQuery = `
SELECT
    imei,
    toUInt64(toUnixTimestamp64Milli(timestamp)) AS ts,
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
			elements  = make(map[uint16]int64)
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
		for elementID, value := range elements {
			lastPoint.IoElements = append(lastPoint.IoElements, &devicepb.IOElement{
				ElementId: int32(elementID),
				Value:     value,
			})
		}
		lastPoints = append(lastPoints, lastPoint)
	}
	return lastPoints, nil
}
