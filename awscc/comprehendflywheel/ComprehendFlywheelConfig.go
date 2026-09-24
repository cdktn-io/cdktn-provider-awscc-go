// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package comprehendflywheel

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ComprehendFlywheelConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/comprehend_flywheel#data_access_role_arn ComprehendFlywheel#data_access_role_arn}.
	DataAccessRoleArn *string `field:"required" json:"dataAccessRoleArn" yaml:"dataAccessRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/comprehend_flywheel#data_lake_s3_uri ComprehendFlywheel#data_lake_s3_uri}.
	DataLakeS3Uri *string `field:"required" json:"dataLakeS3Uri" yaml:"dataLakeS3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/comprehend_flywheel#flywheel_name ComprehendFlywheel#flywheel_name}.
	FlywheelName *string `field:"required" json:"flywheelName" yaml:"flywheelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/comprehend_flywheel#active_model_arn ComprehendFlywheel#active_model_arn}.
	ActiveModelArn *string `field:"optional" json:"activeModelArn" yaml:"activeModelArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/comprehend_flywheel#data_security_config ComprehendFlywheel#data_security_config}.
	DataSecurityConfig *ComprehendFlywheelDataSecurityConfig `field:"optional" json:"dataSecurityConfig" yaml:"dataSecurityConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/comprehend_flywheel#model_type ComprehendFlywheel#model_type}.
	ModelType *string `field:"optional" json:"modelType" yaml:"modelType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/comprehend_flywheel#tags ComprehendFlywheel#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/comprehend_flywheel#task_config ComprehendFlywheel#task_config}.
	TaskConfig *ComprehendFlywheelTaskConfig `field:"optional" json:"taskConfig" yaml:"taskConfig"`
}

