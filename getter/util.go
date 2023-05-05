package getter

import (
	"encoding/json"
)

func UnmarshallJson(body []byte, data any) {
	err := json.Unmarshal(body, &data)
	if err != nil {
		//panic("Could not unmarshall json response!")
		panic(err)
	}
}