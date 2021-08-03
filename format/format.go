package format

import (
	"github.com/kerberos-io/joy4/av/avutil"
	"github.com/kerberos-io/joy4/format/aac"
	"github.com/kerberos-io/joy4/format/flv"
	"github.com/kerberos-io/joy4/format/fmp4"
	"github.com/kerberos-io/joy4/format/mp4"
	"github.com/kerberos-io/joy4/format/rtmp"
	"github.com/kerberos-io/joy4/format/rtsp"
	"github.com/kerberos-io/joy4/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(fmp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
