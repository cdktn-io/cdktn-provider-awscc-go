// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package amplifyapp

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AmplifyAppEnvironmentVariablesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AmplifyAppEnvironmentVariablesList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AmplifyAppEnvironmentVariablesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AmplifyAppEnvironmentVariablesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AmplifyAppEnvironmentVariablesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AmplifyAppEnvironmentVariablesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AmplifyAppEnvironmentVariablesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAmplifyAppEnvironmentVariablesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

