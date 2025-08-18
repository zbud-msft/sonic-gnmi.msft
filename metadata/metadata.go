package metadata

import (
	"os"
)

const (
	defaultVersion string = "dev"
)

var metadataVersion string

func init() {
	if containerVersion := os.Getenv("CONTAINER_VERSION"); buildVersion != "" {
		metadataVersion = containerVersion
	} else {
		metadataVersion = defaultBuildVersion
	}
}

func Version() string { return metadataVersion }
