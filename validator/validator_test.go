package validator

import (
	"fmt"
	"testing"
)

func TestVerifyNRC(t *testing.T) {
	validator := NewNRCValidator()

	if validator.VerifyNRC("၄/ကပလ(နိုင်)၀၉၄၁၈၆") != true {
		t.Log("it Should be correct tsp code")
		t.Fail()
	}

	if validator.VerifyNRC("၁၂/ကပလ(နိုင်)၀၉၄၁၈၆") == true {
		t.Log("it Should be correct tsp code")
		t.Fail()
	}
}

func TestValidator(t *testing.T) {
	validator := NewNRCValidator()

	if validator.ValidateNRC("၄/ဝကမ(နိုင်)၀၉၄၁၈၆") != true {
		t.Log("it Should be correct NRC")
		t.Fail()
	}

	if validator.ValidateNRC("၁၂/မဂတ(နိုင်)၀၉၄၁၈၆") != true {
		t.Log("it Should be correct NRC Format")
		t.Fail()
	}
}

func TestExtractor(t *testing.T) {
	validator := NewNRCValidator()
	fmt.Println(validator.ExtractNrcInfo("၁၂/မဂတ(နိုင်)၀၉၄၁၈၆"))
	nrcInfo := NRCInfo{
		SRCodeEN:          "12",
		SRCodeMM:          "၁၂",
		CitizenShipStatus: "(နိုင်)",
		TspName:           "မင်္ဂလာတောင်ညွန့်",
		TspCode:           "မဂတ",
	}
	if nrcInfo != validator.ExtractNrcInfo("၁၂/မဂတ(နိုင်)၀၉၄၁၈၆") {
		t.Log("It should be correct NRC Info")
		t.Fail()
	}
	nullNrcInfo := NRCInfo{}
	if nullNrcInfo.TspName == validator.ExtractNrcInfo("၁၂/မဂတ(နိုင်)၀၉၄၁၈၆").TspName {
		t.Log("It should be empty response")
		t.Fail()
	}
}
