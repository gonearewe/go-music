// Requests are sent through channels for communication.
// It works like a mail system, every routine shares a channel
// which is regarded by themselves as a outport and they send
// out their requests for routines outside through it.
// And every routine holds a channel from where they receive outside
// requests. Thus a mail routine(usually main routine) is required
// to listen to the shared channel and dispatch requests there to
// their destination.
package request

type (
	RequestType int
	Destination int
)

type Request struct {
	Req         RequestType
	Destination Destination
	Attachments interface{}
}

const (
	// request sent to player
	RequestRandomMode RequestType = iota
	RequestRepeatMode
	RequestSequentialMode
	RequestNextTrack
	RequestPrevTrack
	RequestStop

	// request sent to panel
	RequestShowLOGO
	RequestClearAndStop

	// request sent to main routine
	RequestSetNewLibrary
)

const (
	PANEL Destination = iota
	PLAYER
	MAIN // main routine
)

func NewRequestToPanel(req RequestType) Request {
	return Request{
		Req:         req,
		Destination: PANEL,
	}
}

func NewRequestToPlayer(req RequestType) Request {
	return Request{
		Req:         req,
		Destination: PLAYER,
	}
}

func NewRequestToMain(req RequestType, attachments interface{}) Request {
	return Request{
		Req:         req,
		Destination: MAIN,
		Attachments: attachments,
	}
}
