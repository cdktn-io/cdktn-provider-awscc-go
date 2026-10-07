// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package scndataintegrationflow


type ScnDataIntegrationFlowTransformationSqlTransformation struct {
	// The transformation SQL query body based on SparkSQL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/scn_data_integration_flow#query ScnDataIntegrationFlow#query}
	Query *string `field:"optional" json:"query" yaml:"query"`
}

