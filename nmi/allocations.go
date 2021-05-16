package nmi

import (
	"fmt"
	"regexp"
)

var (
	allocationsUndefined = Allocations{}

	allocationsACTEWP = Allocations{
		Pattern(`^(NGGG[A-HJ-NP-VX-Z\d][A-HJ-NP-Z\d]{5})$`),
		Pattern(`^(7001\d{6})$`),
	}

	allocationsAEMORESERVED = Allocations{
		Pattern(`^(880[1-5]\d{6})$`),
		Pattern(`^(9\d{9})$`),
	}

	allocationsAURORAP = Allocations{
		Pattern(`^(T000000(([0-4]\d{3})|(500[01])))$`),
		Pattern(`^(8000\d{6})$`),
		Pattern(`^(8590[23]\d{5})$`),
	}

	allocationsCITIPP = Allocations{
		Pattern(`^(VAAA[A-HJ-NP-VX-Z\d][A-HJ-NP-Z\d]{5})$`),
		Pattern(`^(610[23]\d{6})$`),
	}

	allocationsCNRGYP = Allocations{
		Pattern(`^(NAAA[A-HJ-NP-VX-Z\d][A-HJ-NP-Z\d]{5})$`),
		Pattern(`^(NBBB[A-HJ-NP-VX-Z\d][A-HJ-NP-Z\d]{5})$`),
		Pattern(`^(NDDD[A-HJ-NP-VX-Z\d][A-HJ-NP-Z\d]{5})$`),
		Pattern(`^(NFFF[A-HJ-NP-VX-Z\d][A-HJ-NP-Z\d]{5})$`),
		Pattern(`^(4001\d{6})$`),
		Pattern(`^(4508\d{6})$`),
		Pattern(`^(4204\d{6})$`),
		Pattern(`^(4407\d{6})$`),
	}

	allocationsEASTERN = Allocations{
		Pattern(`^(VBBB[A-HJ-NP-VX-Z\d][A-HJ-NP-Z\d]{5})$`),
		Pattern(`^(630[56]\d{6})$`),
	}

	allocationsENERGEXP = Allocations{
		Pattern(`^(QB\d{2}[A-HJ-NP-VX-Z\d][A-HJ-NP-Z\d]{5})$`),
		Pattern(`^(31\d{8})$`),
	}

	allocationsENERGYAP = Allocations{
		Pattern(`^(NCCC[A-HJ-NP-VX-Z\d][A-HJ-NP-Z\d]{5})$`),
		Pattern(`^(410[234]\d{6})$`),
	}

	allocationsERGONETP = Allocations{
		Pattern(`^(QAAA[A-HJ-NP-VX-Z\d][A-HJ-NP-Z\d]{5})$`),
		Pattern(`^(QCCC[A-HJ-NP-VX-Z\d][A-HJ-NP-Z\d]{5})$`),
		Pattern(`^(QDDD[A-HJ-NP-VX-Z\d][A-HJ-NP-Z\d]{5})$`),
		Pattern(`^(QEEE[A-HJ-NP-VX-Z\d][A-HJ-NP-Z\d]{5})$`),
		Pattern(`^(QFFF[A-HJ-NP-VX-Z\d][A-HJ-NP-Z\d]{5})$`),
		Pattern(`^(QGGG[A-HJ-NP-VX-Z\d][A-HJ-NP-Z\d]{5})$`),
		Pattern(`^(30\d{8})$`),
	}

	allocationsETSATP = Allocations{
		Pattern(`^(S[A-HJ-NP-Z\d]{3}W[A-HJ-NP-Z\d]{5})$`),
		Pattern(`^(210200\d{4})$`),
	}

	allocationsEXEMPTNETWORKS = Allocations{
		Pattern(`^(NKKK[A-HJ-NP-VX-Z\d][A-HJ-NP-Z\d]{5})$`),
		Pattern(`^(7102\d{6})$`),
	}

	allocationsFEDAIRPORTS = Allocations{
		Pattern(`^(NJJJNR[A-HJ-NP-Z\d]{4})$`),
	}

	allocationsGAS_NSW = Allocations{
		Pattern(`^(52\d{8})$`),
	}

	allocationsGAS_QLD = Allocations{
		Pattern(`^(54\d{8})$`),
	}

	allocationsGAS_SA = Allocations{
		Pattern(`^(55\d{8})$`),
	}

	allocationsGAS_TAS = Allocations{
		Pattern(`^(57\d{8})$`),
	}

	allocationsGAS_VIC = Allocations{
		Pattern(`^(53\d{8})$`),
	}

	allocationsGAS_WA = Allocations{
		Pattern(`^(56\d{8})$`),
	}

	allocationsGPUPP = Allocations{
		Pattern(`^(V[A-HJ-NP-Z\d]{3}W[A-HJ-NP-Z\d]{5})$`),
		Pattern(`^(650900\d{4})$`),
	}

	allocationsHORIZONPOWER = Allocations{
		Pattern(`^(8021\d{6})$`),
	}

	allocationsINTEGP = Allocations{
		Pattern(`^(NEEE[A-HJ-NP-VX-Z\d][A-HJ-NP-Z\d]{5})$`),
		Pattern(`^(431\d{7})$`),
	}

	allocationsNT_RESERVED = Allocations{
		Pattern(`^(250\d{7})$`),
	}

	allocationsPLINKP = Allocations{
		Pattern(`^(Q[A-HJ-NP-Z\d]{3}W[A-HJ-NP-Z\d]{5})$`),
		Pattern(`^(320200\d{4})$`),
	}

	allocationsPOWCP = Allocations{
		Pattern(`^(VCCC[A-HJ-NP-VX-Z\d][A-HJ-NP-Z\d]{5})$`),
		Pattern(`^(620[34]\d{6})$`),
	}

	allocationsSNOWY = Allocations{
		Pattern(`^(4708109\d{3})$`),
	}

	allocationsSOLARISP = Allocations{
		Pattern(`^(VDDD[A-HJ-NP-VX-Z\d][A-HJ-NP-Z\d]{5})$`),
		Pattern(`^(6001\d{6})$`),
	}

	allocationsTRANSEND = Allocations{
		Pattern(`^(T[A-HJ-NP-Z\d]{3}W[A-HJ-NP-Z\d]{5})$`),
	}

	allocationsTRANSGP = Allocations{
		Pattern(`^(NTTT[A-HJ-NP-VX-Z\d][A-HJ-NP-Z\d]{5})$`),
		Pattern(`^(460810[0-8]\d{3})$`),
	}

	allocationsUMPLP = Allocations{
		Pattern(`^(SAAA[A-HJ-NP-VX-Z\d][A-HJ-NP-Z\d]{5})$`),
		Pattern(`^(SASMPL[\d]{4})$`),
		Pattern(`^(200[12]\d{6})$`),
	}

	allocationsUNITED = Allocations{
		Pattern(`^(VEEE[A-HJ-NP-VX-Z\d][A-HJ-NP-Z\d]{5})$`),
		Pattern(`^(640[78]\d{6})$`),
	}

	allocationsWESTERNPOWER = Allocations{
		Pattern(`^(WAAA[A-HJ-NP-VX-Z\d][A-HJ-NP-Z\d]{5})$`),
		Pattern(`^(800[1-9]\d{6})$`),
		Pattern(`^(801\d{7})$`),
		Pattern(`^(8020\d{6})$`),
	}
)

// Allocations is an allocation of the NMI
type Allocations []Pattern

// Compile compiles all patterns.
func (a Allocations) Compile() ([]*regexp.Regexp, error) {
	resp := make([]*regexp.Regexp, len(a))

	for i, p := range a {
		pc, err := p.Compile()
		if err != nil {
			return nil, fmt.Errorf("pattern[%d]: '%s': %w", i, p.String(), ErrPatternInvalid)
		}

		resp[i] = pc
	}

	return resp, nil
}
