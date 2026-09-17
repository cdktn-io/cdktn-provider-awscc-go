// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneenvironmentprofile

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatazoneEnvironmentProfileConfig struct {
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
	// The AWS account in which the Amazon DataZone environment is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_environment_profile#aws_account_id DatazoneEnvironmentProfile#aws_account_id}
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// The AWS region in which this environment profile is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_environment_profile#aws_account_region DatazoneEnvironmentProfile#aws_account_region}
	AwsAccountRegion *string `field:"required" json:"awsAccountRegion" yaml:"awsAccountRegion"`
	// The ID of the Amazon DataZone domain in which this environment profile is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_environment_profile#domain_identifier DatazoneEnvironmentProfile#domain_identifier}
	DomainIdentifier *string `field:"required" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The ID of the blueprint with which this environment profile is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_environment_profile#environment_blueprint_identifier DatazoneEnvironmentProfile#environment_blueprint_identifier}
	EnvironmentBlueprintIdentifier *string `field:"required" json:"environmentBlueprintIdentifier" yaml:"environmentBlueprintIdentifier"`
	// The name of this Amazon DataZone environment profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_environment_profile#name DatazoneEnvironmentProfile#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The identifier of the project in which to create the environment profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_environment_profile#project_identifier DatazoneEnvironmentProfile#project_identifier}
	ProjectIdentifier *string `field:"required" json:"projectIdentifier" yaml:"projectIdentifier"`
	// The description of this Amazon DataZone environment profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_environment_profile#description DatazoneEnvironmentProfile#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The user parameters of this Amazon DataZone environment profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_environment_profile#user_parameters DatazoneEnvironmentProfile#user_parameters}
	UserParameters interface{} `field:"optional" json:"userParameters" yaml:"userParameters"`
}

