# go-aemo changelog

## [v0.2.0] (2022-09-20)

### Added

*   New method flags:

    | Type | `EST` and/or `SUB` | Descriptor |
    | --- | --- | --- |
    | Type 20 | `SUB` | Churn Correction (Like Day) |
    | Type 21 | `SUB` | Five-minute No Historical Data |
    | Type 59 | `SUB`, `EST` | Five-minute No Historical Data |
    | Type 69 | `SUB` | Linear extrapolation |

## [v0.1.1] (2022-02-04)

### Fixed

*   `MDMDataStreamIdentifier` field updated due to variation by AEMO in their documenation
    (Ref: <https://www.aemo.com.au/-/media/files/stakeholder_consultation/consultations/nem-consultations/2019/5ms-metering-package-3/final/standing-data-for-msats-final-clean.pdf>)
    (#4).

### Changed

*   CI moved from Travis to Github Actions (#6)

## [v0.1.0] (2021-12-17)

Initial release of the `go-aemo` package.

### Added

*   [`nem12`](./nem12/doc.go) package, for NEM12 file parsing.
*   [`nmi`](./nmi/doc.go) package, delivering National Meter Identifier information.
*   [`region`](./region/region.go) package, delivering AEMO regions.
