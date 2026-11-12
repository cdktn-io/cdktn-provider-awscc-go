// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lakeformationresource

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LakeformationResourceConfig struct {
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
	// The Amazon Resource Name (ARN) of the Data Source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/lakeformation_resource#resource_arn LakeformationResource#resource_arn}
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// Designates a trusted caller, an IAM principal, by registering this caller with the Data Catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/lakeformation_resource#use_service_linked_role LakeformationResource#use_service_linked_role}
	UseServiceLinkedRole interface{} `field:"required" json:"useServiceLinkedRole" yaml:"useServiceLinkedRole"`
	// Indicates whether the data access of tables pointing to the location can be managed by both Lake Formation permissions as well as Amazon S3 bucket policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/lakeformation_resource#hybrid_access_enabled LakeformationResource#hybrid_access_enabled}
	HybridAccessEnabled interface{} `field:"optional" json:"hybridAccessEnabled" yaml:"hybridAccessEnabled"`
	// The IAM role that registered a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/lakeformation_resource#role_arn LakeformationResource#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Allows Lake Formation to assume a role to access tables in a federated database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/lakeformation_resource#with_federation LakeformationResource#with_federation}
	WithFederation interface{} `field:"optional" json:"withFederation" yaml:"withFederation"`
}

