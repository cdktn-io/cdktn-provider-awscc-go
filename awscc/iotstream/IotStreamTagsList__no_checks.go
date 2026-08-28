// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package iotstream

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IotStreamTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IotStreamTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IotStreamTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IotStreamTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IotStreamTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IotStreamTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IotStreamTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIotStreamTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

