package utils

import (
	"github.com/google/uuid"
	"strings"
)

func UUID() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")
}

// Generator uuid generator
var Generator = UUID
