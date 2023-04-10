package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-template/rules/compute"
)

var manualRules = []tflint.Rule{
	NewYandexAlbLoadBalancerInvalidZoneIdRule(),
	compute.NewComputeInstanceInvalidPlatformIdRule(),
	NewYandexVpcSubnetInvalidZoneRule(),
	NewYandexCmCertificateInvalidChallengeTypeRule(),
	compute.NewYandexComputeDiskInvalidTypeRule(),
	compute.NewYandexComputeDiskInvalidZoneRule(),
}

var Rules []tflint.Rule

func init() {
	Rules = append(Rules, manualRules...)
}