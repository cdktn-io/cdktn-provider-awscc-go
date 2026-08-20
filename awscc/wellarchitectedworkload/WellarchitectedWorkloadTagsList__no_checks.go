// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package wellarchitectedworkload

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WellarchitectedWorkloadTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WellarchitectedWorkloadTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WellarchitectedWorkloadTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WellarchitectedWorkloadTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WellarchitectedWorkloadTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WellarchitectedWorkloadTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WellarchitectedWorkloadTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWellarchitectedWorkloadTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

