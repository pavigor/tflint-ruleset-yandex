package rules

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"net"
)

type YandexComputeInstanceInvalidPrivateIpAddressRule struct {
	tflint.DefaultRule

	resourceType  string
	blockName     string
	attributeName string
}

func NewYandexComputeInstanceInvalidPrivateIpAddressRule() *YandexComputeInstanceInvalidPrivateIpAddressRule {
	return &YandexComputeInstanceInvalidPrivateIpAddressRule{
		resourceType:  "yandex_compute_instance",
		blockName:     "network_interface",
		attributeName: "ip_address",
	}
}

func (r *YandexComputeInstanceInvalidPrivateIpAddressRule) Name() string {
	return "yandex_compute_instance_invalid_private_ip_address"
}

func (r *YandexComputeInstanceInvalidPrivateIpAddressRule) Enabled() bool {
	return true
}

func (r *YandexComputeInstanceInvalidPrivateIpAddressRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *YandexComputeInstanceInvalidPrivateIpAddressRule) Link() string {
	return ""
}

func (r *YandexComputeInstanceInvalidPrivateIpAddressRule) Check(runner tflint.Runner) error {
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
			var ipv4Address string
			err := runner.EvaluateExpr(attribute.Expr, &ipv4Address, nil)

			err = runner.EnsureNoError(err, func() error {
				if !net.IP.IsPrivate(net.ParseIP(ipv4Address)) {
					runner.EmitIssue(r, fmt.Sprintf("\"%s\" is not a private IP address", ipv4Address), attribute.Expr.Range())
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
