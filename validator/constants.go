package validator

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

const nrcCodeJSON = `{
	"data": [
	  {
		"ID": 58,
		"nrc_township_code": "၀သန",
		"nrc_township_name": "ဝန်းသို",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 224,
		"nrc_township_code": "ကကက",
		"nrc_township_name": "ကိုကိုးကျွန်း",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 296,
		"nrc_township_code": "ကကထ",
		"nrc_township_name": "ကန်ကြီးထောင့်",
		"nrc_sr_code": "၁၄/"
	  },
	  {
		"ID": 100,
		"nrc_township_code": "ကကန",
		"nrc_township_name": "ကျောက်ကြီး",
		"nrc_sr_code": "၇/"
	  },
	  {
		"ID": 293,
		"nrc_township_code": "ကကန",
		"nrc_township_name": "ကုန်းကြမ်း",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 300,
		"nrc_township_code": "ကကန",
		"nrc_township_name": "ကျောင်းကုန်း",
		"nrc_sr_code": "၁၄/"
	  },
	  {
		"ID": 31,
		"nrc_township_code": "ကကရ",
		"nrc_township_name": "ကော့ကရိတ်",
		"nrc_sr_code": "၃/"
	  },
	  {
		"ID": 223,
		"nrc_township_code": "ကခက",
		"nrc_township_name": "ကွမ်းခြံကုန်း",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 268,
		"nrc_township_code": "ကခန",
		"nrc_township_name": "ကွတ်ခိုင်",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 308,
		"nrc_township_code": "ကခန",
		"nrc_township_name": "ကြံခင်း",
		"nrc_sr_code": "၁၄/"
	  },
	  {
		"ID": 85,
		"nrc_township_code": "ကစန",
		"nrc_township_name": "ကျွန်းစု",
		"nrc_sr_code": "၆/"
	  },
	  {
		"ID": 32,
		"nrc_township_code": "ကဆက",
		"nrc_township_name": "ကြာအင်းဆိပ်ကြီး",
		"nrc_sr_code": "၃/"
	  },
	  {
		"ID": 151,
		"nrc_township_code": "ကဆန",
		"nrc_township_name": "ကျောက်ဆည်",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 95,
		"nrc_township_code": "ကတခ",
		"nrc_township_name": "ကျောက်တံခါး",
		"nrc_sr_code": "၇/"
	  },
	  {
		"ID": 233,
		"nrc_township_code": "ကတတ",
		"nrc_township_name": "ကျောက်တံတား",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 194,
		"nrc_township_code": "ကတန",
		"nrc_township_name": "ကျောက်တော်",
		"nrc_sr_code": "၁၁/"
	  },
	  {
		"ID": 222,
		"nrc_township_code": "ကတန",
		"nrc_township_name": "ကျောက်တန်း",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 255,
		"nrc_township_code": "ကတန",
		"nrc_township_name": "ကျိုင်းတုံ",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 181,
		"nrc_township_code": "ကထန",
		"nrc_township_name": "ကျိုက်ထို",
		"nrc_sr_code": "၁၀/"
	  },
	  {
		"ID": 68,
		"nrc_township_code": "ကနန",
		"nrc_township_name": "ကနီ",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 117,
		"nrc_township_code": "ကပက",
		"nrc_township_name": "ကြို့ပင်ကောက်",
		"nrc_sr_code": "၇/"
	  },
	  {
		"ID": 164,
		"nrc_township_code": "ကပတ",
		"nrc_township_name": "ကျောက်ပန်းတောင်း",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 297,
		"nrc_township_code": "ကပန",
		"nrc_township_name": "ကျုံပျော်",
		"nrc_sr_code": "၁၄/"
	  },
	  {
		"ID": 38,
		"nrc_township_code": "ကပလ",
		"nrc_township_name": "ကန်ပက်လက်",
		"nrc_sr_code": "၄/"
	  },
	  {
		"ID": 187,
		"nrc_township_code": "ကဖန",
		"nrc_township_name": "ကျောက်ဖြူ",
		"nrc_sr_code": "၁၁/"
	  },
	  {
		"ID": 77,
		"nrc_township_code": "ကဘလ",
		"nrc_township_name": "ကန့်ဘလူ",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 234,
		"nrc_township_code": "ကမတ",
		"nrc_township_name": "ကြည့်မြင်တိုင်",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 9,
		"nrc_township_code": "ကမန",
		"nrc_township_name": "ဖားကန့်",
		"nrc_sr_code": "၁/"
	  },
	  {
		"ID": 135,
		"nrc_township_code": "ကမန",
		"nrc_township_name": "ကံမ",
		"nrc_sr_code": "၈/"
	  },
	  {
		"ID": 225,
		"nrc_township_code": "ကမန",
		"nrc_township_name": "ကော့မှူး",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 286,
		"nrc_township_code": "ကမန",
		"nrc_township_name": "ကျောက်မဲ",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 176,
		"nrc_township_code": "ကမရ",
		"nrc_township_name": "ကျိုက်မရော",
		"nrc_sr_code": "၁၀/"
	  },
	  {
		"ID": 232,
		"nrc_township_code": "ကမရ",
		"nrc_township_name": "ကမာရွတ်",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 57,
		"nrc_township_code": "ကလတ",
		"nrc_township_name": "ကောလင်း",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 50,
		"nrc_township_code": "ကလထ",
		"nrc_township_name": "ကလေး",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 78,
		"nrc_township_code": "ကလန",
		"nrc_township_name": "ကျွန်းလှ",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 247,
		"nrc_township_code": "ကလန",
		"nrc_township_name": "ကလော",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 283,
		"nrc_township_code": "ကလန",
		"nrc_township_name": "ကွမ်းလုံ",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 318,
		"nrc_township_code": "ကလန",
		"nrc_township_name": "ကျိုက်လတ်",
		"nrc_sr_code": "၁၄/"
	  },
	  {
		"ID": 51,
		"nrc_township_code": "ကလဝ",
		"nrc_township_name": "ကလေးဝ",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 92,
		"nrc_township_code": "ကဝန",
		"nrc_township_name": "ကဝ",
		"nrc_sr_code": "၇/"
	  },
	  {
		"ID": 53,
		"nrc_township_code": "ကသန",
		"nrc_township_name": "ကသာ",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 88,
		"nrc_township_code": "ကသန",
		"nrc_township_name": "ကော့သောင်း",
		"nrc_sr_code": "၆/"
	  },
	  {
		"ID": 279,
		"nrc_township_code": "ကသန",
		"nrc_township_name": "ကျေးသီး",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 275,
		"nrc_township_code": "ကဟန",
		"nrc_township_name": "ကွန်ဟိန်း",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 175,
		"nrc_township_code": "ခဆန",
		"nrc_township_name": "ချောင်းဆုံ",
		"nrc_sr_code": "၁၀/"
	  },
	  {
		"ID": 60,
		"nrc_township_code": "ခတန",
		"nrc_township_name": "ခန္တီး",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 3,
		"nrc_township_code": "ခဖန",
		"nrc_township_name": "ချီဖွေ",
		"nrc_sr_code": "၁/"
	  },
	  {
		"ID": 18,
		"nrc_township_code": "ခဘဒ",
		"nrc_township_name": "ခေါင်လန်ဖူး",
		"nrc_sr_code": "၁/"
	  },
	  {
		"ID": 146,
		"nrc_township_code": "ခမစ",
		"nrc_township_name": "ချမ်းမြသာစည်",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 120,
		"nrc_township_code": "ခမန",
		"nrc_township_name": "ချောက်",
		"nrc_sr_code": "၈/"
	  },
	  {
		"ID": 226,
		"nrc_township_code": "ခရန",
		"nrc_township_name": "ခရမ်း",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 145,
		"nrc_township_code": "ခအဇ",
		"nrc_township_name": "ချမ်းအေးသာဇံ",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 45,
		"nrc_township_code": "ခဥတ",
		"nrc_township_name": "ချောင်းဦး",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 72,
		"nrc_township_code": "ခဥန",
		"nrc_township_name": "ခင်ဦး",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 129,
		"nrc_township_code": "ဂဂန",
		"nrc_township_name": "ဂန့်ဂေါ",
		"nrc_sr_code": "၈/"
	  },
	  {
		"ID": 197,
		"nrc_township_code": "ဂမန",
		"nrc_township_name": "ဂွ",
		"nrc_sr_code": "၁၁/"
	  },
	  {
		"ID": 161,
		"nrc_township_code": "ငဇန",
		"nrc_township_name": "ငါန်းဇွန်",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 298,
		"nrc_township_code": "ငပတ",
		"nrc_township_name": "ငပုတော",
		"nrc_sr_code": "၁၄/"
	  },
	  {
		"ID": 126,
		"nrc_township_code": "ငဖန",
		"nrc_township_name": "ငဖဲ",
		"nrc_sr_code": "၈/"
	  },
	  {
		"ID": 153,
		"nrc_township_code": "စကတ",
		"nrc_township_name": "စဉ့်ကိုင်",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 47,
		"nrc_township_code": "စကန",
		"nrc_township_name": "စစ်ကိုင်း",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 168,
		"nrc_township_code": "စကန",
		"nrc_township_name": "စဉ့်ကူး",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 235,
		"nrc_township_code": "စခန",
		"nrc_township_name": "စမ်းချောင်း",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 183,
		"nrc_township_code": "စတန",
		"nrc_township_name": "စစ်တွေ",
		"nrc_sr_code": "၁၁/"
	  },
	  {
		"ID": 128,
		"nrc_township_code": "စတရ",
		"nrc_township_name": "စေတုတ္ထရာ",
		"nrc_sr_code": "၈/"
	  },
	  {
		"ID": 127,
		"nrc_township_code": "စလန",
		"nrc_township_name": "စလင်း",
		"nrc_sr_code": "၈/"
	  },
	  {
		"ID": 227,
		"nrc_township_code": "ဆကခ",
		"nrc_township_name": "ဆိပ်ကြီးခနောင်တို",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 236,
		"nrc_township_code": "ဆကန",
		"nrc_township_name": "ဆိပ်ကမ်း",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 246,
		"nrc_township_code": "ဆဆန",
		"nrc_township_name": "ဆီဆိုင်",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 15,
		"nrc_township_code": "ဆပဘ",
		"nrc_township_name": "ဆွမ်ပရာဘွမ်",
		"nrc_sr_code": "၁/"
	  },
	  {
		"ID": 138,
		"nrc_township_code": "ဆပဝ",
		"nrc_township_name": "ဆင်ပေါင်ဝဲ",
		"nrc_sr_code": "၈/"
	  },
	  {
		"ID": 143,
		"nrc_township_code": "ဆဖန",
		"nrc_township_name": "ဆိပ်ဖြူ",
		"nrc_sr_code": "၈/"
	  },
	  {
		"ID": 131,
		"nrc_township_code": "ဆမန",
		"nrc_township_name": "ဆော",
		"nrc_sr_code": "၈/"
	  },
	  {
		"ID": 69,
		"nrc_township_code": "ဆလက",
		"nrc_township_name": "ဆားလင်းကြီး",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 4,
		"nrc_township_code": "ဆလန",
		"nrc_township_name": "ဆော့လော်",
		"nrc_sr_code": "၁/"
	  },
	  {
		"ID": 114,
		"nrc_township_code": "ဇကန",
		"nrc_township_name": "ဇီးကုန်း",
		"nrc_sr_code": "၇/"
	  },
	  {
		"ID": 324,
		"nrc_township_code": "ဇဗသ",
		"nrc_township_name": "ဇမ္ဗုသီရိ",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 328,
		"nrc_township_code": "ဇယသ",
		"nrc_township_name": "ဇေယျာသီရိ",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 307,
		"nrc_township_code": "ဇလန",
		"nrc_township_name": "ဇလွန်",
		"nrc_sr_code": "၁၄/"
	  },
	  {
		"ID": 305,
		"nrc_township_code": "ညတန",
		"nrc_township_name": "ညောင်တုန်း",
		"nrc_sr_code": "၁၄/"
	  },
	  {
		"ID": 249,
		"nrc_township_code": "ညရန",
		"nrc_township_name": "ညောင်ရွှေ",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 94,
		"nrc_township_code": "ညလပ",
		"nrc_township_name": "ညောင်လေးပင်",
		"nrc_sr_code": "၇/"
	  },
	  {
		"ID": 163,
		"nrc_township_code": "ညဥန",
		"nrc_township_name": "ညောင်ဦး",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 196,
		"nrc_township_code": "တကန",
		"nrc_township_name": "တောင်ကုတ်",
		"nrc_sr_code": "၁၁/"
	  },
	  {
		"ID": 214,
		"nrc_township_code": "တကန",
		"nrc_township_name": "တိုက်ကြီး",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 245,
		"nrc_township_code": "တကန",
		"nrc_township_name": "တောင်ကြီး",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 327,
		"nrc_township_code": "တကန",
		"nrc_township_name": "တပ်ကုန်း",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 259,
		"nrc_township_code": "တခလ",
		"nrc_township_name": "တာချီလိတ်",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 98,
		"nrc_township_code": "တငန",
		"nrc_township_name": "တောင်ငူ",
		"nrc_sr_code": "၇/"
	  },
	  {
		"ID": 76,
		"nrc_township_code": "တဆန",
		"nrc_township_name": "တန့်ဆည်",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 36,
		"nrc_township_code": "တဇန",
		"nrc_township_name": "တွန်းဇန်",
		"nrc_sr_code": "၄/"
	  },
	  {
		"ID": 121,
		"nrc_township_code": "တတက",
		"nrc_township_name": "တောင်တွင်းကြီး",
		"nrc_sr_code": "၈/"
	  },
	  {
		"ID": 35,
		"nrc_township_code": "တတန",
		"nrc_township_name": "တီးတိန်",
		"nrc_sr_code": "၄/"
	  },
	  {
		"ID": 228,
		"nrc_township_code": "တတန",
		"nrc_township_name": "တွံတွေး",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 154,
		"nrc_township_code": "တတဥ",
		"nrc_township_name": "တံတာဦး",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 5,
		"nrc_township_code": "တနန",
		"nrc_township_name": "တနိုင်း",
		"nrc_sr_code": "၁/"
	  },
	  {
		"ID": 79,
		"nrc_township_code": "တမန",
		"nrc_township_name": "တမူး",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 200,
		"nrc_township_code": "တမန",
		"nrc_township_name": "တာမွေ",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 285,
		"nrc_township_code": "တယန",
		"nrc_township_name": "တန့်ယန်း",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 162,
		"nrc_township_code": "တသန",
		"nrc_township_name": "တောင်သာ",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 87,
		"nrc_township_code": "တသရ",
		"nrc_township_name": "တနင်္သာရီ",
		"nrc_sr_code": "၆/"
	  },
	  {
		"ID": 55,
		"nrc_township_code": "ထခန",
		"nrc_township_name": "ထီးချိုင့်",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 103,
		"nrc_township_code": "ထတပ",
		"nrc_township_name": "ထန်းတပင်",
		"nrc_sr_code": "၇/"
	  },
	  {
		"ID": 215,
		"nrc_township_code": "ထတပ",
		"nrc_township_name": "ထန်းတပင်",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 40,
		"nrc_township_code": "ထတလ",
		"nrc_township_name": "ထန်တလန်",
		"nrc_sr_code": "၄/"
	  },
	  {
		"ID": 130,
		"nrc_township_code": "ထလန",
		"nrc_township_name": "ထီးလင်း",
		"nrc_sr_code": "၈/"
	  },
	  {
		"ID": 80,
		"nrc_township_code": "ထဝန",
		"nrc_township_name": "ထားဝယ်",
		"nrc_sr_code": "၆/"
	  },
	  {
		"ID": 321,
		"nrc_township_code": "ဒခသ",
		"nrc_township_name": "ဒက္ခိဏသီရိ",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 202,
		"nrc_township_code": "ဒဂဆ",
		"nrc_township_name": "ဒဂုံမြို့သစ်(ဆိပ်ကမ်း)",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 203,
		"nrc_township_code": "ဒဂတ",
		"nrc_township_name": "ဒဂုံမြို့သစ်(တောင်ပိုင်း)",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 237,
		"nrc_township_code": "ဒဂန",
		"nrc_township_name": "ဒဂုံ",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 204,
		"nrc_township_code": "ဒဂမ",
		"nrc_township_name": "ဒဂုံမြို့သစ်(မြောက်ပိုင်း)",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 205,
		"nrc_township_code": "ဒဂရ",
		"nrc_township_name": "ဒဂုံမြို့သစ်(အရှေ့ပိုင်း)",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 317,
		"nrc_township_code": "ဒဒရ",
		"nrc_township_name": "ဒေးဒရဲ",
		"nrc_sr_code": "၁၄/"
	  },
	  {
		"ID": 206,
		"nrc_township_code": "ဒပန",
		"nrc_township_name": "ဒေါပုံ",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 75,
		"nrc_township_code": "ဒပယ",
		"nrc_township_name": "ဒီပဲယင်း",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 20,
		"nrc_township_code": "ဒမဆ",
		"nrc_township_name": "ဒီးမော့ဆို",
		"nrc_sr_code": "၂/"
	  },
	  {
		"ID": 229,
		"nrc_township_code": "ဒလန",
		"nrc_township_name": "ဒလ",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 96,
		"nrc_township_code": "ဒဥန",
		"nrc_township_name": "ဒိုက်ဦး",
		"nrc_sr_code": "၇/"
	  },
	  {
		"ID": 303,
		"nrc_township_code": "ဓနဖ",
		"nrc_township_name": "ဓနုဖြူ",
		"nrc_sr_code": "၁၄/"
	  },
	  {
		"ID": 291,
		"nrc_township_code": "နခတ",
		"nrc_township_name": "နောင်ချို",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 269,
		"nrc_township_code": "နခန",
		"nrc_township_name": "နမ့်ခမ်း",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 280,
		"nrc_township_code": "နစန",
		"nrc_township_name": "နမ့်စန်",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 289,
		"nrc_township_code": "နဆန",
		"nrc_township_name": "နမ့်ဆန်",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 115,
		"nrc_township_code": "နတလ",
		"nrc_township_name": "နတ်တလင်း",
		"nrc_sr_code": "၇/"
	  },
	  {
		"ID": 160,
		"nrc_township_code": "နထက",
		"nrc_township_name": "နွားထိုးကြီး",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 290,
		"nrc_township_code": "နမတ",
		"nrc_township_name": "နမ္မတူ",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 17,
		"nrc_township_code": "နမန",
		"nrc_township_name": "နောင်မွန်း",
		"nrc_sr_code": "၁/"
	  },
	  {
		"ID": 123,
		"nrc_township_code": "နမန",
		"nrc_township_name": "နတ်မောက်",
		"nrc_sr_code": "၈/"
	  },
	  {
		"ID": 64,
		"nrc_township_code": "နယန",
		"nrc_township_name": "နန်းယွန်း",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 148,
		"nrc_township_code": "ပကခ",
		"nrc_township_name": "ပြည်ကြီးတံခွန်",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 139,
		"nrc_township_code": "ပခက",
		"nrc_township_name": "ပခုက္ကူ",
		"nrc_sr_code": "၈/"
	  },
	  {
		"ID": 105,
		"nrc_township_code": "ပခတ",
		"nrc_township_name": "ပေါက်ခေါင်း",
		"nrc_sr_code": "၇/"
	  },
	  {
		"ID": 90,
		"nrc_township_code": "ပခန",
		"nrc_township_name": "ပဲခူး",
		"nrc_sr_code": "၇/"
	  },
	  {
		"ID": 207,
		"nrc_township_code": "ပဇတ",
		"nrc_township_name": "ပုဇွန်တောင်",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 185,
		"nrc_township_code": "ပဏတ",
		"nrc_township_name": "ပုဏ္ဏားကျွန်း",
		"nrc_sr_code": "၁၁/"
	  },
	  {
		"ID": 107,
		"nrc_township_code": "ပတတ",
		"nrc_township_name": "ပေါင်းတည်",
		"nrc_sr_code": "၇/"
	  },
	  {
		"ID": 106,
		"nrc_township_code": "ပတန",
		"nrc_township_name": "ပန်းတောင်း",
		"nrc_sr_code": "၇/"
	  },
	  {
		"ID": 184,
		"nrc_township_code": "ပတန",
		"nrc_township_name": "ပေါက်တော",
		"nrc_sr_code": "၁၁/"
	  },
	  {
		"ID": 304,
		"nrc_township_code": "ပတန",
		"nrc_township_name": "ပန်းတနော်",
		"nrc_sr_code": "၁၄/"
	  },
	  {
		"ID": 250,
		"nrc_township_code": "ပတယ",
		"nrc_township_name": "ပင်းတယ",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 14,
		"nrc_township_code": "ပတအ",
		"nrc_township_name": "ပူတာအို",
		"nrc_sr_code": "၁/"
	  },
	  {
		"ID": 125,
		"nrc_township_code": "ပဖန",
		"nrc_township_name": "ပွင့်ဖြူ",
		"nrc_sr_code": "၈/"
	  },
	  {
		"ID": 326,
		"nrc_township_code": "ပဗသ",
		"nrc_township_name": "ပုဗ္ဗသီရိ",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 244,
		"nrc_township_code": "ပဘတ",
		"nrc_township_name": "ပန်းဘဲတန်း",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 172,
		"nrc_township_code": "ပဘန",
		"nrc_township_name": "ပျော်ဘွယ်",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 104,
		"nrc_township_code": "ပမန",
		"nrc_township_name": "ပြည်",
		"nrc_sr_code": "၇/"
	  },
	  {
		"ID": 142,
		"nrc_township_code": "ပမန",
		"nrc_township_name": "ပေါက်",
		"nrc_sr_code": "၈/"
	  },
	  {
		"ID": 182,
		"nrc_township_code": "ပမန",
		"nrc_township_name": "ပေါင်",
		"nrc_sr_code": "၁၀/"
	  },
	  {
		"ID": 323,
		"nrc_township_code": "ပမန",
		"nrc_township_name": "ပျဉ်းမနား",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 262,
		"nrc_township_code": "ပယန",
		"nrc_township_name": "မက်မန်း",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 70,
		"nrc_township_code": "ပလန",
		"nrc_township_name": "ပုလဲ",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 86,
		"nrc_township_code": "ပလန",
		"nrc_township_name": "ပုလော",
		"nrc_sr_code": "၆/"
	  },
	  {
		"ID": 251,
		"nrc_township_code": "ပလန",
		"nrc_township_name": "ပင်လောင်း",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 59,
		"nrc_township_code": "ပလဘ",
		"nrc_township_name": "ပင်လည်ဘူး",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 42,
		"nrc_township_code": "ပလဝ",
		"nrc_township_name": "ပလက်ဝ",
		"nrc_sr_code": "၄/"
	  },
	  {
		"ID": 149,
		"nrc_township_code": "ပသက",
		"nrc_township_name": "ပုသိမ်ကြီး",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 295,
		"nrc_township_code": "ပသန",
		"nrc_township_name": "ပုသိမ်",
		"nrc_sr_code": "၁၄/"
	  },
	  {
		"ID": 166,
		"nrc_township_code": "ပဥလ",
		"nrc_township_name": "ပြင်ဦးလွင်",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 253,
		"nrc_township_code": "ဖခန",
		"nrc_township_name": "ဖယ်ခုံ",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 24,
		"nrc_township_code": "ဖဆန",
		"nrc_township_name": "ဖားဆောင်း",
		"nrc_sr_code": "၂/"
	  },
	  {
		"ID": 33,
		"nrc_township_code": "ဖပန",
		"nrc_township_name": "ဖာပွန်",
		"nrc_sr_code": "၃/"
	  },
	  {
		"ID": 66,
		"nrc_township_code": "ဖပန",
		"nrc_township_name": "ဖောင်းပြင်",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 315,
		"nrc_township_code": "ဖပန",
		"nrc_township_name": "ဖျာပုံ",
		"nrc_sr_code": "၁၄/"
	  },
	  {
		"ID": 101,
		"nrc_township_code": "ဖမန",
		"nrc_township_name": "ဖြူး",
		"nrc_sr_code": "၇/"
	  },
	  {
		"ID": 21,
		"nrc_township_code": "ဖရဆ",
		"nrc_township_name": "ဖရူဆို",
		"nrc_sr_code": "၂/"
	  },
	  {
		"ID": 34,
		"nrc_township_code": "ဖလန",
		"nrc_township_name": "ဖလမ်း",
		"nrc_sr_code": "၄/"
	  },
	  {
		"ID": 208,
		"nrc_township_code": "ဗတထ",
		"nrc_township_name": "ဗိုလ်တထောင်",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 10,
		"nrc_township_code": "ဗမန",
		"nrc_township_name": "ဗန်းမော်",
		"nrc_sr_code": "၁/"
	  },
	  {
		"ID": 56,
		"nrc_township_code": "ဗမန",
		"nrc_township_name": "ဗန်းမောက်",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 238,
		"nrc_township_code": "ဗဟန",
		"nrc_township_name": "ဗဟန်း",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 316,
		"nrc_township_code": "ဘကလ",
		"nrc_township_name": "ဘိုကလေး",
		"nrc_sr_code": "၁၄/"
	  },
	  {
		"ID": 46,
		"nrc_township_code": "ဘတလ",
		"nrc_township_name": "ဘုတလင်",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 89,
		"nrc_township_code": "ဘပန",
		"nrc_township_name": "ဘုတ်ပြင်း",
		"nrc_sr_code": "၆/"
	  },
	  {
		"ID": 23,
		"nrc_township_code": "ဘလခ",
		"nrc_township_name": "ဘောလခဲ",
		"nrc_sr_code": "၂/"
	  },
	  {
		"ID": 180,
		"nrc_township_code": "ဘလန",
		"nrc_township_name": "ဘီးလင်း",
		"nrc_sr_code": "၁၀/"
	  },
	  {
		"ID": 199,
		"nrc_township_code": "ဘသတ",
		"nrc_township_name": "ဘူးသီးတောင်",
		"nrc_sr_code": "၁၁/"
	  },
	  {
		"ID": 27,
		"nrc_township_code": "ဘအန",
		"nrc_township_name": "ဘားအံ",
		"nrc_sr_code": "၃/"
	  },
	  {
		"ID": 1,
		"nrc_township_code": "မကန",
		"nrc_township_name": "မြစ်ကြီးနား",
		"nrc_sr_code": "၁/"
	  },
	  {
		"ID": 8,
		"nrc_township_code": "မကန",
		"nrc_township_name": "မိုးကောင်း",
		"nrc_sr_code": "၁/"
	  },
	  {
		"ID": 52,
		"nrc_township_code": "မကန",
		"nrc_township_name": "မင်းကင်း",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 118,
		"nrc_township_code": "မကန",
		"nrc_township_name": "မကွေး",
		"nrc_sr_code": "၈/"
	  },
	  {
		"ID": 167,
		"nrc_township_code": "မကန",
		"nrc_township_name": "မိုးကုတ်",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 277,
		"nrc_township_code": "မကန",
		"nrc_township_name": "မိုင်းကိုင်",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 159,
		"nrc_township_code": "မခန",
		"nrc_township_name": "မြင်းခြံ",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 256,
		"nrc_township_code": "မခန",
		"nrc_township_name": "မိုင်းခတ်",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 16,
		"nrc_township_code": "မခဘ",
		"nrc_township_name": "မချမ်းဘော",
		"nrc_sr_code": "၁/"
	  },
	  {
		"ID": 213,
		"nrc_township_code": "မဂတ",
		"nrc_township_name": "မင်္ဂလာတောင်ညွန့်",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 216,
		"nrc_township_code": "မဂဒ",
		"nrc_township_name": "မင်္ဂလာဒုံ",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 12,
		"nrc_township_code": "မစန",
		"nrc_township_name": "မံစီ",
		"nrc_sr_code": "၁/"
	  },
	  {
		"ID": 25,
		"nrc_township_code": "မစန",
		"nrc_township_name": "မယ်စဲ",
		"nrc_sr_code": "၂/"
	  },
	  {
		"ID": 267,
		"nrc_township_code": "မဆတ",
		"nrc_township_name": "မူဆယ်",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 265,
		"nrc_township_code": "မဆန",
		"nrc_township_name": "မိုင်းဆတ်",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 7,
		"nrc_township_code": "မညန",
		"nrc_township_name": "မိုးညှင်း",
		"nrc_sr_code": "၁/"
	  },
	  {
		"ID": 116,
		"nrc_township_code": "မညန",
		"nrc_township_name": "မိုးညို",
		"nrc_sr_code": "၇/"
	  },
	  {
		"ID": 288,
		"nrc_township_code": "မတတ",
		"nrc_township_name": "မန်တုံ",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 37,
		"nrc_township_code": "မတန",
		"nrc_township_name": "မင်းတပ်",
		"nrc_sr_code": "၄/"
	  },
	  {
		"ID": 134,
		"nrc_township_code": "မတန",
		"nrc_township_name": "မင်းတုန်း",
		"nrc_sr_code": "၈/"
	  },
	  {
		"ID": 198,
		"nrc_township_code": "မတန",
		"nrc_township_name": "မောင်တော",
		"nrc_sr_code": "၁၁/"
	  },
	  {
		"ID": 266,
		"nrc_township_code": "မတန",
		"nrc_township_name": "မိုင်းတုံ",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 41,
		"nrc_township_code": "မတပ",
		"nrc_township_name": "မတူပီ",
		"nrc_sr_code": "၄/"
	  },
	  {
		"ID": 170,
		"nrc_township_code": "မတရ",
		"nrc_township_name": "မတ္တရာ",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 137,
		"nrc_township_code": "မထန",
		"nrc_township_name": "အောင်လံ",
		"nrc_sr_code": "၈/"
	  },
	  {
		"ID": 155,
		"nrc_township_code": "မထလ",
		"nrc_township_name": "မိတ္ထီလာ",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 177,
		"nrc_township_code": "မဒန",
		"nrc_township_name": "မုဒုံ",
		"nrc_sr_code": "၁၀/"
	  },
	  {
		"ID": 332,
		"nrc_township_code": "မနတ",
		"nrc_township_name": "အောင်မြေသာဇံ",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 271,
		"nrc_township_code": "မနန",
		"nrc_township_name": "မိုးနဲ",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 330,
		"nrc_township_code": "မနမ",
		"nrc_township_name": "အောင်မြေသာဇံ",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 192,
		"nrc_township_code": "မပတ",
		"nrc_township_name": "မြေပုံ",
		"nrc_sr_code": "၁၁/"
	  },
	  {
		"ID": 258,
		"nrc_township_code": "မပတ",
		"nrc_township_name": "မိုင်းပျဉ်း",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 193,
		"nrc_township_code": "မပန",
		"nrc_township_name": "မင်းပြား",
		"nrc_sr_code": "၁၁/"
	  },
	  {
		"ID": 272,
		"nrc_township_code": "မပန",
		"nrc_township_name": "မိုင်းပန်",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 260,
		"nrc_township_code": "မဖန",
		"nrc_township_name": "မိုင်းဖြတ်",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 124,
		"nrc_township_code": "မဘန",
		"nrc_township_name": "မင်းဘူး",
		"nrc_sr_code": "၈/"
	  },
	  {
		"ID": 217,
		"nrc_township_code": "မဘန",
		"nrc_township_name": "မှော်ဘီ",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 264,
		"nrc_township_code": "မဘန",
		"nrc_township_name": "မဘိမ်း",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 320,
		"nrc_township_code": "မမက",
		"nrc_township_name": "မော်လမြိုင်ကျွန်း",
		"nrc_sr_code": "၁၄/"
	  },
	  {
		"ID": 48,
		"nrc_township_code": "မမတ",
		"nrc_township_name": "မြင်းမူ",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 263,
		"nrc_township_code": "မမတ",
		"nrc_township_name": "မိုးမိတ်",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 13,
		"nrc_township_code": "မမန",
		"nrc_township_name": "မိုးမောက်",
		"nrc_sr_code": "၁/"
	  },
	  {
		"ID": 49,
		"nrc_township_code": "မမန",
		"nrc_township_name": "မြောင်",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 84,
		"nrc_township_code": "မမန",
		"nrc_township_name": "မြိတ်",
		"nrc_sr_code": "၆/"
	  },
	  {
		"ID": 141,
		"nrc_township_code": "မမန",
		"nrc_township_name": "မြိုင်",
		"nrc_sr_code": "၈/"
	  },
	  {
		"ID": 165,
		"nrc_township_code": "မမန",
		"nrc_township_name": "ပြင်ဦးလွင်",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 273,
		"nrc_township_code": "မမန",
		"nrc_township_name": "မောက်မယ်",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 312,
		"nrc_township_code": "မမန",
		"nrc_township_name": "မြောင်းမြ",
		"nrc_sr_code": "၁၄/"
	  },
	  {
		"ID": 261,
		"nrc_township_code": "မယတ",
		"nrc_township_name": "မိုင်းယောင်း",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 257,
		"nrc_township_code": "မယန",
		"nrc_township_name": "မိုင်းယန်း",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 239,
		"nrc_township_code": "မရက",
		"nrc_township_name": "မရမ်းကုန်း",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 284,
		"nrc_township_code": "မရတ",
		"nrc_township_name": "မိုင်းရယ်",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 329,
		"nrc_township_code": "မရတ",
		"nrc_township_name": "အောင်မြေသာဇံ",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 43,
		"nrc_township_code": "မရန",
		"nrc_township_name": "မုံရွာ",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 278,
		"nrc_township_code": "မရန",
		"nrc_township_name": "မိုင်းရှူး",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 331,
		"nrc_township_code": "မရမ",
		"nrc_township_name": "အောင်မြေသာဇံ",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 65,
		"nrc_township_code": "မလန",
		"nrc_township_name": "မော်လိုက်",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 112,
		"nrc_township_code": "မလန",
		"nrc_township_name": "မင်းလှ",
		"nrc_sr_code": "၇/"
	  },
	  {
		"ID": 133,
		"nrc_township_code": "မလန",
		"nrc_township_name": "မင်းလှ",
		"nrc_sr_code": "၈/"
	  },
	  {
		"ID": 156,
		"nrc_township_code": "မလန",
		"nrc_township_name": "မလှိုင်",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 173,
		"nrc_township_code": "မလမ",
		"nrc_township_name": "မော်လမြိုင်",
		"nrc_sr_code": "၁၀/"
	  },
	  {
		"ID": 30,
		"nrc_township_code": "မဝတ",
		"nrc_township_name": "မြဝတီ",
		"nrc_sr_code": "၃/"
	  },
	  {
		"ID": 122,
		"nrc_township_code": "မသန",
		"nrc_township_name": "မြို့သစ်",
		"nrc_sr_code": "၈/"
	  },
	  {
		"ID": 152,
		"nrc_township_code": "မသန",
		"nrc_township_name": "မြစ်သား",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 147,
		"nrc_township_code": "မဟမ",
		"nrc_township_name": "မဟာအောင်မြေ",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 190,
		"nrc_township_code": "မအန",
		"nrc_township_name": "မာန်အောင်",
		"nrc_sr_code": "၁၁/"
	  },
	  {
		"ID": 309,
		"nrc_township_code": "မအန",
		"nrc_township_name": "မြန်အောင်",
		"nrc_sr_code": "၁၄/"
	  },
	  {
		"ID": 302,
		"nrc_township_code": "မအပ",
		"nrc_township_name": "မအူပင်",
		"nrc_sr_code": "၁၄/"
	  },
	  {
		"ID": 191,
		"nrc_township_code": "မဥန",
		"nrc_township_name": "မြောက်ဦး",
		"nrc_sr_code": "၁၁/"
	  },
	  {
		"ID": 67,
		"nrc_township_code": "ယမပ",
		"nrc_township_name": "ယင်းမာပင်",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 11,
		"nrc_township_code": "ရကန",
		"nrc_township_name": "ရွှေကူ",
		"nrc_sr_code": "၁/"
	  },
	  {
		"ID": 97,
		"nrc_township_code": "ရကန",
		"nrc_township_name": "ရွှေကျင်",
		"nrc_sr_code": "၇/"
	  },
	  {
		"ID": 210,
		"nrc_township_code": "ရကန",
		"nrc_township_name": "ရန်ကင်း",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 301,
		"nrc_township_code": "ရကန",
		"nrc_township_name": "ရေကြည်",
		"nrc_sr_code": "၁၄/"
	  },
	  {
		"ID": 252,
		"nrc_township_code": "ရငန",
		"nrc_township_name": "ရွာငံ",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 140,
		"nrc_township_code": "ရစက",
		"nrc_township_name": "ရေစကြို",
		"nrc_sr_code": "၈/"
	  },
	  {
		"ID": 248,
		"nrc_township_code": "ရစန",
		"nrc_township_name": "ရပ်စောက်",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 22,
		"nrc_township_code": "ရတန",
		"nrc_township_name": "ရှားတော",
		"nrc_sr_code": "၂/"
	  },
	  {
		"ID": 109,
		"nrc_township_code": "ရတန",
		"nrc_township_name": "ရွှေတောင်",
		"nrc_sr_code": "၇/"
	  },
	  {
		"ID": 99,
		"nrc_township_code": "ရတရ",
		"nrc_township_name": "ရေတာရှည်",
		"nrc_sr_code": "၇/"
	  },
	  {
		"ID": 119,
		"nrc_township_code": "ရနခ",
		"nrc_township_name": "ရေနံချောင်း",
		"nrc_sr_code": "၈/"
	  },
	  {
		"ID": 218,
		"nrc_township_code": "ရပသ",
		"nrc_township_name": "ရွှေပြည်သာ",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 83,
		"nrc_township_code": "ရဖန",
		"nrc_township_name": "ရေဖြူ",
		"nrc_sr_code": "၆/"
	  },
	  {
		"ID": 188,
		"nrc_township_code": "ရဗန",
		"nrc_township_name": "ရမ်းဗြဲ",
		"nrc_sr_code": "၁၁/"
	  },
	  {
		"ID": 71,
		"nrc_township_code": "ရဘန",
		"nrc_township_name": "ရွှေဘို",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 174,
		"nrc_township_code": "ရမန",
		"nrc_township_name": "ရေး",
		"nrc_sr_code": "၁၀/"
	  },
	  {
		"ID": 171,
		"nrc_township_code": "ရမသ",
		"nrc_township_name": "ရမည်းသင်း",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 186,
		"nrc_township_code": "ရသတ",
		"nrc_township_name": "ရသေ့တောင်",
		"nrc_sr_code": "၁၁/"
	  },
	  {
		"ID": 26,
		"nrc_township_code": "ရသန",
		"nrc_township_name": "ရွာသစ်",
		"nrc_sr_code": "၂/"
	  },
	  {
		"ID": 74,
		"nrc_township_code": "ရဥန",
		"nrc_township_name": "ရေဦး",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 19,
		"nrc_township_code": "လကန",
		"nrc_township_name": "လွိုင်ကော်",
		"nrc_sr_code": "၂/"
	  },
	  {
		"ID": 219,
		"nrc_township_code": "လကန",
		"nrc_township_name": "လှည်းကူး",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 292,
		"nrc_township_code": "လကန",
		"nrc_township_name": "လောက်ကိုင်",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 270,
		"nrc_township_code": "လခတ",
		"nrc_township_name": "လင်းခေး",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 276,
		"nrc_township_code": "လခန",
		"nrc_township_name": "လဲချား",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 111,
		"nrc_township_code": "လပတ",
		"nrc_township_name": "လက်ပံတန်း",
		"nrc_sr_code": "၇/"
	  },
	  {
		"ID": 319,
		"nrc_township_code": "လပတ",
		"nrc_township_name": "လပွတ္တာ",
		"nrc_sr_code": "၁၄/"
	  },
	  {
		"ID": 28,
		"nrc_township_code": "လဘန",
		"nrc_township_name": "လှိုင်းဘွဲ့",
		"nrc_sr_code": "၃/"
	  },
	  {
		"ID": 240,
		"nrc_township_code": "လမတ",
		"nrc_township_name": "လမ်းမတော်",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 242,
		"nrc_township_code": "လမန",
		"nrc_township_name": "လှိုင်",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 311,
		"nrc_township_code": "လမန",
		"nrc_township_name": "လေးမျက်နှာ",
		"nrc_sr_code": "၁၄/"
	  },
	  {
		"ID": 62,
		"nrc_township_code": "လရန",
		"nrc_township_name": "လေရှီး",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 281,
		"nrc_township_code": "လရန",
		"nrc_township_name": "လားရှိူး",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 81,
		"nrc_township_code": "လလန",
		"nrc_township_name": "လောင်းလုံး",
		"nrc_sr_code": "၆/"
	  },
	  {
		"ID": 274,
		"nrc_township_code": "လလန",
		"nrc_township_name": "လွိုင်လင်",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 322,
		"nrc_township_code": "လဝန",
		"nrc_township_name": "လယ်ဝေး",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 241,
		"nrc_township_code": "လသန",
		"nrc_township_name": "လသာ",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 220,
		"nrc_township_code": "လသယ",
		"nrc_township_name": "လှိုင်သာယာ",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 63,
		"nrc_township_code": "လဟန",
		"nrc_township_name": "လဟယ်",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 313,
		"nrc_township_code": "ဝခမ",
		"nrc_township_name": "ဝါးခယ်မ",
		"nrc_sr_code": "၁၄/"
	  },
	  {
		"ID": 158,
		"nrc_township_code": "ဝတန",
		"nrc_township_name": "ဝမ်းတွင်း",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 2,
		"nrc_township_code": "ဝမန",
		"nrc_township_name": "ဝိုင်းမော်ြမို့",
		"nrc_sr_code": "၁/"
	  },
	  {
		"ID": 93,
		"nrc_township_code": "ဝမန",
		"nrc_township_name": "ဝေါ",
		"nrc_sr_code": "၇/"
	  },
	  {
		"ID": 73,
		"nrc_township_code": "ဝလန",
		"nrc_township_name": "ဝက်လက်",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 212,
		"nrc_township_code": "သကတ",
		"nrc_township_name": "သာကေတ",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 108,
		"nrc_township_code": "သကန",
		"nrc_township_name": "သဲကုန်း",
		"nrc_sr_code": "၇/"
	  },
	  {
		"ID": 231,
		"nrc_township_code": "သခန",
		"nrc_township_name": "သုံးခွ",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 211,
		"nrc_township_code": "သဃက",
		"nrc_township_name": "သင်္ဃန်းကျွန်း",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 157,
		"nrc_township_code": "သစန",
		"nrc_township_name": "သာစည်",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 29,
		"nrc_township_code": "သတန",
		"nrc_township_name": "သံတောင်ကြီး",
		"nrc_sr_code": "၃/"
	  },
	  {
		"ID": 195,
		"nrc_township_code": "သတန",
		"nrc_township_name": "သံတွဲ",
		"nrc_sr_code": "၁၁/"
	  },
	  {
		"ID": 179,
		"nrc_township_code": "သထန",
		"nrc_township_name": "သထုံ",
		"nrc_sr_code": "၁၀/"
	  },
	  {
		"ID": 282,
		"nrc_township_code": "သနန",
		"nrc_township_name": "သိန္နီ",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 91,
		"nrc_township_code": "သနပ",
		"nrc_township_name": "သနပ်ပင်",
		"nrc_sr_code": "၇/"
	  },
	  {
		"ID": 169,
		"nrc_township_code": "သပက",
		"nrc_township_name": "သပိတ်ကျင်း",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 287,
		"nrc_township_code": "သပန",
		"nrc_township_name": "သီပေါ",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 299,
		"nrc_township_code": "သပန",
		"nrc_township_name": "သာပေါင်း",
		"nrc_sr_code": "၁၄/"
	  },
	  {
		"ID": 178,
		"nrc_township_code": "သဖရ",
		"nrc_township_name": "သံဖြူဇရပ်",
		"nrc_sr_code": "၁၀/"
	  },
	  {
		"ID": 82,
		"nrc_township_code": "သရခ",
		"nrc_township_name": "သရက်ချောင်း",
		"nrc_sr_code": "၆/"
	  },
	  {
		"ID": 132,
		"nrc_township_code": "သရန",
		"nrc_township_name": "သရက်",
		"nrc_sr_code": "၈/"
	  },
	  {
		"ID": 230,
		"nrc_township_code": "သလန",
		"nrc_township_name": "သန်လျင်",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 110,
		"nrc_township_code": "သဝတ",
		"nrc_township_name": "သာယာဝတီ",
		"nrc_sr_code": "၇/"
	  },
	  {
		"ID": 39,
		"nrc_township_code": "ဟခန",
		"nrc_township_name": "ဟားခါး",
		"nrc_sr_code": "၄/"
	  },
	  {
		"ID": 254,
		"nrc_township_code": "ဟပန",
		"nrc_township_name": "ဟိုပုံး",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 294,
		"nrc_township_code": "ဟပန",
		"nrc_township_name": "ဟိုပန်",
		"nrc_sr_code": "၁၃/"
	  },
	  {
		"ID": 61,
		"nrc_township_code": "ဟမလ",
		"nrc_township_name": "ဟုမ္မလင်း",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 306,
		"nrc_township_code": "ဟသတ",
		"nrc_township_name": "ဟင်္သာတ",
		"nrc_sr_code": "၁၄/"
	  },
	  {
		"ID": 310,
		"nrc_township_code": "အဂပ",
		"nrc_township_name": "အင်္ဂပူ",
		"nrc_sr_code": "၁၄/"
	  },
	  {
		"ID": 6,
		"nrc_township_code": "အဂယ",
		"nrc_township_name": "အင်ဂျန်းယန်",
		"nrc_sr_code": "၁/"
	  },
	  {
		"ID": 221,
		"nrc_township_code": "အစန",
		"nrc_township_name": "အင်းစိန်",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 54,
		"nrc_township_code": "အတန",
		"nrc_township_name": "အင်းတော်",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 102,
		"nrc_township_code": "အတန",
		"nrc_township_name": "အုတ်တွင်း",
		"nrc_sr_code": "၇/"
	  },
	  {
		"ID": 113,
		"nrc_township_code": "အဖန",
		"nrc_township_name": "အုတ်ဖို",
		"nrc_sr_code": "၇/"
	  },
	  {
		"ID": 144,
		"nrc_township_code": "အမဇ",
		"nrc_township_name": "အောင်မြေသာဇံ",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 189,
		"nrc_township_code": "အမန",
		"nrc_township_name": "အမ်း",
		"nrc_sr_code": "၁၁/"
	  },
	  {
		"ID": 314,
		"nrc_township_code": "အမန",
		"nrc_township_name": "အိမ်မဲ",
		"nrc_sr_code": "၁၄/"
	  },
	  {
		"ID": 150,
		"nrc_township_code": "အမရ",
		"nrc_township_name": "အမရပူရ",
		"nrc_sr_code": "၉/"
	  },
	  {
		"ID": 44,
		"nrc_township_code": "အရတ",
		"nrc_township_name": "အရာတော်",
		"nrc_sr_code": "၅/"
	  },
	  {
		"ID": 136,
		"nrc_township_code": "အလန",
		"nrc_township_name": "အောင်လံ",
		"nrc_sr_code": "၈/"
	  },
	  {
		"ID": 243,
		"nrc_township_code": "အလန",
		"nrc_township_name": "အလုံ",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 201,
		"nrc_township_code": "ဥကတ",
		"nrc_township_name": "တောင်ဥက္ကလာပ",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 209,
		"nrc_township_code": "ဥကမ",
		"nrc_township_name": "မြောက်ဥက္လလာပ",
		"nrc_sr_code": "၁၂/"
	  },
	  {
		"ID": 325,
		"nrc_township_code": "ဥတသ",
		"nrc_township_name": "ဥတ္တရသီရိ",
		"nrc_sr_code": "၉/"
	  }
	]
  }
  `
