package nmi

import (
	"fmt"
	"strings"

	"github.com/jufemaiz/go-aemo/region"
)

const (
	// ParticipantUndefined when participant is undefined.
	ParticipantUndefined Participant = iota
	// ParticipantACTEWP for the participant code ACTEWP.
	ParticipantACTEWP
	// ParticipantAEMORESERVED for the participant code AEMORESERVED.
	ParticipantAEMORESERVED
	// ParticipantAURORAP for the participant code AURORAP.
	ParticipantAURORAP
	// ParticipantCITIPP for the participant code CITIPP.
	ParticipantCITIPP
	// ParticipantCNRGYP for the participant code CNRGYP.
	ParticipantCNRGYP
	// ParticipantEASTERN for the participant code EASTERN.
	ParticipantEASTERN
	// ParticipantENERGEXP for the participant code ENERGEXP.
	ParticipantENERGEXP
	// ParticipantENERGYAP for the participant code ENERGYAP.
	ParticipantENERGYAP
	// ParticipantERGONETP for the participant code ERGONETP.
	ParticipantERGONETP
	// ParticipantETSATP for the participant code ETSATP.
	ParticipantETSATP
	// ParticipantEXEMPTNETWORKS for the participant code EXEMPTNETWORKS.
	ParticipantEXEMPTNETWORKS
	// ParticipantFEDAIRPORTS for the participant code FEDAIRPORTS.
	ParticipantFEDAIRPORTS
	// ParticipantGASNSW for the participant code GASNSW.
	ParticipantGASNSW
	// ParticipantGASQLD for the participant code GASQLD.
	ParticipantGASQLD
	// ParticipantGASSA for the participant code GASSA.
	ParticipantGASSA
	// ParticipantGASTAS for the participant code GASTAS.
	ParticipantGASTAS
	// ParticipantGASVIC for the participant code GASVIC.
	ParticipantGASVIC
	// ParticipantGASWA for the participant code GASWA.
	ParticipantGASWA
	// ParticipantGPUPP for the participant code GPUPP.
	ParticipantGPUPP
	// ParticipantHORIZONPOWER for the participant code HORIZONPOWER.
	ParticipantHORIZONPOWER
	// ParticipantINTEGP for the participant code INTEGP.
	ParticipantINTEGP
	// ParticipantNTRESERVED for the participant code NTRESERVED.
	ParticipantNTRESERVED
	// ParticipantPLINKP for the participant code PLINKP.
	ParticipantPLINKP
	// ParticipantPOWCP for the participant code POWCP.
	ParticipantPOWCP
	// ParticipantSNOWY for the participant code SNOWY.
	ParticipantSNOWY
	// ParticipantSOLARISP for the participant code SOLARISP.
	ParticipantSOLARISP
	// ParticipantTRANSEND for the participant code TRANSEND.
	ParticipantTRANSEND
	// ParticipantTRANSGP for the participant code TRANSGP.
	ParticipantTRANSGP
	// ParticipantUMPLP for the participant code UMPLP.
	ParticipantUMPLP
	// ParticipantUNITED for the participant code UNITED.
	ParticipantUNITED
	// ParticipantWESTERNPOWER for the participant code WESTERNPOWER.
	ParticipantWESTERNPOWER
)

