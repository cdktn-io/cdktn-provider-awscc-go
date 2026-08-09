// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ssopermissionset

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SsoPermissionSetTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SsoPermissionSetTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SsoPermissionSetTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SsoPermissionSetTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SsoPermissionSetTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SsoPermissionSetTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SsoPermissionSetTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSsoPermissionSetTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

