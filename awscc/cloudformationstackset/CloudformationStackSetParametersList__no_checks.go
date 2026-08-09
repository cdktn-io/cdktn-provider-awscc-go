// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cloudformationstackset

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudformationStackSetParametersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CloudformationStackSetParametersList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudformationStackSetParametersList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudformationStackSetParametersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CloudformationStackSetParametersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudformationStackSetParametersList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudformationStackSetParametersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudformationStackSetParametersListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

