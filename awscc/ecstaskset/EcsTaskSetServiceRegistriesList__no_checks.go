// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ecstaskset

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcsTaskSetServiceRegistriesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EcsTaskSetServiceRegistriesList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EcsTaskSetServiceRegistriesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EcsTaskSetServiceRegistriesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EcsTaskSetServiceRegistriesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EcsTaskSetServiceRegistriesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EcsTaskSetServiceRegistriesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEcsTaskSetServiceRegistriesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

