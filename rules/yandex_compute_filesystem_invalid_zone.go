package rules

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type YandexComputeFilesystemInvalidZoneRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

func NewYandexComputeFilesystemInvalidZoneRule() *YandexComputeFilesystemInvalidZoneRule {
	return &YandexComputeFilesystemInvalidZoneRule{
		resourceType:  "yandex_compute_filesystem",
		attributeName: "zone",
	}
}

func (r *YandexComputeFilesystemInvalidZoneRule) Name() string {
	return "yandex_compute_filesystem_invalid_zone"
}

func (r *YandexComputeFilesystemInvalidZoneRule) Enabled() bool {
	return true
}

func (r *YandexComputeFilesystemInvalidZoneRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *YandexComputeFilesystemInvalidZoneRule) Link() string {
	return ""
}

func (r *YandexComputeFilesystemInvalidZoneRule) Check(runner tflint.Runner) error {
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
		var zone_id string
		err := runner.EvaluateExpr(attribute.Expr, &zone_id, nil)

		err = runner.EnsureNoError(err, func() error {
			if !ValidAvailabilityZones[zone_id] {
				runner.EmitIssue(r, fmt.Sprintf("Invalid zone %s", zone_id), attribute.Expr.Range())
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}
