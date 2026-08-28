// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package identitystoreuser

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IdentitystoreUserRolesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IdentitystoreUserRolesList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IdentitystoreUserRolesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IdentitystoreUserRolesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IdentitystoreUserRolesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IdentitystoreUserRolesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IdentitystoreUserRolesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIdentitystoreUserRolesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

