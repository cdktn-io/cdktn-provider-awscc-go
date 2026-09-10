// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudformationstackset


type CloudformationStackSetManagedExecution struct {
	// When true, StackSets performs non-conflicting operations concurrently and queues conflicting operations.
	//
	// After conflicting operations finish, StackSets starts queued operations in request order.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cloudformation_stack_set#active CloudformationStackSet#active}
	Active interface{} `field:"optional" json:"active" yaml:"active"`
}

