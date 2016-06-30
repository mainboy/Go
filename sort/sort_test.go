package sort

import "testing"

func TestLen(t *testing.T) {
	IntLen := IntArray{1, 2, 3, 4, 5}
	if IntLen.Len() != 5 {
		t.Errorf("error type int count")
	}

	StringLen := StringArray{"1", "2", "3"}
	if StringLen.Len() != 3 {
		t.Errorf("error type string count")
	}

}
