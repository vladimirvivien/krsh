// Package container provides constructor types for *coreV1.Container
package container

import (
	coreV1 "k8s.io/api/core/v1"
)

type Constructor struct {
	container coreV1.Container
}

// Name initializer function for *Constructor
func Name(name string) Constructor {
	return Constructor{container: coreV1.Container{Name: name}}
}

// Image setter for container image name
func (c Constructor) Image(image string) Constructor {
	c.container.Image = image
	return c
}

// Args setter method for container args
func (c Constructor) Args(args ...string) Constructor {
	c.container.Args = args
	return c
}

// Commands setter method for container commands
func (c Constructor) Commands(cmds ...string) Constructor {
	c.container.Command = cmds
	return c
}

// Build finalizer method that returns the built *coreV1.Container
func (c Constructor) Build() coreV1.Container {
	return c.container
}
