// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package appconfigextension

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppconfigExtensionParametersMap) validateGetParameters(key *string) error {
	return nil
}

func (a *jsiiProxy_AppconfigExtensionParametersMap) validateInterpolationForAttributeParameters(property *string) error {
	return nil
}

func (a *jsiiProxy_AppconfigExtensionParametersMap) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppconfigExtensionParametersMap) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppconfigExtensionParametersMap) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppconfigExtensionParametersMap) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func validateNewAppconfigExtensionParametersMapParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	return nil
}

