package rules

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type YandexDnsRecordsetInvalidRecordTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

func NewYandexDnsRecordsetInvalidRecordTypeRule() *YandexDnsRecordsetInvalidRecordTypeRule {
	return &YandexDnsRecordsetInvalidRecordTypeRule{
		resourceType:  "yandex_dns_recordset",
		attributeName: "type",
	}
}

func (r *YandexDnsRecordsetInvalidRecordTypeRule) Name() string {
	return "yandex_dns_recordset_invalid_record_type"
}

func (r *YandexDnsRecordsetInvalidRecordTypeRule) Enabled() bool {
	return true
}

func (r *YandexDnsRecordsetInvalidRecordTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *YandexDnsRecordsetInvalidRecordTypeRule) Link() string {
	return ""
}

func (r *YandexDnsRecordsetInvalidRecordTypeRule) Check(runner tflint.Runner) error {
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
		var recordType string
		err := runner.EvaluateExpr(attribute.Expr, &recordType, nil)

		err = runner.EnsureNoError(err, func() error {
			if !ValidDnsRecordTypes[recordType] {
				runner.EmitIssue(r, fmt.Sprintf("\"%s\" is incorrect value for 'type'", recordType), attribute.Expr.Range())
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}
