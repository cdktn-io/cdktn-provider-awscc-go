// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sqsqueue


type SqsQueueTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sqs_queue#key SqsQueue#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sqs_queue#value SqsQueue#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

