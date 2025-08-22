package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

var RuleDefinitionToApply = []tflint.Rule{
	NewAwsRdbInstanceMustBeRedundant(),
}
