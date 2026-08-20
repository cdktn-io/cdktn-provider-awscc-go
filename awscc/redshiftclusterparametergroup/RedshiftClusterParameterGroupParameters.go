// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package redshiftclusterparametergroup


type RedshiftClusterParameterGroupParameters struct {
	// The name of the parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/redshift_cluster_parameter_group#parameter_name RedshiftClusterParameterGroup#parameter_name}
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
	// The value of the parameter. If `ParameterName` is `wlm_json_configuration`, then the maximum size of `ParameterValue` is 8000 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/redshift_cluster_parameter_group#parameter_value RedshiftClusterParameterGroup#parameter_value}
	ParameterValue *string `field:"optional" json:"parameterValue" yaml:"parameterValue"`
}

