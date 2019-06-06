package layer

import (
	"github.com/pedafy/pedafy-tasks/src/database"
	"github.com/pedafy/pedafy-tasks/src/database/dbv1"
	"github.com/pedafy/pedafy-tasks/src/version"
)

// NewDatabaseHandler will create a new database handler depending on
// the given API version
func NewDatabaseHandler(currentVersion version.Version) database.DatabaseHandler {
	switch currentVersion {
	case version.Version1:
		return &dbv1.DBv1{}
	default:
		return nil
	}
}
