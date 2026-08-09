// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package rumappmonitor

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RumAppMonitorTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RumAppMonitorTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RumAppMonitorTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RumAppMonitorTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RumAppMonitorTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RumAppMonitorTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RumAppMonitorTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRumAppMonitorTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

