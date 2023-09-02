package h264parser

import (
	"fmt"

	"github.com/RealKeyboardWarrior/joy4/utils/bits/pio"
)

var ErrInvalidAnnexB = fmt.Errorf("bytestream b was not AnnexB")

func DecodeAnnexB(b []byte) ([][]byte, error) {
	if len(b) < 4 {
		return nil, ErrInvalidAnnexB
	}

	val3 := pio.U24BE(b)
	val4 := pio.U32BE(b)

	nalus := [][]byte{}

	if val3 == 1 || val4 == 1 {
		_val3 := val3
		_val4 := val4
		start := 0
		pos := 0
		for {
			if start != pos {
				nalus = append(nalus, b[start:pos])
			}
			if _val3 == 1 {
				pos += 3
			} else if _val4 == 1 {
				pos += 4
			}
			start = pos
			if start == len(b) {
				break
			}
			_val3 = 0
			_val4 = 0
			for pos < len(b) {
				if pos+2 < len(b) && b[pos] == 0 {
					_val3 = pio.U24BE(b[pos:])
					if _val3 == 0 {
						if pos+3 < len(b) {
							_val4 = uint32(b[pos+3])
							if _val4 == 1 {
								break
							}
						}
					} else if _val3 == 1 {
						break
					}
					pos++
				} else {
					pos++
				}
			}
		}
		return nalus, nil
	}

	return nil, ErrInvalidAnnexB
}
