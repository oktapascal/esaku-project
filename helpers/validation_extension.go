package helpers

import (
	"errors"
	"strings"
)

var magicTable = map[string]string{
	"\xff\xd8\xff":      "image/jpeg",
	"\x89PNG\r\n\x1a\n": "image/png",
}

func CheckOnlyImage(file []byte) (string, error) {
	extString := []byte(file)

	for magic, mime := range magicTable {
		if strings.HasPrefix(string(extString), magic) {
			return mime, nil
		}
	}

	return "", errors.New("image type is invalid")
}
