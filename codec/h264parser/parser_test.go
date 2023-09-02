package h264parser

import (
	"encoding/hex"
	"testing"

	"github.com/RealKeyboardWarrior/joy4/av"
)

var h264NALunit = "0000010605145a4f4f4d0780043802800170000003000003000003000a0000030000030280016880000000016764001eac1b1a78280bde59a0100000000168c9a3c30000000168526cf0c00000000168726cf2c00000000168269b3c3000000001682e9b3cb00000000168369a3c30"

func TestSplitNALUs(t *testing.T) {
	var typ int
	var nalus [][]byte

	annexbFrame, _ := hex.DecodeString("00000001223322330000000122332233223300000133000001000001")
	nalus, typ = SplitNALUs(annexbFrame)
	if typ != NALU_ANNEXB {
		t.Errorf("expected type NALU_ANNEXB")
	}
	if len(nalus) != 3 {
		t.Errorf("expected 3 NAL units")
	}

	if nal := hex.EncodeToString(nalus[0]); nal != "22332233" {
		t.Errorf("Expected %v", nal)
	}
	if nal := hex.EncodeToString(nalus[1]); nal != "223322332233" {
		t.Errorf("Expected %v", nal)
	}
	if nal := hex.EncodeToString(nalus[2]); nal != "33" {
		t.Errorf("Expected %v", nal)
	}

	avccFrame, _ := hex.DecodeString(
		"00000008aabbccaabbccaabb00000001aa",
	)
	nalus, typ = SplitNALUs(avccFrame)
	if typ != NALU_AVCC {
		t.Errorf("expected type NALU_AVCC")
	}

	if len(nalus) != 2 {
		t.Errorf("expected 2 NAL unit")
	}
	if nal := hex.EncodeToString(nalus[0]); nal != "aabbccaabbccaabb" {
		t.Errorf("Expected %v", nal)
	}
	if nal := hex.EncodeToString(nalus[1]); nal != "aa" {
		t.Errorf("Expected %v", nal)
	}
}

func TestSplitNALUsCrash(t *testing.T) {
	var typ int
	var nalus [][]byte

	annexbFrame, _ := hex.DecodeString("00000101ca048008fff26423fffa5800000300140e596e979403183257cda0f047bf5755ffed1c853f637fbfbaf607d3fe5ed9073be977bf583a9db48a03a09badae0690f9c98c618844a549e0f6fada8b1cd574484f7f4c6aec7c8ef48c39d8ac8bad3767c9088a0395c8fe3d7ea3bd988194efd3a000ba79cbe3fb7b77dd16155116e30ea9d54367b56e3e2a12132556a3d0c285d6ce86f796231384bfbf6a88f81b284e2dd085f0786ffe443d7b9009a79d188f1d0000b616f0983b3484b31ba157a2f1ae4dde78dd1b12f79cc5df0d7cf215066901090e1599266b571843649e73d80fbaf3c700e7467a79563cc05c3501f0821954c762e4cdfdca87ea3a58")
	nalus, typ = SplitNALUs(annexbFrame)
	if typ != NALU_ANNEXB {
		t.Errorf("expected type NALU_ANNEXB")
	}
	if len(nalus) != 1 {
		t.Errorf("expected 1 NAL units")
	}

	if nal := hex.EncodeToString(nalus[0]); nal != "01ca048008fff26423fffa5800000300140e596e979403183257cda0f047bf5755ffed1c853f637fbfbaf607d3fe5ed9073be977bf583a9db48a03a09badae0690f9c98c618844a549e0f6fada8b1cd574484f7f4c6aec7c8ef48c39d8ac8bad3767c9088a0395c8fe3d7ea3bd988194efd3a000ba79cbe3fb7b77dd16155116e30ea9d54367b56e3e2a12132556a3d0c285d6ce86f796231384bfbf6a88f81b284e2dd085f0786ffe443d7b9009a79d188f1d0000b616f0983b3484b31ba157a2f1ae4dde78dd1b12f79cc5df0d7cf215066901090e1599266b571843649e73d80fbaf3c700e7467a79563cc05c3501f0821954c762e4cdfdca87ea3a58" {
		t.Errorf("Unexpected %v", nal)
	}
}

func TestPktToCodecData(t *testing.T) {
	nalUnitBytes, _ := hex.DecodeString(h264NALunit)
	pkt := av.Packet{
		IsKeyFrame: true,
		Data:       nalUnitBytes,
	}
	decoded, err := PktToCodecData(pkt)
	if err != nil {
		t.Error(err)
	}

	sps := decoded.(CodecData).SPS()
	pps := decoded.(CodecData).PPS()

	if spsHex := hex.EncodeToString(sps); spsHex != "6764001eac1b1a78280bde59a010" {
		t.Errorf("unexpected sps = %v", spsHex)
	}

	if ppsHex := hex.EncodeToString(pps); ppsHex != "68369a3c30" {
		t.Errorf("unexpected pps = %v", ppsHex)
	}
}

func TestParseSliceHeaderFromNALU(t *testing.T) {
	nalUnitBytes, _ := hex.DecodeString(h264NALunit)

	typ, err := ParseSliceHeaderFromNALU(nalUnitBytes[4:])
	if err != nil {
		t.Error(err)
	}
	if typ != SLICE_I {
		t.Errorf("didn't find I keyframe %v", typ)
	}
}
