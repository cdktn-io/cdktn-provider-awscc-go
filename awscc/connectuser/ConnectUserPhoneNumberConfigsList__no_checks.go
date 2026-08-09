// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package connectuser

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConnectUserPhoneNumberConfigsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ConnectUserPhoneNumberConfigsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ConnectUserPhoneNumberConfigsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ConnectUserPhoneNumberConfigsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ConnectUserPhoneNumberConfigsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConnectUserPhoneNumberConfigsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ConnectUserPhoneNumberConfigsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewConnectUserPhoneNumberConfigsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

