package model

import "strings"

func GetCoinSuffix() map[string]string {
	var suffix = map[string]string{
		"btc": "H/s",
		"bcc": "H/s",
		"bsv": "H/s",
		"bch": "H/s",
		"ltc": "H/s",
		"sbtc": "H/s",
		"ubtc": "H/s",
		"zec": "sol/s",
		"eth": "H/s",
		"etc": "H/s",
		"dcr": "H/s",
		"xmc": "H/s",
		"dash": "H/s",
		"grin": "g/s",
		"grin_c29": "g/s",
		"grin_c31": "g/s",
		"beam": "sol/s",
		"ae": "sol/s",
		"ckb": "H/s",
		"smart_sha256": "H/s",
		"trb": "H/s",
	}
	return suffix
}

func GetCoinSuffixByCoinType(coinType string) string {
	return GetCoinSuffix()[strings.ToLower(coinType)]
}
