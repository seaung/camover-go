package templates

import (
	"seaung/camover-go/internal/pkg/model"

	"gopkg.in/yaml.v2"
)

type Template struct {
	ID          string            `json:"id" yaml:"id"`
	Information model.Information `json:"info" yaml:"info"`
	Port        []string          `json:"port" yaml:"port"`
	Verify      bool              `json:"verify" yaml:"verify"`
}

func (t *Template) MarshalTemplate2Yaml() ([]byte, error) {
	out, err := yaml.Marshal(t)
	return out, err
}

func (t *Template) UnMarshalTemplate2Yaml() error {
	return nil
}
