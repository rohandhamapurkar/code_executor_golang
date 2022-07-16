package utils

import "time"

func BuildUpdateQuery(update map[string]interface{}) {
	for k, v := range update {
		if IsZero(v) == true {
			delete(update, k)
		}
	}
	update["UpdatedAt"] = time.Now()
}
