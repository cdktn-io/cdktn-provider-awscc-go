// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package networkmanagerlink

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkmanagerLinkTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NetworkmanagerLinkTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkmanagerLinkTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkmanagerLinkTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NetworkmanagerLinkTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkmanagerLinkTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkmanagerLinkTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkmanagerLinkTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

