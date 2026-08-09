// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package transferconnector

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TransferConnectorTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TransferConnectorTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TransferConnectorTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TransferConnectorTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TransferConnectorTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TransferConnectorTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TransferConnectorTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTransferConnectorTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

