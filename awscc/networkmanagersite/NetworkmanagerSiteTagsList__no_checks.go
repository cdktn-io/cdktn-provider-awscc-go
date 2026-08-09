// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package networkmanagersite

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkmanagerSiteTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NetworkmanagerSiteTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkmanagerSiteTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkmanagerSiteTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NetworkmanagerSiteTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkmanagerSiteTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkmanagerSiteTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkmanagerSiteTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

