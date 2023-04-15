package rules

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type YandexComputeFilesystemInvalidDiskTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

func NewYandexComputeFilesystemInvalidDiskTypeRule() *YandexComputeFilesystemInvalidDiskTypeRule {
	return &YandexComputeFilesystemInvalidDiskTypeRule{
		resourceType:  "yandex_compute_filesystem",
		attributeName: "type",
	}
}

func (r *YandexComputeFilesystemInvalidDiskTypeRule) Name() string {
	return "yandex_compute_filesystem_invalid_disk_type"
}

func (r *YandexComputeFilesystemInvalidDiskTypeRule) Enabled() bool {
	return true
}

func (r *YandexComputeFilesystemInvalidDiskTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *YandexComputeFilesystemInvalidDiskTypeRule) Link() string {
	return ""
}

func (r *YandexComputeFilesystemInvalidDiskTypeRule) Check(runner tflint.Runner) error {
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
		var diskType string
		err := runner.EvaluateExpr(attribute.Expr, &diskType, nil)

		err = runner.EnsureNoError(err, func() error {
			if !ValidDiskTypes[diskType] {
				runner.EmitIssue(r, fmt.Sprintf("Invalid disk type %s", diskType), attribute.Expr.Range())
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}
