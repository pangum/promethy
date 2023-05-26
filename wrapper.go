package promethy

type wrapper struct {
	Prometheus *config `json:"prometheus" yaml:"prometheus" xml:"prometheus" toml:"prometheus" validate:"required"`
}
