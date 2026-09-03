// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appflowflow


type AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoData struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appflow_flow#object_path AppflowFlow#object_path}.
	ObjectPath *string `field:"optional" json:"objectPath" yaml:"objectPath"`
	// SAP Source connector page size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appflow_flow#pagination_config AppflowFlow#pagination_config}
	PaginationConfig *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataPaginationConfig `field:"optional" json:"paginationConfig" yaml:"paginationConfig"`
	// SAP Source connector parallelism factor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appflow_flow#parallelism_config AppflowFlow#parallelism_config}
	ParallelismConfig *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataParallelismConfig `field:"optional" json:"parallelismConfig" yaml:"parallelismConfig"`
}

