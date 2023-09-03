package h264parser

import (
	"fmt"

	"github.com/RealKeyboardWarrior/joy4/utils/bits/pio"
)

var ErrInvalidAVCC = fmt.Errorf("bytestream b was not AVCC")

func DecodeAVCC(b []byte) ([][]byte, error) {
	if len(b) < 4 {
		return nil, ErrInvalidAVCC
	}

	// TODO: AVCC can be 1, 2, 3 or 4 bytes length encoding
	// At the moment we just force expect it to be 4
	val4 := pio.U32BE(b)
	if val4 > uint32(len(b)-4) {
		return nil, ErrInvalidAVCC
	}

	_val4 := val4
	_b := b[4:]
	nalus := [][]byte{}
	for {
		nalus = append(nalus, _b[:_val4])
		if _val4 > uint32(len(_b)) {
			break
		}
		_b = _b[_val4:]
		if len(_b) < 4 {
			break
		}
		_val4 = pio.U32BE(_b)
		_b = _b[4:]
		if _val4 > uint32(len(_b)) {
			break
		}
	}
	if len(_b) == 0 {
		return nalus, nil
	}

	return nil, ErrInvalidAVCC
}

func EncodeAVCC(nalus [][]byte) []byte {
	b := make([]byte, 0)

	for _, nalu := range nalus {
		// Always encode AVCC with 4 bytes length
		size := make([]byte, 4)
		pio.PutU32BE(size, uint32(len(nalu)))
		b = append(b, size...)
		b = append(b, nalu...)
	}

	return b
}
