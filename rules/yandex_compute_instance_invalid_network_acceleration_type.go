package rules

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type YandexComputeInstanceInvalidNetworkAccelerationTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

func NewYandexComputeInstanceInvalidNetworkAccelerationTypeRule() *YandexComputeInstanceInvalidNetworkAccelerationTypeRule {
	return &YandexComputeInstanceInvalidNetworkAccelerationTypeRule{
		resourceType:  "yandex_compute_instance",
		attributeName: "network_acceleration_type",
	}
}

func (r *YandexComputeInstanceInvalidNetworkAccelerationTypeRule) Name() string {
	return "yandex_compute_instance_invalid_network_acceleration_type"
}

func (r *YandexComputeInstanceInvalidNetworkAccelerationTypeRule) Enabled() bool {
	return true
}

func (r *YandexComputeInstanceInvalidNetworkAccelerationTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *YandexComputeInstanceInvalidNetworkAccelerationTypeRule) Link() string {
	return ""
}

func (r *YandexComputeInstanceInvalidNetworkAccelerationTypeRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)

	if err != nil {
		return err
	}
	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}
		var networkAccelerationType string
		err := runner.EvaluateExpr(attribute.Expr, &networkAccelerationType, nil)

		err = runner.EnsureNoError(err, func() error {
			if !ValidNetworkAccelecrationTypes[networkAccelerationType] {
				runner.EmitIssue(r, fmt.Sprintf("\"%s\" is incorrect value for network_acceleration_type", networkAccelerationType), attribute.Expr.Range())
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}
