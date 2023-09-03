package h264parser

import (
	"encoding/hex"
	"testing"
)

func FuzzDecodeAVCC(f *testing.F) {
	var (
		invalidPayload, _ = hex.DecodeString("00000101ca048008fff26423fffa5800000300140e596e979403183257cda0f047bf5755ffed1c853f637fbfbaf607d3fe5ed9073be977bf583a9db48a03a09badae0690f9c98c618844a549e0f6fada8b1cd574484f7f4c6aec7c8ef48c39d8ac8bad3767c9088a0395c8fe3d7ea3bd988194efd3a000ba79cbe3fb7b77dd16155116e30ea9d54367b56e3e2a12132556a3d0c285d6ce86f796231384bfbf6a88f81b284e2dd085f0786ffe443d7b9009a79d188f1d0000b616f0983b3484b31ba157a2f1ae4dde78dd1b12f79cc5df0d7cf215066901090e1599266b571843649e73d80fbaf3c700e7467a79563cc05c3501f0821954c762e4cdfdca87ea3a58")
		validPayload, _   = hex.DecodeString("00000008aabbccaabbccaabb00000001aa")
	)
	f.Add(invalidPayload)
	f.Add(validPayload)
	f.Fuzz(func(t *testing.T, payload []byte) {
		DecodeAVCC(payload)
	})
}

func TestDecodeAVCC(t *testing.T) {
	var err error
	var nalus [][]byte

	avccFrame, _ := hex.DecodeString(
		"00000008aabbccaabbccaabb00000001aa",
	)
	nalus, err = DecodeAVCC(avccFrame)
	if err != nil {
		t.Errorf("did not expect error")
	}

	if len(nalus) != 2 {
		t.Errorf("expected 2 NAL unit")
	}
	if nal := hex.EncodeToString(nalus[0]); nal != "aabbccaabbccaabb" {
		t.Errorf("Unexpected %v", nal)
	}
	if nal := hex.EncodeToString(nalus[1]); nal != "aa" {
		t.Errorf("Unexpected %v", nal)
	}
}

func TestEncodeAVCC(t *testing.T) {
	nalOne, _ := hex.DecodeString("aabbccaabbccaabb")
	nalTwo, _ := hex.DecodeString("aa")

	avccFrame := EncodeAVCC([][]byte{nalOne, nalTwo})

	if avccFrameHex := hex.EncodeToString(avccFrame); avccFrameHex != "00000008aabbccaabbccaabb00000001aa" {
		t.Errorf("Unexpected %v", avccFrameHex)
	}
}
