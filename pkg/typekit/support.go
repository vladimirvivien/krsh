package typekit

import (
	"fmt"
	"regexp"

	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
)

var nameSanitizeRegex = regexp.MustCompile(`[^a-zA-Z0-9]`)

func SanitizeIdentifier(str string) string {
	return nameSanitizeRegex.ReplaceAllString(str, "_")
}

func BuiltinError(funcName string, err error) (starlark.Value, error) {
	return starlark.None, fmt.Errorf("%s: failed: %s", funcName, err)
}

func AsStarlarkStruct(value interface{}) (*starlarkstruct.Struct, error) {
	starStruct := new(starlarkstruct.Struct)
	if err := Go(value).Starlark(starStruct); err != nil {
		return nil, fmt.Errorf("starlark type conversion failed: %w", err)
	}
	return starStruct, nil
}

func BuiltinResult(funcName string, result interface{}) (starlark.Value, error) {
	starResult, err := AsStarlarkStruct(result)
	if err != nil {
		return BuiltinError(funcName, fmt.Errorf("conversion error: %v", err))
	}
	return starResult, nil
}
