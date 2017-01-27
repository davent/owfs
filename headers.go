package owfs

const (
	TYPE_ERROR       uint32 = 0
	TYPE_NOOP        uint32 = 1
	TYPE_READ        uint32 = 2
	TYPE_WRITE       uint32 = 3
	TYPE_DIR         uint32 = 4
	TYPE_SIZE        uint32 = 5
	TYPE_PRESENT     uint32 = 6
	TYPE_DIRALL      uint32 = 7
	TYPE_GET         uint32 = 8
	TYPE_DIRALLSLASH uint32 = 9
	TYPE_GETSLASH    uint32 = 10
)

type RequestHeader struct {
	Version       uint32
	PayloadLength uint32
	Type          uint32
	ControlFlags  uint32
	Size          uint32
	Offset        uint32
}

type ResponseHeader struct {
	Version       uint32
	PayloadLength uint32
	Ret           uint32
	ControlFlags  uint32
	Size          uint32
	Offset        uint32
}
