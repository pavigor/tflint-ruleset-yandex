package rules

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type YandexComputeInstanceInvalidFilesystemModeRule struct {
	tflint.DefaultRule

	resourceType  string
	blockName     string
	attributeName string
}

func NewYandexComputeInstanceInvalidFilesystemModeRule() *YandexComputeInstanceInvalidFilesystemModeRule {
	return &YandexComputeInstanceInvalidFilesystemModeRule{
		resourceType:  "yandex_compute_instance",
		blockName:     "filesystem",
		attributeName: "mode",
	}
}

func (r *YandexComputeInstanceInvalidFilesystemModeRule) Name() string {
	return "yandex_compute_instance_invalid_filesystem_mode"
}

func (r *YandexComputeInstanceInvalidFilesystemModeRule) Enabled() bool {
	return true
}

func (r *YandexComputeInstanceInvalidFilesystemModeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *YandexComputeInstanceInvalidFilesystemModeRule) Link() string {
	return ""
}

func (r *YandexComputeInstanceInvalidFilesystemModeRule) Check(runner tflint.Runner) error {
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
					runner.EmitIssue(r, fmt.Sprintf("\"%s\" is incorrect value for 'filesystem.mode'", diskMode), attribute.Expr.Range())
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
