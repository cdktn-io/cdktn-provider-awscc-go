// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ecscapacityprovider

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcsCapacityProviderTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EcsCapacityProviderTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EcsCapacityProviderTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EcsCapacityProviderTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EcsCapacityProviderTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EcsCapacityProviderTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EcsCapacityProviderTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEcsCapacityProviderTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

