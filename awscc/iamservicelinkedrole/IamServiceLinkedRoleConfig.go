// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iamservicelinkedrole

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IamServiceLinkedRoleConfig struct {
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
	// The service principal for the AWS service to which this role is attached.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iam_service_linked_role#aws_service_name IamServiceLinkedRole#aws_service_name}
	AwsServiceName *string `field:"optional" json:"awsServiceName" yaml:"awsServiceName"`
	// A string that you provide, which is combined with the service-provided prefix to form the complete role name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iam_service_linked_role#custom_suffix IamServiceLinkedRole#custom_suffix}
	CustomSuffix *string `field:"optional" json:"customSuffix" yaml:"customSuffix"`
	// The description of the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iam_service_linked_role#description IamServiceLinkedRole#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

