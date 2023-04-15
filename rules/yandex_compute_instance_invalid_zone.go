package rules

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type YandexComputeInstanceZoneRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

func NewYandexComputeInstanceInvalidZoneRule() *YandexComputeInstanceZoneRule {
	return &YandexComputeInstanceZoneRule{
		resourceType:  "yandex_compute_instance",
		attributeName: "zone",
	}
}

func (r *YandexComputeInstanceZoneRule) Name() string {
	return "yandex_compute_instance_invalid_zone"
}

func (r *YandexComputeInstanceZoneRule) Enabled() bool {
	return true
}

func (r *YandexComputeInstanceZoneRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *YandexComputeInstanceZoneRule) Link() string {
	return ""
}

func (r *YandexComputeInstanceZoneRule) Check(runner tflint.Runner) error {
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
		var zone string
		err := runner.EvaluateExpr(attribute.Expr, &zone, nil)

		err = runner.EnsureNoError(err, func() error {
			if !ValidAvailabilityZones[zone] {
				runner.EmitIssue(r, fmt.Sprintf("Invalid zone %s\n", zone), attribute.Expr.Range())
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}
