package urldb

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

type Urldb map[string]string

func (db Urldb) Create(url string) (string, error) {
	id := GetMD5Hash(url)[:8]
	db[id] = url
	return id, nil
}

func (db Urldb) Get(id string) (string, error) {
	url, found := db[id]
	if found {
		return url, nil
	}
	return "", fmt.Errorf("ID Not found")
}
