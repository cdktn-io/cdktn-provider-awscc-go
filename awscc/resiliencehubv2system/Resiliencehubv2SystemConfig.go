// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubv2system

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Resiliencehubv2SystemConfig struct {
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
	// The name of the system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/resiliencehubv2_system#name Resiliencehubv2System#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/resiliencehubv2_system#description Resiliencehubv2System#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The KMS key ID for encrypting system data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/resiliencehubv2_system#kms_key_id Resiliencehubv2System#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Whether the system is enabled to be shared with other members of the Organization.
	//
	// Only applicable if the system owner is a management account or delegated admin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/resiliencehubv2_system#sharing_enabled Resiliencehubv2System#sharing_enabled}
	SharingEnabled interface{} `field:"optional" json:"sharingEnabled" yaml:"sharingEnabled"`
	// Tags assigned to the system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/resiliencehubv2_system#tags Resiliencehubv2System#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

