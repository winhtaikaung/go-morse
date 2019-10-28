package main

import (
	"morsecode/morse"
	"testing"
)

const RAW_STRING = string("hello world")
const ENCODED_STRING = string("....  .  .-..  .-..  ---  /  .--  ---  .-.  .-..  -.. ")

func TestEncode(t *testing.T) {
	if morse.Encode(RAW_STRING) != ENCODED_STRING {
		t.Log("Should match with morse Encoded string Fail")
		t.Fail()
	}

}

func TestDecode(t *testing.T) {
	if morse.Decode(ENCODED_STRING) != RAW_STRING {
		t.Log("Should match with morse Decoded string Fail")
		t.Fail()
	}
}
