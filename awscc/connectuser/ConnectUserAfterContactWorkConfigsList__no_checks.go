// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package connectuser

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConnectUserAfterContactWorkConfigsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ConnectUserAfterContactWorkConfigsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ConnectUserAfterContactWorkConfigsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ConnectUserAfterContactWorkConfigsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ConnectUserAfterContactWorkConfigsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConnectUserAfterContactWorkConfigsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ConnectUserAfterContactWorkConfigsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewConnectUserAfterContactWorkConfigsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

