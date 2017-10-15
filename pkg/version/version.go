package version

import (
	"fmt"
	"runtime"
)

// Info contains git, build, runtime information
var Info = &info{
	AppVersion: AppVersion,
	GitCommit:  GitCommit,
	GitBranch:  GitBranch,
	GitState:   GitState,
	GitSummary: GitSummary,
	BuildDate:  BuildDate,
	GoVersion:  runtime.Version(),
	Compiler:   runtime.Compiler,
	Platform:   fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
}
