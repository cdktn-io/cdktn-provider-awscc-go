// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package drslaunchconfigurationtemplate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DrsLaunchConfigurationTemplateConfig struct {
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
	// Copy private IP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/drs_launch_configuration_template#copy_private_ip DrsLaunchConfigurationTemplate#copy_private_ip}
	CopyPrivateIp interface{} `field:"optional" json:"copyPrivateIp" yaml:"copyPrivateIp"`
	// Copy tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/drs_launch_configuration_template#copy_tags DrsLaunchConfigurationTemplate#copy_tags}
	CopyTags interface{} `field:"optional" json:"copyTags" yaml:"copyTags"`
	// S3 bucket ARN to export Source Network templates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/drs_launch_configuration_template#export_bucket_arn DrsLaunchConfigurationTemplate#export_bucket_arn}
	ExportBucketArn *string `field:"optional" json:"exportBucketArn" yaml:"exportBucketArn"`
	// Launch disposition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/drs_launch_configuration_template#launch_disposition DrsLaunchConfigurationTemplate#launch_disposition}
	LaunchDisposition *string `field:"optional" json:"launchDisposition" yaml:"launchDisposition"`
	// DRS will set the 'launch into instance ID' of any source server when performing a drill, recovery or failback to the previous region or availability zone, using the instance ID of the source instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/drs_launch_configuration_template#launch_into_source_instance DrsLaunchConfigurationTemplate#launch_into_source_instance}
	LaunchIntoSourceInstance interface{} `field:"optional" json:"launchIntoSourceInstance" yaml:"launchIntoSourceInstance"`
	// Configuration of a machine's license.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/drs_launch_configuration_template#licensing DrsLaunchConfigurationTemplate#licensing}
	Licensing *DrsLaunchConfigurationTemplateLicensing `field:"optional" json:"licensing" yaml:"licensing"`
	// Whether we want to activate post-launch actions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/drs_launch_configuration_template#post_launch_enabled DrsLaunchConfigurationTemplate#post_launch_enabled}
	PostLaunchEnabled interface{} `field:"optional" json:"postLaunchEnabled" yaml:"postLaunchEnabled"`
	// A set of tags associated with the Launch Configuration Template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/drs_launch_configuration_template#tags DrsLaunchConfigurationTemplate#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Target instance type right-sizing method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/drs_launch_configuration_template#target_instance_type_right_sizing_method DrsLaunchConfigurationTemplate#target_instance_type_right_sizing_method}
	TargetInstanceTypeRightSizingMethod *string `field:"optional" json:"targetInstanceTypeRightSizingMethod" yaml:"targetInstanceTypeRightSizingMethod"`
}

