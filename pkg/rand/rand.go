package rand

import (
	cryptoRand "crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/gofrs/uuid"
	"math/rand"
	"time"
)

const alphaCharset = "abcdefghijklmnopqrstuvwyxz"
const emailCharset = "abcdefghijklmnopqrstuvwxyz0123456789"
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const base64Charset = "abcdefghijklmnopqrstuvwxyz0123456789"

const unambiguousCharset = "abcdefghjkmnprstvwxy2346789"

var seededRand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func ApiKey() string {
	//4-4-4-4
	newUuid := fmt.Sprintf(
		"%v-%v-%v-%v",
		StringWithCharset(4, unambiguousCharset),
		StringWithCharset(4, unambiguousCharset),
		StringWithCharset(4, unambiguousCharset),
		StringWithCharset(4, unambiguousCharset),
	)
	return newUuid
}

func Uuid() string {
	//8-4-4-4-12
	newUuid := fmt.Sprintf(
		"%v-%v-%v-%v-%v",
		StringWithCharset(8, emailCharset),
		StringWithCharset(4, emailCharset),
		StringWithCharset(4, emailCharset),
		StringWithCharset(4, emailCharset),
		StringWithCharset(12, emailCharset),
	)
	return newUuid
}

func Email() string {
	name := StringWithCharset(10, emailCharset)
	return fmt.Sprintf("%v@purplevisits.com", name)
}

func StringFixed(length int) string {
	return StringWithCharset(length, charset)
}

func StringBase64Fixed(length int) string {
	return StringWithCharset(length, base64Charset)
}

func String() string {
	return StringWithCharset(10, charset)
}

func StringAlpha() string {
	return StringWithCharset(10, alphaCharset)
}

// ClientSecret generates a random client secret. In base64 encoding it can only be 72 bytes as that's the maximum allowed in bcrypt
func ClientSecret() (string, error) {
	// 54 bytes is a 432 bit key. In url encoding this produces a 72 character string, which is the maximum allowed in bcrypt
	secret := make([]byte, 54)
	_, err := cryptoRand.Read(secret)
	if err != nil {
		return "", err
	}
	// Encode the secret in URL-safe base64
	encodedSecret := base64.RawURLEncoding.EncodeToString(secret)
	return encodedSecret, nil
}

func UidV4() string {
	newUuid, _ := uuid.NewV4()
	return newUuid.String()
}
