package timeutil

import (
	"testing"
)

func TestRFC3339TotimeLayout(t *testing.T) {
	t.Log(RFC3339TotimeLayout("2020-11-08T08:18:46+08:00"))
}

func TestNowLayoutString(t *testing.T) {
	t.Log(NowLayoutString())
}


func TestParseInLocation(t *testing.T) {
	t.Log(ParseInLocation(NowLayoutString()))
}

func TestLayoutStringToUnix(t *testing.T) {
	t.Log(LayoutStringToUnix(NowLayoutString()))
}

func TestGMTLayoutString(t *testing.T) {
	t.Log(GMTLayoutString())
}

func TestParseGMTInLocation(t *testing.T) {
	t.Log(ParseGMTInLocation(GMTLayoutString()))
}

func TestSubInLocation(t *testing.T) {
	ts, err := ParseGMTInLocation(GMTLayoutString())
	if err != nil {
		t.Error(err)
	}
	t.Log(SubInLocation(ts))
}