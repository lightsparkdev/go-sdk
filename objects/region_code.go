
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
    "encoding/json"
    "strings"
)

// RegionCode The alpha-2 representation of a country, as defined by the ISO 3166-1 standard.
type RegionCode int
const(
    RegionCodeUndefined RegionCode = iota

    // RegionCodeAf The code representing the country of Afghanistan.
    RegionCodeAf
    // RegionCodeAx The code representing the country of Åland Islands.
    RegionCodeAx
    // RegionCodeAl The code representing the country of Albania.
    RegionCodeAl
    // RegionCodeDz The code representing the country of Algeria.
    RegionCodeDz
    // RegionCodeAs The code representing the country of American Samoa.
    RegionCodeAs
    // RegionCodeAd The code representing the country of Andorra.
    RegionCodeAd
    // RegionCodeAo The code representing the country of Angola.
    RegionCodeAo
    // RegionCodeAi The code representing the country of Anguilla.
    RegionCodeAi
    // RegionCodeAq The code representing the country of Antarctica.
    RegionCodeAq
    // RegionCodeAg The code representing the country of Antigua and Barbuda.
    RegionCodeAg
    // RegionCodeAr The code representing the country of Argentina.
    RegionCodeAr
    // RegionCodeAm The code representing the country of Armenia.
    RegionCodeAm
    // RegionCodeAw The code representing the country of Aruba.
    RegionCodeAw
    // RegionCodeAu The code representing the country of Australia.
    RegionCodeAu
    // RegionCodeAt The code representing the country of Austria.
    RegionCodeAt
    // RegionCodeAz The code representing the country of Azerbaijan.
    RegionCodeAz
    // RegionCodeBs The code representing the country of Bahamas.
    RegionCodeBs
    // RegionCodeBh The code representing the country of Bahrain.
    RegionCodeBh
    // RegionCodeBd The code representing the country of Bangladesh.
    RegionCodeBd
    // RegionCodeBb The code representing the country of Barbados.
    RegionCodeBb
    // RegionCodeBy The code representing the country of Belarus.
    RegionCodeBy
    // RegionCodeBe The code representing the country of Belgium.
    RegionCodeBe
    // RegionCodeBz The code representing the country of Belize.
    RegionCodeBz
    // RegionCodeBj The code representing the country of Benin.
    RegionCodeBj
    // RegionCodeBm The code representing the country of Bermuda.
    RegionCodeBm
    // RegionCodeBt The code representing the country of Bhutan.
    RegionCodeBt
    // RegionCodeBo The code representing the country of The Plurinational State of Bolivia.
    RegionCodeBo
    // RegionCodeBq The code representing the country of Bonaire, Sint Eustatius, and Saba.
    RegionCodeBq
    // RegionCodeBa The code representing the country of Bosnia and Herzegovina.
    RegionCodeBa
    // RegionCodeBw The code representing the country of Botswana.
    RegionCodeBw
    // RegionCodeBv The code representing the country of Bouvet Island.
    RegionCodeBv
    // RegionCodeBr The code representing the country of Brazil.
    RegionCodeBr
    // RegionCodeIo The code representing the country of British Indian Ocean Territory.
    RegionCodeIo
    // RegionCodeBn The code representing the country of Brunei Darussalam.
    RegionCodeBn
    // RegionCodeBg The code representing the country of Bulgaria.
    RegionCodeBg
    // RegionCodeBf The code representing the country of Burkina Faso.
    RegionCodeBf
    // RegionCodeBi The code representing the country of Burundi.
    RegionCodeBi
    // RegionCodeKh The code representing the country of Cambodia.
    RegionCodeKh
    // RegionCodeCm The code representing the country of Cameroon.
    RegionCodeCm
    // RegionCodeCa The code representing the country of Canada.
    RegionCodeCa
    // RegionCodeCv The code representing the country of Cape Verde.
    RegionCodeCv
    // RegionCodeKy The code representing the country of Cayman Islands.
    RegionCodeKy
    // RegionCodeCf The code representing the country of Central African Republic.
    RegionCodeCf
    // RegionCodeTd The code representing the country of Chad.
    RegionCodeTd
    // RegionCodeCl The code representing the country of Chile.
    RegionCodeCl
    // RegionCodeCn The code representing the country of China.
    RegionCodeCn
    // RegionCodeCx The code representing the country of Christmas Island.
    RegionCodeCx
    // RegionCodeCc The code representing the country of Cocos (Keeling) Islands.
    RegionCodeCc
    // RegionCodeCo The code representing the country of Colombia.
    RegionCodeCo
    // RegionCodeKm The code representing the country of Comoros.
    RegionCodeKm
    // RegionCodeCg The code representing the country of Congo.
    RegionCodeCg
    // RegionCodeCd The code representing the country of The Democratic Republic of the Congo.
    RegionCodeCd
    // RegionCodeCk The code representing the country of Cook Islands.
    RegionCodeCk
    // RegionCodeCr The code representing the country of Costa Rica.
    RegionCodeCr
    // RegionCodeCi The code representing the country of Côte d'Ivoire.
    RegionCodeCi
    // RegionCodeHr The code representing the country of Croatia.
    RegionCodeHr
    // RegionCodeCu The code representing the country of Cuba.
    RegionCodeCu
    // RegionCodeCw The code representing the country of Curaçao.
    RegionCodeCw
    // RegionCodeCy The code representing the country of Cyprus.
    RegionCodeCy
    // RegionCodeCz The code representing the country of Czech Republic.
    RegionCodeCz
    // RegionCodeDk The code representing the country of Denmark.
    RegionCodeDk
    // RegionCodeDj The code representing the country of Djibouti.
    RegionCodeDj
    // RegionCodeDm The code representing the country of Dominica.
    RegionCodeDm
    // RegionCodeDo The code representing the country of Dominican Republic.
    RegionCodeDo
    // RegionCodeEc The code representing the country of Ecuador.
    RegionCodeEc
    // RegionCodeEg The code representing the country of Egypt.
    RegionCodeEg
    // RegionCodeSv The code representing the country of El Salvador.
    RegionCodeSv
    // RegionCodeGq The code representing the country of Equatorial Guinea.
    RegionCodeGq
    // RegionCodeEr The code representing the country of Eritrea.
    RegionCodeEr
    // RegionCodeEe The code representing the country of Estonia.
    RegionCodeEe
    // RegionCodeEt The code representing the country of Ethiopia.
    RegionCodeEt
    // RegionCodeFk The code representing the country of Falkland Islands (Malvinas).
    RegionCodeFk
    // RegionCodeFo The code representing the country of Faroe Islands.
    RegionCodeFo
    // RegionCodeFj The code representing the country of Fiji.
    RegionCodeFj
    // RegionCodeFi The code representing the country of Finland.
    RegionCodeFi
    // RegionCodeFr The code representing the country of France.
    RegionCodeFr
    // RegionCodeGf The code representing the country of French Guiana.
    RegionCodeGf
    // RegionCodePf The code representing the country of French Polynesia.
    RegionCodePf
    // RegionCodeTf The code representing the country of French Southern Territories.
    RegionCodeTf
    // RegionCodeGa The code representing the country of Gabon.
    RegionCodeGa
    // RegionCodeGm The code representing the country of Gambia.
    RegionCodeGm
    // RegionCodeGe The code representing the country of Georgia.
    RegionCodeGe
    // RegionCodeDe The code representing the country of Germany.
    RegionCodeDe
    // RegionCodeGh The code representing the country of Ghana.
    RegionCodeGh
    // RegionCodeGi The code representing the country of Gibraltar.
    RegionCodeGi
    // RegionCodeGr The code representing the country of Greece.
    RegionCodeGr
    // RegionCodeGl The code representing the country of Greenland.
    RegionCodeGl
    // RegionCodeGd The code representing the country of Grenada.
    RegionCodeGd
    // RegionCodeGp The code representing the country of Guadeloupe.
    RegionCodeGp
    // RegionCodeGu The code representing the country of Guam.
    RegionCodeGu
    // RegionCodeGt The code representing the country of Guatemala.
    RegionCodeGt
    // RegionCodeGg The code representing the country of Guernsey.
    RegionCodeGg
    // RegionCodeGn The code representing the country of Guinea.
    RegionCodeGn
    // RegionCodeGw The code representing the country of Guinea-Bissau.
    RegionCodeGw
    // RegionCodeGy The code representing the country of Guyana.
    RegionCodeGy
    // RegionCodeHt The code representing the country of Haiti.
    RegionCodeHt
    // RegionCodeHm The code representing the country of Heard Island and McDonald Islands.
    RegionCodeHm
    // RegionCodeVa The code representing the country of Holy See (Vatican City State).
    RegionCodeVa
    // RegionCodeHn The code representing the country of Honduras.
    RegionCodeHn
    // RegionCodeHk The code representing the country of Hong Kong.
    RegionCodeHk
    // RegionCodeHu The code representing the country of Hungary.
    RegionCodeHu
    // RegionCodeIs The code representing the country of Iceland.
    RegionCodeIs
    // RegionCodeIn The code representing the country of India.
    RegionCodeIn
    // RegionCodeId The code representing the country of Indonesia.
    RegionCodeId
    // RegionCodeIr The code representing the country of Islamic Republic of Iran.
    RegionCodeIr
    // RegionCodeIq The code representing the country of Iraq.
    RegionCodeIq
    // RegionCodeIe The code representing the country of Ireland.
    RegionCodeIe
    // RegionCodeIm The code representing the country of Isle of Man.
    RegionCodeIm
    // RegionCodeIl The code representing the country of Israel.
    RegionCodeIl
    // RegionCodeIt The code representing the country of Italy.
    RegionCodeIt
    // RegionCodeJm The code representing the country of Jamaica.
    RegionCodeJm
    // RegionCodeJp The code representing the country of Japan.
    RegionCodeJp
    // RegionCodeJe The code representing the country of Jersey.
    RegionCodeJe
    // RegionCodeJo The code representing the country of Jordan.
    RegionCodeJo
    // RegionCodeKz The code representing the country of Kazakhstan.
    RegionCodeKz
    // RegionCodeKe The code representing the country of Kenya.
    RegionCodeKe
    // RegionCodeKi The code representing the country of Kiribati.
    RegionCodeKi
    // RegionCodeKp The code representing the country of Democratic People's Republic ofKorea.
    RegionCodeKp
    // RegionCodeKr The code representing the country of Republic of Korea.
    RegionCodeKr
    // RegionCodeKw The code representing the country of Kuwait.
    RegionCodeKw
    // RegionCodeKg The code representing the country of Kyrgyzstan.
    RegionCodeKg
    // RegionCodeLa The code representing the country of Lao People's Democratic Republic.
    RegionCodeLa
    // RegionCodeLv The code representing the country of Latvia.
    RegionCodeLv
    // RegionCodeLb The code representing the country of Lebanon.
    RegionCodeLb
    // RegionCodeLs The code representing the country of Lesotho.
    RegionCodeLs
    // RegionCodeLr The code representing the country of Liberia.
    RegionCodeLr
    // RegionCodeLy The code representing the country of Libya.
    RegionCodeLy
    // RegionCodeLi The code representing the country of Liechtenstein.
    RegionCodeLi
    // RegionCodeLt The code representing the country of Lithuania.
    RegionCodeLt
    // RegionCodeLu The code representing the country of Luxembourg.
    RegionCodeLu
    // RegionCodeMo The code representing the country of Macao.
    RegionCodeMo
    // RegionCodeMk The code representing the country of The Former Yugoslav Republic of Macedonia.
    RegionCodeMk
    // RegionCodeMg The code representing the country of Madagascar.
    RegionCodeMg
    // RegionCodeMw The code representing the country of Malawi.
    RegionCodeMw
    // RegionCodeMy The code representing the country of Malaysia.
    RegionCodeMy
    // RegionCodeMv The code representing the country of Maldives.
    RegionCodeMv
    // RegionCodeMl The code representing the country of Mali.
    RegionCodeMl
    // RegionCodeMt The code representing the country of Malta.
    RegionCodeMt
    // RegionCodeMh The code representing the country of Marshall Islands.
    RegionCodeMh
    // RegionCodeMq The code representing the country of Martinique.
    RegionCodeMq
    // RegionCodeMr The code representing the country of Mauritania.
    RegionCodeMr
    // RegionCodeMu The code representing the country of Mauritius.
    RegionCodeMu
    // RegionCodeYt The code representing the country of Mayotte.
    RegionCodeYt
    // RegionCodeMx The code representing the country of Mexico.
    RegionCodeMx
    // RegionCodeFm The code representing the country of Federated States ofMicronesia.
    RegionCodeFm
    // RegionCodeMd The code representing the country of Republic of Moldova.
    RegionCodeMd
    // RegionCodeMc The code representing the country of Monaco.
    RegionCodeMc
    // RegionCodeMn The code representing the country of Mongolia.
    RegionCodeMn
    // RegionCodeMe The code representing the country of Montenegro.
    RegionCodeMe
    // RegionCodeMs The code representing the country of Montserrat.
    RegionCodeMs
    // RegionCodeMa The code representing the country of Morocco.
    RegionCodeMa
    // RegionCodeMz The code representing the country of Mozambique.
    RegionCodeMz
    // RegionCodeMm The code representing the country of Myanmar.
    RegionCodeMm
    // RegionCodeNa The code representing the country of Namibia.
    RegionCodeNa
    // RegionCodeNr The code representing the country of Nauru.
    RegionCodeNr
    // RegionCodeNp The code representing the country of Nepal.
    RegionCodeNp
    // RegionCodeNl The code representing the country of Netherlands.
    RegionCodeNl
    // RegionCodeNc The code representing the country of New Caledonia.
    RegionCodeNc
    // RegionCodeNz The code representing the country of New Zealand.
    RegionCodeNz
    // RegionCodeNi The code representing the country of Nicaragua.
    RegionCodeNi
    // RegionCodeNe The code representing the country of Niger.
    RegionCodeNe
    // RegionCodeNg The code representing the country of Nigeria.
    RegionCodeNg
    // RegionCodeNu The code representing the country of Niue.
    RegionCodeNu
    // RegionCodeNf The code representing the country of Norfolk Island.
    RegionCodeNf
    // RegionCodeMp The code representing the country of Northern Mariana Islands.
    RegionCodeMp
    // RegionCodeNo The code representing the country of Norway.
    RegionCodeNo
    // RegionCodeOm The code representing the country of Oman.
    RegionCodeOm
    // RegionCodePk The code representing the country of Pakistan.
    RegionCodePk
    // RegionCodePw The code representing the country of Palau.
    RegionCodePw
    // RegionCodePs The code representing the country of State of Palestine.
    RegionCodePs
    // RegionCodePa The code representing the country of Panama.
    RegionCodePa
    // RegionCodePg The code representing the country of Papua New Guinea.
    RegionCodePg
    // RegionCodePy The code representing the country of Paraguay.
    RegionCodePy
    // RegionCodePe The code representing the country of Peru.
    RegionCodePe
    // RegionCodePh The code representing the country of Philippines.
    RegionCodePh
    // RegionCodePn The code representing the country of Pitcairn.
    RegionCodePn
    // RegionCodePl The code representing the country of Poland.
    RegionCodePl
    // RegionCodePt The code representing the country of Portugal.
    RegionCodePt
    // RegionCodePr The code representing the country of Puerto Rico.
    RegionCodePr
    // RegionCodeQa The code representing the country of Qatar.
    RegionCodeQa
    // RegionCodeRe The code representing the country of Réunion.
    RegionCodeRe
    // RegionCodeRo The code representing the country of Romania.
    RegionCodeRo
    // RegionCodeRu The code representing the country of Russian Federation.
    RegionCodeRu
    // RegionCodeRw The code representing the country of Rwanda.
    RegionCodeRw
    // RegionCodeBl The code representing the country of Saint Barthélemy.
    RegionCodeBl
    // RegionCodeSh The code representing the country of Saint Helena  Ascension and Tristan da Cunha.
    RegionCodeSh
    // RegionCodeKn The code representing the country of Saint Kitts and Nevis.
    RegionCodeKn
    // RegionCodeLc The code representing the country of Saint Lucia.
    RegionCodeLc
    // RegionCodeMf The code representing the country of Saint Martin (French part).
    RegionCodeMf
    // RegionCodePm The code representing the country of Saint Pierre and Miquelon.
    RegionCodePm
    // RegionCodeVc The code representing the country of Saint Vincent and the Grenadines.
    RegionCodeVc
    // RegionCodeWs The code representing the country of Samoa.
    RegionCodeWs
    // RegionCodeSm The code representing the country of San Marino.
    RegionCodeSm
    // RegionCodeSt The code representing the country of Sao Tome and Principe.
    RegionCodeSt
    // RegionCodeSa The code representing the country of Saudi Arabia.
    RegionCodeSa
    // RegionCodeSn The code representing the country of Senegal.
    RegionCodeSn
    // RegionCodeRs The code representing the country of Serbia.
    RegionCodeRs
    // RegionCodeSc The code representing the country of Seychelles.
    RegionCodeSc
    // RegionCodeSl The code representing the country of Sierra Leone.
    RegionCodeSl
    // RegionCodeSg The code representing the country of Singapore.
    RegionCodeSg
    // RegionCodeSx The code representing the country of Sint Maarten (Dutch part).
    RegionCodeSx
    // RegionCodeSk The code representing the country of Slovakia.
    RegionCodeSk
    // RegionCodeSi The code representing the country of Slovenia.
    RegionCodeSi
    // RegionCodeSb The code representing the country of Solomon Islands.
    RegionCodeSb
    // RegionCodeSo The code representing the country of Somalia.
    RegionCodeSo
    // RegionCodeZa The code representing the country of South Africa.
    RegionCodeZa
    // RegionCodeGs The code representing the country of South Georgia and the South Sandwich Islands.
    RegionCodeGs
    // RegionCodeSs The code representing the country of South Sudan.
    RegionCodeSs
    // RegionCodeEs The code representing the country of Spain.
    RegionCodeEs
    // RegionCodeLk The code representing the country of Sri Lanka.
    RegionCodeLk
    // RegionCodeSd The code representing the country of Sudan.
    RegionCodeSd
    // RegionCodeSr The code representing the country of Suriname.
    RegionCodeSr
    // RegionCodeSj The code representing the country of Svalbard and Jan Mayen.
    RegionCodeSj
    // RegionCodeSz The code representing the country of Swaziland.
    RegionCodeSz
    // RegionCodeSe The code representing the country of Sweden.
    RegionCodeSe
    // RegionCodeCh The code representing the country of Switzerland.
    RegionCodeCh
    // RegionCodeSy The code representing the country of Syrian Arab Republic.
    RegionCodeSy
    // RegionCodeTw The code representing the country of Taiwan, Province of China.
    RegionCodeTw
    // RegionCodeTj The code representing the country of Tajikistan.
    RegionCodeTj
    // RegionCodeTz The code representing the country of United Republic of Tanzania.
    RegionCodeTz
    // RegionCodeTh The code representing the country of Thailand.
    RegionCodeTh
    // RegionCodeTl The code representing the country of Timor-Leste.
    RegionCodeTl
    // RegionCodeTg The code representing the country of Togo.
    RegionCodeTg
    // RegionCodeTk The code representing the country of Tokelau.
    RegionCodeTk
    // RegionCodeTo The code representing the country of Tonga.
    RegionCodeTo
    // RegionCodeTt The code representing the country of Trinidad and Tobago.
    RegionCodeTt
    // RegionCodeTn The code representing the country of Tunisia.
    RegionCodeTn
    // RegionCodeTr The code representing the country of Turkey.
    RegionCodeTr
    // RegionCodeTm The code representing the country of Turkmenistan.
    RegionCodeTm
    // RegionCodeTc The code representing the country of Turks and Caicos Islands.
    RegionCodeTc
    // RegionCodeTv The code representing the country of Tuvalu.
    RegionCodeTv
    // RegionCodeUg The code representing the country of Uganda.
    RegionCodeUg
    // RegionCodeUa The code representing the country of Ukraine.
    RegionCodeUa
    // RegionCodeAe The code representing the country of United Arab Emirates.
    RegionCodeAe
    // RegionCodeGb The code representing the country of United Kingdom.
    RegionCodeGb
    // RegionCodeUs The code representing the country of United States.
    RegionCodeUs
    // RegionCodeUm The code representing the country of United States Minor Outlying Islands.
    RegionCodeUm
    // RegionCodeUy The code representing the country of Uruguay.
    RegionCodeUy
    // RegionCodeUz The code representing the country of Uzbekistan.
    RegionCodeUz
    // RegionCodeVu The code representing the country of Vanuatu.
    RegionCodeVu
    // RegionCodeVe The code representing the country of Bolivarian Republic of Venezuela.
    RegionCodeVe
    // RegionCodeVn The code representing the country of Viet Nam.
    RegionCodeVn
    // RegionCodeVg The code representing the country of British Virgin Islands.
    RegionCodeVg
    // RegionCodeVi The code representing the country of U.S. Virgin Islands.
    RegionCodeVi
    // RegionCodeWf The code representing the country of Wallis and Futuna.
    RegionCodeWf
    // RegionCodeEh The code representing the country of Western Sahara.
    RegionCodeEh
    // RegionCodeYe The code representing the country of Yemen.
    RegionCodeYe
    // RegionCodeZm The code representing the country of Zambia.
    RegionCodeZm
    // RegionCodeZw The code representing the country of Zimbabwe.
    RegionCodeZw

)

