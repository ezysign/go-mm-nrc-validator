package validator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

var mmNRCformatRegex = `([၀-၉]{1,2})/([ကခဂဃငစဆဇဈညဎဏတထဒဓနပဖဗဘမယရလဝသဟဠဥဧ]{3})(\([နိုင်,ဧည့်,ပြု,စ,သ,သီ]{1,6}\))([၀-၉]{6})$`
var mmNum2En = map[string]string{
	"၀": "0",
	"၁": "1",
	"၂": "2",
	"၃": "3",
	"၄": "4",
	"၅": "5",
	"၆": "6",
	"၇": "7",
	"၈": "8",
	"၉": "9",
}

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

type nrcvalidator interface {
	GetMatchNRC(nrc string) []string
	ValidateNRC(nrc string) bool
	VerifyNRC(nrc string) bool
	ExtractNrcInfo(nrc string) string
	GetTownShipData() []TownShipData
	GetNRCData() []NRCCodeData
}

type validator struct {
	NrcData      []NRCCodeData
	TownShipData []TownShipData
}

func NewNRCValidator() nrcvalidator {
	var nrccodes NRCCode
	json.Unmarshal(parseJson("nrccode.json"), &nrccodes)
	var township TownShip
	json.Unmarshal(parseJson("township.json"), &township)
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

func parseJson(filename string) []byte {
	jsonFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}

func (v *validator) GetMatchNRC(nrc string) []string {
	// re := regexp.MustCompile(mmNRCformatRegex)
	if v.ValidateNRC(strings.Trim(nrc, "")) == false {
		return make([]string, 0)
	}
	// matchArr := re.FindAllStringSubmatch(strings.Trim(nrc, ""), -1)[0]
	// citizenShipStatus := matchArr[3]
	// nrcSRCodebackSlash := fmt.Sprintf("%s/", matchArr[1])
	// nrcSRCode := fmt.Sprintf("%s/", matchArr[1])

	// re.FindAllStringSubmatch(strings.Trim(nrc, ""), -1)[0]

	return make([]string, 3)
}
func (v *validator) ValidateNRC(nrc string) bool {

	re := regexp.MustCompile(mmNRCformatRegex)
	return len(re.FindAllStringSubmatch(strings.Trim(nrc, ""), -1)) > 0
}
func (v *validator) VerifyNRC(nrc string) bool {
	return false
}
func (v *validator) ExtractNrcInfo(nrc string) string {
	return ""
}
