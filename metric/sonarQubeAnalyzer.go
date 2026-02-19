package metric

import (
	"PReQual/compilation"
	"PReQual/helper"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	_ "github.com/joho/godotenv/autoload"
)

type SonarQubeAnalyzer struct{}

func (a *SonarQubeAnalyzer) AnalyzeProject(repoName string, path string) error {
	files := []string{"base.zip", "head.zip"}

	var compiler compilation.Compiler
	compiler = &compilation.JavaCompiler{}

	for _, zipName := range files {
		zipPath := filepath.Join(path, zipName)
		if _, err := os.Stat(zipPath); err != nil {
			return fmt.Errorf("zip not found: %s", zipPath)
		}

		targetDir := filepath.Join(path, zipName[:len(zipName)-4])

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
	}

	return nil
}

func runSonarScanner(path string) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	cmd := exec.Command(
		"docker", "run", "--rm",
		"--network", "sonar",
		"-v", absPath+":/usr/src",
		"-e", "SONAR_TOKEN="+os.Getenv("SONAR_TOKEN"),
		"sonarsource/sonar-scanner-cli",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