func (a *RegionCode) UnmarshalJSON(b []byte) error {
    var s string
    if err := json.Unmarshal(b, &s); err != nil {
        return err
    }
    switch s {
    default:
        *a = RegionCodeUndefined
    case "AF":
        *a = RegionCodeAf
    case "AX":
        *a = RegionCodeAx
    case "AL":
        *a = RegionCodeAl
    case "DZ":
        *a = RegionCodeDz
    case "AS":
        *a = RegionCodeAs
    case "AD":
        *a = RegionCodeAd
    case "AO":
        *a = RegionCodeAo
    case "AI":
        *a = RegionCodeAi
    case "AQ":
        *a = RegionCodeAq
    case "AG":
        *a = RegionCodeAg
    case "AR":
        *a = RegionCodeAr
    case "AM":
        *a = RegionCodeAm
    case "AW":
        *a = RegionCodeAw
    case "AU":
        *a = RegionCodeAu
    case "AT":
        *a = RegionCodeAt
    case "AZ":
        *a = RegionCodeAz
    case "BS":
        *a = RegionCodeBs
    case "BH":
        *a = RegionCodeBh
    case "BD":
        *a = RegionCodeBd
    case "BB":
        *a = RegionCodeBb
    case "BY":
        *a = RegionCodeBy
    case "BE":
        *a = RegionCodeBe
    case "BZ":
        *a = RegionCodeBz
    case "BJ":
        *a = RegionCodeBj
    case "BM":
        *a = RegionCodeBm
    case "BT":
        *a = RegionCodeBt
    case "BO":
        *a = RegionCodeBo
    case "BQ":
        *a = RegionCodeBq
    case "BA":
        *a = RegionCodeBa
    case "BW":
        *a = RegionCodeBw
    case "BV":
        *a = RegionCodeBv
    case "BR":
        *a = RegionCodeBr
    case "IO":
        *a = RegionCodeIo
    case "BN":
        *a = RegionCodeBn
    case "BG":
        *a = RegionCodeBg
    case "BF":
        *a = RegionCodeBf
    case "BI":
        *a = RegionCodeBi
    case "KH":
        *a = RegionCodeKh
    case "CM":
        *a = RegionCodeCm
    case "CA":
        *a = RegionCodeCa
    case "CV":
        *a = RegionCodeCv
    case "KY":
        *a = RegionCodeKy
    case "CF":
        *a = RegionCodeCf
    case "TD":
        *a = RegionCodeTd
    case "CL":
        *a = RegionCodeCl
    case "CN":
        *a = RegionCodeCn
    case "CX":
        *a = RegionCodeCx
    case "CC":
        *a = RegionCodeCc
    case "CO":
        *a = RegionCodeCo
    case "KM":
        *a = RegionCodeKm
    case "CG":
        *a = RegionCodeCg
    case "CD":
        *a = RegionCodeCd
    case "CK":
        *a = RegionCodeCk
    case "CR":
        *a = RegionCodeCr
    case "CI":
        *a = RegionCodeCi
    case "HR":
        *a = RegionCodeHr
    case "CU":
        *a = RegionCodeCu
    case "CW":
        *a = RegionCodeCw
    case "CY":
        *a = RegionCodeCy
    case "CZ":
        *a = RegionCodeCz
    case "DK":
        *a = RegionCodeDk
    case "DJ":
        *a = RegionCodeDj
    case "DM":
        *a = RegionCodeDm
    case "DO":
        *a = RegionCodeDo
    case "EC":
        *a = RegionCodeEc
    case "EG":
        *a = RegionCodeEg
    case "SV":
        *a = RegionCodeSv
    case "GQ":
        *a = RegionCodeGq
    case "ER":
        *a = RegionCodeEr
    case "EE":
        *a = RegionCodeEe
    case "ET":
        *a = RegionCodeEt
    case "FK":
        *a = RegionCodeFk
    case "FO":
        *a = RegionCodeFo
    case "FJ":
        *a = RegionCodeFj
    case "FI":
        *a = RegionCodeFi
    case "FR":
        *a = RegionCodeFr
    case "GF":
        *a = RegionCodeGf
    case "PF":
        *a = RegionCodePf
    case "TF":
        *a = RegionCodeTf
    case "GA":
        *a = RegionCodeGa
    case "GM":
        *a = RegionCodeGm
    case "GE":
        *a = RegionCodeGe
    case "DE":
        *a = RegionCodeDe
    case "GH":
        *a = RegionCodeGh
    case "GI":
        *a = RegionCodeGi
    case "GR":
        *a = RegionCodeGr
    case "GL":
        *a = RegionCodeGl
    case "GD":
        *a = RegionCodeGd
    case "GP":
        *a = RegionCodeGp
    case "GU":
        *a = RegionCodeGu
    case "GT":
        *a = RegionCodeGt
    case "GG":
        *a = RegionCodeGg
    case "GN":
        *a = RegionCodeGn
    case "GW":
        *a = RegionCodeGw
    case "GY":
        *a = RegionCodeGy
    case "HT":
        *a = RegionCodeHt
    case "HM":
        *a = RegionCodeHm
    case "VA":
        *a = RegionCodeVa
    case "HN":
        *a = RegionCodeHn
    case "HK":
        *a = RegionCodeHk
    case "HU":
        *a = RegionCodeHu
    case "IS":
        *a = RegionCodeIs
    case "IN":
        *a = RegionCodeIn
    case "ID":
        *a = RegionCodeId
    case "IR":
        *a = RegionCodeIr
    case "IQ":
        *a = RegionCodeIq
    case "IE":
        *a = RegionCodeIe
    case "IM":
        *a = RegionCodeIm
    case "IL":
        *a = RegionCodeIl
    case "IT":
        *a = RegionCodeIt
    case "JM":
        *a = RegionCodeJm
    case "JP":
        *a = RegionCodeJp
    case "JE":
        *a = RegionCodeJe
    case "JO":
        *a = RegionCodeJo
    case "KZ":
        *a = RegionCodeKz
    case "KE":
        *a = RegionCodeKe
    case "KI":
        *a = RegionCodeKi
    case "KP":
        *a = RegionCodeKp
    case "KR":
        *a = RegionCodeKr
    case "KW":
        *a = RegionCodeKw
    case "KG":
        *a = RegionCodeKg
    case "LA":
        *a = RegionCodeLa
    case "LV":
        *a = RegionCodeLv
    case "LB":
        *a = RegionCodeLb
    case "LS":
        *a = RegionCodeLs
    case "LR":
        *a = RegionCodeLr
    case "LY":
        *a = RegionCodeLy
    case "LI":
        *a = RegionCodeLi
    case "LT":
        *a = RegionCodeLt
    case "LU":
        *a = RegionCodeLu
    case "MO":
        *a = RegionCodeMo
    case "MK":
        *a = RegionCodeMk
    case "MG":
        *a = RegionCodeMg
    case "MW":
        *a = RegionCodeMw
    case "MY":
        *a = RegionCodeMy
    case "MV":
        *a = RegionCodeMv
    case "ML":
        *a = RegionCodeMl
    case "MT":
        *a = RegionCodeMt
    case "MH":
        *a = RegionCodeMh
    case "MQ":
        *a = RegionCodeMq
    case "MR":
        *a = RegionCodeMr
    case "MU":
        *a = RegionCodeMu
    case "YT":
        *a = RegionCodeYt
    case "MX":
        *a = RegionCodeMx
    case "FM":
        *a = RegionCodeFm
    case "MD":
        *a = RegionCodeMd
    case "MC":
        *a = RegionCodeMc
    case "MN":
        *a = RegionCodeMn
    case "ME":
        *a = RegionCodeMe
    case "MS":
        *a = RegionCodeMs
    case "MA":
        *a = RegionCodeMa
    case "MZ":
        *a = RegionCodeMz
    case "MM":
        *a = RegionCodeMm
    case "NA":
        *a = RegionCodeNa
    case "NR":
        *a = RegionCodeNr
    case "NP":
        *a = RegionCodeNp
    case "NL":
        *a = RegionCodeNl
    case "NC":
        *a = RegionCodeNc
    case "NZ":
        *a = RegionCodeNz
    case "NI":
        *a = RegionCodeNi
    case "NE":
        *a = RegionCodeNe
    case "NG":
        *a = RegionCodeNg
    case "NU":
        *a = RegionCodeNu
    case "NF":
        *a = RegionCodeNf
    case "MP":
        *a = RegionCodeMp
    case "NO":
        *a = RegionCodeNo
    case "OM":
        *a = RegionCodeOm
    case "PK":
        *a = RegionCodePk
    case "PW":
        *a = RegionCodePw
    case "PS":
        *a = RegionCodePs
    case "PA":
        *a = RegionCodePa
    case "PG":
        *a = RegionCodePg
    case "PY":
        *a = RegionCodePy
    case "PE":
        *a = RegionCodePe
    case "PH":
        *a = RegionCodePh
    case "PN":
        *a = RegionCodePn
    case "PL":
        *a = RegionCodePl
    case "PT":
        *a = RegionCodePt
    case "PR":
        *a = RegionCodePr
    case "QA":
        *a = RegionCodeQa
    case "RE":
        *a = RegionCodeRe
    case "RO":
        *a = RegionCodeRo
    case "RU":
        *a = RegionCodeRu
    case "RW":
        *a = RegionCodeRw
    case "BL":
        *a = RegionCodeBl
    case "SH":
        *a = RegionCodeSh
    case "KN":
        *a = RegionCodeKn
    case "LC":
        *a = RegionCodeLc
    case "MF":
        *a = RegionCodeMf
    case "PM":
        *a = RegionCodePm
    case "VC":
        *a = RegionCodeVc
    case "WS":
        *a = RegionCodeWs
    case "SM":
        *a = RegionCodeSm
    case "ST":
        *a = RegionCodeSt
    case "SA":
        *a = RegionCodeSa
    case "SN":
        *a = RegionCodeSn
    case "RS":
        *a = RegionCodeRs
    case "SC":
        *a = RegionCodeSc
    case "SL":
        *a = RegionCodeSl
    case "SG":
        *a = RegionCodeSg
    case "SX":
        *a = RegionCodeSx
    case "SK":
        *a = RegionCodeSk
    case "SI":
        *a = RegionCodeSi
    case "SB":
        *a = RegionCodeSb
    case "SO":
        *a = RegionCodeSo
    case "ZA":
        *a = RegionCodeZa
    case "GS":
        *a = RegionCodeGs
    case "SS":
        *a = RegionCodeSs
    case "ES":
        *a = RegionCodeEs
    case "LK":
        *a = RegionCodeLk
    case "SD":
        *a = RegionCodeSd
    case "SR":
        *a = RegionCodeSr
    case "SJ":
        *a = RegionCodeSj
    case "SZ":
        *a = RegionCodeSz
    case "SE":
        *a = RegionCodeSe
    case "CH":
        *a = RegionCodeCh
    case "SY":
        *a = RegionCodeSy
    case "TW":
        *a = RegionCodeTw
    case "TJ":
        *a = RegionCodeTj
    case "TZ":
        *a = RegionCodeTz
    case "TH":
        *a = RegionCodeTh
    case "TL":
        *a = RegionCodeTl
    case "TG":
        *a = RegionCodeTg
    case "TK":
        *a = RegionCodeTk
    case "TO":
        *a = RegionCodeTo
    case "TT":
        *a = RegionCodeTt
    case "TN":
        *a = RegionCodeTn
    case "TR":
        *a = RegionCodeTr
    case "TM":
        *a = RegionCodeTm
    case "TC":
        *a = RegionCodeTc
    case "TV":
        *a = RegionCodeTv
    case "UG":
        *a = RegionCodeUg
    case "UA":
        *a = RegionCodeUa
    case "AE":
        *a = RegionCodeAe
    case "GB":
        *a = RegionCodeGb
    case "US":
        *a = RegionCodeUs
    case "UM":
        *a = RegionCodeUm
    case "UY":
        *a = RegionCodeUy
    case "UZ":
        *a = RegionCodeUz
    case "VU":
        *a = RegionCodeVu
    case "VE":
        *a = RegionCodeVe
    case "VN":
        *a = RegionCodeVn
    case "VG":
        *a = RegionCodeVg
    case "VI":
        *a = RegionCodeVi
    case "WF":
        *a = RegionCodeWf
    case "EH":
        *a = RegionCodeEh
    case "YE":
        *a = RegionCodeYe
    case "ZM":
        *a = RegionCodeZm
    case "ZW":
        *a = RegionCodeZw

    }
    return nil
}

