package promethy

import (
	"net/http"

	"github.com/goexl/promethy"
	"github.com/pangum/logging"
	"github.com/pangum/pangu"
)

func newPrometheus(config *pangu.Config, logger logging.Logger, mux *http.ServeMux) (registry *Registry, err error) {
	wrap := new(wrapper)
	if err = config.Load(wrap); nil != err {
		return
	}

	conf := wrap.Prometheus
	builder := promethy.New().Logger(logger)
	for key, value := range conf.Labels {
		builder.Label(key, value)
	}

	prom := builder.Build()
	registry = prom.Register()
	if handler, he := prom.Handler().Handle(); nil != he {
		err = he
	} else {
		mux.Handle(conf.Path, handler)
	}

	return
}
