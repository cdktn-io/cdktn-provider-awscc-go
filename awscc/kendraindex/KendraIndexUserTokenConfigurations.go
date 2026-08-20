// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendraindex


type KendraIndexUserTokenConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kendra_index#json_token_type_configuration KendraIndex#json_token_type_configuration}.
	JsonTokenTypeConfiguration *KendraIndexUserTokenConfigurationsJsonTokenTypeConfiguration `field:"optional" json:"jsonTokenTypeConfiguration" yaml:"jsonTokenTypeConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kendra_index#jwt_token_type_configuration KendraIndex#jwt_token_type_configuration}.
	JwtTokenTypeConfiguration *KendraIndexUserTokenConfigurationsJwtTokenTypeConfiguration `field:"optional" json:"jwtTokenTypeConfiguration" yaml:"jwtTokenTypeConfiguration"`
}

