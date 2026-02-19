package metric

type ProjectAnalyser interface {
	AnalyzeProject(projectName string, path string) error
}
