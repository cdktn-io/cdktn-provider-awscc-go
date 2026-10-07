// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package lambdawebfunctionendpoint

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaWebFunctionEndpointRegionalEndpointsMap) validateGetParameters(key *string) error {
	return nil
}

func (l *jsiiProxy_LambdaWebFunctionEndpointRegionalEndpointsMap) validateInterpolationForAttributeParameters(property *string) error {
	return nil
}

func (l *jsiiProxy_LambdaWebFunctionEndpointRegionalEndpointsMap) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LambdaWebFunctionEndpointRegionalEndpointsMap) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LambdaWebFunctionEndpointRegionalEndpointsMap) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func validateNewLambdaWebFunctionEndpointRegionalEndpointsMapParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	return nil
}

