// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package iotthingtype

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IotThingTypeTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IotThingTypeTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IotThingTypeTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IotThingTypeTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IotThingTypeTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IotThingTypeTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IotThingTypeTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIotThingTypeTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

