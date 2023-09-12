package types

type Distance struct {
	Value  float64 `json:"value"`
	UnitId int     `json:"unitId"`
	Unix   int64   `json:"unix"`
}

type UnitCoordinate struct {
	UnitId    int     `json:"unitId"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Invoice struct {
	UnitId		int     `json:"unitId"`
	TotalDistance		float64   `json:"total"`
	TotalCharge		float64   `json:"charge"`
}
