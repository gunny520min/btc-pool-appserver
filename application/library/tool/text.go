package tool

import "github.com/bitly/go-simplejson"

var sensitiveWords = []string{
	"password",
	"old_password",
	"new_password",
	"passwd",
	"pass",
	"google_auth_secret",
	"google_auth_code",
	"security_code",
}

func FilterSecret(text []byte) []byte {
	sj, _ := simplejson.NewJson(text)
	for _, s := range sensitiveWords {
		if pwd, _ := sj.Get(s).String(); pwd != "" {
			sj.Set(s, "******")
		}
	}
	res, _ := sj.MarshalJSON()

	return res
}
