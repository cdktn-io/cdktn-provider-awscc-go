// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package directconnectconnection

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DirectconnectConnectionTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DirectconnectConnectionTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DirectconnectConnectionTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DirectconnectConnectionTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DirectconnectConnectionTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DirectconnectConnectionTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DirectconnectConnectionTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDirectconnectConnectionTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

