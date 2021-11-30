package nem12_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/jufemaiz/go-aemo/nem12"
	"github.com/jufemaiz/go-aemo/nmi"
	"github.com/shopspring/decimal"
	. "github.com/smartystreets/goconvey/convey"
)

func TestIntervalSet_Normalize(t *testing.T) {
	Convey("nem12/IntervalSet.Normalize", t, func() {
		tests := map[string]struct {
			set        func() *nem12.IntervalSet
			assertions func(set *nem12.IntervalSet, resp *nem12.IntervalSet, err error)
			err        error
		}{
			"interval set is nil": {
				set: func() *nem12.IntervalSet { return nil },
				assertions: func(set *nem12.IntervalSet, resp *nem12.IntervalSet, err error) {
					So(resp, ShouldBeNil)
				},
			},
			"interval set metadata is nil": {
				set: func() *nem12.IntervalSet {
					return &nem12.IntervalSet{
						Data: []*nem12.Interval{
							{
								Time:           time.Date(2021, 9, 30, 1, 30, 0, 0, time.UTC),
								IntervalLength: 30 * time.Minute,
								Value: nem12.IntervalValue{
									Value:        1.234,
									DecimalValue: decimal.NewFromFloat(1.234),
									QualityFlag:  nem12.QualityActual,
								},
							},
						},
					}
				},
				assertions: func(set *nem12.IntervalSet, resp *nem12.IntervalSet, err error) {},
				err:        nem12.ErrIntervalMetadataNil,
			},
			"uom is nil": {
				set: func() *nem12.IntervalSet {
					sfx, err := nem12.NewSuffix("E1")
					So(err, ShouldBeNil)

					return &nem12.IntervalSet{
						Metadata: &nem12.IntervalMetadata{
							Nmi:    &nmi.Nmi{Identifier: "4123456789"},
							Meter:  &nmi.Meter{Identifier: "1"},
							Suffix: &sfx,
						},
						Data: []*nem12.Interval{
							{
								Time:           time.Date(2021, 9, 30, 1, 30, 0, 0, time.UTC),
								IntervalLength: 30 * time.Minute,
								Value: nem12.IntervalValue{
									Value:        1.234,
									DecimalValue: decimal.NewFromFloat(1.234),
									QualityFlag:  nem12.QualityActual,
								},
							},
						},
					}
				},
				assertions: func(set *nem12.IntervalSet, resp *nem12.IntervalSet, err error) {},
				err:        nem12.ErrUnitOfMeasureNil,
			},
			"uom is invalid": {
				set: func() *nem12.IntervalSet {
					uom := nem12.UnitOfMeasure(-1)
					So(uom.Validate(), ShouldBeError)

					sfx, err := nem12.NewSuffix("E1")
					So(err, ShouldBeNil)

					return &nem12.IntervalSet{
						Metadata: &nem12.IntervalMetadata{
							Nmi:           &nmi.Nmi{Identifier: "4123456789"},
							Meter:         &nmi.Meter{Identifier: "1"},
							Suffix:        &sfx,
							UnitOfMeasure: &uom,
						},
						Data: []*nem12.Interval{
							{
								Time:           time.Date(2021, 9, 30, 1, 30, 0, 0, time.UTC),
								IntervalLength: 30 * time.Minute,
								Value: nem12.IntervalValue{
									Value:        1.234,
									DecimalValue: decimal.NewFromFloat(1.234),
									QualityFlag:  nem12.QualityActual,
								},
							},
						},
					}
				},
				assertions: func(set *nem12.IntervalSet, resp *nem12.IntervalSet, err error) {},
				err:        nem12.ErrUnitOfMeasureInvalid,
			},
			"interval set has nil data": {
				set: func() *nem12.IntervalSet {
					uom, err := nem12.NewUnit("KWH")
					So(err, ShouldBeNil)

					sfx, err := nem12.NewSuffix("E1")
					So(err, ShouldBeNil)

					return &nem12.IntervalSet{
						Metadata: &nem12.IntervalMetadata{
							Nmi:           &nmi.Nmi{Identifier: "4123456789"},
							Meter:         &nmi.Meter{Identifier: "1"},
							Suffix:        &sfx,
							UnitOfMeasure: &uom,
						},
					}
				},
				assertions: func(set *nem12.IntervalSet, resp *nem12.IntervalSet, err error) {
					So(resp, ShouldNotBeNil)

					So(resp.Metadata, ShouldNotBeNil)
					So(resp.Metadata.Nmi, ShouldResemble, set.Metadata.Nmi)
					So(resp.Metadata.Meter, ShouldResemble, set.Metadata.Meter)
					So(resp.Metadata.Suffix, ShouldResemble, set.Metadata.Suffix)

					So(*resp.Metadata.UnitOfMeasure, ShouldResemble, nem12.UnitWattHour)

					So(resp.Data, ShouldHaveLength, 0)
				},
			},
			"interval set has empty data": {
				set: func() *nem12.IntervalSet {
					uom, err := nem12.NewUnit("KWH")
					So(err, ShouldBeNil)

					sfx, err := nem12.NewSuffix("E1")
					So(err, ShouldBeNil)

					return &nem12.IntervalSet{
						Metadata: &nem12.IntervalMetadata{
							Nmi:           &nmi.Nmi{Identifier: "4123456789"},
							Meter:         &nmi.Meter{Identifier: "1"},
							Suffix:        &sfx,
							UnitOfMeasure: &uom,
						},
						Data: []*nem12.Interval{},
					}
				},
				assertions: func(set *nem12.IntervalSet, resp *nem12.IntervalSet, err error) {
					So(resp, ShouldNotBeNil)

					So(resp.Metadata, ShouldNotBeNil)
					So(resp.Metadata.Nmi, ShouldResemble, set.Metadata.Nmi)
					So(resp.Metadata.Meter, ShouldResemble, set.Metadata.Meter)
					So(resp.Metadata.Suffix, ShouldResemble, set.Metadata.Suffix)

					So(*resp.Metadata.UnitOfMeasure, ShouldResemble, nem12.UnitWattHour)

					So(resp.Data, ShouldHaveLength, 0)
				},
			},
			"interval is set and uom is valid and in SI units": {
				set: func() *nem12.IntervalSet {
					uom, err := nem12.NewUnit("WH")
					So(err, ShouldBeNil)

					sfx, err := nem12.NewSuffix("E1")
					So(err, ShouldBeNil)

					return &nem12.IntervalSet{
						Metadata: &nem12.IntervalMetadata{
							Nmi:           &nmi.Nmi{Identifier: "4123456789"},
							Meter:         &nmi.Meter{Identifier: "1"},
							Suffix:        &sfx,
							UnitOfMeasure: &uom,
						},
						Data: []*nem12.Interval{
							{
								Time:           time.Date(2021, 9, 30, 1, 30, 0, 0, time.UTC),
								IntervalLength: 30 * time.Minute,
								Value: nem12.IntervalValue{
									Value:        1.234,
									DecimalValue: decimal.NewFromFloat(1.234),
									QualityFlag:  nem12.QualityActual,
								},
							},
						},
					}
				},
				assertions: func(set *nem12.IntervalSet, resp *nem12.IntervalSet, err error) {
					So(resp, ShouldNotBeNil)

					So(resp.Metadata, ShouldNotBeNil)
					So(resp.Metadata.Nmi, ShouldResemble, set.Metadata.Nmi)
					So(resp.Metadata.Meter, ShouldResemble, set.Metadata.Meter)
					So(resp.Metadata.Suffix, ShouldResemble, set.Metadata.Suffix)

					So(*resp.Metadata.UnitOfMeasure, ShouldResemble, nem12.UnitWattHour)

					So(resp.Data, ShouldHaveLength, 1)

					So(resp.Data[0].Time, ShouldEqual, set.Data[0].Time)
					So(resp.Data[0].IntervalLength, ShouldEqual, set.Data[0].IntervalLength)
					So(resp.Data[0].Value.Value, ShouldEqual, set.Data[0].Value.Value)
					So(resp.Data[0].Value.DecimalValue, ShouldEqual, set.Data[0].Value.DecimalValue)
					So(resp.Data[0].Value.QualityFlag, ShouldEqual, set.Data[0].Value.QualityFlag)
					So(resp.Data[0].Value.MethodFlag, ShouldEqual, set.Data[0].Value.MethodFlag)
					So(resp.Data[0].Value.ReasonCode, ShouldEqual, set.Data[0].Value.ReasonCode)
					So(resp.Data[0].Value.ReasonDescription, ShouldEqual, set.Data[0].Value.ReasonDescription)
					So(resp.Data[0].Value.UpdateDateTime, ShouldEqual, set.Data[0].Value.UpdateDateTime)
					So(resp.Data[0].Value.MSATSLoadDateTime, ShouldEqual, set.Data[0].Value.MSATSLoadDateTime)
					So(resp.Data[0].Metadata, ShouldResemble, set.Data[0].Metadata)
				},
			},
			"interval is set and uom is valid and not in SI units": {
				set: func() *nem12.IntervalSet {
					uom, err := nem12.NewUnit("MWH")
					So(err, ShouldBeNil)

					sfx, err := nem12.NewSuffix("E1")
					So(err, ShouldBeNil)

					return &nem12.IntervalSet{
						Metadata: &nem12.IntervalMetadata{
							Nmi:           &nmi.Nmi{Identifier: "4123456789"},
							Meter:         &nmi.Meter{Identifier: "1"},
							Suffix:        &sfx,
							UnitOfMeasure: &uom,
						},
						Data: []*nem12.Interval{
							{
								Time:           time.Date(2021, 9, 30, 1, 30, 0, 0, time.UTC),
								IntervalLength: 30 * time.Minute,
								Value: nem12.IntervalValue{
									Value:        1.234,
									DecimalValue: decimal.NewFromFloat(1.234),
									QualityFlag:  nem12.QualityActual,
								},
							},
						},
					}
				},
				assertions: func(set *nem12.IntervalSet, resp *nem12.IntervalSet, err error) {
					So(resp, ShouldNotBeNil)

					So(resp.Metadata, ShouldNotBeNil)
					So(resp.Metadata.Nmi, ShouldResemble, set.Metadata.Nmi)
					So(resp.Metadata.Meter, ShouldResemble, set.Metadata.Meter)
					So(resp.Metadata.Suffix, ShouldResemble, set.Metadata.Suffix)

					So(*resp.Metadata.UnitOfMeasure, ShouldResemble, nem12.UnitWattHour)

					So(resp.Data, ShouldHaveLength, 1)

					So(resp.Data[0].Time, ShouldEqual, set.Data[0].Time)
					So(resp.Data[0].IntervalLength, ShouldEqual, set.Data[0].IntervalLength)
					So(resp.Data[0].Value.Value, ShouldEqual, set.Data[0].Value.Value*1e6)
					So(resp.Data[0].Value.DecimalValue, ShouldEqual, set.Data[0].Value.DecimalValue.Mul(decimal.NewFromFloat(1e6)))
					So(resp.Data[0].Value.QualityFlag, ShouldEqual, set.Data[0].Value.QualityFlag)
					So(resp.Data[0].Value.MethodFlag, ShouldEqual, set.Data[0].Value.MethodFlag)
					So(resp.Data[0].Value.ReasonCode, ShouldEqual, set.Data[0].Value.ReasonCode)
					So(resp.Data[0].Value.ReasonDescription, ShouldEqual, set.Data[0].Value.ReasonDescription)
					So(resp.Data[0].Value.UpdateDateTime, ShouldEqual, set.Data[0].Value.UpdateDateTime)
					So(resp.Data[0].Value.MSATSLoadDateTime, ShouldEqual, set.Data[0].Value.MSATSLoadDateTime)
					So(resp.Data[0].Metadata, ShouldResemble, set.Data[0].Metadata)
				},
			},
			"interval set with a nil data value and uom is valid and not in SI units": {
				set: func() *nem12.IntervalSet {
					uom, err := nem12.NewUnit("MWH")
					So(err, ShouldBeNil)

					sfx, err := nem12.NewSuffix("E1")
					So(err, ShouldBeNil)

					return &nem12.IntervalSet{
						Metadata: &nem12.IntervalMetadata{
							Nmi:           &nmi.Nmi{Identifier: "4123456789"},
							Meter:         &nmi.Meter{Identifier: "1"},
							Suffix:        &sfx,
							UnitOfMeasure: &uom,
						},
						Data: []*nem12.Interval{
							nil,
							{
								Time:           time.Date(2021, 9, 30, 1, 30, 0, 0, time.UTC),
								IntervalLength: 30 * time.Minute,
								Value: nem12.IntervalValue{
									Value:        1.234,
									DecimalValue: decimal.NewFromFloat(1.234),
									QualityFlag:  nem12.QualityActual,
								},
							},
							nil,
						},
					}
				},
				assertions: func(set *nem12.IntervalSet, resp *nem12.IntervalSet, err error) {
					So(resp, ShouldNotBeNil)

					So(resp.Metadata, ShouldNotBeNil)
					So(resp.Metadata.Nmi, ShouldResemble, set.Metadata.Nmi)
					So(resp.Metadata.Meter, ShouldResemble, set.Metadata.Meter)
					So(resp.Metadata.Suffix, ShouldResemble, set.Metadata.Suffix)

					So(*resp.Metadata.UnitOfMeasure, ShouldResemble, nem12.UnitWattHour)

					So(resp.Data, ShouldHaveLength, 1)

					So(resp.Data[0].Time, ShouldEqual, set.Data[1].Time)
					So(resp.Data[0].IntervalLength, ShouldEqual, set.Data[1].IntervalLength)
					So(resp.Data[0].Value.Value, ShouldEqual, set.Data[1].Value.Value*1e6)
					So(resp.Data[0].Value.DecimalValue, ShouldEqual, set.Data[1].Value.DecimalValue.Mul(decimal.NewFromFloat(1e6)))
					So(resp.Data[0].Value.QualityFlag, ShouldEqual, set.Data[1].Value.QualityFlag)
					So(resp.Data[0].Value.MethodFlag, ShouldEqual, set.Data[1].Value.MethodFlag)
					So(resp.Data[0].Value.ReasonCode, ShouldEqual, set.Data[1].Value.ReasonCode)
					So(resp.Data[0].Value.ReasonDescription, ShouldEqual, set.Data[1].Value.ReasonDescription)
					So(resp.Data[0].Value.UpdateDateTime, ShouldEqual, set.Data[1].Value.UpdateDateTime)
					So(resp.Data[0].Value.MSATSLoadDateTime, ShouldEqual, set.Data[1].Value.MSATSLoadDateTime)
					So(resp.Data[0].Metadata, ShouldResemble, set.Data[1].Metadata)
				},
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given %s", name), func() {
				set := tc.set()

				resp, err := set.Normalize()

				if tc.err != nil {
					So(err, ShouldWrap, tc.err)
					So(resp, ShouldBeNil)
				} else {
					So(err, ShouldBeNil)
				}

				tc.assertions(set, resp, err)
			})
		}
	})
}

