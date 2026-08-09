// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package lightsaildisk

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LightsailDiskTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LightsailDiskTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LightsailDiskTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LightsailDiskTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LightsailDiskTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LightsailDiskTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LightsailDiskTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLightsailDiskTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