var (
	participants = []Participant{ //nolint:gochecknoglobals
		ParticipantACTEWP,
		ParticipantAEMORESERVED,
		ParticipantAURORAP,
		ParticipantCITIPP,
		ParticipantCNRGYP,
		ParticipantEASTERN,
		ParticipantENERGEXP,
		ParticipantENERGYAP,
		ParticipantERGONETP,
		ParticipantETSATP,
		ParticipantEXEMPTNETWORKS,
		ParticipantFEDAIRPORTS,
		ParticipantGASNSW,
		ParticipantGASQLD,
		ParticipantGASSA,
		ParticipantGASTAS,
		ParticipantGASVIC,
		ParticipantGASWA,
		ParticipantGPUPP,
		ParticipantHORIZONPOWER,
		ParticipantINTEGP,
		ParticipantNTRESERVED,
		ParticipantPLINKP,
		ParticipantPOWCP,
		ParticipantSNOWY,
		ParticipantSOLARISP,
		ParticipantTRANSEND,
		ParticipantTRANSGP,
		ParticipantUMPLP,
		ParticipantUNITED,
		ParticipantWESTERNPOWER,
	}

	// ParticipantName maps Participants to strings.
	ParticipantName = map[Participant]string{ //nolint:gochecknoglobals
		ParticipantUndefined:      "UNDEFINED",
		ParticipantACTEWP:         "ACTEWP",
		ParticipantAEMORESERVED:   "AEMORESERVED",
		ParticipantAURORAP:        "AURORAP",
		ParticipantCITIPP:         "CITIPP",
		ParticipantCNRGYP:         "CNRGYP",
		ParticipantEASTERN:        "EASTERN",
		ParticipantENERGEXP:       "ENERGEXP",
		ParticipantENERGYAP:       "ENERGYAP",
		ParticipantERGONETP:       "ERGONETP",
		ParticipantETSATP:         "ETSATP",
		ParticipantEXEMPTNETWORKS: "EXEMPTNETWORKS",
		ParticipantFEDAIRPORTS:    "FEDAIRPORTS",
		ParticipantGASNSW:         "GASNSW",
		ParticipantGASQLD:         "GASQLD",
		ParticipantGASSA:          "GASSA",
		ParticipantGASTAS:         "GASTAS",
		ParticipantGASVIC:         "GASVIC",
		ParticipantGASWA:          "GASWA",
		ParticipantGPUPP:          "GPUPP",
		ParticipantHORIZONPOWER:   "HORIZONPOWER",
		ParticipantINTEGP:         "INTEGP",
		ParticipantNTRESERVED:     "NTRESERVED",
		ParticipantPLINKP:         "PLINKP",
		ParticipantPOWCP:          "POWCP",
		ParticipantSNOWY:          "SNOWY",
		ParticipantSOLARISP:       "SOLARISP",
		ParticipantTRANSEND:       "TRANSEND",
		ParticipantTRANSGP:        "TRANSGP",
		ParticipantUMPLP:          "UMPLP",
		ParticipantUNITED:         "UNITED",
		ParticipantWESTERNPOWER:   "WESTERNPOWER",
	}

	// ParticipantValue maps strings to Participants.
	ParticipantValue = map[string]Participant{ //nolint:gochecknoglobals
		"UNDEFINED":      ParticipantUndefined,
		"ACTEWP":         ParticipantACTEWP,
		"AEMORESERVED":   ParticipantAEMORESERVED,
		"AURORAP":        ParticipantAURORAP,
		"CITIPP":         ParticipantCITIPP,
		"CNRGYP":         ParticipantCNRGYP,
		"EASTERN":        ParticipantEASTERN,
		"ENERGEXP":       ParticipantENERGEXP,
		"ENERGYAP":       ParticipantENERGYAP,
		"ERGONETP":       ParticipantERGONETP,
		"ETSATP":         ParticipantETSATP,
		"EXEMPTNETWORKS": ParticipantEXEMPTNETWORKS,
		"FEDAIRPORTS":    ParticipantFEDAIRPORTS,
		"GASNSW":         ParticipantGASNSW,
		"GASQLD":         ParticipantGASQLD,
		"GASSA":          ParticipantGASSA,
		"GASTAS":         ParticipantGASTAS,
		"GASVIC":         ParticipantGASVIC,
		"GASWA":          ParticipantGASWA,
		"GPUPP":          ParticipantGPUPP,
		"HORIZONPOWER":   ParticipantHORIZONPOWER,
		"INTEGP":         ParticipantINTEGP,
		"NTRESERVED":     ParticipantNTRESERVED,
		"PLINKP":         ParticipantPLINKP,
		"POWCP":          ParticipantPOWCP,
		"SNOWY":          ParticipantSNOWY,
		"SOLARISP":       ParticipantSOLARISP,
		"TRANSEND":       ParticipantTRANSEND,
		"TRANSGP":        ParticipantTRANSGP,
		"UMPLP":          ParticipantUMPLP,
		"UNITED":         ParticipantUNITED,
		"WESTERNPOWER":   ParticipantWESTERNPOWER,
	}

	participantRegions = []region.Region{ //nolint:gochecknoglobals
		region.RegionUndefined,
		region.RegionACT,
		region.RegionUndefined,
		region.RegionTAS,
		region.RegionVIC,
		region.RegionNSW,
		region.RegionVIC,
		region.RegionQLD,
		region.RegionNSW,
		region.RegionQLD,
		region.RegionSA,
		region.RegionUndefined,
		region.RegionNSW,
		region.RegionNSW,
		region.RegionQLD,
		region.RegionSA,
		region.RegionTAS,
		region.RegionVIC,
		region.RegionWA,
		region.RegionVIC,
		region.RegionWA,
		region.RegionNSW,
		region.RegionNT,
		region.RegionQLD,
		region.RegionVIC,
		region.RegionNSW,
		region.RegionVIC,
		region.RegionTAS,
		region.RegionNSW,
		region.RegionSA,
		region.RegionVIC,
		region.RegionWA,
	}

	// ParticipantIDs lists all participant ids.
	ParticipantIDs = []string{ //nolint:gochecknoglobals
		"UNDEFINED",
		"ACTEWP",
		"AEMORESERVED",
		"AURORAP",
		"CITIPP",
		"CNRGYP",
		"EASTERN",
		"ENERGEXP",
		"ENERGYAP",
		"ERGONETP",
		"ETSATP",
		"EXEMPTNETWORKS",
		"FEDAIRPORTS",
		"GASNSW",
		"GASQLD",
		"GASSA",
		"GASTAS",
		"GASVIC",
		"GASWA",
		"GPUPP",
		"HORIZONPOWER",
		"INTEGP",
		"NTRESERVED",
		"PLINKP",
		"POWCP",
		"SNOWY",
		"SOLARISP",
		"TRANSEND",
		"TRANSGP",
		"UMPLP",
		"UNITED",
		"WESTERNPOWER",
	}

	participantLongNames = []string{ //nolint:gochecknoglobals
		"UNDEFINED",
		"Icon Distribution Investments Limited and Jemena Networks (ACT) Pty Ltd Trading as Evoenergy – DNSP & TNSP",
		"AEMO Reserved",
		"TasNetworks",
		"CitiPower",
		"Essential Energy",
		"SP AusNet",
		"ENERGEX Limited",
		"Ausgrid",
		"Ergon Energy Corporation",
		"ElectraNet SA",
		"Exempt Networks - various",
		"Federal Airports Corporation (Sydney Airport)",
		"GAS NSW",
		"GAS QLD",
		"GAS SA",
		"GAS TAS",
		"GAS VIC",
		"GAS WA",
		"SP AusNet TNSP",
		"Horizon Power",
		"Endeavour Energy",
		"Northern Territory Reserved Block",
		"Qld Electricity Transmission Corp (Powerlink)",
		"PowerCor Australia",
		"Snowy Hydro Ltd",
		"Jemena  Electricity Networks (VIC)",
		"TasNetworks",
		"TransGrid",
		"SA Power Networks",
		"United Energy Distribution",
		"Western Power",
	}

	participantShortNames = []string{ //nolint:gochecknoglobals
		"UNDEFINED",
		"Evoenergy",
		"AEMO Reserved",
		"TasNetworks",
		"CitiPower",
		"Essential Energy",
		"SP AusNet DNSP",
		"Energex",
		"Ausgrid",
		"Ergon Energy",
		"ElectraNet SA",
		"Exempt Networks - various",
		"Sydney Airport",
		"GAS NSW",
		"GAS QLD",
		"GAS SA",
		"GAS TAS",
		"GAS VIC",
		"GAS WA",
		"SP AusNet TNSP",
		"Horizon Power",
		"Endeavour Energy",
		"Northern Territory Reserved Block",
		"Powerlink",
		"PowerCor",
		"Snowy Hydro",
		"Jemena",
		"TasNetworks",
		"TransGrid",
		"SA Power Networks",
		"United Energy",
		"Western Power",
	}

	participantEnergies = []Energy{ //nolint:gochecknoglobals
		EnergyUndefined,
		EnergyElectricity,
		EnergyElectricity,
		EnergyElectricity,
		EnergyElectricity,
		EnergyElectricity,
		EnergyElectricity,
		EnergyElectricity,
		EnergyElectricity,
		EnergyElectricity,
		EnergyElectricity,
		EnergyElectricity,
		EnergyElectricity,
		EnergyGas,
		EnergyGas,
		EnergyGas,
		EnergyGas,
		EnergyGas,
		EnergyGas,
		EnergyElectricity,
		EnergyElectricity,
		EnergyElectricity,
		EnergyElectricity,
		EnergyElectricity,
		EnergyElectricity,
		EnergyElectricity,
		EnergyElectricity,
		EnergyElectricity,
		EnergyElectricity,
		EnergyElectricity,
		EnergyElectricity,
		EnergyElectricity,
	}

	participantAllocations = []Allocations{ //nolint:gochecknoglobals
		allocationsUndefined,
		allocationsACTEWP,
		allocationsAEMORESERVED,
		allocationsAURORAP,
		allocationsCITIPP,
		allocationsCNRGYP,
		allocationsEASTERN,
		allocationsENERGEXP,
		allocationsENERGYAP,
		allocationsERGONETP,
		allocationsETSATP,
		allocationsEXEMPTNETWORKS,
		allocationsFEDAIRPORTS,
		allocationsGASNSW,
		allocationsGASQLD,
		allocationsGASSA,
		allocationsGASTAS,
		allocationsGASVIC,
		allocationsGASWA,
		allocationsGPUPP,
		allocationsHORIZONPOWER,
		allocationsINTEGP,
		allocationsNTRESERVED,
		allocationsPLINKP,
		allocationsPOWCP,
		allocationsSNOWY,
		allocationsSOLARISP,
		allocationsTRANSEND,
		allocationsTRANSGP,
		allocationsUMPLP,
		allocationsUNITED,
		allocationsWESTERNPOWER,
	}
)

