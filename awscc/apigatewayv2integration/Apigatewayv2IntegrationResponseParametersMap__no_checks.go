// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package apigatewayv2integration

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_Apigatewayv2IntegrationResponseParametersMap) validateGetParameters(key *string) error {
	return nil
}

func (a *jsiiProxy_Apigatewayv2IntegrationResponseParametersMap) validateInterpolationForAttributeParameters(property *string) error {
	return nil
}

func (a *jsiiProxy_Apigatewayv2IntegrationResponseParametersMap) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponseParametersMap) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponseParametersMap) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Apigatewayv2IntegrationResponseParametersMap) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func validateNewApigatewayv2IntegrationResponseParametersMapParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	return nil
}

