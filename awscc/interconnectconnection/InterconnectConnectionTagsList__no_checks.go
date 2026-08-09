// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package interconnectconnection

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InterconnectConnectionTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_InterconnectConnectionTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_InterconnectConnectionTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_InterconnectConnectionTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_InterconnectConnectionTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_InterconnectConnectionTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_InterconnectConnectionTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewInterconnectConnectionTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

