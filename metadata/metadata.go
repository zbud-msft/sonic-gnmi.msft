package metadata

import (
	"os"
)

const (
	defaultVersion string = "dev"
)

var metadataVersion string

func init() {
	if containerVersion := os.Getenv("CONTAINER_VERSION"); containerVersion != "" {
		metadataVersion = containerVersion
	} else {
		metadataVersion = defaultVersion
	}
}

func Version() string { return metadataVersion }
