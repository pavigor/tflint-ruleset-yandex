package rules

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type YandexComputeImageInvalidOsTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

func NewYandexComputeImageInvalidOsTypeRule() *YandexComputeImageInvalidOsTypeRule {
	return &YandexComputeImageInvalidOsTypeRule{
		resourceType:  "yandex_compute_image",
		attributeName: "os_type",
	}
}

func (r *YandexComputeImageInvalidOsTypeRule) Name() string {
	return "yandex_compute_image_invalid_os_type"
}

func (r *YandexComputeImageInvalidOsTypeRule) Enabled() bool {
	return true
}

func (r *YandexComputeImageInvalidOsTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *YandexComputeImageInvalidOsTypeRule) Link() string {
	return ""
}

func (r *YandexComputeImageInvalidOsTypeRule) Check(runner tflint.Runner) error {
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
		var osType string
		err := runner.EvaluateExpr(attribute.Expr, &osType, nil)

		err = runner.EnsureNoError(err, func() error {
			if !ValidImageOsTypes[osType] {
				runner.EmitIssue(r, fmt.Sprintf("Invalid OS type %s\n", osType), attribute.Expr.Range())
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}
