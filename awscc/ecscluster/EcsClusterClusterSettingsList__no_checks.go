// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ecscluster

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcsClusterClusterSettingsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EcsClusterClusterSettingsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EcsClusterClusterSettingsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EcsClusterClusterSettingsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EcsClusterClusterSettingsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EcsClusterClusterSettingsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EcsClusterClusterSettingsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEcsClusterClusterSettingsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

