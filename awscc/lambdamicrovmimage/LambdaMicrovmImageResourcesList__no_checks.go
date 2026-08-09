// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package lambdamicrovmimage

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaMicrovmImageResourcesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LambdaMicrovmImageResourcesList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LambdaMicrovmImageResourcesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LambdaMicrovmImageResourcesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LambdaMicrovmImageResourcesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LambdaMicrovmImageResourcesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LambdaMicrovmImageResourcesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLambdaMicrovmImageResourcesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

