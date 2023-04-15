package rules

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type YandexComputeGpuClusterInvalidInterconnectTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

func NewYandexComputeGpuClusterInvalidInterconnectTypeRule() *YandexComputeGpuClusterInvalidInterconnectTypeRule {
	return &YandexComputeGpuClusterInvalidInterconnectTypeRule{
		resourceType:  "yandex_compute_gpu_cluster",
		attributeName: "interconnect_type",
	}
}

func (r *YandexComputeGpuClusterInvalidInterconnectTypeRule) Name() string {
	return "yandex_compute_gpu_cluster_invalid_interconnect_type"
}

func (r *YandexComputeGpuClusterInvalidInterconnectTypeRule) Enabled() bool {
	return true
}

func (r *YandexComputeGpuClusterInvalidInterconnectTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *YandexComputeGpuClusterInvalidInterconnectTypeRule) Link() string {
	return ""
}

func (r *YandexComputeGpuClusterInvalidInterconnectTypeRule) Check(runner tflint.Runner) error {
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
		var interconnectType string
		err := runner.EvaluateExpr(attribute.Expr, &interconnectType, nil)

		err = runner.EnsureNoError(err, func() error {
			if !ValidGpuInterconnectType[interconnectType] {
				runner.EmitIssue(r, fmt.Sprintf("Incorrect GPU interconnect tyep %s\n", interconnectType), attribute.Expr.Range())
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}
