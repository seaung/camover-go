package model

type Information struct {
	Name        string     `json:"name,omitempty" yaml:"name,omitempty"`
	Author      string     `json:"author,omitempty" yaml:"author,omitempty"`
	Description string     `json:"description,omitempty" yaml:"description,omitempty"`
	Reference   string     `json:"reference,omitempty" yaml:"reference,omitempty"`
	Classifiy   *Classifiy `json:"classifiy,omitempty" yaml:"classifiy,omitempty"`
}

type Classifiy struct {
	CVEID string `json:"cve-id,omitempty" yaml:"cve-id,omitempty"`
	CWEID string `json:"cwe-id,omitempty" yaml:"cwe-id,omitempty"`
	CPE   string `json:"cpe,omitempty" yaml:"cpe,omitempty"`
}
