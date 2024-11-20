# go-aemo changelog

## [v0.3.0] (2024-11-21)

### Added

*   [NEM12 Method flags](./lib/aemo/nem12/quality_method.rb) 22, 23, 24, 25 as per [Metrology Procedure: Part B v7.8 @ 2024-11-04](https://aemo.com.au/-/media/files/electricity/nem/retail_and_metering/market_settlement_and_transfer_solutions/2024/metrology-procedure-part-b-v781-clean.pdf?la=en)

### Changed

*   [NEM12 Method flags](./lib/aemo/nem12/quality_method.rb) 14, 15, 16, 20 as per [Metrology Procedure: Part B v7.8 @ 2024-11-04](https://aemo.com.au/-/media/files/electricity/nem/retail_and_metering/market_settlement_and_transfer_solutions/2024/metrology-procedure-part-b-v781-clean.pdf?la=en)


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
