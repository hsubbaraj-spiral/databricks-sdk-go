package databricks

import (
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/useragent"
	"github.com/databricks/databricks-sdk-go/version"
)

type Config config.Config

// Version of this SDK
func Version() string {
	return version.Version
}

// WithProduct is expected to be set by developers to differentiate their app from others.
//
// Example setting is:
//
//	func init() {
//		databricks.WithProduct("your-product", "0.0.1")
//	}
func WithProduct(name, version string) {
	useragent.WithProduct(name, version)
}
