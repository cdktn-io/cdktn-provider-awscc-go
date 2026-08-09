// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package connecttestcase

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConnectTestCaseTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ConnectTestCaseTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ConnectTestCaseTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ConnectTestCaseTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ConnectTestCaseTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConnectTestCaseTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ConnectTestCaseTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewConnectTestCaseTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

