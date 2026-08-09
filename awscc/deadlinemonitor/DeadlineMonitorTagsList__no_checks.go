// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package deadlinemonitor

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DeadlineMonitorTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DeadlineMonitorTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DeadlineMonitorTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DeadlineMonitorTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DeadlineMonitorTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DeadlineMonitorTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DeadlineMonitorTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDeadlineMonitorTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

