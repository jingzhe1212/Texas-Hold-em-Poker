package room

type IOccupant interface {
	GetRoom() IRoom
	WriteMsg(msg interface{})
}
type IHandler interface {
	NewRoom() IRoom
	NoRoomHandler(m interface{}) IRoom
}
type IRoom interface {
	Cap() uint8
	Len() uint8
	GetNumber() string
	SetNumber(string)
	Data() interface{}
	SetData(interface{})
	Close()
	Closed() chan struct{}
	Send(IOccupant, interface{}) error
	WriteMsg(interface{}, ...uint32)
	Regist(interface{}, interface{})
}