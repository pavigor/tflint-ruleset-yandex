package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

var manualRules = []tflint.Rule{
	NewYandexAlbLoadBalancerInvalidZoneIdRule(),
	NewYandexCmCertificateInvalidChallengeTypeRule(),
	NewYandexComputeDiskInvalidTypeRule(),
	NewYandexComputeDiskInvalidZoneRule(),
	NewYandexComputeDiskPlacementGroupInvalidZoneRule(),
	NewYandexComputeFilesystemInvalidDiskTypeRule(),
	NewYandexComputeFilesystemInvalidZoneRule(),
	NewYandexComputeGpuClusterInvalidInterconnectTypeRule(),
	NewYandexComputeGpuClusterInvalidZoneRule(),
	NewYandexComputeImageInvalidOsTypeRule(),
	NewYandexComputeInstanceInvalidDiskModeRule(),
	NewYandexComputeInstanceInvalidFilesystemModeRule(),
	NewYandexComputeInstanceInvalidNetworkAccelerationTypeRule(),
	NewYandexComputeInstanceInvalidPlatformIdRule(),
	NewYandexComputeInstanceInvalidPrivateIpAddressRule(),
	NewYandexComputeInstanceInvalidSecondaryDiskModeRule(),
	NewYandexComputeInstanceInvalidZoneRule(),
	NewYandexContainerRegistryIamBindingInvalidRoleRule(),
	NewYandexContainerRepositoryIamBindingInvalidRoleRule(),
	NewYandexVpcSubnetInvalidZoneRule(),
}

var Rules []tflint.Rule

func init() {
	Rules = append(Rules, manualRules...)
}
