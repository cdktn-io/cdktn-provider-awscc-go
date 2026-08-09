// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package connectworkspace

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConnectWorkspacePagesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ConnectWorkspacePagesList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ConnectWorkspacePagesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ConnectWorkspacePagesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ConnectWorkspacePagesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConnectWorkspacePagesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ConnectWorkspacePagesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewConnectWorkspacePagesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

