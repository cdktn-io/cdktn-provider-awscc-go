// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncresolver


type AppsyncResolverSyncConfig struct {
	// The Conflict Detection strategy to use.
	//
	// +  *VERSION*: Detect conflicts based on object versions for this resolver.
	//   +  *NONE*: Do not detect conflicts when invoking this resolver.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appsync_resolver#conflict_detection AppsyncResolver#conflict_detection}
	ConflictDetection *string `field:"optional" json:"conflictDetection" yaml:"conflictDetection"`
	// The Conflict Resolution strategy to perform in the event of a conflict.
	//
	// +  *OPTIMISTIC_CONCURRENCY*: Resolve conflicts by rejecting mutations when versions don't match the latest version at the server.
	//   +  *AUTOMERGE*: Resolve conflicts with the Automerge conflict resolution strategy.
	//   +  *LAMBDA*: Resolve conflicts with an LAMlong function supplied in the ``LambdaConflictHandlerConfig``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appsync_resolver#conflict_handler AppsyncResolver#conflict_handler}
	ConflictHandler *string `field:"optional" json:"conflictHandler" yaml:"conflictHandler"`
	// The ``LambdaConflictHandlerConfig`` when configuring ``LAMBDA`` as the Conflict Handler.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appsync_resolver#lambda_conflict_handler_config AppsyncResolver#lambda_conflict_handler_config}
	LambdaConflictHandlerConfig *AppsyncResolverSyncConfigLambdaConflictHandlerConfig `field:"optional" json:"lambdaConflictHandlerConfig" yaml:"lambdaConflictHandlerConfig"`
}

