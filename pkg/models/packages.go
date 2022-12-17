package models

type PackageType string

type Package struct {
	Type    PackageType `json:"type"`
	Name    string
	Epoch   string
	Version string
	Release string
	Arch    string
}
