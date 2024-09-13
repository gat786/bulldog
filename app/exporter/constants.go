package exporter

import (
	"time"
)

var (
	OUTPUTS_DIRECTORY_ENV_VAR = "OUTPUTS_DIRECTORY"
	RUNTIME_STAMP             = time.Now().Format("20060102150405")
)
