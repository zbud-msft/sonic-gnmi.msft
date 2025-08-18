package gnmi

import (
	"os"
	"strconv"
	"time"

	"github.com/sonic-net/sonic-gnmi/internal/buildver"
	gnmipb "github.com/openconfig/gnmi/proto/gnmi"
)

const (
	metadataPrefix string = "sonic_server_meta"
	versionPath string = "version"
)

func versionMetaDisabled() bool {
	value := os.Getenv("DISABLE_METAVERSION")
	return value == "1"
}

func buildVersionMetadataNotification() *gnmipb.Notification {
	if versionMetaDisabled() {
		return nil
	}
	versionData := []byte(strconv.Quote(buildver.Version()))

	return &gnmipb.Notification{
		Timestamp: time.Now().UnixNano(),
		Prefix: &gnmipb.Path{Origin: metadataPrefix},
		Update: []*gnmipb.Update{{
			Path: &gnmipb.Path{
				Elem: []*gnmipb.PathElem{{Name: versionPath}},
			},
			Val: &gnmipb.TypedValue{
				Value: &gnmipb.TypedValue_JsonIetfVal{JsonIetfVal: versionData},
			},
		}},
	}
}
