package onefs

import (
	"testing"
)

func TestVoiceMarsahl(t *testing.T) {

	v := voice{}
	v.Sid = "hello world"
	v.Format = "opus"
	v.Voice = "test data"

	data, err := v.marshal()
	if err != nil {
		t.Fataf("unmarshal:%s\n", err)
	}

	v1 := voice{}
	err = v1.unmarshal(data)
	if err != nil {
		t.Fataf("marshal:%s\n", err)
	}

}
