package layer

import (
	"github.com/pedafy/pedafy-tasks/src/api"
	"github.com/pedafy/pedafy-tasks/src/api/apiv1"
	"github.com/pedafy/pedafy-tasks/src/version"
)

// NewAPIHandler will return an APIHandler matching the given
// version
func NewAPIHandler(currentVersion version.Version) api.APIHandler {
	switch currentVersion {
	case version.Version1:
		return &apiv1.APIv1{}
	default:
		return nil
	}
}
