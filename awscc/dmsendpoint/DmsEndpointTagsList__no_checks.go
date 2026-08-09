// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dmsendpoint

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DmsEndpointTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DmsEndpointTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DmsEndpointTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DmsEndpointTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DmsEndpointTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DmsEndpointTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DmsEndpointTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDmsEndpointTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

