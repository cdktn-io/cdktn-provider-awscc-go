// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ecsgateway

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcsGatewayTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EcsGatewayTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EcsGatewayTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EcsGatewayTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EcsGatewayTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EcsGatewayTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EcsGatewayTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEcsGatewayTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

