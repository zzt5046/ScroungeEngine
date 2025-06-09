package api

import (
	"encoding/json"
)

// Recipe schema
func RecipeSchema() []byte {
	schema := map[string]interface{}{
        "$schema":       "http://json-schema.org/draft-07/schema#",
        "title":         "Recipe",
        "type":          "object",
        "properties": map[string]interface{}{
            "name": map[string]interface{}{
                "type": "string",
            },
            "description": map[string]interface{}{
                "type": "string",
            },
			"ingredients": map[string]interface{}{
                "type": "array",
                "items": map[string]interface{}{
					"type": "string",
				},
            },
			"directions": map[string]interface{}{
                "type": "array",
                "items": map[string]interface{}{
					"type": "string",
				},
            },
			"notes": map[string]interface{}{
                "type": "string",
            },
        },
    }

	return toByteArray(schema)
}

func toByteArray(schema map[string]interface{}) []byte {
	jsonSchema, err := json.Marshal(schema)
    if err != nil {
        panic(err)
    } else {
		return jsonSchema
	}
}