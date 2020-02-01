package request

type (
	RequestType int
	Destination int
)

type Request struct{
	Req RequestType
	Destination Destination
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
)

const (
	PANEL Destination=iota
	PLAYER
)

func NewRequestToPanel(req RequestType)Request{
	return Request{
		Req:req,
		Destination:PANEL,
	}
}

func NewRequestToPlayer(req RequestType)Request{
	return Request{
		Req:req,
		Destination:PLAYER,
	}
}