package plugin

type Wrapper struct {
	Prometheus Config `json:"prometheus" yaml:"prometheus" xml:"prometheus" toml:"prometheus"`
}
