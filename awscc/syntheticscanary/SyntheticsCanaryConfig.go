// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package syntheticscanary

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SyntheticsCanaryConfig struct {
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
	// Provide the s3 bucket output location for test results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/synthetics_canary#artifact_s3_location SyntheticsCanary#artifact_s3_location}
	ArtifactS3Location *string `field:"required" json:"artifactS3Location" yaml:"artifactS3Location"`
	// Provide the canary script source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/synthetics_canary#code SyntheticsCanary#code}
	Code *SyntheticsCanaryCode `field:"required" json:"code" yaml:"code"`
	// Lambda Execution role used to run your canaries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/synthetics_canary#execution_role_arn SyntheticsCanary#execution_role_arn}
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// Name of the canary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/synthetics_canary#name SyntheticsCanary#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Runtime version of Synthetics Library.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/synthetics_canary#runtime_version SyntheticsCanary#runtime_version}
	RuntimeVersion *string `field:"required" json:"runtimeVersion" yaml:"runtimeVersion"`
	// Frequency to run your canaries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/synthetics_canary#schedule SyntheticsCanary#schedule}
	Schedule *SyntheticsCanarySchedule `field:"required" json:"schedule" yaml:"schedule"`
	// Provide artifact configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/synthetics_canary#artifact_config SyntheticsCanary#artifact_config}
	ArtifactConfig *SyntheticsCanaryArtifactConfig `field:"optional" json:"artifactConfig" yaml:"artifactConfig"`
	// List of browser configurations for the canary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/synthetics_canary#browser_configs SyntheticsCanary#browser_configs}
	BrowserConfigs interface{} `field:"optional" json:"browserConfigs" yaml:"browserConfigs"`
	// Deletes associated lambda resources created by Synthetics if set to True. Default is False.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/synthetics_canary#delete_lambda_resources_on_canary_deletion SyntheticsCanary#delete_lambda_resources_on_canary_deletion}
	DeleteLambdaResourcesOnCanaryDeletion interface{} `field:"optional" json:"deleteLambdaResourcesOnCanaryDeletion" yaml:"deleteLambdaResourcesOnCanaryDeletion"`
	// Setting to control if UpdateCanary will perform a DryRun and validate it is PASSING before performing the Update.
	//
	// Default is FALSE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/synthetics_canary#dry_run_and_update SyntheticsCanary#dry_run_and_update}
	DryRunAndUpdate interface{} `field:"optional" json:"dryRunAndUpdate" yaml:"dryRunAndUpdate"`
	// Retention period of failed canary runs represented in number of days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/synthetics_canary#failure_retention_period SyntheticsCanary#failure_retention_period}
	FailureRetentionPeriod *float64 `field:"optional" json:"failureRetentionPeriod" yaml:"failureRetentionPeriod"`
	// Setting to control if provisioned resources created by Synthetics are deleted alongside the canary. Default is AUTOMATIC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/synthetics_canary#provisioned_resource_cleanup SyntheticsCanary#provisioned_resource_cleanup}
	ProvisionedResourceCleanup *string `field:"optional" json:"provisionedResourceCleanup" yaml:"provisionedResourceCleanup"`
	// List of resources which canary tags should be replicated to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/synthetics_canary#resources_to_replicate_tags SyntheticsCanary#resources_to_replicate_tags}
	ResourcesToReplicateTags *[]*string `field:"optional" json:"resourcesToReplicateTags" yaml:"resourcesToReplicateTags"`
	// Provide canary run configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/synthetics_canary#run_config SyntheticsCanary#run_config}
	RunConfig *SyntheticsCanaryRunConfig `field:"optional" json:"runConfig" yaml:"runConfig"`
	// Runs canary if set to True. Default is False.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/synthetics_canary#start_canary_after_creation SyntheticsCanary#start_canary_after_creation}
	StartCanaryAfterCreation interface{} `field:"optional" json:"startCanaryAfterCreation" yaml:"startCanaryAfterCreation"`
	// Retention period of successful canary runs represented in number of days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/synthetics_canary#success_retention_period SyntheticsCanary#success_retention_period}
	SuccessRetentionPeriod *float64 `field:"optional" json:"successRetentionPeriod" yaml:"successRetentionPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/synthetics_canary#tags SyntheticsCanary#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Visual reference configuration for visual testing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/synthetics_canary#visual_reference SyntheticsCanary#visual_reference}
	VisualReference *SyntheticsCanaryVisualReference `field:"optional" json:"visualReference" yaml:"visualReference"`
	// List of visual references for the canary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/synthetics_canary#visual_references SyntheticsCanary#visual_references}
	VisualReferences interface{} `field:"optional" json:"visualReferences" yaml:"visualReferences"`
	// Provide VPC Configuration if enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/synthetics_canary#vpc_config SyntheticsCanary#vpc_config}
	VpcConfig *SyntheticsCanaryVpcConfig `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

