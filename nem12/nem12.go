package nem12

// Nem12Reader provides the ability to read nem12 files.
type Nem12Reader interface {
	// ReadDataStreams(ds []DataStream) (n int, err error)
	// ReadDays(i int) (n int, err error)
	// ReadData(i int) (n int, err error)

	// StreamDataStreams() (<-chan DataStream, <-chan error)
	// StreamDays() (<-chan DataStream, <-chan error)
	// StreamData() (<-chan DataStream, <-chan error)
}

// Nem12Writer provides the ability to write nem12 files.
type Nem12Writer interface{}

// Nem12ReaderWriter delivers reading and writing of NEM12 files.
type Nem12ReaderWriter interface {
	Nem12Reader
	Nem12Writer
}