func TestInterval_Normalize(t *testing.T) {
	Convey("nem12/Interval.Normalize", t, func() {
		tests := map[string]struct {
			interval   func() *nem12.Interval
			arg        func() *nem12.UnitOfMeasure
			assertions func(interval *nem12.Interval, arg *nem12.UnitOfMeasure, resp *nem12.Interval, err error)
			err        error
		}{
			"interval is nil": {
				interval: func() *nem12.Interval { return nil },
				arg: func() *nem12.UnitOfMeasure {
					uom, err := nem12.NewUnit("WH")
					So(err, ShouldBeNil)

					return &uom
				},
				assertions: func(interval *nem12.Interval, arg *nem12.UnitOfMeasure, resp *nem12.Interval, err error) {},
				err:        nem12.ErrIntervalNil,
			},
			"uom is invalid": {
				interval: func() *nem12.Interval {
					return &nem12.Interval{
						Time:           time.Date(2021, 9, 30, 1, 30, 0, 0, time.UTC),
						IntervalLength: 30 * time.Minute,
						Value: nem12.IntervalValue{
							Value:        1.234,
							DecimalValue: decimal.NewFromFloat(1.234),
							QualityFlag:  nem12.QualityActual,
						},
					}
				},
				arg: func() *nem12.UnitOfMeasure {
					uom := nem12.UnitOfMeasure(-5)
					So(uom.Validate(), ShouldBeError)

					return &uom
				},
				assertions: func(interval *nem12.Interval, arg *nem12.UnitOfMeasure, resp *nem12.Interval, err error) {},
				err:        nem12.ErrUnitOfMeasureInvalid,
			},
			"interval is set and uom is nil": {
				interval: func() *nem12.Interval {
					return &nem12.Interval{
						Time:           time.Date(2021, 9, 30, 1, 30, 0, 0, time.UTC),
						IntervalLength: 30 * time.Minute,
						Value: nem12.IntervalValue{
							Value:        1.234,
							DecimalValue: decimal.NewFromFloat(1.234),
							QualityFlag:  nem12.QualityActual,
						},
					}
				},
				arg: func() *nem12.UnitOfMeasure {
					return nil
				},
				assertions: func(interval *nem12.Interval, arg *nem12.UnitOfMeasure, resp *nem12.Interval, err error) {
					So(resp, ShouldResemble, interval)
				},
			},
			"interval is set and uom is valid and in SI units": {
				interval: func() *nem12.Interval {
					return &nem12.Interval{
						Time:           time.Date(2021, 9, 30, 1, 30, 0, 0, time.UTC),
						IntervalLength: 30 * time.Minute,
						Value: nem12.IntervalValue{
							Value:        1.234,
							DecimalValue: decimal.NewFromFloat(1.234),
							QualityFlag:  nem12.QualityActual,
						},
					}
				},
				arg: func() *nem12.UnitOfMeasure {
					uom, err := nem12.NewUnit("WH")
					So(err, ShouldBeNil)

					return &uom
				},
				assertions: func(interval *nem12.Interval, arg *nem12.UnitOfMeasure, resp *nem12.Interval, err error) {
					So(resp, ShouldResemble, interval)
				},
			},
			"interval is set and uom is valid and not in SI units": {
				interval: func() *nem12.Interval {
					return &nem12.Interval{
						Time:           time.Date(2021, 9, 30, 1, 30, 0, 0, time.UTC),
						IntervalLength: 30 * time.Minute,
						Value: nem12.IntervalValue{
							Value:        1.234,
							DecimalValue: decimal.NewFromFloat(1.234),
							QualityFlag:  nem12.QualityActual,
						},
					}
				},
				arg: func() *nem12.UnitOfMeasure {
					uom, err := nem12.NewUnit("MWH")
					So(err, ShouldBeNil)

					return &uom
				},
				assertions: func(interval *nem12.Interval, arg *nem12.UnitOfMeasure, resp *nem12.Interval, err error) {
					So(resp.Time, ShouldEqual, interval.Time)
					So(resp.IntervalLength, ShouldEqual, interval.IntervalLength)
					So(resp.Value.Value, ShouldEqual, interval.Value.Value*1e6)
					So(resp.Value.DecimalValue, ShouldEqual, interval.Value.DecimalValue.Mul(decimal.NewFromFloat(1e6)))
					So(resp.Value.QualityFlag, ShouldEqual, interval.Value.QualityFlag)
					So(resp.Value.MethodFlag, ShouldEqual, interval.Value.MethodFlag)
					So(resp.Value.ReasonCode, ShouldEqual, interval.Value.ReasonCode)
					So(resp.Value.ReasonDescription, ShouldEqual, interval.Value.ReasonDescription)
					So(resp.Value.UpdateDateTime, ShouldEqual, interval.Value.UpdateDateTime)
					So(resp.Value.MSATSLoadDateTime, ShouldEqual, interval.Value.MSATSLoadDateTime)
					So(resp.Metadata, ShouldResemble, interval.Metadata)
				},
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given %s", name), func() {
				interval := tc.interval()

				arg := tc.arg()
				resp, err := interval.Normalize(arg)

				if tc.err != nil {
					So(err, ShouldWrap, tc.err)
					So(resp, ShouldBeNil)
				} else {
					So(resp, ShouldNotBeNil)
					So(err, ShouldBeNil)
				}

				tc.assertions(interval, arg, resp, err)
			})
		}
	})
}
