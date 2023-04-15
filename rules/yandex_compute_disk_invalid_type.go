package rules

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type YandexComputeDiskInvalidTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

func NewYandexComputeDiskInvalidTypeRule() *YandexComputeDiskInvalidTypeRule {
	return &YandexComputeDiskInvalidTypeRule{
		resourceType:  "yandex_compute_disk",
		attributeName: "type",
	}
}

func (r *YandexComputeDiskInvalidTypeRule) Name() string {
	return "yandex_compute_disk_invalid_type"
}

func (r *YandexComputeDiskInvalidTypeRule) Enabled() bool {
	return true
}

func (r *YandexComputeDiskInvalidTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *YandexComputeDiskInvalidTypeRule) Link() string {
	return ""
}

func (r *YandexComputeDiskInvalidTypeRule) Check(runner tflint.Runner) error {
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
				runner.EmitIssue(r, fmt.Sprintf("Invalid disk type %s\n", diskType), attribute.Expr.Range())
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}
