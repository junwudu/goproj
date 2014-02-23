package errors

import (
	"testing"
	"strings"
)


func TestBaiduError_Parse(t *testing.T) {
	var parser BaiduError
	s := `{"Error":{"code":"1","Message":"signature errors","LogId":"11111111"}}`
	parser.Parse(strings.NewReader(s))

	if parser.ErrorCode != 1 {
		t.Error(parser.Description)
	}

	if parser.Description != "signature errors" {
		t.Error("Description is not equal message: signature errors")
	}
}

func TestBaiduError_error(t *testing.T) {
	var e error
	bErr := new(BaiduError)

	bErr.Description = "testing"
	bErr.ErrorCode = 0
	bErr.StatusCode = 400

	e = bErr

	if e.Error() != bErr.Description {
		t.Error(e.Error())
	}
}


func TestBaiduCodeList(t *testing.T) {

	list := BaiduErrorCodeList()

	if len(list) < 1 {
		t.Error("list is empty")
	}

	if list[0].StatusCode != 400 {
		t.Error("list[0].StatusCode = " + string(list[0].StatusCode))
	}

	if list[0].ErrorCode != 1 {
		t.Error("list[0].ErrorCode = " + string(list[0].ErrorCode))
	}

	if list[len(list) - 1].StatusCode != 504 {
		t.Error("list[-1] = " + string(list[len(list) - 1].StatusCode))
	}

	if list[len(list) - 1].ErrorCode != 22 {
		t.Error("list[-1] = " + string(list[len(list) - 1].ErrorCode))
	}
}


func TestBaiduCodeMap(t *testing.T) {
	codeMap := BaiduErrorCodeMap()

	if codeMap["16"].StatusCode != codeMap["12"].StatusCode {
		t.Error(codeMap["16"], codeMap["12"])
	}
}


func TestGetBaiduError(t *testing.T) {
	err := GetBaiduError(16)

	if err.ErrorCode == 16 {
		t.Error(err)
	}

	if err.StatusCode == 503 {
		t.Error(err)
	}

}
