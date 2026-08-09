// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package appconfigextension

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppconfigExtensionActionsMap) validateGetParameters(key *string) error {
	return nil
}

func (a *jsiiProxy_AppconfigExtensionActionsMap) validateInterpolationForAttributeParameters(property *string) error {
	return nil
}

func (a *jsiiProxy_AppconfigExtensionActionsMap) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppconfigExtensionActionsMap) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppconfigExtensionActionsMap) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppconfigExtensionActionsMap) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func validateNewAppconfigExtensionActionsMapParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	return nil
}

