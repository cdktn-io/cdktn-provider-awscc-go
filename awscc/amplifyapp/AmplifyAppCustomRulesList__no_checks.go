// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package amplifyapp

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AmplifyAppCustomRulesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AmplifyAppCustomRulesList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AmplifyAppCustomRulesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AmplifyAppCustomRulesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AmplifyAppCustomRulesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AmplifyAppCustomRulesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AmplifyAppCustomRulesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAmplifyAppCustomRulesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

