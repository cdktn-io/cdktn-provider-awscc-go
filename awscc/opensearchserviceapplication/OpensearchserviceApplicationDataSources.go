// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchserviceapplication


type OpensearchserviceApplicationDataSources struct {
	// The ARN of the data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/opensearchservice_application#data_source_arn OpensearchserviceApplication#data_source_arn}
	DataSourceArn *string `field:"optional" json:"dataSourceArn" yaml:"dataSourceArn"`
	// Description of the data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/opensearchservice_application#data_source_description OpensearchserviceApplication#data_source_description}
	DataSourceDescription *string `field:"optional" json:"dataSourceDescription" yaml:"dataSourceDescription"`
}

