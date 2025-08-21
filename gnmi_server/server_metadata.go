package gnmi

import (
	"encoding/json"
	log "github.com/golang/glog"
	gnmipb "github.com/openconfig/gnmi/proto/gnmi"
	"github.com/sonic-net/sonic-gnmi/metadata"
	"os"
	"time"
)

const (
	metadataPrefix        string = "metadata"
	metadataEnvVar        string = "ENABLE_METADATA"
	versionPath           string = "version"
	versionKey            string = "container_version"
	versionMetadataEnvVar string = "ENABLE_METADATA_VERSION"
)

func metadataEnabled() bool {
	value := os.Getenv(metadataEnvVar)
	return value == "true"
}

func versionMetadataEnabled() bool {
	value := os.Getenv(versionMetadataEnvVar)
	return value == "true"
}

func buildVersionMetadataUpdate() *gnmipb.Update {
	if !versionMetadataEnabled() {
		log.V(4).Infof("Version metadata is disabled")
		return nil
	}

	versionPayload := map[string]string{
		versionKey: metadata.Version(),
	}

	versionData, err := json.Marshal(versionPayload)
	if err != nil {
		log.Warningf("failed to marshal version metadata: %v", err)
		return nil
	}

	return &gnmipb.Update{
		Path: &gnmipb.Path{
			Elem: []*gnmipb.PathElem{{Name: versionPath}},
		},
		Val: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_JsonIetfVal{JsonIetfVal: versionData},
		},
	}
}

func buildMetadataNotification() *gnmipb.Notification {
	if !metadataEnabled() {
		log.V(4).Infof("metadata is disabled")
		return nil
	}

	var updates []*gnmipb.Update

	if versionUpdate := buildVersionMetadataUpdate(); versionUpdate != nil {
		updates = append(updates, versionUpdate)
	}

	if len(updates) == 0 {
		return nil
	}

	return &gnmipb.Notification{
		Timestamp: time.Now().UnixNano(),
		Prefix:    &gnmipb.Path{Origin: metadataPrefix},
		Update:    updates,
	}
}
