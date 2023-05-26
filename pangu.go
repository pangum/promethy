package promethy

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Dependencies(
		newPrometheus,
	)
}
