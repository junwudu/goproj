package bcs

import (
	"testing"
)

func TestSign(t *testing.T) {
	bcs := new(BCS)
	bcs.host, bcs.accessKey, bcs.accessSecret = "http://bcs.duapp.com", "mzx6uUfGhzNiidxNuRjaEmTc","TNmHdLcEPBUI1cmNr2GD1tL5YRTvb72l"

	getSignedCp := bcs.sign("GET", "mytestoss", "/test", "1390838400", "", "100")

	getSigned := "http://bcs.duapp.com/mytestoss/test?sign=MBOITS:mzx6uUfGhzNiidxNuRjaEmTc:mqOmoUs6xoVwKc%2B6nmaaJh12DJ8%3D&time=1390838400&size=100"

	if getSigned != getSignedCp {
		t.Error("error" + getSignedCp)
	}



}
