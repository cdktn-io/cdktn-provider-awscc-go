// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipespipe


type PipesPipeSourceParametersActiveMqBrokerParametersCredentials struct {
	// Optional SecretManager ARN which stores the database credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/pipes_pipe#basic_auth PipesPipe#basic_auth}
	BasicAuth *string `field:"optional" json:"basicAuth" yaml:"basicAuth"`
}

