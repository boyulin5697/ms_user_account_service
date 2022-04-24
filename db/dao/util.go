//Dao utils, Author @by. since 2022/4/23

package dao

import (
	"github.com/satori/go.uuid"
	"strings"
)

// GenerateUUID , randomly generate a UUID without "-"
func GenerateUUID() string {
	originUUID := uuid.NewV4()
	return strings.ReplaceAll(originUUID.String(), "-", "")
}
