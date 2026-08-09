// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cloudfrontfunction

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudfrontFunctionTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CloudfrontFunctionTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudfrontFunctionTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudfrontFunctionTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CloudfrontFunctionTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudfrontFunctionTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudfrontFunctionTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudfrontFunctionTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

