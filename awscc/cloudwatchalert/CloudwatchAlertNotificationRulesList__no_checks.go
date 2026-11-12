// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cloudwatchalert

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudwatchAlertNotificationRulesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CloudwatchAlertNotificationRulesList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudwatchAlertNotificationRulesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudwatchAlertNotificationRulesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CloudwatchAlertNotificationRulesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudwatchAlertNotificationRulesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudwatchAlertNotificationRulesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudwatchAlertNotificationRulesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

