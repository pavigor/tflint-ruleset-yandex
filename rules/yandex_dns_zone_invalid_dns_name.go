package rules

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	s "strings"
)

type YandexDnsZoneInvalidDnsNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

func NewYandexDnsZoneInvalidDnsNameRule() *YandexDnsZoneInvalidDnsNameRule {
	return &YandexDnsZoneInvalidDnsNameRule{
		resourceType:  "yandex_dns_zone",
		attributeName: "name",
	}
}

func (r *YandexDnsZoneInvalidDnsNameRule) Name() string {
	return "yandex_dns_zone_invalid_dns_name"
}

func (r *YandexDnsZoneInvalidDnsNameRule) Enabled() bool {
	return true
}

func (r *YandexDnsZoneInvalidDnsNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *YandexDnsZoneInvalidDnsNameRule) Link() string {
	return ""
}

func (r *YandexDnsZoneInvalidDnsNameRule) Check(runner tflint.Runner) error {
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
		var dnsName string
		err := runner.EvaluateExpr(attribute.Expr, &dnsName, nil)

		err = runner.EnsureNoError(err, func() error {
			if !s.HasSuffix(dnsName, ".") {
				runner.EmitIssue(r, fmt.Sprintf("\"%s\" is incorrect value for 'name'. Must ends with dot", dnsName), attribute.Expr.Range())
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}