// Participant is an index for the participant.
type Participant int32

// ParticipantInfo provides structure for participant information.
type ParticipantInfo struct {
	Participant   Participant   `json:"participant"`
	ParticipantID string        `json:"participantID"` //nolint:tagliatelle
	Region        region.Region `json:"region"`
	LongName      string        `json:"longName"`
	ShortName     string        `json:"shortName"`
	Energy        Energy        `json:"energy"`
	Allocations   Allocations   `json:"allocations"`
}

// Participants returns all valid participants.
func Participants() []Participant {
	return participants
}

// ParticipantsForRegion returns the pariticipants that operate in a region.
func ParticipantsForRegion(r region.Region) []Participant {
	resp := []Participant{}

	for _, p := range participants {
		if p.Region() == r {
			resp = append(resp, p)
		}
	}

	return resp
}

// NewParticipant returns a participant for a string (AEMO Participant ID).
func NewParticipant(s string) (Participant, error) {
	p, ok := ParticipantValue[strings.ToUpper(s)]
	if !ok {
		return ParticipantUndefined, ErrParticipantInvalid
	}

	return p, nil
}

// Valid returns true if a valid participant.
func (p Participant) Valid() bool {
	if _, ok := ParticipantName[p]; !ok {
		return false
	}

	return true
}

