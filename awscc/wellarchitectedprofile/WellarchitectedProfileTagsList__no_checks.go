// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package wellarchitectedprofile

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WellarchitectedProfileTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WellarchitectedProfileTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WellarchitectedProfileTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WellarchitectedProfileTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WellarchitectedProfileTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WellarchitectedProfileTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WellarchitectedProfileTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWellarchitectedProfileTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

