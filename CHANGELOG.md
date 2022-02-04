# go-aemo changelog

## [v0.1.1] (2022-02-04)

### Fixed

*   `MDMDataStreamIdentifier` field updated due to variation by AEMO in their documenation.
    (Ref: <https://www.aemo.com.au/-/media/files/stakeholder_consultation/consultations/nem-consultations/2019/5ms-metering-package-3/final/standing-data-for-msats-final-clean.pdf>)
    (#4).

## [v0.1.0] (2021-12-17)

Initial release of the `go-aemo` package.

### Added

*   [`nem12`](./nem12/doc.go) package, for NEM12 file parsing.
*   [`nmi`](./nmi/doc.go) package, delivering National Meter Identifier information.
*   [`region`](./region/region.go) package, delivering AEMO regions.
