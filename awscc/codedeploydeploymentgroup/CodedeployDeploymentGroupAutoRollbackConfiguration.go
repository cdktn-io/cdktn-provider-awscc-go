// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codedeploydeploymentgroup


type CodedeployDeploymentGroupAutoRollbackConfiguration struct {
	// Indicates whether a defined automatic rollback configuration is currently enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/codedeploy_deployment_group#enabled CodedeployDeploymentGroup#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The event type or types that trigger a rollback.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/codedeploy_deployment_group#events CodedeployDeploymentGroup#events}
	Events *[]*string `field:"optional" json:"events" yaml:"events"`
}

