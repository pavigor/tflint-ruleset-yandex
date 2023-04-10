package compute

import "github.com/terraform-linters/tflint-plugin-sdk/tflint"

type ComputeInstanceZoneRule struct {
	tflint.DefaultRule

	resourceType string
	attributeName string
}

func NewComputeInstanceInvalidZoneRule() *ComputeInstanceZoneRule {
	return &ComputeInstanceZoneRule{
		resourceType:  "yandex_compute_instance",
		attributeName: "zone",
	}
}


