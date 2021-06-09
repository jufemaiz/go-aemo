// Package nmi delivers nmi capabilities to the aemo library.
//
// From: https://aemo.com.au/-/media/files/electricity/nem/retail_and_metering/metering-procedures/2016/0610-0008-pdf.pdf
//
// NMI Structure
//
// The NMI is a ten (10) character identifier assigned by Local Network Service
// Providers (LNSPs) in accordance with this procedure. The publication of this
// procedure and assignment of NMIs is authorised by the National Electricity Rules
// at clause 7.3.1(d), (da), and (db).
//
// The NMI may be used in conjunction with other identifiers or suffixes. These
// include:
//
// • The NMI checksum, a single numeral used to assist with data validation when the NMI is
// manually entered into a computer system.
//
// • The NMI data stream suffix used to identify a particular data stream associated with a
// connection point.
//
// Generally the NMI is an all numeric identifier; the limited circumstances under
// which alpha characters may be used are listed later in this procedure.
//
// The key attributes of the NMI are:
//
// • The NMI must embody only numeric characters, except as explicitly provided within this
// document, and must not contain spaces.
//
// • Character letters ‘O’ and ‘I’ are not permitted in order to avoid confusion with numbers 0
// and 1.
//
// • ‘W’ is a reserved character to be used as the fifth digit of the Allocated Identifier for
// wholesale transmission connection metering points only. It may only be used if the NMI is
// allocated from an alphanumeric block.
//
// • Embedded characters or meanings should not be used in allocating NMIs.
//
// • Where AEMO has allocated a block of NMIs to an LNSP, the LNSP must only use
// numeric characters in the NMIs allocated to the market unless AEMO has directed the
// block to be alphanumeric.
//
// • Where AEMO has allocated a block of NMIs to an LNSP, and directed the block to be
// alphanumeric, the LNSP may use all-numeric or alphanumeric characters in the NMIs
// allocated to the market.
//
// Network Service Providers must maintain a register of all NMIs released. AEMO
// maintains a Register of all ‘on-market’ NMIs within AEMO’s market systems.
//
// The base NMI is ten characters. In some circumstances the NMI checksum is
// appended to the NMI to form an eleven-character NMI, or the two character NMI
// data stream suffix may be appended to form a twelve-character NMI. NMI
// checksum is not used with the data stream suffix because the data stream suffix
// is intended for use only with electronic data transfer.
//
// In the initial allocation of alphanumeric NMIs the first character was a jurisdiction
// indicator. Jurisdiction indicators were abolished in October.
//
// AEMO may allocate blocks of NMIs to LNSPs from any unused range.
//
// The range 5 XXX XXX XXX has been reserved for use within the gas industry. To
// avoid the risk of confusion, AEMO has agreed not to issue NMIs commencing
// with 5.
//
// The range 9 XXX XXX XXX has been reserved as a “break-out” if it becomes
// necessary to move to an 11 character NMI.
package nmi
