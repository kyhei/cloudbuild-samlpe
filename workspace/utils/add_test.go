package utils

import "testing"

func TestSuccessAdd(t *testing.T) {
	expect := 3
	result := Add(2, 1)

	if expect != result {
		t.Fatalf("result is expected %d, got %d ", expect, result)
	}
}

func TestFailureAdd(t *testing.T) {
	//t.Fatal("FAILURE!!")
}
