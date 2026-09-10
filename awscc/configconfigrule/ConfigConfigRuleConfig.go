// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package configconfigrule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConfigConfigRuleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Provides the rule owner (```` for managed rules, ``CUSTOM_POLICY`` for Custom Policy rules, and ``CUSTOM_LAMBDA`` for Custom Lambda rules), the rule identifier, and the notifications that cause the function to evaluate your AWS resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/config_config_rule#source ConfigConfigRule#source}
	Source *ConfigConfigRuleSource `field:"required" json:"source" yaml:"source"`
	// Indicates whether an AWS resource or CC rule is compliant and provides the number of contributors that affect the compliance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/config_config_rule#compliance ConfigConfigRule#compliance}
	Compliance *ConfigConfigRuleCompliance `field:"optional" json:"compliance" yaml:"compliance"`
	// A name for the CC rule.
	//
	// If you don't specify a name, CFN generates a unique physical ID and uses that ID for the rule name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/config_config_rule#config_rule_name ConfigConfigRule#config_rule_name}
	ConfigRuleName *string `field:"optional" json:"configRuleName" yaml:"configRuleName"`
	// The description that you provide for the CC rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/config_config_rule#description ConfigConfigRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The modes the CC rule can be evaluated in.
	//
	// The valid values are distinct objects. By default, the value is Detective evaluation mode only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/config_config_rule#evaluation_modes ConfigConfigRule#evaluation_modes}
	EvaluationModes interface{} `field:"optional" json:"evaluationModes" yaml:"evaluationModes"`
	// A string, in JSON format, that is passed to the CC rule Lambda function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/config_config_rule#input_parameters ConfigConfigRule#input_parameters}
	InputParameters *string `field:"optional" json:"inputParameters" yaml:"inputParameters"`
	// The maximum frequency with which CC runs evaluations for a rule.
	//
	// You can specify a value for ``MaximumExecutionFrequency`` when:
	//   +  You are using an AWS managed rule that is triggered at a periodic frequency.
	//   +  Your custom rule is triggered when CC delivers the configuration snapshot. For more information, see [ConfigSnapshotDeliveryProperties](https://docs.aws.amazon.com/config/latest/APIReference/API_ConfigSnapshotDeliveryProperties.html).
	//
	//   By default, rules with a periodic trigger are evaluated every 24 hours. To change the frequency, specify a valid value for the ``MaximumExecutionFrequency`` parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/config_config_rule#maximum_execution_frequency ConfigConfigRule#maximum_execution_frequency}
	MaximumExecutionFrequency *string `field:"optional" json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
	// Defines which resources can trigger an evaluation for the rule.
	//
	// The scope can include one or more resource types, a combination of one resource type and one resource ID, or a combination of a tag key and value. Specify a scope to constrain the resources that can trigger an evaluation for the rule. If you do not specify a scope, evaluations are triggered when any resource in the recording group changes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/config_config_rule#scope ConfigConfigRule#scope}
	Scope *ConfigConfigRuleScope `field:"optional" json:"scope" yaml:"scope"`
}