const townShipJSON = `
  {"data":[    {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "194",
		"Tsp_Name": "ကနီ"
	  },
	  {
		"SR_Code": "14",
		"SR_Name": "ဧရာဝတီတိုင်းဒေသကြီး",
		"Tsp_Code": "7",
		"Tsp_Name": "ကန်ကြီးထောင့်"
	  },
	  {
		"SR_Code": "4",
		"SR_Name": "ချင်းပြည်နယ်",
		"Tsp_Code": "57",
		"Tsp_Name": "ကန်ပက်လက်"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "193",
		"Tsp_Name": "ကန့်ဘလူ"
	  },
	  {
		"SR_Code": "8",
		"SR_Name": "မကွေးတိုင်းဒေသကြီး",
		"Tsp_Code": "99",
		"Tsp_Name": "ကံမ"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "304",
		"Tsp_Name": "ကမာရွတ်"
	  },
	  {
		"SR_Code": "6",
		"SR_Name": "တနင်္သာရီတိုင်းဒေသကြီး",
		"Tsp_Code": "280",
		"Tsp_Name": "ကျွန်းစု"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "198",
		"Tsp_Name": "ကျွန်းလှ"
	  },
	  {
		"SR_Code": "10",
		"SR_Name": "မွန်ပြည်နယ်",
		"Tsp_Code": "152",
		"Tsp_Name": "ကျိုက်ထို"
	  },
	  {
		"SR_Code": "10",
		"SR_Name": "မွန်ပြည်နယ်",
		"Tsp_Code": "151",
		"Tsp_Name": "ကျိုက်မရော"
	  },
	  {
		"SR_Code": "14",
		"SR_Name": "ဧရာဝတီတိုင်းဒေသကြီး",
		"Tsp_Code": "8",
		"Tsp_Name": "ကျိုက်လတ်"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "221",
		"Tsp_Name": "ကျိုင်းတုံ"
	  },
	  {
		"SR_Code": "14",
		"SR_Name": "ဧရာဝတီတိုင်းဒေသကြီး",
		"Tsp_Code": "11",
		"Tsp_Name": "ကျုံပျော်"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "260",
		"Tsp_Name": "ကျေးသီး"
	  },
	  {
		"SR_Code": "7",
		"SR_Name": "ပဲခူးတိုင်းဒေသကြီး",
		"Tsp_Code": "32",
		"Tsp_Name": "ကျောက်ကြီး"
	  },
	  {
		"SR_Code": "9",
		"SR_Name": "မန္တလေးတိုင်းဒေသကြီး",
		"Tsp_Code": "127",
		"Tsp_Name": "ကျောက်ဆည်"
	  },
	  {
		"SR_Code": "7",
		"SR_Name": "ပဲခူးတိုင်းဒေသကြီး",
		"Tsp_Code": "33",
		"Tsp_Name": "ကျောက်တံခါး"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "308",
		"Tsp_Name": "ကျောက်တံတား"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "309",
		"Tsp_Name": "ကျောက်တန်း"
	  },
	  {
		"SR_Code": "11",
		"SR_Name": "ရခိုင်ပြည်နယ်",
		"Tsp_Code": "171",
		"Tsp_Name": "ကျောက်တော်"
	  },
	  {
		"SR_Code": "9",
		"SR_Name": "မန္တလေးတိုင်းဒေသကြီး",
		"Tsp_Code": "126",
		"Tsp_Name": "ကျောက်ပန်းတောင်း"
	  },
	  {
		"SR_Code": "11",
		"SR_Name": "ရခိုင်ပြည်နယ်",
		"Tsp_Code": "170",
		"Tsp_Name": "ကျောက်ဖြူ"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "237",
		"Tsp_Name": "ကျောက်မဲ"
	  },
	  {
		"SR_Code": "14",
		"SR_Name": "ဧရာဝတီတိုင်းဒေသကြီး",
		"Tsp_Code": "10",
		"Tsp_Name": "ကျောင်းကုန်း"
	  },
	  {
		"SR_Code": "14",
		"SR_Name": "ဧရာဝတီတိုင်းဒေသကြီး",
		"Tsp_Code": "9",
		"Tsp_Name": "ကြံခင်း"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "310",
		"Tsp_Name": "ကြည့်မြင်တိုင်"
	  },
	  {
		"SR_Code": "3",
		"SR_Name": "ကရင်ပြည်နယ်",
		"Tsp_Code": "93",
		"Tsp_Name": "ကြာအင်းဆိပ်ကြီး"
	  },
	  {
		"SR_Code": "7",
		"SR_Name": "ပဲခူးတိုင်းဒေသကြီး",
		"Tsp_Code": "41",
		"Tsp_Name": "ကြို့ပင်ကောက်"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "191",
		"Tsp_Name": "ကလေး"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "192",
		"Tsp_Name": "ကလေးဝ"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "258",
		"Tsp_Name": "ကလော"
	  },
	  {
		"SR_Code": "7",
		"SR_Name": "ပဲခူးတိုင်းဒေသကြီး",
		"Tsp_Code": "31",
		"Tsp_Name": "ကဝ"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "236",
		"Tsp_Name": "ကွတ်ခိုင်"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "259",
		"Tsp_Name": "ကွန်ဟိန်း"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "307",
		"Tsp_Name": "ကွမ်းခြံကုန်း"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "235",
		"Tsp_Name": "ကွမ်းလုံ"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "195",
		"Tsp_Name": "ကသာ"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "289",
		"Tsp_Name": "ကိုကိုးကျွန်း"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "234",
		"Tsp_Name": "ကုန်းကြမ်း"
	  },
	  {
		"SR_Code": "3",
		"SR_Name": "ကရင်ပြည်နယ်",
		"Tsp_Code": "92",
		"Tsp_Name": "ကော့ကရိတ်"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "305",
		"Tsp_Name": "ကော့မှူး"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "196",
		"Tsp_Name": "ကောလင်း"
	  },
	  {
		"SR_Code": "6",
		"SR_Name": "တနင်္သာရီတိုင်းဒေသကြီး",
		"Tsp_Code": "279",
		"Tsp_Name": "ကော့သောင်း"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "197",
		"Tsp_Name": "ခင်ဦး"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "188",
		"Tsp_Name": "ခန္တီး"
	  },
	  {
		"SR_Code": "9",
		"SR_Name": "မန္တလေးတိုင်းဒေသကြီး",
		"Tsp_Code": "125",
		"Tsp_Name": "ချမ်းမြသာစည်"
	  },
	  {
		"SR_Code": "9",
		"SR_Name": "မန္တလေးတိုင်းဒေသကြီး",
		"Tsp_Code": "124",
		"Tsp_Name": "ချမ်းအေးသာစံ"
	  },
	  {
		"SR_Code": "1",
		"SR_Name": "ကချင်ပြည်နယ်",
		"Tsp_Code": "65",
		"Tsp_Name": "ချီဖွေ"
	  },
	  {
		"SR_Code": "8",
		"SR_Name": "မကွေးတိုင်းဒေသကြီး",
		"Tsp_Code": "97",
		"Tsp_Name": "ချောက်"
	  },
	  {
		"SR_Code": "10",
		"SR_Name": "မွန်ပြည်နယ်",
		"Tsp_Code": "150",
		"Tsp_Name": "ချောင်းဆုံ"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "187",
		"Tsp_Name": "ချောင်းဦး"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "306",
		"Tsp_Name": "ခရမ်း"
	  },
	  {
		"SR_Code": "1",
		"SR_Name": "ကချင်ပြည်နယ်",
		"Tsp_Code": "68",
		"Tsp_Name": "ခေါင်လန်ဖူး"
	  },
	  {
		"SR_Code": "8",
		"SR_Name": "မကွေးတိုင်းဒေသကြီး",
		"Tsp_Code": "98",
		"Tsp_Name": "ဂန့်ဂေါ"
	  },
	  {
		"SR_Code": "11",
		"SR_Name": "ရခိုင်ပြည်နယ်",
		"Tsp_Code": "169",
		"Tsp_Name": "ဂွ"
	  },
	  {
		"SR_Code": "14",
		"SR_Name": "ဧရာဝတီတိုင်းဒေသကြီး",
		"Tsp_Code": "18",
		"Tsp_Name": "ငပုတော"
	  },
	  {
		"SR_Code": "8",
		"SR_Name": "မကွေးတိုင်းဒေသကြီး",
		"Tsp_Code": "107",
		"Tsp_Name": "ငဖဲ"
	  },
	  {
		"SR_Code": "9",
		"SR_Name": "မန္တလေးတိုင်းဒေသကြီး",
		"Tsp_Code": "136",
		"Tsp_Name": "ငါန်းဇွန်"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "210",
		"Tsp_Name": "စစ်ကိုင်း"
	  },
	  {
		"SR_Code": "11",
		"SR_Name": "ရခိုင်ပြည်နယ်",
		"Tsp_Code": "181",
		"Tsp_Name": "စစ်တွေ"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "319",
		"Tsp_Name": "စမ်းချောင်း"
	  },
	  {
		"SR_Code": "8",
		"SR_Name": "မကွေးတိုင်းဒေသကြီး",
		"Tsp_Code": "111",
		"Tsp_Name": "စလင်း"
	  },
	  {
		"SR_Code": "9",
		"SR_Name": "မန္တလေးတိုင်းဒေသကြီး",
		"Tsp_Code": "142",
		"Tsp_Name": "စဥ့်ကိုင်"
	  },
	  {
		"SR_Code": "9",
		"SR_Name": "မန္တလေးတိုင်းဒေသကြီး",
		"Tsp_Code": "141",
		"Tsp_Name": "စဥ့်ကူး"
	  },
	  {
		"SR_Code": "8",
		"SR_Name": "မကွေးတိုင်းဒေသကြီး",
		"Tsp_Code": "114",
		"Tsp_Name": "စေတုတ္တရာ"
	  },
	  {
		"SR_Code": "8",
		"SR_Name": "မကွေးတိုင်းဒေသကြီး",
		"Tsp_Code": "115",
		"Tsp_Name": "ဆင်ပေါင်ဝဲ"
	  },
	  {
		"SR_Code": "1",
		"SR_Name": "ကချင်ပြည်နယ်",
		"Tsp_Code": "78",
		"Tsp_Name": "ဆွမ်ပရာဘွမ်"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "211",
		"Tsp_Name": "ဆားလင်းကြီး"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "320",
		"Tsp_Name": "ဆိပ်ကြီး/ခနောင်တို"
	  },
	  {
		"SR_Code": "8",
		"SR_Name": "မကွေးတိုင်းဒေသကြီး",
		"Tsp_Code": "113",
		"Tsp_Name": "ဆိပ်ဖြူ"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "257",
		"Tsp_Name": "ဆီဆိုင်"
	  },
	  {
		"SR_Code": "8",
		"SR_Name": "မကွေးတိုင်းဒေသကြီး",
		"Tsp_Code": "112",
		"Tsp_Name": "ဆော"
	  },
	  {
		"SR_Code": "1",
		"SR_Name": "ကချင်ပြည်နယ်",
		"Tsp_Code": "80",
		"Tsp_Name": "ဆော့လော်"
	  },
	  {
		"SR_Code": "15",
		"SR_Name": "နေပြည်တော်",
		"Tsp_Code": "165",
		"Tsp_Name": "ဇမ္ဗူသီရိ"
	  },
	  {
		"SR_Code": "14",
		"SR_Name": "ဧရာဝတီတိုင်းဒေသကြီး",
		"Tsp_Code": "24",
		"Tsp_Name": "ဇလွန်"
	  },
	  {
		"SR_Code": "7",
		"SR_Name": "ပဲခူးတိုင်းဒေသကြီး",
		"Tsp_Code": "53",
		"Tsp_Name": "ဇီးကုန်း"
	  },
	  {
		"SR_Code": "15",
		"SR_Name": "နေပြည်တော်",
		"Tsp_Code": "166",
		"Tsp_Name": "ဇေယျာသီရိ"
	  },
	  {
		"SR_Code": "14",
		"SR_Name": "ဧရာဝတီတိုင်းဒေသကြီး",
		"Tsp_Code": "25",
		"Tsp_Name": "ညောင်တုန်း"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "255",
		"Tsp_Name": "ညောင်ရွှေ"
	  },
	  {
		"SR_Code": "7",
		"SR_Name": "ပဲခူးတိုင်းဒေသကြီး",
		"Tsp_Code": "40",
		"Tsp_Name": "ညောင်လေးပင်"
	  },
	  {
		"SR_Code": "9",
		"SR_Name": "မန္တလေးတိုင်းဒေသကြီး",
		"Tsp_Code": "121",
		"Tsp_Name": "ညောင်ဦး"
	  },
	  {
		"SR_Code": "9",
		"SR_Name": "မန္တလေးတိုင်းဒေသကြီး",
		"Tsp_Code": "143",
		"Tsp_Name": "တံတားဦး"
	  },
	  {
		"SR_Code": "6",
		"SR_Name": "တနင်္သာရီတိုင်းဒေသကြီး",
		"Tsp_Code": "284",
		"Tsp_Name": "တနင်္သာရီ"
	  },
	  {
		"SR_Code": "1",
		"SR_Name": "ကချင်ပြည်နယ်",
		"Tsp_Code": "79",
		"Tsp_Name": "တနိုင်း"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "215",
		"Tsp_Name": "တန့်ဆည်"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "254",
		"Tsp_Name": "တန့်ယန်း"
	  },
	  {
		"SR_Code": "15",
		"SR_Name": "နေပြည်တော်",
		"Tsp_Code": "164",
		"Tsp_Name": "တပ်ကုန်း"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "214",
		"Tsp_Name": "တမူး"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "329",
		"Tsp_Name": "တွံတေး"
	  },
	  {
		"SR_Code": "4",
		"SR_Name": "ချင်းပြည်နယ်",
		"Tsp_Code": "63",
		"Tsp_Name": "တွန်းဇံ"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "230",
		"Tsp_Name": "တာချီလိတ်"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "324",
		"Tsp_Name": "တာမွေ"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "323",
		"Tsp_Name": "တိုက်ကြီး"
	  },
	  {
		"SR_Code": "4",
		"SR_Name": "ချင်းပြည်နယ်",
		"Tsp_Code": "61",
		"Tsp_Name": "တီးတိန်"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "274",
		"Tsp_Name": "တောင်ကြီး"
	  },
	  {
		"SR_Code": "11",
		"SR_Name": "ရခိုင်ပြည်နယ်",
		"Tsp_Code": "183",
		"Tsp_Name": "တောင်ကုတ်"
	  },
	  {
		"SR_Code": "7",
		"SR_Name": "ပဲခူးတိုင်းဒေသကြီး",
		"Tsp_Code": "37",
		"Tsp_Name": "တောင်ငူ"
	  },
	  {
		"SR_Code": "8",
		"SR_Name": "မကွေးတိုင်းဒေသကြီး",
		"Tsp_Code": "116",
		"Tsp_Name": "တောင်တွင်းကြီး"
	  },
	  {
		"SR_Code": "9",
		"SR_Name": "မန္တလေးတိုင်းဒေသကြီး",
		"Tsp_Code": "144",
		"Tsp_Name": "တောင်သာ"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "322",
		"Tsp_Name": "တောင်ဥက္ကလာပ"
	  },
	  {
		"SR_Code": "7",
		"SR_Name": "ပဲခူးတိုင်းဒေသကြီး",
		"Tsp_Code": "30",
		"Tsp_Name": "ထန်းတပင်"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "302",
		"Tsp_Name": "ထန်းတပင်"
	  },
	  {
		"SR_Code": "4",
		"SR_Name": "ချင်းပြည်နယ်",
		"Tsp_Code": "62",
		"Tsp_Name": "ထန်တလန်"
	  },
	  {
		"SR_Code": "6",
		"SR_Name": "တနင်္သာရီတိုင်းဒေသကြီး",
		"Tsp_Code": "278",
		"Tsp_Name": "ထားဝယ်"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "216",
		"Tsp_Name": "ထီးချိုင့်"
	  },
	  {
		"SR_Code": "8",
		"SR_Name": "မကွေးတိုင်းဒေသကြီး",
		"Tsp_Code": "118",
		"Tsp_Name": "ထီးလင်း"
	  },
	  {
		"SR_Code": "15",
		"SR_Name": "နေပြည်တော်",
		"Tsp_Code": "160",
		"Tsp_Name": "ဒက္ခိဏသီရိ"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "290",
		"Tsp_Name": "ဒဂုံ"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "293",
		"Tsp_Name": "ဒဂုံမြို့သစ်(ဆိပ်ကမ်း)"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "294",
		"Tsp_Name": "ဒဂုံမြို့သစ်(တောင်ပိုင်း)"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "292",
		"Tsp_Name": "ဒဂုံမြို့သစ်(မြောက်ပိုင်း)"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "291",
		"Tsp_Name": "ဒဂုံမြို့သစ်(အရှေ့ပိုင်း)"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "295",
		"Tsp_Name": "ဒလ"
	  },
	  {
		"SR_Code": "7",
		"SR_Name": "ပဲခူးတိုင်းဒေသကြီး",
		"Tsp_Code": "29",
		"Tsp_Name": "ဒိုက်ဦး"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "213",
		"Tsp_Name": "ဒီပဲယင်း"
	  },
	  {
		"SR_Code": "2",
		"SR_Name": "ကယားပြည်နယ်",
		"Tsp_Code": "83",
		"Tsp_Name": "ဒီးမော့ဆို"
	  },
	  {
		"SR_Code": "14",
		"SR_Name": "ဧရာဝတီတိုင်းဒေသကြီး",
		"Tsp_Code": "3",
		"Tsp_Name": "ဒေးဒရဲ"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "296",
		"Tsp_Name": "ဒေါပုံ"
	  },
	  {
		"SR_Code": "14",
		"SR_Name": "ဧရာဝတီတိုင်းဒေသကြီး",
		"Tsp_Code": "2",
		"Tsp_Name": "ဓနုဖြူ"
	  },
	  {
		"SR_Code": "7",
		"SR_Name": "ပဲခူးတိုင်းဒေသကြီး",
		"Tsp_Code": "45",
		"Tsp_Name": "နတ်တလင်း"
	  },
	  {
		"SR_Code": "8",
		"SR_Name": "မကွေးတိုင်းဒေသကြီး",
		"Tsp_Code": "106",
		"Tsp_Name": "နတ်မောက်"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "206",
		"Tsp_Name": "နန်းယွန်း"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "249",
		"Tsp_Name": "နမ္မတူ"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "247",
		"Tsp_Name": "နမ့်ခမ်း"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "270",
		"Tsp_Name": "နမ့်စန်"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "248",
		"Tsp_Name": "နမ့်ဆန်"
	  },
	  {
		"SR_Code": "9",
		"SR_Name": "မန္တလေးတိုင်းဒေသကြီး",
		"Tsp_Code": "135",
		"Tsp_Name": "နွားထိုးကြီး"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "250",
		"Tsp_Name": "နားဖန်"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "251",
		"Tsp_Name": "နောင်ချို"
	  },
	  {
		"SR_Code": "1",
		"SR_Name": "ကချင်ပြည်နယ်",
		"Tsp_Code": "75",
		"Tsp_Name": "နောင်မွန်း"
	  },
	  {
		"SR_Code": "8",
		"SR_Name": "မကွေးတိုင်းဒေသကြီး",
		"Tsp_Code": "108",
		"Tsp_Name": "ပခုက္ကူ"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "272",
		"Tsp_Name": "ပင်းတယ"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "209",
		"Tsp_Name": "ပင်လည်ဘူး"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "273",
		"Tsp_Name": "ပင်လောင်း"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "252",
		"Tsp_Name": "ပန်ဆန်း"
	  },
	  {
		"SR_Code": "14",
		"SR_Name": "ဧရာဝတီတိုင်းဒေသကြီး",
		"Tsp_Code": "19",
		"Tsp_Name": "ပန်းတနော်"
	  },
	  {
		"SR_Code": "7",
		"SR_Name": "ပဲခူးတိုင်းဒေသကြီး",
		"Tsp_Code": "46",
		"Tsp_Name": "ပန်းတောင်း"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "317",
		"Tsp_Name": "ပန်းဘဲတန်း"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "253",
		"Tsp_Name": "ပန်ဝိုင်း"
	  },
	  {
		"SR_Code": "15",
		"SR_Name": "နေပြည်တော်",
		"Tsp_Code": "163",
		"Tsp_Name": "ပျဥ်းမနား"
	  },
	  {
		"SR_Code": "9",
		"SR_Name": "မန္တလေးတိုင်းဒေသကြီး",
		"Tsp_Code": "138",
		"Tsp_Name": "ပျော်ဘွယ်"
	  },
	  {
		"SR_Code": "9",
		"SR_Name": "မန္တလေးတိုင်းဒေသကြီး",
		"Tsp_Code": "140",
		"Tsp_Name": "ပြင်ဦးလွင်"
	  },
	  {
		"SR_Code": "7",
		"SR_Name": "ပဲခူးတိုင်းဒေသကြီး",
		"Tsp_Code": "49",
		"Tsp_Name": "ပြည်"
	  },
	  {
		"SR_Code": "9",
		"SR_Name": "မန္တလေးတိုင်းဒေသကြီး",
		"Tsp_Code": "139",
		"Tsp_Name": "ပြည်ကြီးတံခွန်"
	  },
	  {
		"SR_Code": "4",
		"SR_Name": "ချင်းပြည်နယ်",
		"Tsp_Code": "60",
		"Tsp_Name": "ပလက်ဝ"
	  },
	  {
		"SR_Code": "8",
		"SR_Name": "မကွေးတိုင်းဒေသကြီး",
		"Tsp_Code": "110",
		"Tsp_Name": "ပွင့်ဖြူ"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "318",
		"Tsp_Name": "ပုဇွန်တောင်"
	  },
	  {
		"SR_Code": "11",
		"SR_Name": "ရခိုင်ပြည်နယ်",
		"Tsp_Code": "178",
		"Tsp_Name": "ပုဏ္ဏားကျွန်း"
	  },
	  {
		"SR_Code": "15",
		"SR_Name": "နေပြည်တော်",
		"Tsp_Code": "162",
		"Tsp_Name": "ပုဗ္ဗသီရိ"
	  },
	  {
		"SR_Code": "6",
		"SR_Name": "တနင်္သာရီတိုင်းဒေသကြီး",
		"Tsp_Code": "283",
		"Tsp_Name": "ပုလော"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "207",
		"Tsp_Name": "ပုလဲ"
	  },
	  {
		"SR_Code": "14",
		"SR_Name": "ဧရာဝတီတိုင်းဒေသကြီး",
		"Tsp_Code": "20",
		"Tsp_Name": "ပုသိမ်"
	  },
	  {
		"SR_Code": "9",
		"SR_Name": "မန္တလေးတိုင်းဒေသကြီး",
		"Tsp_Code": "137",
		"Tsp_Name": "ပုသိမ်ကြီး"
	  },
	  {
		"SR_Code": "1",
		"SR_Name": "ကချင်ပြည်နယ်",
		"Tsp_Code": "76",
		"Tsp_Name": "ပူတာအို"
	  },
	  {
		"SR_Code": "8",
		"SR_Name": "မကွေးတိုင်းဒေသကြီး",
		"Tsp_Code": "109",
		"Tsp_Name": "ပေါက်"
	  },
	  {
		"SR_Code": "7",
		"SR_Name": "ပဲခူးတိုင်းဒေသကြီး",
		"Tsp_Code": "47",
		"Tsp_Name": "ပေါက်ခေါင်း"
	  },
	  {
		"SR_Code": "11",
		"SR_Name": "ရခိုင်ပြည်နယ်",
		"Tsp_Code": "177",
		"Tsp_Name": "ပေါက်တော"
	  },
	  {
		"SR_Code": "10",
		"SR_Name": "မွန်ပြည်နယ်",
		"Tsp_Code": "155",
		"Tsp_Name": "ပေါင်"
	  },
	  {
		"SR_Code": "7",
		"SR_Name": "ပဲခူးတိုင်းဒေသကြီး",
		"Tsp_Code": "48",
		"Tsp_Name": "ပေါင်းတည်"
	  },
	  {
		"SR_Code": "7",
		"SR_Name": "ပဲခူးတိုင်းဒေသကြီး",
		"Tsp_Code": "28",
		"Tsp_Name": "ပဲခူး"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "271",
		"Tsp_Name": "ဖယ်ခုံ"
	  },
	  {
		"SR_Code": "14",
		"SR_Name": "ဧရာဝတီတိုင်းဒေသကြီး",
		"Tsp_Code": "21",
		"Tsp_Name": "ဖျာပုံ"
	  },
	  {
		"SR_Code": "2",
		"SR_Name": "ကယားပြည်နယ်",
		"Tsp_Code": "85",
		"Tsp_Name": "ဖရူဆို"
	  },
	  {
		"SR_Code": "7",
		"SR_Name": "ပဲခူးတိုင်းဒေသကြီး",
		"Tsp_Code": "35",
		"Tsp_Name": "ဖြူး"
	  },
	  {
		"SR_Code": "4",
		"SR_Name": "ချင်းပြည်နယ်",
		"Tsp_Code": "55",
		"Tsp_Name": "ဖလမ်း"
	  },
	  {
		"SR_Code": "1",
		"SR_Name": "ကချင်ပြည်နယ်",
		"Tsp_Code": "66",
		"Tsp_Name": "ဖားကန့်"
	  },
	  {
		"SR_Code": "2",
		"SR_Name": "ကယားပြည်နယ်",
		"Tsp_Code": "84",
		"Tsp_Name": "ဖာဆောင်း"
	  },
	  {
		"SR_Code": "3",
		"SR_Name": "ကရင်ပြည်နယ်",
		"Tsp_Code": "91",
		"Tsp_Name": "ဖာပွန်"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "208",
		"Tsp_Name": "ဖောင်းပြင်"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "185",
		"Tsp_Name": "ဗန်းမောက်"
	  },
	  {
		"SR_Code": "1",
		"SR_Name": "ကချင်ပြည်နယ်",
		"Tsp_Code": "64",
		"Tsp_Name": "ဗန်းမော်"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "287",
		"Tsp_Name": "ဗဟန်း"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "288",
		"Tsp_Name": "ဗိုလ်တထောင်"
	  },
	  {
		"SR_Code": "3",
		"SR_Name": "ကရင်ပြည်နယ်",
		"Tsp_Code": "90",
		"Tsp_Name": "ဘားအံ"
	  },
	  {
		"SR_Code": "14",
		"SR_Name": "ဧရာဝတီတိုင်းဒေသကြီး",
		"Tsp_Code": "1",
		"Tsp_Name": "ဘိုကလေး"
	  },
	  {
		"SR_Code": "10",
		"SR_Name": "မွန်ပြည်နယ်",
		"Tsp_Code": "149",
		"Tsp_Name": "ဘီးလင်း"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "186",
		"Tsp_Name": "ဘုတလင်"
	  },
	  {
		"SR_Code": "6",
		"SR_Name": "တနင်္သာရီတိုင်းဒေသကြီး",
		"Tsp_Code": "277",
		"Tsp_Name": "ဘုတ်ပြင်း"
	  },
	  {
		"SR_Code": "11",
		"SR_Name": "ရခိုင်ပြည်နယ်",
		"Tsp_Code": "168",
		"Tsp_Name": "ဘူးသီးတောင်"
	  },
	  {
		"SR_Code": "2",
		"SR_Name": "ကယားပြည်နယ်",
		"Tsp_Code": "82",
		"Tsp_Name": "ဘောလခဲ"
	  },
	  {
		"SR_Code": "8",
		"SR_Name": "မကွေးတိုင်းဒေသကြီး",
		"Tsp_Code": "100",
		"Tsp_Name": "မကွေး"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "242",
		"Tsp_Name": "မက်မန်း"
	  },
	  {
		"SR_Code": "1",
		"SR_Name": "ကချင်ပြည်နယ်",
		"Tsp_Code": "69",
		"Tsp_Name": "မချမ်းဘော"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "202",
		"Tsp_Name": "မင်းကင်း"
	  },
	  {
		"SR_Code": "4",
		"SR_Name": "ချင်းပြည်နယ်",
		"Tsp_Code": "59",
		"Tsp_Name": "မင်းတပ်"
	  },
	  {
		"SR_Code": "8",
		"SR_Name": "မကွေးတိုင်းဒေသကြီး",
		"Tsp_Code": "102",
		"Tsp_Name": "မင်းတုန်း"
	  },
	  {
		"SR_Code": "11",
		"SR_Name": "ရခိုင်ပြည်နယ်",
		"Tsp_Code": "173",
		"Tsp_Name": "မင်းပြား"
	  },
	  {
		"SR_Code": "8",
		"SR_Name": "မကွေးတိုင်းဒေသကြီး",
		"Tsp_Code": "101",
		"Tsp_Name": "မင်းဘူး"
	  },
	  {
		"SR_Code": "7",
		"SR_Name": "ပဲခူးတိုင်းဒေသကြီး",
		"Tsp_Code": "43",
		"Tsp_Name": "မင်းလှ"
	  },
	  {
		"SR_Code": "8",
		"SR_Name": "မကွေးတိုင်းဒေသကြီး",
		"Tsp_Code": "103",
		"Tsp_Name": "မင်းလှ"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "315",
		"Tsp_Name": "မင်္ဂလာတောင်ညွန့်"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "314",
		"Tsp_Name": "မင်္ဂလာဒုံ"
	  },
	  {
		"SR_Code": "1",
		"SR_Name": "ကချင်ပြည်နယ်",
		"Tsp_Code": "70",
		"Tsp_Name": "မံစီ"
	  },
	  {
		"SR_Code": "4",
		"SR_Name": "ချင်းပြည်နယ်",
		"Tsp_Code": "58",
		"Tsp_Name": "မတူပီ"
	  },
	  {
		"SR_Code": "9",
		"SR_Name": "မန္တလေးတိုင်းဒေသကြီး",
		"Tsp_Code": "128",
		"Tsp_Name": "မတ္တရာ"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "241",
		"Tsp_Name": "မန်တုံ"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "240",
		"Tsp_Name": "မဘိမ်း"
	  },
	  {
		"SR_Code": "2",
		"SR_Name": "ကယားပြည်နယ်",
		"Tsp_Code": "87",
		"Tsp_Name": "မယ်စဲ"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "313",
		"Tsp_Name": "မရမ်းကုန်း"
	  },
	  {
		"SR_Code": "9",
		"SR_Name": "မန္တလေးတိုင်းဒေသကြီး",
		"Tsp_Code": "133",
		"Tsp_Name": "မြင်းခြံ"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "205",
		"Tsp_Name": "မြင်းမူ"
	  },
	  {
		"SR_Code": "1",
		"SR_Name": "ကချင်ပြည်နယ်",
		"Tsp_Code": "74",
		"Tsp_Name": "မြစ်ကြီးနား"
	  },
	  {
		"SR_Code": "9",
		"SR_Name": "မန္တလေးတိုင်းဒေသကြီး",
		"Tsp_Code": "134",
		"Tsp_Name": "မြစ်သား"
	  },
	  {
		"SR_Code": "14",
		"SR_Name": "ဧရာဝတီတိုင်းဒေသကြီး",
		"Tsp_Code": "16",
		"Tsp_Name": "မြန်အောင်"
	  },
	  {
		"SR_Code": "3",
		"SR_Name": "ကရင်ပြည်နယ်",
		"Tsp_Code": "94",
		"Tsp_Name": "မြဝတီ"
	  },
	  {
		"SR_Code": "6",
		"SR_Name": "တနင်္သာရီတိုင်းဒေသကြီး",
		"Tsp_Code": "282",
		"Tsp_Name": "မြိတ်မြို့"
	  },
	  {
		"SR_Code": "8",
		"SR_Name": "မကွေးတိုင်းဒေသကြီး",
		"Tsp_Code": "104",
		"Tsp_Name": "မြိုင်"
	  },
	  {
		"SR_Code": "8",
		"SR_Name": "မကွေးတိုင်းဒေသကြီး",
		"Tsp_Code": "105",
		"Tsp_Name": "မြို့သစ်"
	  },
	  {
		"SR_Code": "11",
		"SR_Name": "ရခိုင်ပြည်နယ်",
		"Tsp_Code": "176",
		"Tsp_Name": "မြေပုံ"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "316",
		"Tsp_Name": "မြောက်ဥက္ကလာပ"
	  },
	  {
		"SR_Code": "11",
		"SR_Name": "ရခိုင်ပြည်နယ်",
		"Tsp_Code": "174",
		"Tsp_Name": "မြောက်ဦး"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "204",
		"Tsp_Name": "မြောင်"
	  },
	  {
		"SR_Code": "14",
		"SR_Name": "ဧရာဝတီတိုင်းဒေသကြီး",
		"Tsp_Code": "17",
		"Tsp_Name": "မြောင်းမြ"
	  },
	  {
		"SR_Code": "9",
		"SR_Name": "မန္တလေးတိုင်းဒေသကြီး",
		"Tsp_Code": "130",
		"Tsp_Name": "မလှိုင်"
	  },
	  {
		"SR_Code": "9",
		"SR_Name": "မန္တလေးတိုင်းဒေသကြီး",
		"Tsp_Code": "129",
		"Tsp_Name": "မဟာအောင်မြေ"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "301",
		"Tsp_Name": "မှော်ဘီ"
	  },
	  {
		"SR_Code": "14",
		"SR_Name": "ဧရာဝတီတိုင်းဒေသကြီး",
		"Tsp_Code": "14",
		"Tsp_Name": "မအူပင်"
	  },
	  {
		"SR_Code": "11",
		"SR_Name": "ရခိုင်ပြည်နယ်",
		"Tsp_Code": "175",
		"Tsp_Name": "မာန်အောင်"
	  },
	  {
		"SR_Code": "9",
		"SR_Name": "မန္တလေးတိုင်းဒေသကြီး",
		"Tsp_Code": "131",
		"Tsp_Name": "မိတ္ထီလာ"
	  },
	  {
		"SR_Code": "9",
		"SR_Name": "မန္တလေးတိုင်းဒေသကြီး",
		"Tsp_Code": "132",
		"Tsp_Name": "မိုးကုတ်"
	  },
	  {
		"SR_Code": "1",
		"SR_Name": "ကချင်ပြည်နယ်",
		"Tsp_Code": "71",
		"Tsp_Name": "မိုးကောင်း"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "266",
		"Tsp_Name": "မိုင်းကိုင်"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "224",
		"Tsp_Name": "မိုင်းခတ်"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "223",
		"Tsp_Name": "မိုင်းဆတ်"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "227",
		"Tsp_Name": "မိုင်းတုံ"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "269",
		"Tsp_Name": "မိုင်းပန်"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "226",
		"Tsp_Name": "မိုင်းပျဉ်း"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "222",
		"Tsp_Name": "မိုင်းဖြတ်"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "243",
		"Tsp_Name": "မိုင်းမော"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "228",
		"Tsp_Name": "မိုင်းယန်း"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "229",
		"Tsp_Name": "မိုင်းယောင်း"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "245",
		"Tsp_Name": "မိုင်းရယ်"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "267",
		"Tsp_Name": "မိုင်းရှူး"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "225",
		"Tsp_Name": "မိုင်းလား"
	  },
	  {
		"SR_Code": "1",
		"SR_Name": "ကချင်ပြည်နယ်",
		"Tsp_Code": "72",
		"Tsp_Name": "မိုးညှင်း"
	  },
	  {
		"SR_Code": "7",
		"SR_Name": "ပဲခူးတိုင်းဒေသကြီး",
		"Tsp_Code": "44",
		"Tsp_Name": "မိုးညို"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "268",
		"Tsp_Name": "မိုးနဲ"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "244",
		"Tsp_Name": "မိုးမိတ်"
	  },
	  {
		"SR_Code": "1",
		"SR_Name": "ကချင်ပြည်နယ်",
		"Tsp_Code": "73",
		"Tsp_Name": "မိုးမောက်"
	  },
	  {
		"SR_Code": "10",
		"SR_Name": "မွန်ပြည်နယ်",
		"Tsp_Code": "154",
		"Tsp_Name": "မုဒုံ"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "203",
		"Tsp_Name": "မုံရွာ"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "246",
		"Tsp_Name": "မူဆယ်"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "265",
		"Tsp_Name": "မောက်မယ်"
	  },
	  {
		"SR_Code": "11",
		"SR_Name": "ရခိုင်ပြည်နယ်",
		"Tsp_Code": "172",
		"Tsp_Name": "မောင်တော"
	  },
	  {
		"SR_Code": "10",
		"SR_Name": "မွန်ပြည်နယ်",
		"Tsp_Code": "153",
		"Tsp_Name": "မော်လမြိုင်"
	  },
	  {
		"SR_Code": "14",
		"SR_Name": "ဧရာဝတီတိုင်းဒေသကြီး",
		"Tsp_Code": "15",
		"Tsp_Name": "မော်လမြိုင်ကျွန်း"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "201",
		"Tsp_Name": "မော်လိုက်"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "219",
		"Tsp_Name": "ယင်းမာပင်"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "330",
		"Tsp_Name": "ရန်ကင်း"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "263",
		"Tsp_Name": "ရပ်စောက်"
	  },
	  {
		"SR_Code": "9",
		"SR_Name": "မန္တလေးတိုင်းဒေသကြီး",
		"Tsp_Code": "148",
		"Tsp_Name": "ရမည်းသင်း"
	  },
	  {
		"SR_Code": "11",
		"SR_Name": "ရခိုင်ပြည်နယ်",
		"Tsp_Code": "179",
		"Tsp_Name": "ရမ်းဗြဲ"
	  },
	  {
		"SR_Code": "7",
		"SR_Name": "ပဲခူးတိုင်းဒေသကြီး",
		"Tsp_Code": "36",
		"Tsp_Name": "ရွှေကျင်"
	  },
	  {
		"SR_Code": "1",
		"SR_Name": "ကချင်ပြည်နယ်",
		"Tsp_Code": "77",
		"Tsp_Name": "ရွှေကူ"
	  },
	  {
		"SR_Code": "7",
		"SR_Name": "ပဲခူးတိုင်းဒေသကြီး",
		"Tsp_Code": "50",
		"Tsp_Name": "ရွှေတောင်"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "321",
		"Tsp_Name": "ရွှေပြည်သာ"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "212",
		"Tsp_Name": "ရွှေဘို"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "275",
		"Tsp_Name": "ရွာငံ"
	  },
	  {
		"SR_Code": "11",
		"SR_Name": "ရခိုင်ပြည်နယ်",
		"Tsp_Code": "180",
		"Tsp_Name": "ရသေ့တောင်"
	  },
	  {
		"SR_Code": "2",
		"SR_Name": "ကယားပြည်နယ်",
		"Tsp_Code": "88",
		"Tsp_Name": "ရှားတော"
	  },
	  {
		"SR_Code": "10",
		"SR_Name": "မွန်ပြည်နယ်",
		"Tsp_Code": "158",
		"Tsp_Name": "ရေး"
	  },
	  {
		"SR_Code": "14",
		"SR_Name": "ဧရာဝတီတိုင်းဒေသကြီး",
		"Tsp_Code": "26",
		"Tsp_Name": "ရေကြည်"
	  },
	  {
		"SR_Code": "8",
		"SR_Name": "မကွေးတိုင်းဒေသကြီး",
		"Tsp_Code": "120",
		"Tsp_Name": "ရေစကြို"
	  },
	  {
		"SR_Code": "7",
		"SR_Name": "ပဲခူးတိုင်းဒေသကြီး",
		"Tsp_Code": "27",
		"Tsp_Name": "ရေတာရှည်"
	  },
	  {
		"SR_Code": "8",
		"SR_Name": "မကွေးတိုင်းဒေသကြီး",
		"Tsp_Code": "119",
		"Tsp_Name": "ရေနံချောင်း"
	  },
	  {
		"SR_Code": "6",
		"SR_Name": "တနင်္သာရီတိုင်းဒေသကြီး",
		"Tsp_Code": "276",
		"Tsp_Name": "ရေဖြူ"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "220",
		"Tsp_Name": "ရေဦး"
	  },
	  {
		"SR_Code": "7",
		"SR_Name": "ပဲခူးတိုင်းဒေသကြီး",
		"Tsp_Code": "42",
		"Tsp_Name": "လက်ပံတန်း"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "262",
		"Tsp_Name": "လင်းခေး"
	  },
	  {
		"SR_Code": "14",
		"SR_Name": "ဧရာဝတီတိုင်းဒေသကြီး",
		"Tsp_Code": "12",
		"Tsp_Name": "လပွတ္တာ"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "311",
		"Tsp_Name": "လမ်းမတော်"
	  },
	  {
		"SR_Code": "15",
		"SR_Name": "နေပြည်တော်",
		"Tsp_Code": "161",
		"Tsp_Name": "လယ်ဝေး"
	  },
	  {
		"SR_Code": "2",
		"SR_Name": "ကယားပြည်နယ်",
		"Tsp_Code": "86",
		"Tsp_Name": "လွိုင်ကော်"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "264",
		"Tsp_Name": "လွိုင်လင်"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "312",
		"Tsp_Name": "လသာ"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "199",
		"Tsp_Name": "လဟယ်"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "300",
		"Tsp_Name": "လှည်းကူး"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "297",
		"Tsp_Name": "လှိုင်"
	  },
	  {
		"SR_Code": "3",
		"SR_Name": "ကရင်ပြည်နယ်",
		"Tsp_Code": "89",
		"Tsp_Name": "လှိုင်းဘွဲ့"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "299",
		"Tsp_Name": "လှိုင်သာယာ (အနောက်ပိုင်း)"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "298",
		"Tsp_Name": "လှိုင်သာယာ (အရှေ့ပိုင်း)"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "238",
		"Tsp_Name": "လားရှိုး"
	  },
	  {
		"SR_Code": "14",
		"SR_Name": "ဧရာဝတီတိုင်းဒေသကြီး",
		"Tsp_Code": "13",
		"Tsp_Name": "လေးမျက်နှာ"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "200",
		"Tsp_Name": "လေရှီး"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "239",
		"Tsp_Name": "လောက်ကိုင်"
	  },
	  {
		"SR_Code": "6",
		"SR_Name": "တနင်္သာရီတိုင်းဒေသကြီး",
		"Tsp_Code": "281",
		"Tsp_Name": "လောင်းလုံး"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "261",
		"Tsp_Name": "လဲချား"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "217",
		"Tsp_Name": "ဝက်လက်"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "218",
		"Tsp_Name": "ဝန်းသို"
	  },
	  {
		"SR_Code": "9",
		"SR_Name": "မန္တလေးတိုင်းဒေသကြီး",
		"Tsp_Code": "147",
		"Tsp_Name": "ဝမ်းတွင်း"
	  },
	  {
		"SR_Code": "14",
		"SR_Name": "ဧရာဝတီတိုင်းဒေသကြီး",
		"Tsp_Code": "23",
		"Tsp_Name": "ဝါးခယ်မ"
	  },
	  {
		"SR_Code": "1",
		"SR_Name": "ကချင်ပြည်နယ်",
		"Tsp_Code": "81",
		"Tsp_Name": "ဝိုင်းမော်"
	  },
	  {
		"SR_Code": "7",
		"SR_Name": "ပဲခူးတိုင်းဒေသကြီး",
		"Tsp_Code": "39",
		"Tsp_Name": "ဝေါ"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "327",
		"Tsp_Name": "သင်္ဃန်းကျွန်း"
	  },
	  {
		"SR_Code": "11",
		"SR_Name": "ရခိုင်ပြည်နယ်",
		"Tsp_Code": "182",
		"Tsp_Name": "သံတွဲ"
	  },
	  {
		"SR_Code": "3",
		"SR_Name": "ကရင်ပြည်နယ်",
		"Tsp_Code": "95",
		"Tsp_Name": "သံတောင်ကြီး"
	  },
	  {
		"SR_Code": "10",
		"SR_Name": "မွန်ပြည်နယ်",
		"Tsp_Code": "157",
		"Tsp_Name": "သထုံ"
	  },
	  {
		"SR_Code": "7",
		"SR_Name": "ပဲခူးတိုင်းဒေသကြီး",
		"Tsp_Code": "38",
		"Tsp_Name": "သနပ်ပင်"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "326",
		"Tsp_Name": "သန်လျင်"
	  },
	  {
		"SR_Code": "9",
		"SR_Name": "မန္တလေးတိုင်းဒေသကြီး",
		"Tsp_Code": "145",
		"Tsp_Name": "သပိတ်ကျင်း"
	  },
	  {
		"SR_Code": "10",
		"SR_Name": "မွန်ပြည်နယ်",
		"Tsp_Code": "156",
		"Tsp_Name": "သံဖြူဇရပ်"
	  },
	  {
		"SR_Code": "8",
		"SR_Name": "မကွေးတိုင်းဒေသကြီး",
		"Tsp_Code": "117",
		"Tsp_Name": "သရက်"
	  },
	  {
		"SR_Code": "6",
		"SR_Name": "တနင်္သာရီတိုင်းဒေသကြီး",
		"Tsp_Code": "285",
		"Tsp_Name": "သရက်ချောင်း"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "325",
		"Tsp_Name": "သာကေတ"
	  },
	  {
		"SR_Code": "9",
		"SR_Name": "မန္တလေးတိုင်းဒေသကြီး",
		"Tsp_Code": "146",
		"Tsp_Name": "သာစည်"
	  },
	  {
		"SR_Code": "14",
		"SR_Name": "ဧရာဝတီတိုင်းဒေသကြီး",
		"Tsp_Code": "22",
		"Tsp_Name": "သာပေါင်း"
	  },
	  {
		"SR_Code": "7",
		"SR_Name": "ပဲခူးတိုင်းဒေသကြီး",
		"Tsp_Code": "51",
		"Tsp_Name": "သာယာဝတီ"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "232",
		"Tsp_Name": "သိန္နီ"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "233",
		"Tsp_Name": "သီပေါ"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "328",
		"Tsp_Name": "သုံးခွ"
	  },
	  {
		"SR_Code": "7",
		"SR_Name": "ပဲခူးတိုင်းဒေသကြီး",
		"Tsp_Code": "52",
		"Tsp_Name": "သဲကုန်း"
	  },
	  {
		"SR_Code": "14",
		"SR_Name": "ဧရာဝတီတိုင်းဒေသကြီး",
		"Tsp_Code": "5",
		"Tsp_Name": "ဟင်္သာတ"
	  },
	  {
		"SR_Code": "4",
		"SR_Name": "ချင်းပြည်နယ်",
		"Tsp_Code": "56",
		"Tsp_Name": "ဟားခါး"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "231",
		"Tsp_Name": "ဟိုပန်"
	  },
	  {
		"SR_Code": "13",
		"SR_Name": "ရှမ်းပြည်နယ်",
		"Tsp_Code": "256",
		"Tsp_Name": "ဟိုပုံး"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "189",
		"Tsp_Name": "ဟုမ္မလင်း"
	  },
	  {
		"SR_Code": "1",
		"SR_Name": "ကချင်ပြည်နယ်",
		"Tsp_Code": "67",
		"Tsp_Name": "အင်ဂျန်းယန်"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "303",
		"Tsp_Name": "အင်းစိန်"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "190",
		"Tsp_Name": "အင်းတော်"
	  },
	  {
		"SR_Code": "14",
		"SR_Name": "ဧရာဝတီတိုင်းဒေသကြီး",
		"Tsp_Code": "6",
		"Tsp_Name": "အင်္ဂပူ"
	  },
	  {
		"SR_Code": "9",
		"SR_Name": "မန္တလေးတိုင်းဒေသကြီး",
		"Tsp_Code": "122",
		"Tsp_Name": "အမရပူရ"
	  },
	  {
		"SR_Code": "11",
		"SR_Name": "ရခိုင်ပြည်နယ်",
		"Tsp_Code": "167",
		"Tsp_Name": "အမ်း"
	  },
	  {
		"SR_Code": "5",
		"SR_Name": "စစ်ကိုင်းတိုင်းဒေသကြီး",
		"Tsp_Code": "184",
		"Tsp_Name": "အရာတော်"
	  },
	  {
		"SR_Code": "12",
		"SR_Name": "ရန်ကုန်တိုင်းဒေသကြီး",
		"Tsp_Code": "286",
		"Tsp_Name": "အလုံ"
	  },
	  {
		"SR_Code": "14",
		"SR_Name": "ဧရာဝတီတိုင်းဒေသကြီး",
		"Tsp_Code": "4",
		"Tsp_Name": "အိမ်မဲ"
	  },
	  {
		"SR_Code": "7",
		"SR_Name": "ပဲခူးတိုင်းဒေသကြီး",
		"Tsp_Code": "34",
		"Tsp_Name": "အုတ်တွင်း"
	  },
	  {
		"SR_Code": "7",
		"SR_Name": "ပဲခူးတိုင်းဒေသကြီး",
		"Tsp_Code": "54",
		"Tsp_Name": "အုတ်ဖို"
	  },
	  {
		"SR_Code": "9",
		"SR_Name": "မန္တလေးတိုင်းဒေသကြီး",
		"Tsp_Code": "123",
		"Tsp_Name": "အောင်မြေသာစံ"
	  },
	  {
		"SR_Code": "8",
		"SR_Name": "မကွေးတိုင်းဒေသကြီး",
		"Tsp_Code": "96",
		"Tsp_Name": "အောင်လံ"
	  },
	  {
		"SR_Code": "15",
		"SR_Name": "နေပြည်တော်",
		"Tsp_Code": "159",
		"Tsp_Name": "ဥတ္တရသီရိ"
	  }]
  }`
