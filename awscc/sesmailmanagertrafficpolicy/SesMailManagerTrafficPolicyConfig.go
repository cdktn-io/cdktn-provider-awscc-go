// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesmailmanagertrafficpolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SesMailManagerTrafficPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ses_mail_manager_traffic_policy#default_action SesMailManagerTrafficPolicy#default_action}.
	DefaultAction *string `field:"required" json:"defaultAction" yaml:"defaultAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ses_mail_manager_traffic_policy#policy_statements SesMailManagerTrafficPolicy#policy_statements}.
	PolicyStatements interface{} `field:"required" json:"policyStatements" yaml:"policyStatements"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ses_mail_manager_traffic_policy#max_message_size_bytes SesMailManagerTrafficPolicy#max_message_size_bytes}.
	MaxMessageSizeBytes *float64 `field:"optional" json:"maxMessageSizeBytes" yaml:"maxMessageSizeBytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ses_mail_manager_traffic_policy#tags SesMailManagerTrafficPolicy#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ses_mail_manager_traffic_policy#traffic_policy_name SesMailManagerTrafficPolicy#traffic_policy_name}.
	TrafficPolicyName *string `field:"optional" json:"trafficPolicyName" yaml:"trafficPolicyName"`
}

