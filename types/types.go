package types

type OBUData struct {
	OBUID int     `json:"obuID"`
	Lat   float64 `json:"lat"`
	Long  float64 `json:"long"`
}

type Distance struct {
	Distance float64 `json:"distance"`
	OBUID    int     `json:"obuID"`
	Unix     int64   `json:"unix"`
}
