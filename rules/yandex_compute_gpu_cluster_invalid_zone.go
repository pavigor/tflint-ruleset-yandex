package rules

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type YandexComputeGpuClusterInvalidZoneRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

func NewYandexComputeGpuClusterInvalidZoneRule() *YandexComputeGpuClusterInvalidZoneRule {
	return &YandexComputeGpuClusterInvalidZoneRule{
		resourceType:  "yandex_compute_gpu_cluster",
		attributeName: "zone",
	}
}

func (r *YandexComputeGpuClusterInvalidZoneRule) Name() string {
	return "yandex_compute_gpu_cluster_invalid_zone"
}

func (r *YandexComputeGpuClusterInvalidZoneRule) Enabled() bool {
	return true
}

func (r *YandexComputeGpuClusterInvalidZoneRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *YandexComputeGpuClusterInvalidZoneRule) Link() string {
	return ""
}

func (r *YandexComputeGpuClusterInvalidZoneRule) Check(runner tflint.Runner) error {
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
		var zoneId string
		err := runner.EvaluateExpr(attribute.Expr, &zoneId, nil)

		err = runner.EnsureNoError(err, func() error {
			if !ValidAvailabilityZones[zoneId] {
				runner.EmitIssue(r, fmt.Sprintf("Invalid zone %s", zoneId), attribute.Expr.Range())
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}
