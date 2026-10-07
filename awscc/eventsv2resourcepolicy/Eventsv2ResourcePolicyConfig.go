// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2resourcepolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Eventsv2ResourcePolicyConfig struct {
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
	// The Amazon Resource Name (ARN) of the event bus whose resource policy this is.
	//
	// The bus must already exist. This resource does not create it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/eventsv2_resource_policy#event_bus_arn Eventsv2ResourcePolicy#event_bus_arn}
	EventBusArn *string `field:"required" json:"eventBusArn" yaml:"eventBusArn"`
	// The resource policy document, as a JSON object.
	//
	// The document can be up to 20 KB. This quota is adjustable. An empty object is not a valid policy. To remove the policy, delete this resource. The principals in the document must exist and be visible to the service when the policy is written. When you create a new IAM role or user, that principal might not be immediately visible to the service. You might need to enforce a delay before you include it in the document. For more information, see "Changes that I make are not always immediately visible" in the IAM User Guide. Declare Version. Write AWS account and role principals as ARNs rather than as account IDs. Write a single Action, Resource, or principal value as a scalar rather than as a one-element list. The service returns these forms as you wrote them. It normalizes other forms, and a normalized value can appear as drift. A stack update replaces the whole policy with this document, including any change made outside CloudFormation. For more information about event bus resource policies, see the Amazon EventBridge User Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/eventsv2_resource_policy#policy_document Eventsv2ResourcePolicy#policy_document}
	PolicyDocument *string `field:"required" json:"policyDocument" yaml:"policyDocument"`
}

