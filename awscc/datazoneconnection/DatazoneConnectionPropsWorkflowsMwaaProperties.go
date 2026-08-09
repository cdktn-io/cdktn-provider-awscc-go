// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneconnection


type DatazoneConnectionPropsWorkflowsMwaaProperties struct {
	// The name of the MWAA environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datazone_connection#mwaa_environment_name DatazoneConnection#mwaa_environment_name}
	MwaaEnvironmentName *string `field:"optional" json:"mwaaEnvironmentName" yaml:"mwaaEnvironmentName"`
}

