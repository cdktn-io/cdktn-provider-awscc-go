// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cognitouserpooldomain


type CognitoUserPoolDomainRouting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cognito_user_pool_domain#failover CognitoUserPoolDomain#failover}.
	Failover *CognitoUserPoolDomainRoutingFailover `field:"optional" json:"failover" yaml:"failover"`
}

