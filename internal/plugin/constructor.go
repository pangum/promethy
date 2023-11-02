package plugin

import (
	"net/http"
	"os"
	"strings"

	"github.com/goexl/log"
	"github.com/goexl/promethy"
	"github.com/pangum/pangu"
	"github.com/pangum/promethy/internal/constant"
)

type Constructor struct {
	// 构造方法
}

func (c *Constructor) New(config *pangu.Config, mux *http.ServeMux, logger log.Logger) (registry *promethy.Registry, err error) {
	wrapper := new(Wrapper)
	if ge := config.Build().Get(wrapper); nil != ge {
		err = ge
	} else {
		registry, err = c.new(&wrapper.Prometheus, mux, logger)
	}

	return
}

func (c *Constructor) new(config *Config, mux *http.ServeMux, logger log.Logger) (registry *promethy.Registry, err error) {
	builder := promethy.New()
	builder.Logger(logger)
	for key, value := range config.Labels {
		builder.Label(key, value)
	}
	// 加载特殊的环境变量
	for _, env := range os.Environ() {
		if key, value, checked := c.checkEvn(env); checked {
			builder.Label(key, value)
		}
	}

	prom := builder.Build()
	registry = prom.Register()
	if handler, he := prom.Handler().Handle(); nil != he {
		err = he
	} else {
		mux.Handle(config.Path, handler)
	}

	return
}

func (c *Constructor) checkEvn(env string) (key string, value string, checked bool) {
	values := strings.Split(env, constant.Equal)
	if strings.HasPrefix(values[0], constant.LabelPrometheusKey) {
		key = values[1]
		value = os.Getenv(strings.ReplaceAll(values[0], constant.LabelPrometheusKey, constant.LabelPrometheusValue))
		if "" != value {
			checked = true
		}
	}

	return
}
