package file

import "testing"

func TestFDesc_AppendStr(t *testing.T) {
	tmpFile := FDesc{
		Name: "/tmp/test_f1",
	}

	err := tmpFile.AppendStr("test str line3\n")
	if err != nil {
		t.Fatalf("failed: %s", err)
	}
}
