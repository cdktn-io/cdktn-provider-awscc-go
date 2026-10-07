// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cloudwatchview

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudwatchViewTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CloudwatchViewTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudwatchViewTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudwatchViewTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CloudwatchViewTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudwatchViewTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudwatchViewTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudwatchViewTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

