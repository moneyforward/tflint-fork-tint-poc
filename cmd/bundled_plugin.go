package cmd

import (
	"fmt"
	"github.com/moneyforward/tflint-fork-tint-poc/rules"
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-terraform/project"
	"github.com/terraform-linters/tflint-ruleset-terraform/terraform"
)

func (cli *CLI) actAsBundledPlugin() int {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &terraform.RuleSet{
			BuiltinRuleSet: tflint.BuiltinRuleSet{
				Name:    "terraform",
				Version: fmt.Sprintf("%s-bundled", project.Version),
			},
			PresetRules: map[string][]tflint.Rule{
				"all":         rules.RuleDefinitionToApply,
				"recommended": rules.RuleDefinitionToApply, // Original TFLint logic use this value
			},
		},
	})
	return ExitCodeOK
}
