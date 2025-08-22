package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type AwsRdbInstanceMustBeRedundant struct {
	tflint.DefaultRule
}

func NewAwsRdbInstanceMustBeRedundant() *AwsRdbInstanceMustBeRedundant {
	return &AwsRdbInstanceMustBeRedundant{}
}

// Name the name of rule
func (r *AwsRdbInstanceMustBeRedundant) Name() string {
	return "aws_rdb_instance_must_be_redundant"
}

// Enabled is enable automatically when the rule is loaded
func (r *AwsRdbInstanceMustBeRedundant) Enabled() bool {
	return true
}

// Severity the severity of error (ERROR, WARNING, NOTICE)
func (r *AwsRdbInstanceMustBeRedundant) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link (Optional) the link for reference
func (r *AwsRdbInstanceMustBeRedundant) Link() string {
	return "https://moneyforward.com"
}

// Check the logic of checking
func (r *AwsRdbInstanceMustBeRedundant) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(
		"aws_db_instance", // resource name to get
		&hclext.BodySchema{
			Attributes: []hclext.AttributeSchema{
				{Name: "multi_az"},
			},
		},
		nil,
	)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Invalid if no multi_az in resource
		attribute, exisits := resource.Body.Attributes["multi_az"]
		if !exisits {
			err := runner.EmitIssue(
				r,
				"multi_az attributes not found",
				resource.DefRange,
			)

			if err != nil {
				return err
			}
			continue
		}

		// Check is multi_az=true
		var isMultiAZTrue bool
		err := runner.EvaluateExpr(
			attribute.Expr,
			&isMultiAZTrue,
			nil,
		)
		if err != nil {
			return err
		}
		err = runner.EnsureNoError(err, func() error {
			if isMultiAZTrue {
				return nil
			}
			return runner.EmitIssue(
				r,
				"multi_az must be true. RDB instance must be redundant.",
				attribute.Expr.Range(),
			)
		})
		if err != nil {
			return err
		}
	}

	return nil
}
