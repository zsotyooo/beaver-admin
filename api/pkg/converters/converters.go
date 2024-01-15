package converters

import "encoding/json"

func StructToMap[T interface{}](structure T) (result map[string]interface{}, err error) {
	data, err := json.Marshal(structure)
	json.Unmarshal(data, &result)
	return
}
