package googleauthenticator

import (
	"crypto/rand"
	"encoding/base32"
	"regexp"
	"strings"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

type Authenticator struct {
	key *otp.Key
}

func NewAuthenticator(issuer string, accountName string, formattedKey string) *Authenticator {
	rx := regexp.MustCompile(`\W+`)
	secret := []byte(rx.ReplaceAllString(formattedKey, ""))
	ret, _ := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: accountName,
		SecretSize:  uint(len(secret)),
		Secret:      secret,
	})
	return &Authenticator{key: ret}
}

func (a *Authenticator) VerifyToken(passcode string) bool {
	rv, _ := totp.ValidateCustom(
		passcode,
		a.key.Secret(),
		time.Now().UTC(),
		totp.ValidateOpts{
			Period:    30,
			Skew:      1,
			Digits:    otp.DigitsSix,
			Algorithm: otp.AlgorithmSHA1,
		},
	)
	return rv
}

func (a *Authenticator) GenerateToken() (passcode string) {
	passcode, _ = totp.GenerateCode(a.key.Secret(), time.Now().UTC())
	return
}

func (a *Authenticator) GenerateTotpUri() string {
	return a.key.URL()
}

func GenerateKey() (formattedKey string) {
	formattedKey = encodeGoogleAuthKey(generateOtpKey())
	return
}

// Generate a key
func generateOtpKey() []byte {
	// 20 cryptographically random binary bytes (160-bit key)
	key := make([]byte, 20)
	_, _ = rand.Read(key)
	return key
}

// Text-encode the key as base32 (in the style of Google Authenticator - same as Facebook, Microsoft, etc)
func encodeGoogleAuthKey(bin []byte) string {
	// 32 ascii characters without trailing '='s
	rx := regexp.MustCompile(`=`)
	base32 := rx.ReplaceAllString(base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(bin), "")
	base32 = strings.ToLower(base32)

	// lowercase with a space every 4 characters
	rx = regexp.MustCompile(`(\w{4})`)
	key := strings.TrimSpace(rx.ReplaceAllString(base32, "$1 "))

	return key
}

// Binary-decode the key from base32 (Google Authenticator, FB, M$, etc)
// func decodeGoogleAuthKey(key string) []byte {
// 	// decode base32 google auth key to binary
// 	rx := regexp.MustCompile(`\W+`)
// 	unformatted := strings.ToUpper(rx.ReplaceAllString(key, ""))
// 	bin, _ := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(unformatted)
// 	return bin
// }

// func GenerateTotpUri(secret, accountName, issuer string) string {
// 	// Full OTPAUTH URI spec as explained at
// 	// https://github.com/google/google-authenticator/wiki/Key-Uri-Format
// 	u := url.URL{}
// 	v := url.Values{}
// 	u.Scheme = "otpauth"
// 	u.Host = "totp"
// 	u.Path = fmt.Sprintf("%s:%s", issuer, accountName)
// 	v.Add("secret", secret)
// 	v.Add("issuer", issuer)
// 	v.Add("algorithm", "SHA1")
// 	v.Add("digits", strconv.Itoa(6))
// 	v.Add("period", strconv.Itoa(30))
// 	u.RawQuery = v.Encode()
// 	return u.String()
// }
