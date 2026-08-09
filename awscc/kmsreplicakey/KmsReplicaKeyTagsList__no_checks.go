// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package kmsreplicakey

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KmsReplicaKeyTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (k *jsiiProxy_KmsReplicaKeyTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_KmsReplicaKeyTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_KmsReplicaKeyTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KmsReplicaKeyTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KmsReplicaKeyTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_KmsReplicaKeyTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewKmsReplicaKeyTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

