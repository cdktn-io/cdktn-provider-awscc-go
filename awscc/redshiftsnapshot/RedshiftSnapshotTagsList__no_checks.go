// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package redshiftsnapshot

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RedshiftSnapshotTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RedshiftSnapshotTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RedshiftSnapshotTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RedshiftSnapshotTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RedshiftSnapshotTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RedshiftSnapshotTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RedshiftSnapshotTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRedshiftSnapshotTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