func (a RegionCode) StringValue() string {
    var s string
    switch a {
    default:
        s = "undefined"
    case RegionCodeAf:
        s = "AF"
    case RegionCodeAx:
        s = "AX"
    case RegionCodeAl:
        s = "AL"
    case RegionCodeDz:
        s = "DZ"
    case RegionCodeAs:
        s = "AS"
    case RegionCodeAd:
        s = "AD"
    case RegionCodeAo:
        s = "AO"
    case RegionCodeAi:
        s = "AI"
    case RegionCodeAq:
        s = "AQ"
    case RegionCodeAg:
        s = "AG"
    case RegionCodeAr:
        s = "AR"
    case RegionCodeAm:
        s = "AM"
    case RegionCodeAw:
        s = "AW"
    case RegionCodeAu:
        s = "AU"
    case RegionCodeAt:
        s = "AT"
    case RegionCodeAz:
        s = "AZ"
    case RegionCodeBs:
        s = "BS"
    case RegionCodeBh:
        s = "BH"
    case RegionCodeBd:
        s = "BD"
    case RegionCodeBb:
        s = "BB"
    case RegionCodeBy:
        s = "BY"
    case RegionCodeBe:
        s = "BE"
    case RegionCodeBz:
        s = "BZ"
    case RegionCodeBj:
        s = "BJ"
    case RegionCodeBm:
        s = "BM"
    case RegionCodeBt:
        s = "BT"
    case RegionCodeBo:
        s = "BO"
    case RegionCodeBq:
        s = "BQ"
    case RegionCodeBa:
        s = "BA"
    case RegionCodeBw:
        s = "BW"
    case RegionCodeBv:
        s = "BV"
    case RegionCodeBr:
        s = "BR"
    case RegionCodeIo:
        s = "IO"
    case RegionCodeBn:
        s = "BN"
    case RegionCodeBg:
        s = "BG"
    case RegionCodeBf:
        s = "BF"
    case RegionCodeBi:
        s = "BI"
    case RegionCodeKh:
        s = "KH"
    case RegionCodeCm:
        s = "CM"
    case RegionCodeCa:
        s = "CA"
    case RegionCodeCv:
        s = "CV"
    case RegionCodeKy:
        s = "KY"
    case RegionCodeCf:
        s = "CF"
    case RegionCodeTd:
        s = "TD"
    case RegionCodeCl:
        s = "CL"
    case RegionCodeCn:
        s = "CN"
    case RegionCodeCx:
        s = "CX"
    case RegionCodeCc:
        s = "CC"
    case RegionCodeCo:
        s = "CO"
    case RegionCodeKm:
        s = "KM"
    case RegionCodeCg:
        s = "CG"
    case RegionCodeCd:
        s = "CD"
    case RegionCodeCk:
        s = "CK"
    case RegionCodeCr:
        s = "CR"
    case RegionCodeCi:
        s = "CI"
    case RegionCodeHr:
        s = "HR"
    case RegionCodeCu:
        s = "CU"
    case RegionCodeCw:
        s = "CW"
    case RegionCodeCy:
        s = "CY"
    case RegionCodeCz:
        s = "CZ"
    case RegionCodeDk:
        s = "DK"
    case RegionCodeDj:
        s = "DJ"
    case RegionCodeDm:
        s = "DM"
    case RegionCodeDo:
        s = "DO"
    case RegionCodeEc:
        s = "EC"
    case RegionCodeEg:
        s = "EG"
    case RegionCodeSv:
        s = "SV"
    case RegionCodeGq:
        s = "GQ"
    case RegionCodeEr:
        s = "ER"
    case RegionCodeEe:
        s = "EE"
    case RegionCodeEt:
        s = "ET"
    case RegionCodeFk:
        s = "FK"
    case RegionCodeFo:
        s = "FO"
    case RegionCodeFj:
        s = "FJ"
    case RegionCodeFi:
        s = "FI"
    case RegionCodeFr:
        s = "FR"
    case RegionCodeGf:
        s = "GF"
    case RegionCodePf:
        s = "PF"
    case RegionCodeTf:
        s = "TF"
    case RegionCodeGa:
        s = "GA"
    case RegionCodeGm:
        s = "GM"
    case RegionCodeGe:
        s = "GE"
    case RegionCodeDe:
        s = "DE"
    case RegionCodeGh:
        s = "GH"
    case RegionCodeGi:
        s = "GI"
    case RegionCodeGr:
        s = "GR"
    case RegionCodeGl:
        s = "GL"
    case RegionCodeGd:
        s = "GD"
    case RegionCodeGp:
        s = "GP"
    case RegionCodeGu:
        s = "GU"
    case RegionCodeGt:
        s = "GT"
    case RegionCodeGg:
        s = "GG"
    case RegionCodeGn:
        s = "GN"
    case RegionCodeGw:
        s = "GW"
    case RegionCodeGy:
        s = "GY"
    case RegionCodeHt:
        s = "HT"
    case RegionCodeHm:
        s = "HM"
    case RegionCodeVa:
        s = "VA"
    case RegionCodeHn:
        s = "HN"
    case RegionCodeHk:
        s = "HK"
    case RegionCodeHu:
        s = "HU"
    case RegionCodeIs:
        s = "IS"
    case RegionCodeIn:
        s = "IN"
    case RegionCodeId:
        s = "ID"
    case RegionCodeIr:
        s = "IR"
    case RegionCodeIq:
        s = "IQ"
    case RegionCodeIe:
        s = "IE"
    case RegionCodeIm:
        s = "IM"
    case RegionCodeIl:
        s = "IL"
    case RegionCodeIt:
        s = "IT"
    case RegionCodeJm:
        s = "JM"
    case RegionCodeJp:
        s = "JP"
    case RegionCodeJe:
        s = "JE"
    case RegionCodeJo:
        s = "JO"
    case RegionCodeKz:
        s = "KZ"
    case RegionCodeKe:
        s = "KE"
    case RegionCodeKi:
        s = "KI"
    case RegionCodeKp:
        s = "KP"
    case RegionCodeKr:
        s = "KR"
    case RegionCodeKw:
        s = "KW"
    case RegionCodeKg:
        s = "KG"
    case RegionCodeLa:
        s = "LA"
    case RegionCodeLv:
        s = "LV"
    case RegionCodeLb:
        s = "LB"
    case RegionCodeLs:
        s = "LS"
    case RegionCodeLr:
        s = "LR"
    case RegionCodeLy:
        s = "LY"
    case RegionCodeLi:
        s = "LI"
    case RegionCodeLt:
        s = "LT"
    case RegionCodeLu:
        s = "LU"
    case RegionCodeMo:
        s = "MO"
    case RegionCodeMk:
        s = "MK"
    case RegionCodeMg:
        s = "MG"
    case RegionCodeMw:
        s = "MW"
    case RegionCodeMy:
        s = "MY"
    case RegionCodeMv:
        s = "MV"
    case RegionCodeMl:
        s = "ML"
    case RegionCodeMt:
        s = "MT"
    case RegionCodeMh:
        s = "MH"
    case RegionCodeMq:
        s = "MQ"
    case RegionCodeMr:
        s = "MR"
    case RegionCodeMu:
        s = "MU"
    case RegionCodeYt:
        s = "YT"
    case RegionCodeMx:
        s = "MX"
    case RegionCodeFm:
        s = "FM"
    case RegionCodeMd:
        s = "MD"
    case RegionCodeMc:
        s = "MC"
    case RegionCodeMn:
        s = "MN"
    case RegionCodeMe:
        s = "ME"
    case RegionCodeMs:
        s = "MS"
    case RegionCodeMa:
        s = "MA"
    case RegionCodeMz:
        s = "MZ"
    case RegionCodeMm:
        s = "MM"
    case RegionCodeNa:
        s = "NA"
    case RegionCodeNr:
        s = "NR"
    case RegionCodeNp:
        s = "NP"
    case RegionCodeNl:
        s = "NL"
    case RegionCodeNc:
        s = "NC"
    case RegionCodeNz:
        s = "NZ"
    case RegionCodeNi:
        s = "NI"
    case RegionCodeNe:
        s = "NE"
    case RegionCodeNg:
        s = "NG"
    case RegionCodeNu:
        s = "NU"
    case RegionCodeNf:
        s = "NF"
    case RegionCodeMp:
        s = "MP"
    case RegionCodeNo:
        s = "NO"
    case RegionCodeOm:
        s = "OM"
    case RegionCodePk:
        s = "PK"
    case RegionCodePw:
        s = "PW"
    case RegionCodePs:
        s = "PS"
    case RegionCodePa:
        s = "PA"
    case RegionCodePg:
        s = "PG"
    case RegionCodePy:
        s = "PY"
    case RegionCodePe:
        s = "PE"
    case RegionCodePh:
        s = "PH"
    case RegionCodePn:
        s = "PN"
    case RegionCodePl:
        s = "PL"
    case RegionCodePt:
        s = "PT"
    case RegionCodePr:
        s = "PR"
    case RegionCodeQa:
        s = "QA"
    case RegionCodeRe:
        s = "RE"
    case RegionCodeRo:
        s = "RO"
    case RegionCodeRu:
        s = "RU"
    case RegionCodeRw:
        s = "RW"
    case RegionCodeBl:
        s = "BL"
    case RegionCodeSh:
        s = "SH"
    case RegionCodeKn:
        s = "KN"
    case RegionCodeLc:
        s = "LC"
    case RegionCodeMf:
        s = "MF"
    case RegionCodePm:
        s = "PM"
    case RegionCodeVc:
        s = "VC"
    case RegionCodeWs:
        s = "WS"
    case RegionCodeSm:
        s = "SM"
    case RegionCodeSt:
        s = "ST"
    case RegionCodeSa:
        s = "SA"
    case RegionCodeSn:
        s = "SN"
    case RegionCodeRs:
        s = "RS"
    case RegionCodeSc:
        s = "SC"
    case RegionCodeSl:
        s = "SL"
    case RegionCodeSg:
        s = "SG"
    case RegionCodeSx:
        s = "SX"
    case RegionCodeSk:
        s = "SK"
    case RegionCodeSi:
        s = "SI"
    case RegionCodeSb:
        s = "SB"
    case RegionCodeSo:
        s = "SO"
    case RegionCodeZa:
        s = "ZA"
    case RegionCodeGs:
        s = "GS"
    case RegionCodeSs:
        s = "SS"
    case RegionCodeEs:
        s = "ES"
    case RegionCodeLk:
        s = "LK"
    case RegionCodeSd:
        s = "SD"
    case RegionCodeSr:
        s = "SR"
    case RegionCodeSj:
        s = "SJ"
    case RegionCodeSz:
        s = "SZ"
    case RegionCodeSe:
        s = "SE"
    case RegionCodeCh:
        s = "CH"
    case RegionCodeSy:
        s = "SY"
    case RegionCodeTw:
        s = "TW"
    case RegionCodeTj:
        s = "TJ"
    case RegionCodeTz:
        s = "TZ"
    case RegionCodeTh:
        s = "TH"
    case RegionCodeTl:
        s = "TL"
    case RegionCodeTg:
        s = "TG"
    case RegionCodeTk:
        s = "TK"
    case RegionCodeTo:
        s = "TO"
    case RegionCodeTt:
        s = "TT"
    case RegionCodeTn:
        s = "TN"
    case RegionCodeTr:
        s = "TR"
    case RegionCodeTm:
        s = "TM"
    case RegionCodeTc:
        s = "TC"
    case RegionCodeTv:
        s = "TV"
    case RegionCodeUg:
        s = "UG"
    case RegionCodeUa:
        s = "UA"
    case RegionCodeAe:
        s = "AE"
    case RegionCodeGb:
        s = "GB"
    case RegionCodeUs:
        s = "US"
    case RegionCodeUm:
        s = "UM"
    case RegionCodeUy:
        s = "UY"
    case RegionCodeUz:
        s = "UZ"
    case RegionCodeVu:
        s = "VU"
    case RegionCodeVe:
        s = "VE"
    case RegionCodeVn:
        s = "VN"
    case RegionCodeVg:
        s = "VG"
    case RegionCodeVi:
        s = "VI"
    case RegionCodeWf:
        s = "WF"
    case RegionCodeEh:
        s = "EH"
    case RegionCodeYe:
        s = "YE"
    case RegionCodeZm:
        s = "ZM"
    case RegionCodeZw:
        s = "ZW"

    }
    return s
}

func (a RegionCode) MarshalJSON() ([]byte, error) {
    s := a.StringValue()
    return json.Marshal(s)
}
