package validator

import (
	"github.com/go-playground/validator/v10"
	"gin-web-init/utils"
	"regexp"
	"strings"
)

const DockerImageURNRegex = `^[a-zA-Z0-9.-]+(:\d+)?/[a-zA-Z0-9._-]+/[a-zA-Z0-9._-]+(:[a-zA-Z0-9._-]+)?$`
const DomainNameRegex = `^([a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`
const PathRegex = `^[-a-zA-Z0-9._/]*$`

var _ = RegisterValidator

var validatorFuncRegistry = map[string]validator.Func{
	"image":     ImageRepositoryValidatorFunc,
	"domain":    DomainNameValidatorFunc,
	"version":   VersionValidatorFunc,
	"cpu":       CpuResourceValidatorFunc,
	"mem":       MemoryResourceValidatorFunc,
	"ord":       OrderValidatorFunc,
	"path":      PathValidatorFunc,
	"validchar": NameValidatorFunc,
}

func RegistryList() map[string]validator.Func {
	return validatorFuncRegistry
}

// RegisterValidator register validator.Func to gin.
func RegisterValidator(name string, fn validator.Func) {
	if _, ok := validatorFuncRegistry[name]; ok {
		return
	}

	validatorFuncRegistry[name] = fn
}

func NameValidatorFunc(f1 validator.FieldLevel) bool {
	var name = f1.Field().String()
	for _, ch := range name {
		if !(ch >= '0' && ch <= '9' || ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '-' || ch == '_') {
			return false
		}
	}
	return true
}

func PathValidatorFunc(f1 validator.FieldLevel) bool {
	var re = regexp.MustCompile(PathRegex)
	return re.MatchString(f1.Field().String())
}

func OrderValidatorFunc(f1 validator.FieldLevel) bool {
	ord := map[string]struct{}{
		"asc":  {},
		"desc": {},
		"ASC":  {},
		"DESC": {},
	}
	if f1.Field().String() == "" {
		return true
	}
	if _, ok := ord[f1.Field().String()]; ok {
		return true
	}
	return false
}

func ImageRepositoryValidatorFunc(f1 validator.FieldLevel) bool {
	return imageRepositoryValidator(f1.Field().String())
}

func DomainNameValidatorFunc(f1 validator.FieldLevel) bool {
	return domainNameValidator(f1.Field().String())
}

func domainNameValidator(domain string) bool {
	var re = regexp.MustCompile(DomainNameRegex)
	return re.MatchString(domain) || domain == ""
}

func imageRepositoryValidator(urn string) bool {
	var re = regexp.MustCompile(DockerImageURNRegex)
	return re.MatchString(urn)
}

func VersionValidatorFunc(f1 validator.FieldLevel) bool {
	return versionValidator(f1.Field().String()) != nil
}

func CpuResourceValidatorFunc(f1 validator.FieldLevel) bool {
	resource := f1.Field().String()
	var token, n = resource, len(resource)
	if strings.HasSuffix(resource, "m") {
		token = resource[0 : n-1]
	}
	for _, ch := range token {
		if ch > '9' || ch < '0' {
			return false
		}
	}
	return true
}

func MemoryResourceValidatorFunc(f1 validator.FieldLevel) bool {
	var resource = []byte(f1.Field().String())
	n := len(resource)
	var code, gauge = make([]byte, 0), ""
	if resource[n-1] == 'i' {
		code = resource[0 : n-2]
		gauge = string(resource[n-2:])
	} else {
		code = resource[0 : n-1]
		gauge = string(resource[n-1:])
	}
	for _, ch := range code {
		if ch < '0' || ch > '9' {
			return false
		}
	}
	return utils.LegalMemoryGauge(gauge)
}

func versionValidator(version string) []int {
	var dot, major, minor, patch int
	for _, ch := range version {
		if byte(ch) >= '0' && byte(ch) <= '9' {
			n := int(byte(ch) - '0')
			switch dot {
			case 0:
				major = (major * 10) + n
			case 1:
				minor = (minor * 10) + n
			case 2:
				patch = (patch * 10) + n
			default:
				return nil
			}
		} else if byte(ch) == '.' {
			dot++
		} else {
			return nil
		}
	}
	return []int{major, minor, patch}
}
