// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ceanomalysubscription


type CeAnomalySubscriptionSubscribers struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ce_anomaly_subscription#address CeAnomalySubscription#address}.
	Address *string `field:"required" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ce_anomaly_subscription#type CeAnomalySubscription#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ce_anomaly_subscription#status CeAnomalySubscription#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

