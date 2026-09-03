// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codedeploydeploymentgroup


type CodedeployDeploymentGroupDeployment struct {
	// A description of the deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/codedeploy_deployment_group#description CodedeployDeploymentGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If true, then if an ApplicationStop, BeforeBlockTraffic, or AfterBlockTraffic deployment lifecycle event to an instance fails, then the deployment continues to the next deployment lifecycle event.
	//
	// If false or not specified, then if a lifecycle event fails during a deployment to an instance, that deployment fails. If deployment to that instance is part of an overall deployment and the number of healthy hosts is not less than the minimum number of healthy hosts, then a deployment to the next instance is attempted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/codedeploy_deployment_group#ignore_application_stop_failures CodedeployDeploymentGroup#ignore_application_stop_failures}
	IgnoreApplicationStopFailures interface{} `field:"optional" json:"ignoreApplicationStopFailures" yaml:"ignoreApplicationStopFailures"`
	// Information about the location of stored application artifacts and the service from which to retrieve them.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/codedeploy_deployment_group#revision CodedeployDeploymentGroup#revision}
	Revision *CodedeployDeploymentGroupDeploymentRevision `field:"optional" json:"revision" yaml:"revision"`
}

