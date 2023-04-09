package rules

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type YandexCmCertificateInvalidChallengeTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

func NewYandexCmCertificateInvalidChallengeTypeRule() *YandexCmCertificateInvalidChallengeTypeRule {
	return &YandexCmCertificateInvalidChallengeTypeRule{
		resourceType:  "yandex_cm_certificate",
		attributeName: "challenge_type",
	}
}

func (r *YandexCmCertificateInvalidChallengeTypeRule) Name() string {
	return "YandexCmCertificateInvalidChallengeTypeRule"
}

func (r *YandexCmCertificateInvalidChallengeTypeRule) Enabled() bool {
	return true
}

func (r *YandexCmCertificateInvalidChallengeTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *YandexCmCertificateInvalidChallengeTypeRule) Link() string {
	return ""
}

func (r *YandexCmCertificateInvalidChallengeTypeRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type: "managed",
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

			var challenge_type string
			err := runner.EvaluateExpr(attribute.Expr, &challenge_type, nil)

			err = runner.EnsureNoError(err, func() error {
				if !ValidChallengeType[challenge_type] {
					runner.EmitIssue(r, fmt.Sprintf("Invalid challenge type %s", challenge_type), attribute.Expr.Range())
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
