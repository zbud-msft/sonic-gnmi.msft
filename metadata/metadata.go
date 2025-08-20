package metadata

import (
	"os"
)

const (
	defaultVersion string = "dev"
	versionEnvVar  string = "CONTAINER_VERSION"
)

func Version() string {
	if containerVersion := os.Getenv(versionEnvVar); containerVersion != "" {
		return containerVersion
	} else {
		return defaultVersion
	}
}
