// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package nimblestudiostreamingimage

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NimblestudioStreamingImageConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/nimblestudio_streaming_image#ec_2_image_id NimblestudioStreamingImage#ec_2_image_id}.
	Ec2ImageId *string `field:"required" json:"ec2ImageId" yaml:"ec2ImageId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/nimblestudio_streaming_image#name NimblestudioStreamingImage#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/nimblestudio_streaming_image#studio_id NimblestudioStreamingImage#studio_id}.
	StudioId *string `field:"required" json:"studioId" yaml:"studioId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/nimblestudio_streaming_image#description NimblestudioStreamingImage#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/nimblestudio_streaming_image#encryption_configuration_key_arn NimblestudioStreamingImage#encryption_configuration_key_arn}.
	EncryptionConfigurationKeyArn *string `field:"optional" json:"encryptionConfigurationKeyArn" yaml:"encryptionConfigurationKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/nimblestudio_streaming_image#encryption_configuration_key_type NimblestudioStreamingImage#encryption_configuration_key_type}.
	EncryptionConfigurationKeyType *string `field:"optional" json:"encryptionConfigurationKeyType" yaml:"encryptionConfigurationKeyType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/nimblestudio_streaming_image#tags NimblestudioStreamingImage#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

