// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotdomainconfiguration


type IotDomainConfigurationServerCertificateConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_domain_configuration#enable_ocsp_check IotDomainConfiguration#enable_ocsp_check}.
	EnableOcspCheck interface{} `field:"optional" json:"enableOcspCheck" yaml:"enableOcspCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_domain_configuration#ocsp_authorized_responder_arn IotDomainConfiguration#ocsp_authorized_responder_arn}.
	OcspAuthorizedResponderArn *string `field:"optional" json:"ocspAuthorizedResponderArn" yaml:"ocspAuthorizedResponderArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_domain_configuration#ocsp_lambda_arn IotDomainConfiguration#ocsp_lambda_arn}.
	OcspLambdaArn *string `field:"optional" json:"ocspLambdaArn" yaml:"ocspLambdaArn"`
}

