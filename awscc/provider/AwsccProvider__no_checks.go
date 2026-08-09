// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsccProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (a *jsiiProxy_AwsccProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (a *jsiiProxy_AwsccProvider) validateRegisterProviderFeatureUsageParameters(feature cdktn.ProviderFeature) error {
	return nil
}

func validateAwsccProvider_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateAwsccProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAwsccProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateAwsccProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsccProvider) validateSetAssumeRoleParameters(val *AwsccProviderAssumeRole) error {
	return nil
}

func (j *jsiiProxy_AwsccProvider) validateSetAssumeRoleWithWebIdentityParameters(val *AwsccProviderAssumeRoleWithWebIdentity) error {
	return nil
}

func (j *jsiiProxy_AwsccProvider) validateSetEndpointsParameters(val *AwsccProviderEndpoints) error {
	return nil
}

func (j *jsiiProxy_AwsccProvider) validateSetInsecureParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsccProvider) validateSetSkipMedatadataApiCheckParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsccProvider) validateSetSkipMetadataApiCheckParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsccProvider) validateSetUserAgentParameters(val interface{}) error {
	return nil
}

func validateNewAwsccProviderParameters(scope constructs.Construct, id *string, config *AwsccProviderConfig) error {
	return nil
}

