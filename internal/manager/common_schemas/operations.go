package common_schemas

type OperationSchema struct {
	Name    string `json:"name" db:"name"`
	Latency int    `json:"latency" db:"latency"`
}
