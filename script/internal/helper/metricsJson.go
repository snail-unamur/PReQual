package helper

import (
	"PReQual/internal/domain"
	"strconv"
)

func ConvertMeasuresToMap(measures domain.SonarMeasures) map[string]interface{} {
	result := make(map[string]interface{})

	for _, m := range measures.Component.Measures {
		if f, err := strconv.ParseFloat(m.Value, 64); err == nil {
			result[m.Metric] = f
		} else {
			result[m.Metric] = m.Value
		}
	}

	return result
}
