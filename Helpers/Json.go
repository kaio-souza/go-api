package Helpers

import "encoding/json"

func JsonEncode(obj interface{}) []byte {
	jsonResponse, err := json.Marshal(obj)
	ValidateErr(err)
	return jsonResponse
}
