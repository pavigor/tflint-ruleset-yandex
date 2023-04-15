package rules

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	s "strings"
)

type YandexDnsRecordsetInvalidDnsNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

func NewYandexDnsRecordsetInvalidDnsNameRule() *YandexDnsRecordsetInvalidDnsNameRule {
	return &YandexDnsRecordsetInvalidDnsNameRule{
		resourceType:  "yandex_dns_recordset",
		attributeName: "name",
	}
}

func (r *YandexDnsRecordsetInvalidDnsNameRule) Name() string {
	return "yandex_dns_recordset_invalid_dns_name"
}

func (r *YandexDnsRecordsetInvalidDnsNameRule) Enabled() bool {
	return true
}

func (r *YandexDnsRecordsetInvalidDnsNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *YandexDnsRecordsetInvalidDnsNameRule) Link() string {
	return ""
}

func (r *YandexDnsRecordsetInvalidDnsNameRule) Check(runner tflint.Runner) error {
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
