// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networksecuritymanagerdeployment


type NetworksecuritymanagerDeploymentDeploymentConfiguration struct {
	// Whether cross-account visibility is enabled for the deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/networksecuritymanager_deployment#enable_cross_account_visibility NetworksecuritymanagerDeployment#enable_cross_account_visibility}
	EnableCrossAccountVisibility interface{} `field:"optional" json:"enableCrossAccountVisibility" yaml:"enableCrossAccountVisibility"`
}

