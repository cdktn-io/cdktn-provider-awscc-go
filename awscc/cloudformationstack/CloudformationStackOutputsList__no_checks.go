// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cloudformationstack

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudformationStackOutputsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CloudformationStackOutputsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudformationStackOutputsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudformationStackOutputsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudformationStackOutputsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudformationStackOutputsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudformationStackOutputsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

