package uuidutils

import (
	"crypto/md5"

	uuid "github.com/pborman/uuid"
)

func DeriveFactsetUUID(factsetID string) string {
	h := md5.New()
	h.Write([]byte(factsetID))
	md5One := h.Sum(nil)
	return uuid.NewMD5(uuid.UUID{}, md5One).String()
}
