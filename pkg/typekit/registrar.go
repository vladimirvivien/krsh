package typekit

import (
	"sync"

	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
)

var (
	mutex    sync.Mutex
	registry map[string]starlark.StringDict
)

func init() {
	registry = map[string]starlark.StringDict{
		"os": make(starlark.StringDict),
	}
}

// Register registers a Starlark built-in value in a package-wide registry
func Register(namespace, name string, builtin starlark.Value) {
	mutex.Lock()
	defer mutex.Unlock()
	registry[namespace][name] = builtin
}

// Registry returns the starlark registry
func Registry() starlark.StringDict {
	result := make(starlark.StringDict)
	for namespace, value := range registry {
		result[namespace] = starlarkstruct.FromStringDict(starlark.String(namespace), value)
	}
	return result
}
