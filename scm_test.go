package scm

import(
	"testing"
)

func Test_scm(t *testing.T){
	if GetStrokeCount("龍") != 16{
		t.Error(`GetStrokeCount("龍")`)
	} 

	if GetStrokeCount("麟") != 23{
		t.Error(`GetStrokeCount("麟")`)
	} 

	if GetStrokeCount("慕") != 14{
		t.Error(`GetStrokeCount("慕")`)
	} 

	if GetStrokeCount("曦") != 20{
		t.Error(`GetStrokeCount("曦")`)
	} 

	if GetStrokeCount("一") != 1{
		t.Error(`GetStrokeCount("一")`)
	} 
}