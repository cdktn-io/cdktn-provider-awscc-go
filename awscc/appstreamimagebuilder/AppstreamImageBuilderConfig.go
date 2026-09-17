// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appstreamimagebuilder

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppstreamImageBuilderConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appstream_image_builder#instance_type AppstreamImageBuilder#instance_type}.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appstream_image_builder#name AppstreamImageBuilder#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appstream_image_builder#access_endpoints AppstreamImageBuilder#access_endpoints}.
	AccessEndpoints interface{} `field:"optional" json:"accessEndpoints" yaml:"accessEndpoints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appstream_image_builder#appstream_agent_version AppstreamImageBuilder#appstream_agent_version}.
	AppstreamAgentVersion *string `field:"optional" json:"appstreamAgentVersion" yaml:"appstreamAgentVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appstream_image_builder#description AppstreamImageBuilder#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appstream_image_builder#disable_imdsv1 AppstreamImageBuilder#disable_imdsv1}.
	DisableImdsv1 interface{} `field:"optional" json:"disableImdsv1" yaml:"disableImdsv1"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appstream_image_builder#display_name AppstreamImageBuilder#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appstream_image_builder#domain_join_info AppstreamImageBuilder#domain_join_info}.
	DomainJoinInfo *AppstreamImageBuilderDomainJoinInfo `field:"optional" json:"domainJoinInfo" yaml:"domainJoinInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appstream_image_builder#enable_default_internet_access AppstreamImageBuilder#enable_default_internet_access}.
	EnableDefaultInternetAccess interface{} `field:"optional" json:"enableDefaultInternetAccess" yaml:"enableDefaultInternetAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appstream_image_builder#iam_role_arn AppstreamImageBuilder#iam_role_arn}.
	IamRoleArn *string `field:"optional" json:"iamRoleArn" yaml:"iamRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appstream_image_builder#image_arn AppstreamImageBuilder#image_arn}.
	ImageArn *string `field:"optional" json:"imageArn" yaml:"imageArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appstream_image_builder#image_name AppstreamImageBuilder#image_name}.
	ImageName *string `field:"optional" json:"imageName" yaml:"imageName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appstream_image_builder#root_volume_config AppstreamImageBuilder#root_volume_config}.
	RootVolumeConfig *AppstreamImageBuilderRootVolumeConfig `field:"optional" json:"rootVolumeConfig" yaml:"rootVolumeConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appstream_image_builder#softwares_to_install AppstreamImageBuilder#softwares_to_install}.
	SoftwaresToInstall *[]*string `field:"optional" json:"softwaresToInstall" yaml:"softwaresToInstall"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appstream_image_builder#softwares_to_uninstall AppstreamImageBuilder#softwares_to_uninstall}.
	SoftwaresToUninstall *[]*string `field:"optional" json:"softwaresToUninstall" yaml:"softwaresToUninstall"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appstream_image_builder#tags AppstreamImageBuilder#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appstream_image_builder#vpc_config AppstreamImageBuilder#vpc_config}.
	VpcConfig *AppstreamImageBuilderVpcConfig `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

