// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package iotauthorizer

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IotAuthorizerTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IotAuthorizerTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IotAuthorizerTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IotAuthorizerTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IotAuthorizerTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IotAuthorizerTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IotAuthorizerTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIotAuthorizerTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

