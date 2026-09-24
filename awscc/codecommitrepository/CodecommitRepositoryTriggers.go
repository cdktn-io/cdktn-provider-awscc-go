// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codecommitrepository


type CodecommitRepositoryTriggers struct {
	// The branches to be included in the trigger configuration.
	//
	// If you specify an empty array, the trigger applies to all branches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/codecommit_repository#branches CodecommitRepository#branches}
	Branches *[]*string `field:"optional" json:"branches" yaml:"branches"`
	// Any custom data associated with the trigger to be included in the information sent to the target of the trigger.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/codecommit_repository#custom_data CodecommitRepository#custom_data}
	CustomData *string `field:"optional" json:"customData" yaml:"customData"`
	// The ARN of the resource that is the target for a trigger (for example, the ARN of a topic in Amazon SNS).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/codecommit_repository#destination_arn CodecommitRepository#destination_arn}
	DestinationArn *string `field:"optional" json:"destinationArn" yaml:"destinationArn"`
	// The repository events that cause the trigger to run actions in another service, such as sending a notification through Amazon SNS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/codecommit_repository#events CodecommitRepository#events}
	Events *[]*string `field:"optional" json:"events" yaml:"events"`
	// The name of the trigger.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/codecommit_repository#name CodecommitRepository#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

