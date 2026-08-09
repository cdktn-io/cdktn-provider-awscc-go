// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package locationplaceindex

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LocationPlaceIndexTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LocationPlaceIndexTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LocationPlaceIndexTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LocationPlaceIndexTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LocationPlaceIndexTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LocationPlaceIndexTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LocationPlaceIndexTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLocationPlaceIndexTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

