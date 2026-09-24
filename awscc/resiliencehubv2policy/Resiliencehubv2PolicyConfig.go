// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubv2policy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Resiliencehubv2PolicyConfig struct {
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
	// The name of the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/resiliencehubv2_policy#name Resiliencehubv2Policy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/resiliencehubv2_policy#availability_slo Resiliencehubv2Policy#availability_slo}.
	AvailabilitySlo *Resiliencehubv2PolicyAvailabilitySlo `field:"optional" json:"availabilitySlo" yaml:"availabilitySlo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/resiliencehubv2_policy#data_recovery Resiliencehubv2Policy#data_recovery}.
	DataRecovery *Resiliencehubv2PolicyDataRecovery `field:"optional" json:"dataRecovery" yaml:"dataRecovery"`
	// The description of the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/resiliencehubv2_policy#description Resiliencehubv2Policy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The KMS key ID for encrypting policy data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/resiliencehubv2_policy#kms_key_id Resiliencehubv2Policy#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/resiliencehubv2_policy#multi_az Resiliencehubv2Policy#multi_az}.
	MultiAz *Resiliencehubv2PolicyMultiAz `field:"optional" json:"multiAz" yaml:"multiAz"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/resiliencehubv2_policy#multi_region Resiliencehubv2Policy#multi_region}.
	MultiRegion *Resiliencehubv2PolicyMultiRegion `field:"optional" json:"multiRegion" yaml:"multiRegion"`
	// Whether the policy is enabled to be shared with other members of the Organization.
	//
	// Only applicable if the policy owner is a management account or delegated admin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/resiliencehubv2_policy#sharing_enabled Resiliencehubv2Policy#sharing_enabled}
	SharingEnabled interface{} `field:"optional" json:"sharingEnabled" yaml:"sharingEnabled"`
	// Tags assigned to the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/resiliencehubv2_policy#tags Resiliencehubv2Policy#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

