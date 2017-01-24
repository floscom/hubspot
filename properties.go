package hubspot

type Property struct {
	Property string      `json:"name"`
	Value    interface{} `json:"value"`
}