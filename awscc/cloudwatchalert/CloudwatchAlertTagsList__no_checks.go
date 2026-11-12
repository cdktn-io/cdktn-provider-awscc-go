// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cloudwatchalert

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudwatchAlertTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CloudwatchAlertTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudwatchAlertTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudwatchAlertTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CloudwatchAlertTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudwatchAlertTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudwatchAlertTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudwatchAlertTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

