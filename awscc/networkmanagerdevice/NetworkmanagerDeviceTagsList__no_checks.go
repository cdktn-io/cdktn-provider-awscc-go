// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package networkmanagerdevice

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkmanagerDeviceTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NetworkmanagerDeviceTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkmanagerDeviceTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkmanagerDeviceTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NetworkmanagerDeviceTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkmanagerDeviceTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkmanagerDeviceTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkmanagerDeviceTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

