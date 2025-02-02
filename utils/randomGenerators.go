package utils

import (
	"strings"
	"fmt"
	"github.com/google/uuid"
)

func GeneratorUUID(strLen int) string {
	fmt.Println("apple")
	return strings.Replace(uuid.NewString(), "-", "", -1)[:strLen]
}
