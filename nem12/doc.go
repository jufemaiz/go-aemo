// Package nem12 delivers the AEMO NEM12 data processing capability, with support
// for both the reading and the generating of NEM12 files. This support extends
// to being both canonical (that is, inclusive of the 100 and 900 rows), non-canonical
// (that is, exclusive of the 100 and 900 rows), and with relaxed validation.
//
// The NEM12 reader interface allows for the reading of entire meter data streams,
// daily chunks or individual meter reads at a time.
//
// The NEM12 writer interface delivers the ability to create either a canonical
// or a non-canonical NEM12 file from the provided data.
package nem12
