// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cognitouserpooldomain


type CognitoUserPoolDomainRoutingFailover struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cognito_user_pool_domain#primary_route_53_health_check_id CognitoUserPoolDomain#primary_route_53_health_check_id}.
	PrimaryRoute53HealthCheckId *string `field:"optional" json:"primaryRoute53HealthCheckId" yaml:"primaryRoute53HealthCheckId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cognito_user_pool_domain#secondary_region CognitoUserPoolDomain#secondary_region}.
	SecondaryRegion *string `field:"optional" json:"secondaryRegion" yaml:"secondaryRegion"`
}

