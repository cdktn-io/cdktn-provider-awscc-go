// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package apigatewaymethod

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApigatewayMethodMethodResponsesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ApigatewayMethodMethodResponsesList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApigatewayMethodMethodResponsesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApigatewayMethodMethodResponsesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApigatewayMethodMethodResponsesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApigatewayMethodMethodResponsesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApigatewayMethodMethodResponsesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApigatewayMethodMethodResponsesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

