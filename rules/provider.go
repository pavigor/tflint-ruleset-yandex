package rules

import "github.com/terraform-linters/tflint-plugin-sdk/tflint"

var manualRules = []tflint.Rule{
	NewYandexAlbLoadBalancerInvalidZoneIdRule(),
	NewComputeInstanceInvalidPlatformIdRule(),
	NewYandexVpcSubnetInvalidZoneRule(),
}

var Rules []tflint.Rule

func init() {
	Rules = append(Rules, manualRules...)
}