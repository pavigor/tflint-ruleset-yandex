package rules

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type YandexVpcSubnetInvalidZoneRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

func NewYandexVpcSubnetInvalidZoneRule() *YandexVpcSubnetInvalidZoneRule {
	return &YandexVpcSubnetInvalidZoneRule{
		resourceType:  "yandex_vpc_subnet",
		attributeName: "zone",
	}
}

func (r *YandexVpcSubnetInvalidZoneRule) Name() string {
	return "YandexVpcSubnetInvalidZoneRule"
}

func (r *YandexVpcSubnetInvalidZoneRule) Enabled() bool {
	return true
}

func (r *YandexVpcSubnetInvalidZoneRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *YandexVpcSubnetInvalidZoneRule) Link() string {
	return ""
}

func (r *YandexVpcSubnetInvalidZoneRule) Check(runner tflint.Runner) error {
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
		var value string
		err := runner.EvaluateExpr(attribute.Expr, &value, nil)

		err = runner.EnsureNoError(err, func() error {
			if !validAvailabilityZones[value] {
				runner.EmitIssue(r, fmt.Sprintf("Invalid zone %s", value), attribute.Expr.Range())
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}
