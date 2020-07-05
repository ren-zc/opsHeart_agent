package file

import "testing"

func TestFDesc_ReadStr(t *testing.T) {
	tmpFile := FDesc{
		Name: "/tmp/test_f1",
	}

	s, err := tmpFile.ReadStr()
	if err != nil {
		t.Fatalf("test failed: %s", err)
	}
	t.Logf("file content: \n%s", s)
}
