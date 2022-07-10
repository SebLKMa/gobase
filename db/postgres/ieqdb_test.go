package db

import (
	"testing"
)

func TestReadLastMetricsAvg(t *testing.T) {

	result, err := ReadMetricsAvgByLatestMinutes("awair-omni_18453", 240)
	if err != nil {
		t.Error(err.Error())
	}

	t.Logf("%+v\n", result)
}
