package version

// info contains versioning information.
type info struct {
	AppVersion string `json:"appVersion"`
	GitCommit  string `json:"gitCommit"`
	GitBranch  string `json:"gitBranch"`
	GitState   string `json:"gitState"`
	GitSummary string `json:"gitSummary"`
	BuildDate  string `json:"buildDate"`
	GoVersion  string `json:"goVersion"`
	Compiler   string `json:"compiler"`
	Platform   string `json:"platform"`
}

var (
	GitBranch  string
	GitState   string
	GitSummary string
	BuildDate  string
	AppVersion string
	GitCommit  string
)
