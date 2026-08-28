// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package healthlakedatatransformationprofile

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type HealthlakeDataTransformationProfileConfig struct {
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
	// The human-readable name of the profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/healthlake_data_transformation_profile#profile_name HealthlakeDataTransformationProfile#profile_name}
	ProfileName *string `field:"required" json:"profileName" yaml:"profileName"`
	// The source format that this profile converts from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/healthlake_data_transformation_profile#source_format HealthlakeDataTransformationProfile#source_format}
	SourceFormat *string `field:"required" json:"sourceFormat" yaml:"sourceFormat"`
	// The identifier (key ID or ARN) of a customer-managed KMS key used to encrypt the profile's template content at rest.
	//
	// If omitted, an AWS owned key is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/healthlake_data_transformation_profile#kms_key_id HealthlakeDataTransformationProfile#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// A human-readable description of the profile's purpose.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/healthlake_data_transformation_profile#profile_description HealthlakeDataTransformationProfile#profile_description}
	ProfileDescription *string `field:"optional" json:"profileDescription" yaml:"profileDescription"`
	// The source from which to create the profile's initial template content.
	//
	// Exactly one of the members must be specified. Use StarterProfile (C-CDA only), ProfileMapping (C-CDA or CSV), or ExistingVersionedProfileId to clone an existing profile. Each produces a published profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/healthlake_data_transformation_profile#source HealthlakeDataTransformationProfile#source}
	Source *HealthlakeDataTransformationProfileSource `field:"optional" json:"source" yaml:"source"`
	// An array of key-value pairs to apply to this profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/healthlake_data_transformation_profile#tags HealthlakeDataTransformationProfile#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

