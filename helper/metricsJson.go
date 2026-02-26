package helper

import (
	"PReQual/model"
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"
)

func WriteSonarMeasuresJSON(path string, fileName string, m model.SonarMeasures) {
	sonarData := convertMeasuresToMap(m)

	data, err := json.MarshalIndent(sonarData, "", "  ")
	if err != nil {
		panic(err)
	}

	filePath := filepath.Join(path, fileName+".json")
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		panic(err)
	}
}

func convertMeasuresToMap(measures model.SonarMeasures) map[string]interface{} {
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
