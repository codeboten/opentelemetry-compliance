package validator

import "go.opentelemetry.io/collector/component"

type Config struct {
	Format string `mapstructure:"format"`
}

func createDefaultConfig() component.Config {
	return &Config{
		Format: "text",
	}
}
