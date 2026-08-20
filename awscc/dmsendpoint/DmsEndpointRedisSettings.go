// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsendpoint


type DmsEndpointRedisSettings struct {
	// The password provided with the auth-role and auth-token options of the AuthType setting for a Redis target endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dms_endpoint#auth_password DmsEndpoint#auth_password}
	AuthPassword *string `field:"optional" json:"authPassword" yaml:"authPassword"`
	// The type of authentication to perform when connecting to a Redis target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dms_endpoint#auth_type DmsEndpoint#auth_type}
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// The user name provided with the auth-role option of the AuthType setting for a Redis target endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dms_endpoint#auth_user_name DmsEndpoint#auth_user_name}
	AuthUserName *string `field:"optional" json:"authUserName" yaml:"authUserName"`
	// Transmission Control Protocol (TCP) port for the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dms_endpoint#port DmsEndpoint#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Fully qualified domain name of the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dms_endpoint#server_name DmsEndpoint#server_name}
	ServerName *string `field:"optional" json:"serverName" yaml:"serverName"`
	// The Amazon Resource Name (ARN) for the certificate authority (CA) that DMS uses to connect to your Redis target endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dms_endpoint#ssl_ca_certificate_arn DmsEndpoint#ssl_ca_certificate_arn}
	SslCaCertificateArn *string `field:"optional" json:"sslCaCertificateArn" yaml:"sslCaCertificateArn"`
	// The connection to a Redis target endpoint using Transport Layer Security (TLS). Valid values include plaintext and ssl-encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dms_endpoint#ssl_security_protocol DmsEndpoint#ssl_security_protocol}
	SslSecurityProtocol *string `field:"optional" json:"sslSecurityProtocol" yaml:"sslSecurityProtocol"`
}

