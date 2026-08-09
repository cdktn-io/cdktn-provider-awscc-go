// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package mskreplicator

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MskReplicatorTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_MskReplicatorTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MskReplicatorTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MskReplicatorTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MskReplicatorTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MskReplicatorTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MskReplicatorTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMskReplicatorTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

