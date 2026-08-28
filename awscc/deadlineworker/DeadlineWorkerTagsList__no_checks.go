// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package deadlineworker

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DeadlineWorkerTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DeadlineWorkerTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DeadlineWorkerTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DeadlineWorkerTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DeadlineWorkerTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DeadlineWorkerTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DeadlineWorkerTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDeadlineWorkerTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

