# go-aemo

[![Go Reference](https://pkg.go.dev/badge/github.com/jufemaiz/go-aemo.svg)](https://pkg.go.dev/github.com/jufemaiz/go-aemo) [![codecov](https://codecov.io/gh/jufemaiz/go-aemo/branch/main/graph/badge.svg?token=ztfbFyBr3c)](https://codecov.io/gh/jufemaiz/go-aemo) [![Quality gate](https://sonarcloud.io/api/project_badges/quality_gate?project=jufemaiz_go-aemo)](https://sonarcloud.io/summary/new_code?id=jufemaiz_go-aemo)

A golang library for interfacing with AEMO data. For the most part this is inspired
by my previous Ruby based library [aemo](https://github.com/jufemaiz/aemo). NEM12
parsing has, however, been completely overhauled with nods to a more idiomatic go
approach.

## Packages

### `nem12`

NEM12 is the standard file format for electricity metering data in Australia's National
Electricity Market. The reference document is the
[Meter Data File Format Specification NEM12 & NEM13 v2.4](https://aemo.com.au/-/media/files/electricity/nem/retail_and_metering/metering-procedures/2021/mdff-specification-nem12-nem13-v24.pdf?la=en)
released 2021-10-01.

### `nmi`

The `nmi` package provides validation and information on the National Meter Identifiers
in Australia.

### `region`

The `region` package provides a valid set of regions that operate within the National
Electricity Market.
