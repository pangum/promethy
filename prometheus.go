package promethy

import (
	"net/http"
	"os"
	"strings"

	"github.com/goexl/promethy"
	"github.com/goexl/simaqian"
	"github.com/pangum/pangu"
)

type prometheusIn struct {
	pangu.In

	Config *pangu.Config
	Logger simaqian.Logger
	Mux    *http.ServeMux
}

func newPrometheus(in prometheusIn) (registry *Registry, err error) {
	wrap := new(wrapper)
	if err = in.Config.Load(wrap); nil != err {
		return
	}

	conf := wrap.Prometheus
	builder := promethy.New()
	builder.Logger(in.Logger)
	for key, value := range conf.Labels {
		builder.Label(key, value)
	}
	// 加载特殊的环境变量
	for _, env := range os.Environ() {
		if key, value, checked := checkEvn(env); checked {
			builder.Label(key, value)
		}
	}

	prom := builder.Build()
	registry = prom.Register()
	if handler, he := prom.Handler().Handle(); nil != he {
		err = he
	} else {
		in.Mux.Handle(conf.Path, handler)
	}

	return
}

func checkEvn(env string) (key string, value string, checked bool) {
	values := strings.Split(env, equal)
	if strings.HasPrefix(values[0], prometheusLabelKey) {
		key = values[1]
		value = os.Getenv(strings.ReplaceAll(values[0], prometheusLabelKey, prometheusLabelValue))
		if "" != value {
			checked = true
		}
	}

	return
}
