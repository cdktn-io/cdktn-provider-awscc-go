// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package scndataintegrationflow


type ScnDataIntegrationFlowTransformation struct {
	// The transformation type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/scn_data_integration_flow#transformation_type ScnDataIntegrationFlow#transformation_type}
	TransformationType *string `field:"required" json:"transformationType" yaml:"transformationType"`
	// The SQL transformation configuration parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/scn_data_integration_flow#sql_transformation ScnDataIntegrationFlow#sql_transformation}
	SqlTransformation *ScnDataIntegrationFlowTransformationSqlTransformation `field:"optional" json:"sqlTransformation" yaml:"sqlTransformation"`
}

