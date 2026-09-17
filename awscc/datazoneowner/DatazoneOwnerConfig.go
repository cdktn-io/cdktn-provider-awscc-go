// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneowner

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatazoneOwnerConfig struct {
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
	// The ID of the domain in which you want to add the entity owner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_owner#domain_identifier DatazoneOwner#domain_identifier}
	DomainIdentifier *string `field:"required" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The ID of the entity to which you want to add an owner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_owner#entity_identifier DatazoneOwner#entity_identifier}
	EntityIdentifier *string `field:"required" json:"entityIdentifier" yaml:"entityIdentifier"`
	// The type of an entity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_owner#entity_type DatazoneOwner#entity_type}
	EntityType *string `field:"required" json:"entityType" yaml:"entityType"`
	// The owner that you want to add to the entity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_owner#owner DatazoneOwner#owner}
	Owner *DatazoneOwnerOwner `field:"required" json:"owner" yaml:"owner"`
}

