package helpers

import "strings"

var magicTable = map[string]string{
	"\xff\xd8\xff":      "image/jpeg",
	"\x89PNG\r\n\x1a\n": "image/png",
}

func CheckOnlyImage(file []byte) string {
	extString := []byte(file)

	for magic, mime := range magicTable {
		if strings.HasPrefix(string(extString), magic) {
			return mime
		}
	}

	return ""
}
