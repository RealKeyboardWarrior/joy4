package main

import (
	"github.com/kerberos-io/joy4/av/avutil"
	"github.com/kerberos-io/joy4/format"
)

func init() {
	format.RegisterAll()
}
func main() {
	muxer, err := avutil.Create("sintel_fragmentedz.fmp4")
	if err != nil {
		panic(err)
	}
	demuxer, _ := avutil.Open("sintel.mp4")
	avutil.CopyFile(muxer, demuxer)

	muxer.Close()
	demuxer.Close()
}
