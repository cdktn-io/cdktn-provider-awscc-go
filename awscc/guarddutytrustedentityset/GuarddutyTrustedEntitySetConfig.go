// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package guarddutytrustedentityset

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GuarddutyTrustedEntitySetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/guardduty_trusted_entity_set#format GuarddutyTrustedEntitySet#format}.
	Format *string `field:"required" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/guardduty_trusted_entity_set#location GuarddutyTrustedEntitySet#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/guardduty_trusted_entity_set#activate GuarddutyTrustedEntitySet#activate}.
	Activate interface{} `field:"optional" json:"activate" yaml:"activate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/guardduty_trusted_entity_set#detector_id GuarddutyTrustedEntitySet#detector_id}.
	DetectorId *string `field:"optional" json:"detectorId" yaml:"detectorId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/guardduty_trusted_entity_set#expected_bucket_owner GuarddutyTrustedEntitySet#expected_bucket_owner}.
	ExpectedBucketOwner *string `field:"optional" json:"expectedBucketOwner" yaml:"expectedBucketOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/guardduty_trusted_entity_set#name GuarddutyTrustedEntitySet#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/guardduty_trusted_entity_set#tags GuarddutyTrustedEntitySet#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

