package clickhouse

import (
	"context"
	pb "github.com/openfms/protos/gen/device/v1"
)

const devicesLastPointsQuery = `
SELECT
    imei,
    timestamp,
    priority,
    longitude,
    latitude,
    altitude,
    angle,
    satellites,
    speed,
    io_elements
FROM
    lastpoints
WHERE
    imei IN (?);
`

// GetLastPoints returns last point of devices filtered by imei
func (adb *AVLDataBase) GetLastPoints(ctx context.Context, imeiList []string) ([]*pb.AVLData, error) {
	rows, err := adb.GetChConn().Query(ctx, devicesLastPointsQuery, imeiList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	lastPoints := make([]*pb.AVLData, 0)
	for rows.Next() {
		lastPoint := &pb.AVLData{}
		gps := &pb.GPS{}
		elements := make(map[uint16]int64)
		err := rows.Scan(
			&lastPoint.Imei,
			&lastPoint.Timestamp,
			&lastPoint.Priority,
			&gps.Longitude,
			&gps.Latitude,
			&gps.Altitude,
			&gps.Angle,
			&gps.Satellites,
			&gps.Speed,
			&lastPoint.IoElements,
			&elements,
		)
		if err != nil {
			return nil, err
		}
		lastPoint.Gps = gps
		for elementID, value := range elements {
			lastPoint.IoElements = append(lastPoint.IoElements, &pb.IOElement{
				ElementId: int32(elementID),
				Value:     value,
			})
		}
		lastPoints = append(lastPoints, lastPoint)
	}
	return lastPoints, nil
}
