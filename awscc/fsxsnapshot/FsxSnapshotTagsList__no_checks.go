// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package fsxsnapshot

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FsxSnapshotTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FsxSnapshotTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FsxSnapshotTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FsxSnapshotTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FsxSnapshotTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FsxSnapshotTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FsxSnapshotTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFsxSnapshotTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

