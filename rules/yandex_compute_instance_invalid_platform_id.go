package rules

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type YandexComputeInstancePlatformIdRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

func NewYandexComputeInstanceInvalidPlatformIdRule() *YandexComputeInstancePlatformIdRule {
	return &YandexComputeInstancePlatformIdRule{
		resourceType:  "yandex_compute_instance",
		attributeName: "platform_id",
	}
}

func (r *YandexComputeInstancePlatformIdRule) Name() string {
	return "yandex_compute_instance_invalid_platform_id"
}

func (r *YandexComputeInstancePlatformIdRule) Enabled() bool {
	return true
}

func (r *YandexComputeInstancePlatformIdRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *YandexComputeInstancePlatformIdRule) Link() string {
	return ""
}

func (r *YandexComputeInstancePlatformIdRule) Check(runner tflint.Runner) error {
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
		var platformId string
		err := runner.EvaluateExpr(attribute.Expr, &platformId, nil)

		err = runner.EnsureNoError(err, func() error {
			if !ValidComputePlatforms[platformId] {
				runner.EmitIssue(r, fmt.Sprintf("\"%s\" is invalid platform id", platformId), attribute.Expr.Range())
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}
