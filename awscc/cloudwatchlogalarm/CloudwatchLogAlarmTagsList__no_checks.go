// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cloudwatchlogalarm

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudwatchLogAlarmTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CloudwatchLogAlarmTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudwatchLogAlarmTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudwatchLogAlarmTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CloudwatchLogAlarmTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudwatchLogAlarmTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudwatchLogAlarmTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudwatchLogAlarmTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

