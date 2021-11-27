package nem12_test

import (
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/jufemaiz/go-aemo/nem12"
)

func TestNewParser(t *testing.T) {
	Convey("nem12.NewParser", t, func() {
		tests := map[string]struct {
			reader     func() io.Reader
			assertions func(r io.Reader, p nem12.Parser)
		}{
			"nil reader": {
				reader: func() io.Reader {
					return nil
				},
				assertions: func(r io.Reader, p nem12.Parser) {
					So(r, ShouldBeNil)

					d, err := p.ReadDay()
					So(d, ShouldBeNil)
					So(err, ShouldBeError)
					So(err, ShouldWrap, nem12.ErrReaderNil)
				},
			},
			"invalid nem12 file": {
				reader: func() io.Reader {
					filepath := "./testdata/invalid/NEM12#000000000000021#CNRGYMDP#NEMMCO"

					r, err := os.Open(filepath)
					So(err, ShouldBeNil)

					return r
				},
				assertions: func(r io.Reader, p nem12.Parser) {
					So(r, ShouldNotBeNil)

					d, err := p.ReadDay()
					So(d, ShouldBeNil)
					So(err, ShouldBeError)
					So(err, ShouldWrap, nem12.ErrParseFailed)
				},
			},
			"valid nem12 file with all actual data for two meters at one nmi": {
				reader: func() io.Reader {
					filepath := "./testdata/valid/NEM12#mdffl0000000001#ACTEWM#NEMMCO.mdff"
					// 100,NEM12,200505171229,ACTEWM,NEMMCO
					// 200,NEM1201001,E1E2,1,E1,N1,01001,kWh,15,
					// 300,20041102,2.569,2.442,2.379,2.347,2.379,2.284,2.252,2.315,2.220,2.220,2.252,2.188,2.188,2.252,2.220,2.220,2.315,2.284,2.315,2.410,2.442,2.506,2.664,2.696,2.918,3.076,3.172,3.267,3.489,3.552,3.679,3.806,3.806,3.838,3.933,3.901,3.964,4.060,3.996,4.028,4.091,4.060,4.091,4.155,4.060,4.060,4.123,4.028,4.060,4.091,3.996,3.996,4.028,3.933,3.933,3.964,3.901,3.869,3.933,3.838,3.838,3.869,3.774,3.774,3.838,3.742,3.742,3.806,3.679,3.679,3.774,3.711,3.647,3.679,3.616,3.584,3.584,3.489,3.457,3.457,3.330,3.298,3.298,3.203,3.108,3.140,3.013,2.950,3.013,3.013,2.981,3.013,2.886,2.791,2.759,2.601,A,,,20050517122948,
					// 300,20041103,2.569,2.442,2.379,2.347,2.379,2.284,2.252,2.315,2.220,2.220,2.252,2.188,2.188,2.252,2.220,2.220,2.315,2.284,2.315,2.410,2.442,2.506,2.664,2.696,2.918,3.076,3.172,3.267,3.489,3.552,3.679,3.806,3.806,3.838,3.933,3.901,3.964,4.060,3.996,4.028,4.091,4.060,4.091,4.155,4.060,4.060,4.123,4.028,4.060,4.091,3.996,3.996,4.028,3.933,3.933,3.964,3.901,3.869,3.933,3.838,3.838,3.869,3.774,3.774,3.838,3.742,3.742,3.806,3.679,3.679,3.774,3.711,3.647,3.679,3.616,3.584,3.584,3.489,3.457,3.457,3.330,3.298,3.298,3.203,3.108,3.140,3.013,2.950,3.013,3.013,2.981,3.013,2.886,2.791,2.759,2.601,A,,,20050517122948,
					// 300,20041104,2.569,2.442,2.379,2.347,2.379,2.284,2.252,2.315,2.220,2.220,2.252,2.188,2.188,2.252,2.220,2.220,2.315,2.284,2.315,2.410,2.442,2.506,2.664,2.696,2.918,3.076,3.172,3.267,3.489,3.552,3.679,3.806,3.806,3.838,3.933,3.901,3.964,4.060,3.996,4.028,4.091,4.060,4.091,4.155,4.060,4.060,4.123,4.028,4.060,4.091,3.996,3.996,4.028,3.933,3.933,3.964,3.901,3.869,3.933,3.838,3.838,3.869,3.774,3.774,3.838,3.742,3.742,3.806,3.679,3.679,3.774,3.711,3.647,3.679,3.616,3.584,3.584,3.489,3.457,3.457,3.330,3.298,3.298,3.203,3.108,3.140,3.013,2.950,3.013,3.013,2.981,3.013,2.886,2.791,2.759,2.601,A,,,20050517122948,
					// 300,20041105,2.569,2.442,2.379,2.347,2.379,2.284,2.252,2.315,2.220,2.220,2.252,2.188,2.188,2.252,2.220,2.220,2.315,2.284,2.315,2.410,2.442,2.506,2.664,2.696,2.918,3.076,3.172,3.267,3.489,3.552,3.679,3.806,3.806,3.838,3.933,3.901,3.964,4.060,3.996,4.028,4.091,4.060,4.091,4.155,4.060,4.060,4.123,4.028,4.060,4.091,3.996,3.996,4.028,3.933,3.933,3.964,3.901,3.869,3.933,3.838,3.838,3.869,3.774,3.774,3.838,3.742,3.742,3.806,3.679,3.679,3.774,3.711,3.647,3.679,3.616,3.584,3.584,3.489,3.457,3.457,3.330,3.298,3.298,3.203,3.108,3.140,3.013,2.950,3.013,3.013,2.981,3.013,2.886,2.791,2.759,2.601,A,,,20050517122948,
					// 200,NEM1201001,E1E2,1,E2,N2,01001,kWh,15,
					// 300,20041102,33.600,33.600,32.400,33.000,34.200,33.000,32.400,31.800,31.800,32.400,33.600,31.200,30.600,31.200,30.600,30.600,30.000,31.200,30.600,28.200,28.200,27.600,27.600,26.400,27.600,28.200,28.200,27.600,28.200,28.200,27.600,28.800,28.800,27.600,28.800,34.800,34.200,36.000,36.000,34.800,34.800,36.000,34.200,35.400,36.000,36.600,36.600,34.800,36.000,35.400,35.400,39.600,41.400,42.000,43.200,41.400,42.000,43.200,42.600,39.000,34.800,36.000,33.600,34.800,34.200,32.400,31.800,32.400,32.400,33.600,33.000,33.000,31.800,34.200,30.600,30.600,30.600,28.800,27.600,28.800,29.400,29.400,29.400,30.600,29.400,28.800,29.400,30.000,28.800,28.800,29.400,28.800,29.400,29.400,29.400,28.800,A,,,20050517122949,
					// 300,20041103,28.800,28.800,29.400,28.800,29.400,28.200,28.200,28.200,29.400,29.400,29.400,28.800,30.000,31.800,30.600,30.000,30.600,30.000,43.800,43.800,61.200,72.000,73.200,77.400,94.200,97.200,108.000,118.800,125.400,129.600,147.000,146.400,155.400,164.400,163.800,159.000,159.000,160.800,154.800,157.200,158.400,160.800,155.400,153.600,148.800,143.400,125.400,136.200,136.800,129.600,130.800,135.600,139.800,129.000,124.200,127.800,124.200,121.200,119.400,112.800,96.600,77.400,73.200,70.800,52.200,48.000,48.000,46.800,35.400,33.600,32.400,33.000,33.600,33.000,31.800,34.200,34.200,34.200,33.000,33.600,33.600,33.600,33.600,33.600,32.400,31.800,31.200,31.800,31.200,32.400,31.200,30.000,30.600,30.600,31.200,31.200,A,,,20050517122949,
					// 300,20041104,31.800,32.400,31.200,31.800,31.200,31.800,31.800,31.200,31.200,32.400,30.600,31.800,31.800,31.200,32.400,32.400,33.600,33.600,46.800,46.800,57.000,65.400,67.800,75.000,95.400,97.800,108.000,118.800,126.600,141.600,144.000,157.800,153.600,169.200,165.600,156.600,151.800,161.400,160.200,156.000,156.000,155.400,161.400,154.200,147.600,142.800,125.400,123.600,129.000,133.800,131.400,124.200,124.200,122.400,111.000,96.600,80.400,77.400,76.800,84.600,81.600,80.400,72.000,70.200,59.400,60.000,63.600,57.000,45.000,42.600,43.200,39.000,37.800,37.800,37.200,36.600,37.800,37.800,36.000,36.000,36.000,36.600,36.000,36.000,33.000,36.000,33.600,34.200,34.800,34.800,34.200,34.200,34.200,33.600,34.800,33.600,A,,,20050517122949,
					// 300,20041105,33.600,33.600,33.000,33.000,33.000,34.200,34.800,36.000,35.400,34.800,34.200,36.000,35.400,36.000,36.000,36.000,37.200,36.000,34.800,33.600,59.400,68.400,67.800,73.200,97.200,98.400,114.600,122.400,133.200,135.600,150.600,157.800,158.400,160.200,155.400,152.400,152.400,153.600,158.400,162.000,158.400,161.400,167.400,158.400,150.000,142.800,124.800,125.400,126.600,133.200,129.000,132.600,144.000,141.000,142.800,132.600,137.400,131.400,124.200,111.600,103.800,99.600,93.000,89.400,78.000,84.000,76.200,70.200,61.800,60.000,61.800,58.800,57.000,60.600,57.600,56.400,58.800,59.400,57.600,51.000,43.800,42.000,46.200,45.000,37.200,36.000,34.800,33.600,36.000,35.400,33.000,33.000,33.000,33.000,30.600,30.000,A,,,20050517122949,
					// 900

					r, err := os.Open(filepath)
					So(err, ShouldBeNil)

					return r
				},
				assertions: func(r io.Reader, p nem12.Parser) {
					So(r, ShouldNotBeNil)

					d, err := p.ReadDay()
					So(d, ShouldNotBeNil)
					So(err, ShouldBeNil)

					So(d.Metadata, ShouldNotBeNil)
					So(d.Metadata.Nmi, ShouldNotBeNil)
					So(d.Metadata.Nmi.Identifier, ShouldEqual, "NEM1201001")
					So(d.Metadata.Meter, ShouldNotBeNil)
					So(d.Metadata.Meter.Identifier, ShouldEqual, "1")
					So(d.Metadata.Suffix, ShouldNotBeNil)
					So(d.Metadata.Suffix.Type, ShouldEqual, nem12.SuffixExportWattHourMaster)
					So(d.Metadata.UnitOfMeasure, ShouldNotBeNil)
					So(*d.Metadata.UnitOfMeasure, ShouldEqual, nem12.UnitKilowattHour)

					So(d.Data, ShouldHaveLength, 96)
					ts := time.Date(2004, 11, 2, 0, 0, 0, 0, nem12.NEMTime())
					updatedAt := time.Date(2005, 5, 17, 12, 29, 48, 0, nem12.NEMTime())
					vals := []float64{
						2.569, 2.442, 2.379, 2.347, 2.379, 2.284, 2.252, 2.315, 2.220, 2.220, 2.252, 2.188, 2.188, 2.252, 2.220, 2.220, 2.315, 2.284,
						2.315, 2.410, 2.442, 2.506, 2.664, 2.696, 2.918, 3.076, 3.172, 3.267, 3.489, 3.552, 3.679, 3.806, 3.806, 3.838, 3.933, 3.901,
						3.964, 4.060, 3.996, 4.028, 4.091, 4.060, 4.091, 4.155, 4.060, 4.060, 4.123, 4.028, 4.060, 4.091, 3.996, 3.996, 4.028, 3.933,
						3.933, 3.964, 3.901, 3.869, 3.933, 3.838, 3.838, 3.869, 3.774, 3.774, 3.838, 3.742, 3.742, 3.806, 3.679, 3.679, 3.774, 3.711,
						3.647, 3.679, 3.616, 3.584, 3.584, 3.489, 3.457, 3.457, 3.330, 3.298, 3.298, 3.203, 3.108, 3.140, 3.013, 2.950, 3.013, 3.013,
						2.981, 3.013, 2.886, 2.791, 2.759, 2.601,
					}
					for i, val := range vals {
						So(d.Data[i], ShouldResemble, &nem12.Interval{
							Time:           ts.Add(time.Duration((i+1)*15) * time.Minute),
							IntervalLength: (15 * time.Minute),
							Value: nem12.IntervalValue{
								Value:             val,
								DecimalValue:      decimal.NewFromFloat(val),
								QualityFlag:       nem12.QualityActual,
								MethodFlag:        nil,
								ReasonCode:        nil,
								ReasonDescription: nil,
								UpdateDateTime:    &updatedAt,
								MSATSLoadDateTime: nil,
							},
						})
					}

					d, err = p.ReadDay()
					So(d, ShouldNotBeNil)
					So(err, ShouldBeNil)

					So(d.Metadata, ShouldNotBeNil)
					So(d.Metadata.Nmi, ShouldNotBeNil)
					So(d.Metadata.Nmi.Identifier, ShouldEqual, "NEM1201001")
					So(d.Metadata.Meter, ShouldNotBeNil)
					So(d.Metadata.Meter.Identifier, ShouldEqual, "1")
					So(d.Metadata.Suffix, ShouldNotBeNil)
					So(d.Metadata.Suffix.Type, ShouldEqual, nem12.SuffixExportWattHourMaster)
					So(d.Metadata.UnitOfMeasure, ShouldNotBeNil)
					So(*d.Metadata.UnitOfMeasure, ShouldEqual, nem12.UnitKilowattHour)
					So(d.Data, ShouldHaveLength, 96)

					d, err = p.ReadDay()
					So(d, ShouldNotBeNil)
					So(err, ShouldBeNil)

					So(d.Metadata, ShouldNotBeNil)
					So(d.Metadata.Nmi, ShouldNotBeNil)
					So(d.Metadata.Nmi.Identifier, ShouldEqual, "NEM1201001")
					So(d.Metadata.Meter, ShouldNotBeNil)
					So(d.Metadata.Meter.Identifier, ShouldEqual, "1")
					So(d.Metadata.Suffix, ShouldNotBeNil)
					So(d.Metadata.Suffix.Type, ShouldEqual, nem12.SuffixExportWattHourMaster)
					So(d.Metadata.UnitOfMeasure, ShouldNotBeNil)
					So(*d.Metadata.UnitOfMeasure, ShouldEqual, nem12.UnitKilowattHour)
					So(d.Data, ShouldHaveLength, 96)

					d, err = p.ReadDay()
					So(d, ShouldNotBeNil)
					So(err, ShouldBeNil)

					So(d.Metadata, ShouldNotBeNil)
					So(d.Metadata.Nmi, ShouldNotBeNil)
					So(d.Metadata.Nmi.Identifier, ShouldEqual, "NEM1201001")
					So(d.Metadata.Meter, ShouldNotBeNil)
					So(d.Metadata.Meter.Identifier, ShouldEqual, "1")
					So(d.Metadata.Suffix, ShouldNotBeNil)
					So(d.Metadata.Suffix.Type, ShouldEqual, nem12.SuffixExportWattHourMaster)
					So(d.Metadata.UnitOfMeasure, ShouldNotBeNil)
					So(*d.Metadata.UnitOfMeasure, ShouldEqual, nem12.UnitKilowattHour)
					So(d.Data, ShouldHaveLength, 96)

					d, err = p.ReadDay()
					So(d, ShouldNotBeNil)
					So(err, ShouldBeNil)

					So(d.Metadata, ShouldNotBeNil)
					So(d.Metadata.Nmi, ShouldNotBeNil)
					So(d.Metadata.Nmi.Identifier, ShouldEqual, "NEM1201001")
					So(d.Metadata.Meter, ShouldNotBeNil)
					So(d.Metadata.Meter.Identifier, ShouldEqual, "2")
					So(d.Metadata.Suffix, ShouldNotBeNil)
					So(d.Metadata.Suffix.Type, ShouldEqual, nem12.SuffixExportWattHourMaster)
					So(d.Metadata.UnitOfMeasure, ShouldNotBeNil)
					So(*d.Metadata.UnitOfMeasure, ShouldEqual, nem12.UnitKilowattHour)

					So(d.Data, ShouldHaveLength, 96)
					ts = time.Date(2004, 11, 2, 0, 0, 0, 0, nem12.NEMTime())
					updatedAt = time.Date(2005, 5, 17, 12, 29, 49, 0, nem12.NEMTime())
					vals = []float64{
						33.600, 33.600, 32.400, 33.000, 34.200, 33.000, 32.400, 31.800, 31.800, 32.400, 33.600, 31.200, 30.600, 31.200,
						30.600, 30.600, 30.000, 31.200, 30.600, 28.200, 28.200, 27.600, 27.600, 26.400, 27.600, 28.200, 28.200, 27.600,
						28.200, 28.200, 27.600, 28.800, 28.800, 27.600, 28.800, 34.800, 34.200, 36.000, 36.000, 34.800, 34.800, 36.000,
						34.200, 35.400, 36.000, 36.600, 36.600, 34.800, 36.000, 35.400, 35.400, 39.600, 41.400, 42.000, 43.200, 41.400,
						42.000, 43.200, 42.600, 39.000, 34.800, 36.000, 33.600, 34.800, 34.200, 32.400, 31.800, 32.400, 32.400, 33.600,
						33.000, 33.000, 31.800, 34.200, 30.600, 30.600, 30.600, 28.800, 27.600, 28.800, 29.400, 29.400, 29.400, 30.600,
						29.400, 28.800, 29.400, 30.000, 28.800, 28.800, 29.400, 28.800, 29.400, 29.400, 29.400, 28.800,
					}
					for i, val := range vals {
						So(d.Data[i], ShouldResemble, &nem12.Interval{
							Time:           ts.Add(time.Duration((i+1)*15) * time.Minute),
							IntervalLength: (15 * time.Minute),
							Value: nem12.IntervalValue{
								Value:             val,
								DecimalValue:      decimal.NewFromFloat(val),
								QualityFlag:       nem12.QualityActual,
								MethodFlag:        nil,
								ReasonCode:        nil,
								ReasonDescription: nil,
								UpdateDateTime:    &updatedAt,
								MSATSLoadDateTime: nil,
							},
						})
					}

					d, err = p.ReadDay()
					So(d, ShouldNotBeNil)
					So(err, ShouldBeNil)

					So(d.Metadata, ShouldNotBeNil)
					So(d.Metadata.Nmi, ShouldNotBeNil)
					So(d.Metadata.Nmi.Identifier, ShouldEqual, "NEM1201001")
					So(d.Metadata.Meter, ShouldNotBeNil)
					So(d.Metadata.Meter.Identifier, ShouldEqual, "2")
					So(d.Metadata.Suffix, ShouldNotBeNil)
					So(d.Metadata.Suffix.Type, ShouldEqual, nem12.SuffixExportWattHourMaster)
					So(d.Metadata.UnitOfMeasure, ShouldNotBeNil)
					So(*d.Metadata.UnitOfMeasure, ShouldEqual, nem12.UnitKilowattHour)
					So(d.Data, ShouldHaveLength, 96)

					d, err = p.ReadDay()
					So(d, ShouldNotBeNil)
					So(err, ShouldBeNil)

					So(d.Metadata, ShouldNotBeNil)
					So(d.Metadata.Nmi, ShouldNotBeNil)
					So(d.Metadata.Nmi.Identifier, ShouldEqual, "NEM1201001")
					So(d.Metadata.Meter, ShouldNotBeNil)
					So(d.Metadata.Meter.Identifier, ShouldEqual, "2")
					So(d.Metadata.Suffix, ShouldNotBeNil)
					So(d.Metadata.Suffix.Type, ShouldEqual, nem12.SuffixExportWattHourMaster)
					So(d.Metadata.UnitOfMeasure, ShouldNotBeNil)
					So(*d.Metadata.UnitOfMeasure, ShouldEqual, nem12.UnitKilowattHour)
					So(d.Data, ShouldHaveLength, 96)

					d, err = p.ReadDay()
					So(d, ShouldNotBeNil)
					So(err, ShouldBeNil)

					So(d.Metadata, ShouldNotBeNil)
					So(d.Metadata.Nmi, ShouldNotBeNil)
					So(d.Metadata.Nmi.Identifier, ShouldEqual, "NEM1201001")
					So(d.Metadata.Meter, ShouldNotBeNil)
					So(d.Metadata.Meter.Identifier, ShouldEqual, "2")
					So(d.Metadata.Suffix, ShouldNotBeNil)
					So(d.Metadata.Suffix.Type, ShouldEqual, nem12.SuffixExportWattHourMaster)
					So(d.Metadata.UnitOfMeasure, ShouldNotBeNil)
					So(*d.Metadata.UnitOfMeasure, ShouldEqual, nem12.UnitKilowattHour)
					So(d.Data, ShouldHaveLength, 96)

					d, err = p.ReadDay()
					So(d, ShouldBeNil)
					So(err, ShouldBeNil)

					d, err = p.ReadDay()
					So(d, ShouldBeNil)
					So(err, ShouldBeError)
					So(err, ShouldWrap, io.EOF)
				},
			},
			"valid nem12 file with variable data quality for one meter at one nmi": {
				reader: func() io.Reader {
					filepath := "./testdata/valid/NEM12#Scenario04#ETSAMDP#NEMMCO.csv"
					// 100,NEM12,200505231738,ETSAMDP,NEMMCO
					// 200,NEM1314071,E1,E1,E1,,04071,KWH,30,20050601
					// 300,20050101,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,V,,,20050321164041,
					// 400,1,20,F56,1,
					// 400,21,48,E56,77,
					// 500,G,SONEM1214071,20050101101523,001123.5
					// 300,20050102,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,E56,77,Estimation Forecast,20050321000001,
					// 300,20050103,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,4.583,E56,77,Estimation Forecast,20050321000001,
					// 900

					r, err := os.Open(filepath)
					So(err, ShouldBeNil)

					return r
				},
				assertions: func(r io.Reader, p nem12.Parser) {
					qe := nem12.QualityValue["E"]
					qf := nem12.QualityValue["F"]
					m56 := nem12.Method(56)
					r1 := nem12.Reason(1)
					r77 := nem12.Reason(77)
					r77Desc := "Estimation Forecast"
					update1 := time.Date(2005, 3, 21, 16, 40, 41, 0, nem12.NEMTime()) // 20050321164041
					update2 := time.Date(2005, 3, 21, 0, 0, 1, 0, nem12.NEMTime())    // 20050321000001

					So(r, ShouldNotBeNil)

					d, err := p.ReadDay()
					So(d, ShouldNotBeNil)
					So(err, ShouldBeNil)

					So(d.Metadata, ShouldNotBeNil)
					So(d.Metadata.Nmi, ShouldNotBeNil)
					So(d.Metadata.Nmi.Identifier, ShouldEqual, "NEM1314071")
					So(d.Metadata.Meter, ShouldNotBeNil)
					So(d.Metadata.Meter.Identifier, ShouldEqual, "1")
					So(d.Metadata.Suffix, ShouldNotBeNil)
					So(d.Metadata.Suffix.Type, ShouldEqual, nem12.SuffixExportWattHourMaster)
					So(d.Metadata.UnitOfMeasure, ShouldNotBeNil)
					So(*d.Metadata.UnitOfMeasure, ShouldEqual, nem12.UnitKilowattHour)

					So(d.Data, ShouldHaveLength, 48)
					ts := time.Date(2005, 1, 1, 0, 0, 0, 0, nem12.NEMTime())
					vals := []float64{
						0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583,
						4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583,
						4.583, 4.583, 4.583, 4.583, 4.583,
					}
					qualities := []nem12.Quality{
						qf, qf, qf, qf, qf, qf, qf, qf, qf, qf, qf, qf, qf, qf, qf, qf, qf, qf, qf, qf,
						qe, qe, qe, qe, qe, qe, qe, qe, qe, qe, qe, qe, qe, qe, qe, qe, qe, qe, qe, qe, qe, qe, qe, qe, qe, qe, qe, qe,
					}
					reasons := []*nem12.Reason{
						&r1, &r1, &r1, &r1, &r1, &r1, &r1, &r1, &r1, &r1, &r1, &r1, &r1, &r1, &r1, &r1, &r1, &r1, &r1, &r1,
						&r77, &r77, &r77, &r77, &r77, &r77, &r77, &r77, &r77, &r77, &r77, &r77, &r77, &r77, &r77, &r77, &r77, &r77, &r77, &r77, &r77, &r77, &r77, &r77, &r77, &r77, &r77, &r77,
					}

					for i, val := range vals {
						So(d.Data[i], ShouldResemble, &nem12.Interval{
							Time:           ts.Add(time.Duration((i+1)*30) * time.Minute),
							IntervalLength: (30 * time.Minute),
							Value: nem12.IntervalValue{
								Value:             val,
								DecimalValue:      decimal.NewFromFloat(val),
								QualityFlag:       qualities[i],
								MethodFlag:        &m56,
								ReasonCode:        reasons[i],
								ReasonDescription: nil,
								UpdateDateTime:    &update1,
								MSATSLoadDateTime: nil,
							},
						})
					}

					d, err = p.ReadDay()
					So(d, ShouldNotBeNil)
					So(err, ShouldBeNil)

					So(d.Metadata, ShouldNotBeNil)
					So(d.Metadata.Nmi, ShouldNotBeNil)
					So(d.Metadata.Nmi.Identifier, ShouldEqual, "NEM1314071")
					So(d.Metadata.Meter, ShouldNotBeNil)
					So(d.Metadata.Meter.Identifier, ShouldEqual, "1")
					So(d.Metadata.Suffix, ShouldNotBeNil)
					So(d.Metadata.Suffix.Type, ShouldEqual, nem12.SuffixExportWattHourMaster)
					So(d.Metadata.UnitOfMeasure, ShouldNotBeNil)
					So(*d.Metadata.UnitOfMeasure, ShouldEqual, nem12.UnitKilowattHour)
					So(d.Data, ShouldHaveLength, 48)

					ts = time.Date(2005, 1, 2, 0, 0, 0, 0, nem12.NEMTime())
					vals = []float64{
						4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583,
						4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583,
						4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583, 4.583,
					}
					for i, val := range vals {
						So(d.Data[i], ShouldResemble, &nem12.Interval{
							Time:           ts.Add(time.Duration((i+1)*30) * time.Minute),
							IntervalLength: (30 * time.Minute),
							Value: nem12.IntervalValue{
								Value:             val,
								DecimalValue:      decimal.NewFromFloat(val),
								QualityFlag:       qe,
								MethodFlag:        &m56,
								ReasonCode:        &r77,
								ReasonDescription: &r77Desc,
								UpdateDateTime:    &update2,
								MSATSLoadDateTime: nil,
							},
						})
					}

					d, err = p.ReadDay()
					So(d, ShouldNotBeNil)
					So(err, ShouldBeNil)

					So(d.Metadata, ShouldNotBeNil)
					So(d.Metadata.Nmi, ShouldNotBeNil)
					So(d.Metadata.Nmi.Identifier, ShouldEqual, "NEM1314071")
					So(d.Metadata.Meter, ShouldNotBeNil)
					So(d.Metadata.Meter.Identifier, ShouldEqual, "1")
					So(d.Metadata.Suffix, ShouldNotBeNil)
					So(d.Metadata.Suffix.Type, ShouldEqual, nem12.SuffixExportWattHourMaster)
					So(d.Metadata.UnitOfMeasure, ShouldNotBeNil)
					So(*d.Metadata.UnitOfMeasure, ShouldEqual, nem12.UnitKilowattHour)
					So(d.Data, ShouldHaveLength, 48)

					d, err = p.ReadDay()
					So(d, ShouldBeNil)
					So(err, ShouldBeNil)

					d, err = p.ReadDay()
					So(d, ShouldBeNil)
					So(err, ShouldBeError)
					So(err, ShouldWrap, io.EOF)
				},
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given %s", name), func() {
				r := tc.reader()

				p := nem12.NewParser(r)

				tc.assertions(r, p)
			})
		}
	})
}
