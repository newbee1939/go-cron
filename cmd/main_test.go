package main

import "testing"

// falseならテストを実施
var Debug bool = false

func TestSample_サンプルテスト(t *testing.T) {
	if Debug {
		t.Skip("テストをスキップします")
	}

	result := Sample()
	expected := "success"

	if result != expected {
		t.Error("テスト失敗です")
	}

	t.Log("TestSample終了")
}
