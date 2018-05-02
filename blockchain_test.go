package main

import "testing"

func TestProofOfWork(t *testing.T) {
	bc := NewBlockchain()
	pw := bc.proofOfWork(0)
	if pw != 15194 {
		t.Fail()
	}
}
