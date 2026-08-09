// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package lambdafunction

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaFunctionTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LambdaFunctionTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LambdaFunctionTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LambdaFunctionTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LambdaFunctionTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LambdaFunctionTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LambdaFunctionTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLambdaFunctionTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

