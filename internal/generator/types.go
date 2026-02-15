package generator

// Config describes the Dockerfile/Containerfile generation inputs.
type Config struct {
	Language        string
	OutputFile      string
	AppName         string
	WorkDir         string
	Port            int
	NodeVersion     string
	PythonVersion   string
	PackageManager  string
	BuildCommand    string
	StartCommand    string
	Requirements    string
	ExposePort      bool
	UseAlpineImages bool
}
