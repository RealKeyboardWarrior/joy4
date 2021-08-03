package fmp4

import (
	"time"

	"github.com/RealKeyboardWarrior/joy4/av"
	"github.com/RealKeyboardWarrior/joy4/format/mp4/mp4io"
)

type Stream struct {
	av.CodecData

	trackAtom *mp4io.Track
	Idx       int

	// pkts to be used in MDAT and MOOF > TRAF > TRUN
	lastpkt    *av.Packet
	pkts       []*av.Packet
	hasBFrames bool

	timeScale int64
	duration  int64

	muxer *Muxer
	//demuxer *Demuxer

	sample *mp4io.SampleTable
	dts    int64
}

func timeToTs(tm time.Duration, timeScale int64) int64 {
	return int64(tm * time.Duration(timeScale) / time.Second)
}

func tsToTime(ts int64, timeScale int64) time.Duration {
	return time.Duration(ts) * time.Second / time.Duration(timeScale)
}

func (self *Stream) timeToTs(tm time.Duration) int64 {
	return int64(tm * time.Duration(self.timeScale) / time.Second)
}

func (self *Stream) tsToTime(ts int64) time.Duration {
	return time.Duration(ts) * time.Second / time.Duration(self.timeScale)
}
