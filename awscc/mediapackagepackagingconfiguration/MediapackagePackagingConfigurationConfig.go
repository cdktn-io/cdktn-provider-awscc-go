// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagepackagingconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediapackagePackagingConfigurationConfig struct {
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
	// The ID of the PackagingConfiguration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediapackage_packaging_configuration#packaging_configuration_id MediapackagePackagingConfiguration#packaging_configuration_id}
	PackagingConfigurationId *string `field:"required" json:"packagingConfigurationId" yaml:"packagingConfigurationId"`
	// The ID of a PackagingGroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediapackage_packaging_configuration#packaging_group_id MediapackagePackagingConfiguration#packaging_group_id}
	PackagingGroupId *string `field:"required" json:"packagingGroupId" yaml:"packagingGroupId"`
	// A CMAF packaging configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediapackage_packaging_configuration#cmaf_package MediapackagePackagingConfiguration#cmaf_package}
	CmafPackage *MediapackagePackagingConfigurationCmafPackage `field:"optional" json:"cmafPackage" yaml:"cmafPackage"`
	// A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediapackage_packaging_configuration#dash_package MediapackagePackagingConfiguration#dash_package}
	DashPackage *MediapackagePackagingConfigurationDashPackage `field:"optional" json:"dashPackage" yaml:"dashPackage"`
	// An HTTP Live Streaming (HLS) packaging configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediapackage_packaging_configuration#hls_package MediapackagePackagingConfiguration#hls_package}
	HlsPackage *MediapackagePackagingConfigurationHlsPackage `field:"optional" json:"hlsPackage" yaml:"hlsPackage"`
	// A Microsoft Smooth Streaming (MSS) PackagingConfiguration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediapackage_packaging_configuration#mss_package MediapackagePackagingConfiguration#mss_package}
	MssPackage *MediapackagePackagingConfigurationMssPackage `field:"optional" json:"mssPackage" yaml:"mssPackage"`
	// A collection of tags associated with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediapackage_packaging_configuration#tags MediapackagePackagingConfiguration#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

