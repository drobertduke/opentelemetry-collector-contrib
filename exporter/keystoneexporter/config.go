package keystoneexporter

import (
	"go.opentelemetry.io/collector/config/configmodels"
)

type Config struct {
	configmodels.ExporterSettings `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct.

	KeystoneUrlFormat string `mapstructure:"keystone_url_format"`
}
