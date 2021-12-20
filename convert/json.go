package convert

import "encoding/json"

// WithJSON convert data with json.
func WithJSON(src, dst interface{}) error {
	rawData, err := json.Marshal(src)
	if err != nil {
		return err
	}

	return json.Unmarshal(rawData, dst)
}
