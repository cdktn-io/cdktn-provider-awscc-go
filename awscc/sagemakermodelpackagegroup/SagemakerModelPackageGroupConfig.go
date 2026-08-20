// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelpackagegroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerModelPackageGroupConfig struct {
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
	// The name of the model package group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_model_package_group#model_package_group_name SagemakerModelPackageGroup#model_package_group_name}
	ModelPackageGroupName *string `field:"required" json:"modelPackageGroupName" yaml:"modelPackageGroupName"`
	// The description of the model package group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_model_package_group#model_package_group_description SagemakerModelPackageGroup#model_package_group_description}
	ModelPackageGroupDescription *string `field:"optional" json:"modelPackageGroupDescription" yaml:"modelPackageGroupDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_model_package_group#model_package_group_policy SagemakerModelPackageGroup#model_package_group_policy}.
	ModelPackageGroupPolicy *string `field:"optional" json:"modelPackageGroupPolicy" yaml:"modelPackageGroupPolicy"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_model_package_group#tags SagemakerModelPackageGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

