// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package wellarchitectedlens

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WellarchitectedLensTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WellarchitectedLensTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WellarchitectedLensTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WellarchitectedLensTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WellarchitectedLensTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WellarchitectedLensTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WellarchitectedLensTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWellarchitectedLensTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

