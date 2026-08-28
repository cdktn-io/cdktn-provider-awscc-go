// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package guarddutydetector

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GuarddutyDetectorConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/guardduty_detector#enable GuarddutyDetector#enable}.
	Enable interface{} `field:"required" json:"enable" yaml:"enable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/guardduty_detector#data_sources GuarddutyDetector#data_sources}.
	DataSources *GuarddutyDetectorDataSources `field:"optional" json:"dataSources" yaml:"dataSources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/guardduty_detector#features GuarddutyDetector#features}.
	Features interface{} `field:"optional" json:"features" yaml:"features"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/guardduty_detector#finding_publishing_frequency GuarddutyDetector#finding_publishing_frequency}.
	FindingPublishingFrequency *string `field:"optional" json:"findingPublishingFrequency" yaml:"findingPublishingFrequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/guardduty_detector#tags GuarddutyDetector#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

