package rules

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type YandexAlbLoadBalancerInvalidZoneIdRule struct {
	tflint.DefaultRule

	resourceType  string
	subAttr		  string
	attributeName string
}

func NewYandexAlbLoadBalancerInvalidZoneIdRule() *YandexAlbLoadBalancerInvalidZoneIdRule {
	return &YandexAlbLoadBalancerInvalidZoneIdRule{
		resourceType:  "yandex_alb_load_balancer",
		subAttr: "allocation_policy",
		attributeName: "zone_id",
	}
}

func (r *YandexAlbLoadBalancerInvalidZoneIdRule) Name() string {
	return "YandexAlbLoadBalancerInvalidZoneIdRule"
}

func (r *YandexAlbLoadBalancerInvalidZoneIdRule) Enabled() bool {
	return true
}

func (r *YandexAlbLoadBalancerInvalidZoneIdRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *YandexAlbLoadBalancerInvalidZoneIdRule) Link() string {
	return ""
}

func (r *YandexAlbLoadBalancerInvalidZoneIdRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type: "allocation_policy",
				Body: &hclext.BodySchema{
					Blocks: []hclext.BlockSchema{
						{
							Type: "location",
							Body: &hclext.BodySchema{
								Attributes: []hclext.AttributeSchema{{Name: "zone_id"}},
							},
						},
					},
				},
			},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		for _, ru := range resource.Body.Blocks {
			for _, rule := range ru.Body.Blocks {
				attribute, exists := rule.Body.Attributes[r.attributeName]
				if !exists {
					continue
				}

				var zone_id string
				err := runner.EvaluateExpr(attribute.Expr, &zone_id, nil)

				err = runner.EnsureNoError(err, func() error {
					if !validAvailabilityZones[zone_id] {
						runner.EmitIssue(r, fmt.Sprintf("Invalid zone %s", zone_id), attribute.Expr.Range())
					}
					return nil
				})
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
