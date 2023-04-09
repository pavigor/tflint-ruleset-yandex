package compute

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-template/rules"
)

type YandexComputeDiskInvalidZoneRule struct {
	tflint.DefaultRule

	resourceType string
	attributeName string
}

func NewYandexComputeDiskInvalidZoneRule() *YandexComputeDiskInvalidZoneRule {
	return &YandexComputeDiskInvalidZoneRule{
		resourceType: "yandex_compute_disk",
		attributeName: "type",
	}
}

func (r *YandexComputeDiskInvalidZoneRule) Name() string {
	return "compute_instance_invalid_platform_type"
}

func (r *YandexComputeDiskInvalidZoneRule) Enabled() bool {
	return true
}

func (r *YandexComputeDiskInvalidZoneRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *YandexComputeDiskInvalidZoneRule) Link() string {
	return ""
}

func (r *YandexComputeDiskInvalidZoneRule) Check(runner tflint.Runner) error {
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
			if !rules.ValidAvailabilityZones[zone] {
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
