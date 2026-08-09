// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package securityhubsecuritycontrol

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecurityhubSecurityControlParametersMap) validateGetParameters(key *string) error {
	return nil
}

func (s *jsiiProxy_SecurityhubSecurityControlParametersMap) validateInterpolationForAttributeParameters(property *string) error {
	return nil
}

func (s *jsiiProxy_SecurityhubSecurityControlParametersMap) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersMap) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersMap) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersMap) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func validateNewSecurityhubSecurityControlParametersMapParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	return nil
}

