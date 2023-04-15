package rules

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type YandexContainerRepositoryIamBindingInvalidRoleRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

func NewYandexContainerRepositoryIamBindingInvalidRoleRule() *YandexContainerRepositoryIamBindingInvalidRoleRule {
	return &YandexContainerRepositoryIamBindingInvalidRoleRule{
		resourceType:  "yandex_container_repository_iam_binding",
		attributeName: "role",
	}
}

func (r *YandexContainerRepositoryIamBindingInvalidRoleRule) Name() string {
	return "yandex_container_repository_iam_binding_invalid_role"
}

func (r *YandexContainerRepositoryIamBindingInvalidRoleRule) Enabled() bool {
	return true
}

func (r *YandexContainerRepositoryIamBindingInvalidRoleRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *YandexContainerRepositoryIamBindingInvalidRoleRule) Link() string {
	return ""
}

func (r *YandexContainerRepositoryIamBindingInvalidRoleRule) Check(runner tflint.Runner) error {
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
		var role string
		err := runner.EvaluateExpr(attribute.Expr, &role, nil)

		err = runner.EnsureNoError(err, func() error {
			if !ValidContainerRegistryServiceRoles[role] {
				runner.EmitIssue(r, fmt.Sprintf("\"%s\" is incorrect value for 'role'", role), attribute.Expr.Range())
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}
