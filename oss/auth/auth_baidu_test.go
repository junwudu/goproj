package auth

import (
	"testing"
)


func TestBaiduDo(t *testing.T) {
	a := BaiduAuth{AccessInfo{"mzx6uUfGhzNiidxNuRjaEmTc", []byte("TNmHdLcEPBUI1cmNr2GD1tL5YRTvb72l")}, "bcs.duapp.com", "baidu"}
	p := SignParameter{"PUT", "mytestoss", "/", "", "", ""}

	if url := a.Do2Url(&p);
		url != "http://bcs.duapp.com/mytestoss?sign=MBO:mzx6uUfGhzNiidxNuRjaEmTc:Ov8AyKZBJS94VRDIRuczUxzv2Rg%3D" {
		t.Fatal(url)
	}

	p = SignParameter{"GET", "mytestoss", "/boot", "", "", ""}
	if url := a.Do2Url(&p);
		url != "http://bcs.duapp.com/mytestoss/boot?sign=MBO:mzx6uUfGhzNiidxNuRjaEmTc:Dx6wWn7ryFzrYvKblmmVPOvp1w8%3D" {
		t.Fatal(url)
	}
}
