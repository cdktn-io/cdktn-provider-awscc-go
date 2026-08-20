// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package connectmetric

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConnectMetricTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ConnectMetricTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ConnectMetricTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ConnectMetricTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ConnectMetricTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConnectMetricTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ConnectMetricTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewConnectMetricTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

