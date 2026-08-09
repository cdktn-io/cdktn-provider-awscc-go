// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package configconformancepack

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConfigConformancePackTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ConfigConformancePackTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ConfigConformancePackTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ConfigConformancePackTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ConfigConformancePackTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConfigConformancePackTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ConfigConformancePackTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewConfigConformancePackTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

