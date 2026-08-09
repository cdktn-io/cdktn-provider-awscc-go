// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package appconfigenvironment

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppconfigEnvironmentTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AppconfigEnvironmentTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppconfigEnvironmentTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppconfigEnvironmentTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppconfigEnvironmentTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppconfigEnvironmentTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppconfigEnvironmentTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppconfigEnvironmentTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

