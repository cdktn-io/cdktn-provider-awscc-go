// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetrigger


type GlueTriggerEventBatchingCondition struct {
	// Number of events that must be received from Amazon EventBridge before EventBridge event trigger fires.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_trigger#batch_size GlueTrigger#batch_size}
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// Window of time in seconds after which EventBridge event trigger fires. Window starts when first event is received.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_trigger#batch_window GlueTrigger#batch_window}
	BatchWindow *float64 `field:"optional" json:"batchWindow" yaml:"batchWindow"`
}

