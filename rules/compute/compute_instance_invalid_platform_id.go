package compute

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-template/rules"
)

type ComputeInstancePlatformIdRule struct {
	tflint.DefaultRule

	resourceType string
	attributeName string
}

func NewComputeInstanceInvalidPlatformIdRule() *ComputeInstancePlatformIdRule {
	return &ComputeInstancePlatformIdRule{
		resourceType: "yandex_compute_instance",
		attributeName: "platform_id",
	}
}

func (r *ComputeInstancePlatformIdRule) Name() string {
	return "compute_instance_invalid_platform_type"
}

func (r *ComputeInstancePlatformIdRule) Enabled() bool {
	return true
}

func (r *ComputeInstancePlatformIdRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *ComputeInstancePlatformIdRule) Link() string {
	return ""
}

func (r *ComputeInstancePlatformIdRule) Check(runner tflint.Runner) error {
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
			if !rules.ValidComputePlatforms[platformId] {
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