package version

// Version is a string representing API's versions
type Version = string

const (
	// Version1 First version
	Version1 = "v1.0"
)

// Default returns the default value for the API version
func Default() string {
	return Version1
}
