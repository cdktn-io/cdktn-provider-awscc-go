// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codedeploydeploymentgroup


type CodedeployDeploymentGroupOnPremisesTagSet struct {
	// A list that contains other lists of on-premises instance tag groups.
	//
	// For an instance to be included in the deployment group, it must be identified by all of the tag groups in the list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/codedeploy_deployment_group#on_premises_tag_set_list CodedeployDeploymentGroup#on_premises_tag_set_list}
	OnPremisesTagSetList interface{} `field:"optional" json:"onPremisesTagSetList" yaml:"onPremisesTagSetList"`
}

