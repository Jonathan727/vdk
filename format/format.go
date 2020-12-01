package format

import (
	"github.com/jonathan727/vdk/av/avutil"
	"github.com/jonathan727/vdk/format/aac"
	"github.com/jonathan727/vdk/format/flv"
	"github.com/jonathan727/vdk/format/mp4"
	"github.com/jonathan727/vdk/format/rtmp"
	"github.com/jonathan727/vdk/format/rtsp"
	"github.com/jonathan727/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
