// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package configconfigrule


type ConfigConfigRuleScope struct {
	// The ID of the only AWS resource that you want to trigger an evaluation for the rule.
	//
	// If you specify a resource ID, you must specify one resource type for ``ComplianceResourceTypes``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/config_config_rule#compliance_resource_id ConfigConfigRule#compliance_resource_id}
	ComplianceResourceId *string `field:"optional" json:"complianceResourceId" yaml:"complianceResourceId"`
	// The resource types of only those AWS resources that you want to trigger an evaluation for the rule.
	//
	// You can only specify one type if you also specify a resource ID for ``ComplianceResourceId``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/config_config_rule#compliance_resource_types ConfigConfigRule#compliance_resource_types}
	ComplianceResourceTypes *[]*string `field:"optional" json:"complianceResourceTypes" yaml:"complianceResourceTypes"`
	// The tag key that is applied to only those AWS resources that you want to trigger an evaluation for the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/config_config_rule#tag_key ConfigConfigRule#tag_key}
	TagKey *string `field:"optional" json:"tagKey" yaml:"tagKey"`
	// The tag value applied to only those AWS resources that you want to trigger an evaluation for the rule.
	//
	// If you specify a value for ``TagValue``, you must also specify a value for ``TagKey``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/config_config_rule#tag_value ConfigConfigRule#tag_value}
	TagValue *string `field:"optional" json:"tagValue" yaml:"tagValue"`
}

