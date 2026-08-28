// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcscluster


type PcsClusterSlurmConfigurationCgroupCustomSettings struct {
	// The cgroup.conf parameter name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pcs_cluster#parameter_name PcsCluster#parameter_name}
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
	// The value for the cgroup.conf parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pcs_cluster#parameter_value PcsCluster#parameter_value}
	ParameterValue *string `field:"optional" json:"parameterValue" yaml:"parameterValue"`
}

