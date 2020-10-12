package validator

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestBracketSanitizer(t *testing.T) {
	validator := NewNRCValidator()

	for _, nrc := range validator.GetNRCData() {
		fmt.Println(nrc.ID)
	}
	fmt.Println("done")
	// utils := NewNRCConverter()
	// if utils.sanitizeBrackets("(") != `\(` {
	// 	t.Log("it Should return SanitizedRegex for bracket (")
	// 	t.Fail()
	// }

	// if utils.sanitizeBrackets(")") != `\)` {
	// 	t.Log("it Should return SanitizedRegex for bracket ) ")
	// 	t.Fail()
	// }
}

func convertBurmeseNum2En(code string) string {
	var number string = ""
	for _, v := range strings.Split(code, "") {
		number += mmNum2En[v]

	}
	return number
}

func TestValidator(t *testing.T) {
	var mmNRCformatRegex = `([၀-၉]{1,2})/([ကခဂဃငစဆဇဈညဎဏတထဒဓနပဖဗဘမယရလဝသဟဠဥဧ]{3})(\([နိုင်,ဧည့်,ပြု,စ,သ,သီ]{1,6}\))([၀-၉]{6})$`
	re := regexp.MustCompile(mmNRCformatRegex)
	validator := NewNRCValidator()
	fmt.Println(re.FindAllStringSubmatch(strings.Trim(" ၁၂/မဂတ(နိုင်)၀၉၄၁၈၆", ""), -1)[0][1])
	matchArr := re.FindAllStringSubmatch(strings.Trim(" ၁၂/မဂတ(နိုင်)၀၉၄၁၈၆", ""), -1)[0]

	citizenShipStatus := matchArr[3]
	nrcSRCodebackSlash := fmt.Sprintf("%s/", matchArr[1])
	nrcSRCode := fmt.Sprintf("%s/", matchArr[1])
	townShipCode := fmt.Sprintf("%s", matchArr[2])
	nrcSrEnglish := convertBurmeseNum2En(nrcSRCode)

	nrcCodeIndex := searchNrcCodeIndexBySrcTownShipCode(nrcSRCodebackSlash, townShipCode)

	if nrcCodeIndex == -1 {
		fmt.Println("Error")
	}
	nrcCodeData := validator.GetNRCData()[nrcCodeIndex]
	fmt.Println(nrcCodeData)

	townshipIndex := searchTownshipIndexBySRCode(nrcSrEnglish, nrcCodeData.NrcTownshipName)
	if townshipIndex == -1 {
		fmt.Println("Error")
	}
	townShip := validator.GetTownShipData()[townshipIndex]
	fmt.Println(townShip)
	fmt.Println(citizenShipStatus)
	fmt.Println("Done")
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

func searchTownshipIndexBySRCode(SRCode string, townshipName string) int {
	var codeData int
	validator := NewNRCValidator()
	codeData = -1
	for index, code := range validator.GetTownShipData() {
		if code.SRCode == SRCode && code.TspName == strings.Trim(townshipName, "") {
			codeData = index
			break
		}
	}
	return codeData
}
