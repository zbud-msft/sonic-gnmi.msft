package gnmi

import (
	"os"
	"strconv"
	"time"

	gnmipb "github.com/openconfig/gnmi/proto/gnmi"
	"github.com/sonic-net/sonic-gnmi/metadata"
)

const (
	metadataPrefix string = "metadata"
	versionPath    string = "version"
)

func versionMetadataDisabled() bool {
	value := os.Getenv("DISABLE_METADATA_VERSION")
	return value == "1"
}

func buildVersionMetadataUpdate() *gnmipb.Update {
	if versionMetadataDisabled() {
		return nil
	}
	versionData := []byte(strconv.Quote(metadata.Version()))

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
