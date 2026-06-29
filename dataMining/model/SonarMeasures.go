package model

type SonarMeasures struct {
	Component struct {
		Key      string `json:"key"`
		Name     string `json:"name"`
		Measures []struct {
			Metric string `json:"metric"`
			Value  string `json:"value"`
		} `json:"measures"`
	} `json:"component"`
}
