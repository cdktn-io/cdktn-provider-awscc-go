// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayusageplankey

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApigatewayUsagePlanKeyConfig struct {
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
	// The Id of the UsagePlanKey resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_usage_plan_key#key_id ApigatewayUsagePlanKey#key_id}
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_usage_plan_key#key_type ApigatewayUsagePlanKey#key_type}.
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
	// The Id of the UsagePlan resource representing the usage plan containing the UsagePlanKey resource representing a plan customer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_usage_plan_key#usage_plan_id ApigatewayUsagePlanKey#usage_plan_id}
	UsagePlanId *string `field:"required" json:"usagePlanId" yaml:"usagePlanId"`
}

