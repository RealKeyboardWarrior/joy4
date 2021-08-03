package fmp4

import (
	"bytes"
	"io"
	"log"

	"github.com/RealKeyboardWarrior/joy4/av"
	"github.com/RealKeyboardWarrior/joy4/av/avutil"
	"github.com/RealKeyboardWarrior/joy4/format/mp4/mp4io"
)

var CodecTypes = []av.CodecType{av.H264, av.AAC}

func Handler(h *avutil.RegisterHandler) {
	h.Ext = ".fmp4"

	h.Probe = func(b []byte) bool {
		probe_reader := bytes.NewReader(b)
		var atoms []mp4io.Atom
		var err error
		if atoms, err = mp4io.ReadFileAtoms(probe_reader); err != nil {
			log.Printf("fmp4: Probe(): errored on read file atoms when probing")
			return false
		}

		for _, atom := range atoms {
			if atom.Tag() == mp4io.MOOF {
				return true
			}
		}

		return false
	}

	/*
		h.ReaderDemuxer = func(r io.Reader) av.Demuxer {
			return NewDemuxer(r.(io.ReadSeeker))
		}
	*/
	h.WriterMuxer = func(w io.Writer) av.Muxer {
		return NewMuxer(w.(io.WriteSeeker))
	}

	h.CodecTypes = CodecTypes
}
