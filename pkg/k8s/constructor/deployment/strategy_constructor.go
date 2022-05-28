package deployment

import (
	appsV1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

var (
	//StrategyDefault is set to value "Recreate"
	StrategyDefault  = StrategyConstructor{strat: appsV1.DeploymentStrategy{Type: appsV1.RecreateDeploymentStrategyType}}
	StrategyRecreate = StrategyDefault
)

// StrategyConstructor is type to build values of type appsV1.DeploymentStrategy
type StrategyConstructor struct {
	strat appsV1.DeploymentStrategy
}

// StrategyRollingUpdate is an initializer function that creates a value of type appsV1.DeploymentStrategy
func StrategyRollingUpdate(maxUnavailable string, maxSurge string) StrategyConstructor {
	unavailParsed := intstr.FromString(maxUnavailable)
	surgeParsed := intstr.FromString(maxSurge)
	return StrategyConstructor{
		strat: appsV1.DeploymentStrategy{
			Type: appsV1.RollingUpdateDeploymentStrategyType,
			RollingUpdate: &appsV1.RollingUpdateDeployment{
				MaxUnavailable: &unavailParsed,
				MaxSurge:       &surgeParsed,
			},
		},
	}
}

// Build is a finalizer method that constructs the value of type appsV1.DeploymentStrategy
func (c StrategyConstructor) Build() appsV1.DeploymentStrategy {
	return c.strat
}
