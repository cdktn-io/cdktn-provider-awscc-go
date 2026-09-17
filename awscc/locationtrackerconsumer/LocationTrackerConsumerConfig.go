// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package locationtrackerconsumer

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LocationTrackerConsumerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/location_tracker_consumer#consumer_arn LocationTrackerConsumer#consumer_arn}.
	ConsumerArn *string `field:"required" json:"consumerArn" yaml:"consumerArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/location_tracker_consumer#tracker_name LocationTrackerConsumer#tracker_name}.
	TrackerName *string `field:"required" json:"trackerName" yaml:"trackerName"`
}

