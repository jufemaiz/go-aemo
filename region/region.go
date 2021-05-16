package region

import "fmt"

const (
	RegionUndefined Region = iota
	RegionAAT
	RegionACT
	RegionNSW
	RegionNT
	RegionQLD
	RegionSA
	RegionTAS
	RegionVIC
	RegionWA
)

var (
	Region_name = map[Region]string{
		RegionUndefined: "UNDEFINED",
		RegionAAT:       "AAT",
		RegionACT:       "ACT",
		RegionNSW:       "NSW",
		RegionNT:        "NT",
		RegionQLD:       "QLD",
		RegionSA:        "SA",
		RegionTAS:       "TAS",
		RegionVIC:       "VIC",
		RegionWA:        "WA",
	}

	Region_value = map[string]Region{
		"UNDEFINED": RegionUndefined,
		"AAT":       RegionAAT,
		"ACT":       RegionACT,
		"NSW":       RegionNSW,
		"NT":        RegionNT,
		"QLD":       RegionQLD,
		"SA":        RegionSA,
		"TAS":       RegionTAS,
		"VIC":       RegionVIC,
		"WA":        RegionWA,
	}

	regions = []Region{
		RegionUndefined,
		RegionAAT,
		RegionACT,
		RegionNSW,
		RegionNT,
		RegionQLD,
		RegionSA,
		RegionTAS,
		RegionVIC,
		RegionWA,
	}

	regionLongNames = []string{
		"Undefined",
		"Australian Antarctic Territory",
		"Australian Capital Territory",
		"New South Wales",
		"Northern Territory",
		"Queensland",
		"South Australia",
		"Tasmania",
		"Victoria",
		"Western Australia",
	}

	regionNames = []string{
		"UNDEFINED",
		"AAT",
		"ACT",
		"NSW",
		"NT",
		"QLD",
		"SA",
		"TAS",
		"VIC",
		"WA",
	}

	regionMarketNodes = []Region{
		RegionUndefined,
		RegionUndefined,
		RegionNSW,
		RegionNSW,
		RegionNT,
		RegionQLD,
		RegionSA,
		RegionTAS,
		RegionVIC,
		RegionWA,
	}

	regionISOCodes = []string{
		"UNDEFINED",
		"UNDEFINED",
		"AU-ACT",
		"AU-NSW",
		"AU-NT",
		"AU-QLD",
		"AU-SA",
		"AU-TAS",
		"AU-VIC",
		"AU-WA",
	}
)

// Region represents one of the regions that AEMO operates in.
type Region int32

// RegionInfo holds a structured set of data for a region.
type RegionInfo struct {
	Region     Region `json:"-"`
	MarketNode Region `json:"marketNode"`
	Name       string `json:"name"`
	LongName   string `json:"longName"`
	ISOCode    string `json:"isoCode"`
}

// NewRegion returns a region for a string (matching short name).
func NewRegion(s string) (Region, error) {
	r, ok := Region_value[s]
	if !ok {
		return RegionUndefined, ErrRegionInvalid
	}

	return r, nil
}

// Regions returns all valid regions.
func Regions() []Region {
	return regions[1:]
}

// GoString meets the gostring interface.
func (r Region) GoString() string {
	return fmt.Sprintf(
		"{Region: %d, MarketNode: \"%s\", Name: \"%s\", LongName: \"%s\", ISOCode: \"%s\"}",
		r, r.MarketNode().Name(), r.Name(), r.LongName(), r.ISOCode(),
	)
}

// Info struct for a region.
func (r Region) Info() (*RegionInfo, error) {
	if r == RegionUndefined {
		return nil, fmt.Errorf("region '%d': %w", r, ErrRegionInvalid)
	}

	return &RegionInfo{
		Region:     r,
		MarketNode: regionMarketNodes[r],
		Name:       regionNames[r],
		LongName:   regionLongNames[r],
		ISOCode:    regionISOCodes[r],
	}, nil
}

// MarketNode retuns the market node fo the region.
func (r Region) MarketNode() Region {
	return regions[regionMarketNodes[r]]
}

// Name retuns the name of the region.
func (r Region) Name() string {
	return regionNames[r]
}

// LongName retuns the long (full) name of the region.
func (r Region) LongName() string {
	return regionLongNames[r]
}

// ISOCode retuns the ISO code of the region.
func (r Region) ISOCode() string {
	return regionISOCodes[r]
}
