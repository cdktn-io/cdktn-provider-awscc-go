// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubconnectorv2


type SecurityhubConnectorV2ProviderNameServiceNow struct {
	// The instance name of ServiceNow ITSM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/securityhub_connector_v2#instance_name SecurityhubConnectorV2#instance_name}
	InstanceName *string `field:"optional" json:"instanceName" yaml:"instanceName"`
	// The Amazon Resource Name (ARN) of the AWS Secrets Manager secret that contains the ServiceNow credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/securityhub_connector_v2#secret_arn SecurityhubConnectorV2#secret_arn}
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

