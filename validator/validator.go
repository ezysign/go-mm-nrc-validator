package validator

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

type NRCCode struct {
	NRCCodeData []NRCCodeData `json:"data"`
}

type NRCCodeData struct {
	ID              int    `json:"ID"`
	NrcTownshipCode string `json:"nrc_township_code"`
	NrcTownshipName string `json:"nrc_township_name"`
	NrcSrCode       string `json:"nrc_sr_code"`
}

type TownShip struct {
	TownShipData []TownShipData `json:"data"`
}

type TownShipData struct {
	SRCode  string `json:"SR_Code"`
	SRName  string `json:"SR_Name"`
	TspCode string `json:"Tsp_Code"`
	TspName string `json:"Tsp_Name"`
}

type NRCInfo struct {
	TspCode           string
	TspName           string
	SRCodeMM          string
	SRCodeEN          string
	CitizenShipStatus string
}

type nrcvalidator interface {
	GetMatchNRC(nrc string) []string
	ValidateNRC(nrc string) bool
	VerifyNRC(nrc string) bool
	ExtractNrcInfo(nrc string) NRCInfo
	GetTownShipData() []TownShipData
	GetNRCData() []NRCCodeData
}

type validator struct {
	NrcData      []NRCCodeData
	TownShipData []TownShipData
}

func NewNRCValidator() nrcvalidator {
	var nrccodes NRCCode
	json.Unmarshal([]byte(nrcCodeJSON), &nrccodes)
	var township TownShip
	json.Unmarshal([]byte(townShipJSON), &township)
	return &validator{
		NrcData:      nrccodes.NRCCodeData,
		TownShipData: township.TownShipData,
	}
}

func (v *validator) GetTownShipData() []TownShipData {
	return v.TownShipData
}
func (v *validator) GetNRCData() []NRCCodeData {
	return v.NrcData
}

func (v *validator) GetMatchNRC(nrc string) []string {
	re := regexp.MustCompile(mmNRCformatRegex)
	if v.ValidateNRC(strings.Trim(nrc, "")) == false {
		return make([]string, 0)
	}
	matchArr := re.FindAllStringSubmatch(strings.Trim(nrc, ""), -1)[0]

	return matchArr
}
func (v *validator) ValidateNRC(nrc string) bool {

	re := regexp.MustCompile(mmNRCformatRegex)
	return len(re.FindAllStringSubmatch(strings.Trim(nrc, ""), -1)) > 0
}
func (v *validator) VerifyNRC(nrc string) bool {
	if v.ValidateNRC(strings.Trim(nrc, "")) == false {
		return false
	}

	matchArr := v.GetMatchNRC(nrc)
	nrcSRCodebackSlash := fmt.Sprintf("%s/", matchArr[1])
	townShipCode := fmt.Sprintf("%s", matchArr[2])
	nrcSRCode := fmt.Sprintf("%s/", matchArr[1])
	nrcSrEnglish := convertBurmeseNum2En(nrcSRCode)

	nrcCodeIndex := searchNrcCodeIndexBySrcTownShipCode(nrcSRCodebackSlash, townShipCode)

	if nrcCodeIndex == -1 {
		return false
	}

	townshipIndex := searchTownshipIndexBySRCode(nrcSrEnglish)
	if townshipIndex == -1 {
		return false
	}
	return true
}
func (v *validator) ExtractNrcInfo(nrc string) NRCInfo {
	var nrcInfo NRCInfo
	validator := NewNRCValidator()
	if validator.ValidateNRC(nrc) == true {
		matchArr := validator.GetMatchNRC(nrc)
		if len(matchArr) > 0 {
			nrcSRCodebackSlash := fmt.Sprintf("%s/", matchArr[1])
			townShipCode := fmt.Sprintf("%s", matchArr[2])
			nrcCodeIndex := searchNrcCodeIndexBySrcTownShipCode(nrcSRCodebackSlash, townShipCode)

			if nrcCodeIndex == -1 {
				fmt.Println("Error")
			}
			nrcCodeData := validator.GetNRCData()[nrcCodeIndex]

			nrcInfo = NRCInfo{
				CitizenShipStatus: matchArr[3],
				TspCode:           fmt.Sprintf("%s", matchArr[2]),
				TspName:           nrcCodeData.NrcTownshipName,
				SRCodeMM:          matchArr[1],
				SRCodeEN:          convertBurmeseNum2En(nrcCodeData.NrcSrCode),
			}
			return nrcInfo

		}
	}
	return nrcInfo
}

func convertBurmeseNum2En(code string) string {
	var number string = ""
	for _, v := range strings.Split(code, "") {
		number += mmNum2En[v]

	}
	return number
}

func searchNrcCodeIndexBySrcTownShipCode(nrcSRCodebackSlash string, townshipCode string) int {
	var codeData int
	validator := NewNRCValidator()
	codeData = -1
	for index, code := range validator.GetNRCData() {
		if code.NrcSrCode == nrcSRCodebackSlash && code.NrcTownshipCode == townshipCode {
			codeData = index
			break
		}
	}
	return codeData
}

func searchTownshipIndexBySRCode(SRCode string) int {
	var codeData int
	validator := NewNRCValidator()
	codeData = -1
	for index, code := range validator.GetTownShipData() {
		if code.SRCode == SRCode {
			codeData = index
			break
		}
	}
	return codeData
}
