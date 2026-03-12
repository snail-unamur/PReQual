package metric

import (
	"PReQual/compilation"
	"PReQual/helper"
	"PReQual/model"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type SonarQubeAnalyzer struct{}

func (a *SonarQubeAnalyzer) AnalyzeProject(repoName string, path string, metrics []string) error {
	files := []string{"base.zip", "head.zip"}

	var compiler compilation.Compiler
	compiler = &compilation.JavaCompiler{}

	for _, zipName := range files {
		zipPath := filepath.Join(path, zipName)
		if _, err := os.Stat(zipPath); err != nil {
			return fmt.Errorf("zip not found: %s", zipPath)
		}

		targetDir := filepath.Join(path, zipName[:len(zipName)-4])

		defer func() {
			if err := os.RemoveAll(targetDir); err != nil {
				fmt.Printf("Warning: could not remove %s: %v\n", targetDir, err)
			}
		}()

		if err := helper.Unzip(zipPath, targetDir); err != nil {
			return err
		}

		projectRoot, err := helper.FindProjectRoot(targetDir)
		if err != nil {
			return err
		}

		if err := compiler.CompileProject(projectRoot); err != nil {
			fmt.Printf("Warning: could not compile %s: %v\n", projectRoot, err)
			return err
		}

		projectName := repoName + "-" + filepath.Base(path) + "-" + zipName[:len(zipName)-4]

		if err := compiler.SetSonarProperties(projectRoot, projectName); err != nil {
			return err
		}

		if err := runSonarScanner(projectRoot); err != nil {
			return err
		}

		if err := waitForAnalysisCompletion(projectName, metrics, 5*time.Minute); err != nil {
			return err
		}

		var data model.SonarMeasures

		data, err = retrieveSonarMetrics(projectName, metrics)
		if err != nil {
			return err
		}

		filePath := filepath.Join(path, zipName[:len(zipName)-4]+"_metrics.json")
		helper.WriteSonarMeasuresJSON(filePath, data)

	}

	return nil
}

func waitForAnalysisCompletion(projectName string, metrics []string, timeout time.Duration) error {
	sonarURL := os.Getenv("SONAR_URL")
	sonarToken := os.Getenv("SONAR_TOKEN")
	client := helper.NewHTTPClient(sonarURL, sonarToken)

	deadline := time.Now().Add(timeout)
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			path := fmt.Sprintf("/api/measures/component?metricKeys=%s&component=%s", strings.Join(metrics, ","), projectName)

			var resp model.SonarMeasures
			err := client.DoRequest("GET", path, nil, &resp)

			if err == nil && len(resp.Component.Measures) == len(metrics) {
				return nil
			}

			if time.Now().After(deadline) {
				return fmt.Errorf("timeout after %v waiting for analysis completion", timeout)
			}

		case <-time.After(timeout):
			return fmt.Errorf("timeout waiting for analysis completion")
		}
	}
}

func runSonarScanner(path string) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	cmd := exec.Command(
		"docker", "run", "--rm",
		"--network", "prequal-sonar-net",
		"-v", absPath+":/usr/src",
		"-e", "SONAR_TOKEN="+os.Getenv("SONAR_TOKEN"),
		"sonarsource/sonar-scanner-cli",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func retrieveSonarMetrics(projectName string, metrics []string) (model.SonarMeasures, error) {
	sonarURL := os.Getenv("SONAR_URL")
	sonarToken := os.Getenv("SONAR_TOKEN")

	client := helper.NewHTTPClient(sonarURL, sonarToken)

	path := fmt.Sprintf(
		"/api/measures/component?metricKeys=%s&component=%s",
		strings.Join(metrics, ","),
		projectName,
	)

	var resp model.SonarMeasures

	if err := client.DoRequest("GET", path, nil, &resp); err != nil {
		return model.SonarMeasures{}, err
	}

	return resp, nil
}
