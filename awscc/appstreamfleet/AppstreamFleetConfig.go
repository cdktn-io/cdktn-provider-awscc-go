// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appstreamfleet

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppstreamFleetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/appstream_fleet#instance_type AppstreamFleet#instance_type}.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/appstream_fleet#name AppstreamFleet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/appstream_fleet#attributes_to_delete AppstreamFleet#attributes_to_delete}.
	AttributesToDelete *[]*string `field:"optional" json:"attributesToDelete" yaml:"attributesToDelete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/appstream_fleet#compute_capacity AppstreamFleet#compute_capacity}.
	ComputeCapacity *AppstreamFleetComputeCapacity `field:"optional" json:"computeCapacity" yaml:"computeCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/appstream_fleet#description AppstreamFleet#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/appstream_fleet#disable_imdsv1 AppstreamFleet#disable_imdsv1}.
	DisableImdsv1 interface{} `field:"optional" json:"disableImdsv1" yaml:"disableImdsv1"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/appstream_fleet#disconnect_timeout_in_seconds AppstreamFleet#disconnect_timeout_in_seconds}.
	DisconnectTimeoutInSeconds *float64 `field:"optional" json:"disconnectTimeoutInSeconds" yaml:"disconnectTimeoutInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/appstream_fleet#display_name AppstreamFleet#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/appstream_fleet#domain_join_info AppstreamFleet#domain_join_info}.
	DomainJoinInfo *AppstreamFleetDomainJoinInfo `field:"optional" json:"domainJoinInfo" yaml:"domainJoinInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/appstream_fleet#enable_default_internet_access AppstreamFleet#enable_default_internet_access}.
	EnableDefaultInternetAccess interface{} `field:"optional" json:"enableDefaultInternetAccess" yaml:"enableDefaultInternetAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/appstream_fleet#fleet_type AppstreamFleet#fleet_type}.
	FleetType *string `field:"optional" json:"fleetType" yaml:"fleetType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/appstream_fleet#iam_role_arn AppstreamFleet#iam_role_arn}.
	IamRoleArn *string `field:"optional" json:"iamRoleArn" yaml:"iamRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/appstream_fleet#idle_disconnect_timeout_in_seconds AppstreamFleet#idle_disconnect_timeout_in_seconds}.
	IdleDisconnectTimeoutInSeconds *float64 `field:"optional" json:"idleDisconnectTimeoutInSeconds" yaml:"idleDisconnectTimeoutInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/appstream_fleet#image_arn AppstreamFleet#image_arn}.
	ImageArn *string `field:"optional" json:"imageArn" yaml:"imageArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/appstream_fleet#image_name AppstreamFleet#image_name}.
	ImageName *string `field:"optional" json:"imageName" yaml:"imageName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/appstream_fleet#max_concurrent_sessions AppstreamFleet#max_concurrent_sessions}.
	MaxConcurrentSessions *float64 `field:"optional" json:"maxConcurrentSessions" yaml:"maxConcurrentSessions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/appstream_fleet#max_sessions_per_instance AppstreamFleet#max_sessions_per_instance}.
	MaxSessionsPerInstance *float64 `field:"optional" json:"maxSessionsPerInstance" yaml:"maxSessionsPerInstance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/appstream_fleet#max_user_duration_in_seconds AppstreamFleet#max_user_duration_in_seconds}.
	MaxUserDurationInSeconds *float64 `field:"optional" json:"maxUserDurationInSeconds" yaml:"maxUserDurationInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/appstream_fleet#platform AppstreamFleet#platform}.
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/appstream_fleet#root_volume_config AppstreamFleet#root_volume_config}.
	RootVolumeConfig *AppstreamFleetRootVolumeConfig `field:"optional" json:"rootVolumeConfig" yaml:"rootVolumeConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/appstream_fleet#session_script_s3_location AppstreamFleet#session_script_s3_location}.
	SessionScriptS3Location *AppstreamFleetSessionScriptS3Location `field:"optional" json:"sessionScriptS3Location" yaml:"sessionScriptS3Location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/appstream_fleet#stream_view AppstreamFleet#stream_view}.
	StreamView *string `field:"optional" json:"streamView" yaml:"streamView"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/appstream_fleet#tags AppstreamFleet#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/appstream_fleet#usb_device_filter_strings AppstreamFleet#usb_device_filter_strings}.
	UsbDeviceFilterStrings *[]*string `field:"optional" json:"usbDeviceFilterStrings" yaml:"usbDeviceFilterStrings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/appstream_fleet#vpc_config AppstreamFleet#vpc_config}.
	VpcConfig *AppstreamFleetVpcConfig `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

