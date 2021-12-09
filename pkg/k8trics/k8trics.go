package k8trics

import (
	"github.com/sagacious-labs/kcli/pkg/k8trics/plugin"
	"github.com/sagacious-labs/kcli/pkg/utils"
)

type K8trics struct {
	baseURL string
}

// New returns a pointer to k8trics instance
func New(baseURL string) *K8trics {
	return &K8trics{
		baseURL: utils.ConstructURL(baseURL, "api", "v1"),
	}
}

// Returns handlers to plugin
func (k *K8trics) Plugin() *plugin.Plugin {
	return plugin.New(k.baseURL)
}
