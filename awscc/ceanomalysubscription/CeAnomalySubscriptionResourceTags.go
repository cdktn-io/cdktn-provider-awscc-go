// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ceanomalysubscription


type CeAnomalySubscriptionResourceTags struct {
	// The key name for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ce_anomaly_subscription#key CeAnomalySubscription#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ce_anomaly_subscription#value CeAnomalySubscription#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

