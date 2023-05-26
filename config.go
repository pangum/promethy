package promethy

type config struct {
	// 是否开启
	Enabled *bool `default:"true" json:"enabled" yaml:"enabled" xml:"enabled" toml:"enabled"`
	// 路径
	Path string `default:"/metrics" json:"path" yaml:"path" xml:"path" toml:"path" validate:"required"`
	// 标签
	Labels map[string]string `json:"labels" yaml:"labels" xml:"labels" toml:"labels"`
}
