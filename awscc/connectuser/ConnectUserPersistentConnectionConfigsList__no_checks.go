// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package connectuser

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConnectUserPersistentConnectionConfigsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ConnectUserPersistentConnectionConfigsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ConnectUserPersistentConnectionConfigsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ConnectUserPersistentConnectionConfigsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ConnectUserPersistentConnectionConfigsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConnectUserPersistentConnectionConfigsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ConnectUserPersistentConnectionConfigsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewConnectUserPersistentConnectionConfigsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

