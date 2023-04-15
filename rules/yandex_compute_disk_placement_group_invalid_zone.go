package rules

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type YandexComputeDiskPlacementGroupInvalidZoneRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

func NewYandexComputeDiskPlacementGroupInvalidZoneRule() *YandexComputeDiskPlacementGroupInvalidZoneRule {
	return &YandexComputeDiskPlacementGroupInvalidZoneRule{
		resourceType:  "yandex_compute_disk_placement_group",
		attributeName: "zone",
	}
}

func (r *YandexComputeDiskPlacementGroupInvalidZoneRule) Name() string {
	return "yandex_compute_disk_placement_group_invalid_zone"
}

func (r *YandexComputeDiskPlacementGroupInvalidZoneRule) Enabled() bool {
	return true
}

func (r *YandexComputeDiskPlacementGroupInvalidZoneRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *YandexComputeDiskPlacementGroupInvalidZoneRule) Link() string {
	return ""
}

func (r *YandexComputeDiskPlacementGroupInvalidZoneRule) Check(runner tflint.Runner) error {
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
