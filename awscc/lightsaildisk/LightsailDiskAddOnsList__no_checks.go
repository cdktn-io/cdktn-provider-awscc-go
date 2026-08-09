// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package lightsaildisk

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LightsailDiskAddOnsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LightsailDiskAddOnsList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LightsailDiskAddOnsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LightsailDiskAddOnsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LightsailDiskAddOnsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LightsailDiskAddOnsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LightsailDiskAddOnsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLightsailDiskAddOnsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

