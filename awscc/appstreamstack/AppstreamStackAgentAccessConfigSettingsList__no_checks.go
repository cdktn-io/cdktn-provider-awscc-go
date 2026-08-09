// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package appstreamstack

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppstreamStackAgentAccessConfigSettingsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AppstreamStackAgentAccessConfigSettingsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppstreamStackAgentAccessConfigSettingsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppstreamStackAgentAccessConfigSettingsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppstreamStackAgentAccessConfigSettingsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppstreamStackAgentAccessConfigSettingsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppstreamStackAgentAccessConfigSettingsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppstreamStackAgentAccessConfigSettingsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

