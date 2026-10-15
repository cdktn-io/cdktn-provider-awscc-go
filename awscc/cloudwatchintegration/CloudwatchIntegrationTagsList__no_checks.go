// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cloudwatchintegration

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudwatchIntegrationTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CloudwatchIntegrationTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudwatchIntegrationTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudwatchIntegrationTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CloudwatchIntegrationTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudwatchIntegrationTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudwatchIntegrationTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudwatchIntegrationTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

