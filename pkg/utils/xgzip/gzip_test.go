package xgzip

import "testing"

func Test_Decompress_Success(t *testing.T) {
	testStr := "这是 1 段测试文字，gzip test string..."
	buf, _ := Compress(testStr)

	result, _ := Decompress(buf)

	if result != testStr {
		t.Errorf("expected: %s, actual: %s", testStr, result)
	}
}
