// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package memorydbsnapshot

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MemorydbSnapshotTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_MemorydbSnapshotTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MemorydbSnapshotTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MemorydbSnapshotTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MemorydbSnapshotTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MemorydbSnapshotTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MemorydbSnapshotTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMemorydbSnapshotTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

