// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package configconfigrule


type ConfigConfigRuleSource struct {
	// Indicates whether AWS or the customer owns and manages the CC rule.
	//
	// CC Managed Rules are predefined rules owned by AWS. For more information, see [Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html) in the *developer guide*.
	//  CC Custom Rules are rules that you can develop either with Guard (``CUSTOM_POLICY``) or LAMlong (``CUSTOM_LAMBDA``). For more information, see [Custom Rules](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_develop-rules.html) in the *developer guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/config_config_rule#owner ConfigConfigRule#owner}
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// Provides the runtime system, policy definition, and whether debug logging is enabled. Required when owner is set to ``CUSTOM_POLICY``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/config_config_rule#custom_policy_details ConfigConfigRule#custom_policy_details}
	CustomPolicyDetails *ConfigConfigRuleSourceCustomPolicyDetails `field:"optional" json:"customPolicyDetails" yaml:"customPolicyDetails"`
	// Provides the source and the message types that cause CC to evaluate your AWS resources against a rule.
	//
	// It also provides the frequency with which you want CC to run evaluations for the rule if the trigger type is periodic.
	//  If the owner is set to ``CUSTOM_POLICY``, the only acceptable values for the CC rule trigger message type are ``ConfigurationItemChangeNotification`` and ``OversizedConfigurationItemChangeNotification``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/config_config_rule#source_details ConfigConfigRule#source_details}
	SourceDetails interface{} `field:"optional" json:"sourceDetails" yaml:"sourceDetails"`
	// For CC Managed rules, a predefined identifier from a list.
	//
	// For example, ``IAM_PASSWORD_POLICY`` is a managed rule. To reference a managed rule, see [List of Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html).
	//  For CC Custom Lambda rules, the identifier is the Amazon Resource Name (ARN) of the rule's LAMlong function, such as ``arn:aws:lambda:us-east-2:123456789012:function:custom_rule_name``.
	//  For CC Custom Policy rules, this field will be ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/config_config_rule#source_identifier ConfigConfigRule#source_identifier}
	SourceIdentifier *string `field:"optional" json:"sourceIdentifier" yaml:"sourceIdentifier"`
}

