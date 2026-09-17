// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appstreamstack


type AppstreamStackAgentAccessConfig struct {
	// The Amazon Resource Name (ARN) of the Amazon S3 bucket where agent screenshots are stored.
	//
	// Required when ScreenshotsUploadEnabled is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appstream_stack#s3_bucket_arn AppstreamStack#s3_bucket_arn}
	S3BucketArn *string `field:"optional" json:"s3BucketArn" yaml:"s3BucketArn"`
	// The image format for agent screen captures.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appstream_stack#screen_image_format AppstreamStack#screen_image_format}
	ScreenImageFormat *string `field:"optional" json:"screenImageFormat" yaml:"screenImageFormat"`
	// The screen resolution for the agent streaming environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appstream_stack#screen_resolution AppstreamStack#screen_resolution}
	ScreenResolution *string `field:"optional" json:"screenResolution" yaml:"screenResolution"`
	// Indicates whether screenshot uploads to Amazon S3 are enabled for agent sessions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appstream_stack#screenshots_upload_enabled AppstreamStack#screenshots_upload_enabled}
	ScreenshotsUploadEnabled interface{} `field:"optional" json:"screenshotsUploadEnabled" yaml:"screenshotsUploadEnabled"`
	// The list of agent access settings that define permissions for each agent action.
	//
	// You must specify at least one setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appstream_stack#settings AppstreamStack#settings}
	Settings interface{} `field:"optional" json:"settings" yaml:"settings"`
	// The user control mode for agent sessions.
	//
	// This setting determines how users can interact with agent sessions. Valid values are VIEW_ONLY, VIEW_STOP, and DISABLED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appstream_stack#user_control_mode AppstreamStack#user_control_mode}
	UserControlMode *string `field:"optional" json:"userControlMode" yaml:"userControlMode"`
}

