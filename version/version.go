package version

import (
	"fmt"
	"runtime"
)

// These variables are usually defined via LDFLags at compile-time
var (
	ApplicationName = "unknown"
	CommitHash      = "unknown"
	BuildDate       = "unknown"
	BuildTag        = "v0-dev"
)

// Info contains all the versioning information for a package/tool
type Info struct {
	ApplicationName string `json:"application_name"`
	CommitHash      string `json:"commit_hash"`
	BuildDate       string `json:"build_date"`
	BuildTarget     string `json:"build_target"`
	BuildTag        string `json:"build_number"`
	GoVersion       string `json:"go_version"`
}

func (vi Info) String() string {
	return fmt.Sprintf("%s %s (%s) %s - BuildDate: %s", vi.ApplicationName, vi.BuildTag, vi.CommitHash, vi.BuildTarget, vi.BuildDate)
}

// Version is the default version of the package/tool
var Version Info

func init() {
	Version = Info{
		ApplicationName: ApplicationName,
		CommitHash:      CommitHash,
		BuildDate:       BuildDate,
		BuildTag:        BuildTag,
		BuildTarget:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
		GoVersion:       runtime.Version(),
	}
}