// GoString meets the gostring interface.
func (p Participant) GoString() string {
	if !p.Valid() {
		return ParticipantUndefined.GoString()
	}

	return fmt.Sprintf(
		"Participant{Participant: %d, ParticipantID: %q, Region: %#v, LongName: %q, ShortName: %q, Energy: %q, Allocations: %#v}",
		p, p.ParticipantID(), p.Region(), p.LongName(), p.ShortName(), p.Energy().String(), p.Allocations(),
	)
}

// String meets the stringer interface.
func (p Participant) String() string {
	if !p.Valid() {
		return ParticipantUndefined.String()
	}

	return p.ParticipantID()
}

// Info struct for a participant.
func (p Participant) Info() (*ParticipantInfo, error) {
	if !p.Valid() || p == ParticipantUndefined {
		return nil, fmt.Errorf("participant '%d': %w", p, ErrParticipantInvalid)
	}

	return &ParticipantInfo{
		Participant:   p,
		ParticipantID: p.ParticipantID(),
		Region:        p.Region(),
		LongName:      p.LongName(),
		ShortName:     p.ShortName(),
		Energy:        p.Energy(),
		Allocations:   p.Allocations(),
	}, nil
}

// ParticipantID returns the Participant ID.
func (p Participant) ParticipantID() string {
	return ParticipantIDs[p]
}

// Region returns the Participant's Region.
func (p Participant) Region() region.Region {
	return participantRegions[p]
}

// LongName returns the long (full) name of the participant.
func (p Participant) LongName() string {
	return participantLongNames[p]
}

// ShortName returns the short (useful) name of the participant.
func (p Participant) ShortName() string {
	return participantShortNames[p]
}

// Energy returns the energy type of the participant's allocations.
func (p Participant) Energy() Energy {
	return participantEnergies[p]
}

// Allocations returns the participant's allocations.
func (p Participant) Allocations() Allocations {
	return participantAllocations[p]
}
