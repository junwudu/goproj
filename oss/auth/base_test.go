package auth

import (
	"testing"
	"strings"
)

func testInfo(info *Info, t *testing.T) {
	if (info.Provider == "") {
		t.Fatal(info.Provider)
	}

	if len(strings.TrimSpace(info.Host)) != len(strings.TrimSpace(info.Host)) {
		t.Fatal(info.Host)
	}

	//host strict check
	if (info.Host == "" || strings.Index(info.Host, ":") < 0 || []byte(info.Host)[len(info.Host) - 1] == '/') {
		t.Fatal(info.Host)
	}
}


func TestParse(t *testing.T) {
	var info Info
	info.Parse("baidu")


	if info.Provider == "" || info.Key == "" || info.Secret == "" || info.Host == "" {
		t.Fatal(info, "info is not fullfilled!")
	}

	testInfo(&info, t)
}



func TestGetInfo(t *testing.T) {
	info := GetInfo("baidu")
	if info.Key == "" || info.Secret == "" || info.Host == "" {
		t.Fatal(info, "info is not fullfilled!")
	}

	testInfo(info, t)
}

