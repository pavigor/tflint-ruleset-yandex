package rules

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type YandexComputeInstanceInvalidDiskModeRule struct {
	tflint.DefaultRule

	resourceType  string
	blockName     string
	attributeName string
}

func NewYandexComputeInstanceInvalidDiskModeRule() *YandexComputeInstanceInvalidDiskModeRule {
	return &YandexComputeInstanceInvalidDiskModeRule{
		resourceType:  "yandex_compute_instance",
		blockName:     "boot_disk",
		attributeName: "mode",
	}
}

func (r *YandexComputeInstanceInvalidDiskModeRule) Name() string {
	return "yandex_compute_instance_invalid_disk_mode"
}

func (r *YandexComputeInstanceInvalidDiskModeRule) Enabled() bool {
	return true
}

func (r *YandexComputeInstanceInvalidDiskModeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *YandexComputeInstanceInvalidDiskModeRule) Link() string {
	return ""
}

func (r *YandexComputeInstanceInvalidDiskModeRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type: r.blockName,
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: r.attributeName},
					},
				},
			},
		},
	}, nil)

	if err != nil {
		return err
	}
	for _, resource := range resources.Blocks {
		for _, rule := range resource.Body.Blocks {
			attribute, exists := rule.Body.Attributes[r.attributeName]
			if !exists {
				continue
			}
			var diskMode string
			err := runner.EvaluateExpr(attribute.Expr, &diskMode, nil)

			err = runner.EnsureNoError(err, func() error {
				if !ValidDiskAccessMode[diskMode] {
					runner.EmitIssue(r, fmt.Sprintf("\"%s\" is incorrect value for boot_data.mode", diskMode), attribute.Expr.Range())
				}
				return nil
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}
