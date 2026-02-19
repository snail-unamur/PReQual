package compilation

type Compiler interface {
	CompileProject(path string) error
	SetSonarProperties(path string, projectName string) error
}
