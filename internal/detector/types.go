package detector

// ProjectInfo describes detected project properties.
type ProjectInfo struct {
	Language       string
	Framework      string
	ProjectName    string
	PackageManager string
	NodeVersion    string
	PythonVersion  string
	StartCommand   string
	BuildCommand   string
	DefaultPort    int
	Requirements   string
}
