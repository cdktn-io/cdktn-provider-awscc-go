// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package transferworkflow

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TransferWorkflowTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TransferWorkflowTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TransferWorkflowTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TransferWorkflowTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TransferWorkflowTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TransferWorkflowTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TransferWorkflowTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTransferWorkflowTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

