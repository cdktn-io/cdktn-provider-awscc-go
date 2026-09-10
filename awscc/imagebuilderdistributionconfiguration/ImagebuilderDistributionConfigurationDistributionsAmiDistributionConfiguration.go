// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuilderdistributionconfiguration


type ImagebuilderDistributionConfigurationDistributionsAmiDistributionConfiguration struct {
	// The tags to apply to AMIs distributed to this Region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/imagebuilder_distribution_configuration#ami_tags ImagebuilderDistributionConfiguration#ami_tags}
	AmiTags *map[string]*string `field:"optional" json:"amiTags" yaml:"amiTags"`
	// The description of the AMI distribution configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/imagebuilder_distribution_configuration#description ImagebuilderDistributionConfiguration#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The KMS key identifier used to encrypt the distributed image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/imagebuilder_distribution_configuration#kms_key_id ImagebuilderDistributionConfiguration#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Launch permissions can be used to configure which AWS accounts can use the AMI to launch instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/imagebuilder_distribution_configuration#launch_permission_configuration ImagebuilderDistributionConfiguration#launch_permission_configuration}
	LaunchPermissionConfiguration *ImagebuilderDistributionConfigurationDistributionsAmiDistributionConfigurationLaunchPermissionConfiguration `field:"optional" json:"launchPermissionConfiguration" yaml:"launchPermissionConfiguration"`
	// The name of the AMI distribution configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/imagebuilder_distribution_configuration#name ImagebuilderDistributionConfiguration#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ID of accounts to which you want to distribute an image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/imagebuilder_distribution_configuration#target_account_ids ImagebuilderDistributionConfiguration#target_account_ids}
	TargetAccountIds *[]*string `field:"optional" json:"targetAccountIds" yaml:"targetAccountIds"`
}

