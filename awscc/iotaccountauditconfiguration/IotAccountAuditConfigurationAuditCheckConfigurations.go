// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotaccountauditconfiguration


type IotAccountAuditConfigurationAuditCheckConfigurations struct {
	// The configuration for a specific audit check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_account_audit_configuration#authenticated_cognito_role_overly_permissive_check IotAccountAuditConfiguration#authenticated_cognito_role_overly_permissive_check}
	AuthenticatedCognitoRoleOverlyPermissiveCheck *IotAccountAuditConfigurationAuditCheckConfigurationsAuthenticatedCognitoRoleOverlyPermissiveCheck `field:"optional" json:"authenticatedCognitoRoleOverlyPermissiveCheck" yaml:"authenticatedCognitoRoleOverlyPermissiveCheck"`
	// The configuration for a specific audit check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_account_audit_configuration#ca_certificate_expiring_check IotAccountAuditConfiguration#ca_certificate_expiring_check}
	CaCertificateExpiringCheck *IotAccountAuditConfigurationAuditCheckConfigurationsCaCertificateExpiringCheck `field:"optional" json:"caCertificateExpiringCheck" yaml:"caCertificateExpiringCheck"`
	// The configuration for a specific audit check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_account_audit_configuration#ca_certificate_key_quality_check IotAccountAuditConfiguration#ca_certificate_key_quality_check}
	CaCertificateKeyQualityCheck *IotAccountAuditConfigurationAuditCheckConfigurationsCaCertificateKeyQualityCheck `field:"optional" json:"caCertificateKeyQualityCheck" yaml:"caCertificateKeyQualityCheck"`
	// The configuration for a specific audit check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_account_audit_configuration#conflicting_client_ids_check IotAccountAuditConfiguration#conflicting_client_ids_check}
	ConflictingClientIdsCheck *IotAccountAuditConfigurationAuditCheckConfigurationsConflictingClientIdsCheck `field:"optional" json:"conflictingClientIdsCheck" yaml:"conflictingClientIdsCheck"`
	// A structure containing the configName and corresponding configValue for configuring DeviceCertAgeCheck.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_account_audit_configuration#device_certificate_age_check IotAccountAuditConfiguration#device_certificate_age_check}
	DeviceCertificateAgeCheck *IotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateAgeCheck `field:"optional" json:"deviceCertificateAgeCheck" yaml:"deviceCertificateAgeCheck"`
	// A structure containing the configName and corresponding configValue for configuring DeviceCertExpirationCheck.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_account_audit_configuration#device_certificate_expiring_check IotAccountAuditConfiguration#device_certificate_expiring_check}
	DeviceCertificateExpiringCheck *IotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateExpiringCheck `field:"optional" json:"deviceCertificateExpiringCheck" yaml:"deviceCertificateExpiringCheck"`
	// The configuration for a specific audit check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_account_audit_configuration#device_certificate_key_quality_check IotAccountAuditConfiguration#device_certificate_key_quality_check}
	DeviceCertificateKeyQualityCheck *IotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateKeyQualityCheck `field:"optional" json:"deviceCertificateKeyQualityCheck" yaml:"deviceCertificateKeyQualityCheck"`
	// The configuration for a specific audit check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_account_audit_configuration#device_certificate_shared_check IotAccountAuditConfiguration#device_certificate_shared_check}
	DeviceCertificateSharedCheck *IotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateSharedCheck `field:"optional" json:"deviceCertificateSharedCheck" yaml:"deviceCertificateSharedCheck"`
	// The configuration for a specific audit check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_account_audit_configuration#intermediate_ca_revoked_for_active_device_certificates_check IotAccountAuditConfiguration#intermediate_ca_revoked_for_active_device_certificates_check}
	IntermediateCaRevokedForActiveDeviceCertificatesCheck *IotAccountAuditConfigurationAuditCheckConfigurationsIntermediateCaRevokedForActiveDeviceCertificatesCheck `field:"optional" json:"intermediateCaRevokedForActiveDeviceCertificatesCheck" yaml:"intermediateCaRevokedForActiveDeviceCertificatesCheck"`
	// The configuration for a specific audit check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_account_audit_configuration#iot_policy_overly_permissive_check IotAccountAuditConfiguration#iot_policy_overly_permissive_check}
	IotPolicyOverlyPermissiveCheck *IotAccountAuditConfigurationAuditCheckConfigurationsIotPolicyOverlyPermissiveCheck `field:"optional" json:"iotPolicyOverlyPermissiveCheck" yaml:"iotPolicyOverlyPermissiveCheck"`
	// The configuration for a specific audit check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_account_audit_configuration#io_t_policy_potential_mis_configuration_check IotAccountAuditConfiguration#io_t_policy_potential_mis_configuration_check}
	IoTPolicyPotentialMisConfigurationCheck *IotAccountAuditConfigurationAuditCheckConfigurationsIoTPolicyPotentialMisConfigurationCheck `field:"optional" json:"ioTPolicyPotentialMisConfigurationCheck" yaml:"ioTPolicyPotentialMisConfigurationCheck"`
	// The configuration for a specific audit check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_account_audit_configuration#iot_role_alias_allows_access_to_unused_services_check IotAccountAuditConfiguration#iot_role_alias_allows_access_to_unused_services_check}
	IotRoleAliasAllowsAccessToUnusedServicesCheck *IotAccountAuditConfigurationAuditCheckConfigurationsIotRoleAliasAllowsAccessToUnusedServicesCheck `field:"optional" json:"iotRoleAliasAllowsAccessToUnusedServicesCheck" yaml:"iotRoleAliasAllowsAccessToUnusedServicesCheck"`
	// The configuration for a specific audit check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_account_audit_configuration#iot_role_alias_overly_permissive_check IotAccountAuditConfiguration#iot_role_alias_overly_permissive_check}
	IotRoleAliasOverlyPermissiveCheck *IotAccountAuditConfigurationAuditCheckConfigurationsIotRoleAliasOverlyPermissiveCheck `field:"optional" json:"iotRoleAliasOverlyPermissiveCheck" yaml:"iotRoleAliasOverlyPermissiveCheck"`
	// The configuration for a specific audit check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_account_audit_configuration#logging_disabled_check IotAccountAuditConfiguration#logging_disabled_check}
	LoggingDisabledCheck *IotAccountAuditConfigurationAuditCheckConfigurationsLoggingDisabledCheck `field:"optional" json:"loggingDisabledCheck" yaml:"loggingDisabledCheck"`
	// The configuration for a specific audit check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_account_audit_configuration#revoked_ca_certificate_still_active_check IotAccountAuditConfiguration#revoked_ca_certificate_still_active_check}
	RevokedCaCertificateStillActiveCheck *IotAccountAuditConfigurationAuditCheckConfigurationsRevokedCaCertificateStillActiveCheck `field:"optional" json:"revokedCaCertificateStillActiveCheck" yaml:"revokedCaCertificateStillActiveCheck"`
	// The configuration for a specific audit check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_account_audit_configuration#revoked_device_certificate_still_active_check IotAccountAuditConfiguration#revoked_device_certificate_still_active_check}
	RevokedDeviceCertificateStillActiveCheck *IotAccountAuditConfigurationAuditCheckConfigurationsRevokedDeviceCertificateStillActiveCheck `field:"optional" json:"revokedDeviceCertificateStillActiveCheck" yaml:"revokedDeviceCertificateStillActiveCheck"`
	// The configuration for a specific audit check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_account_audit_configuration#unauthenticated_cognito_role_overly_permissive_check IotAccountAuditConfiguration#unauthenticated_cognito_role_overly_permissive_check}
	UnauthenticatedCognitoRoleOverlyPermissiveCheck *IotAccountAuditConfigurationAuditCheckConfigurationsUnauthenticatedCognitoRoleOverlyPermissiveCheck `field:"optional" json:"unauthenticatedCognitoRoleOverlyPermissiveCheck" yaml:"unauthenticatedCognitoRoleOverlyPermissiveCheck"`
}

