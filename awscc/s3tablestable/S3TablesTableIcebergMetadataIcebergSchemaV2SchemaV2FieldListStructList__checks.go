// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package s3tablestable

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (s *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2SchemaV2FieldListStructList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2SchemaV2FieldListStructList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2SchemaV2FieldListStructList) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2SchemaV2FieldListStructList) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*S3TablesTableIcebergMetadataIcebergSchemaV2SchemaV2FieldListStruct:
		val := val.(*[]*S3TablesTableIcebergMetadataIcebergSchemaV2SchemaV2FieldListStruct)
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	case []*S3TablesTableIcebergMetadataIcebergSchemaV2SchemaV2FieldListStruct:
		val_ := val.([]*S3TablesTableIcebergMetadataIcebergSchemaV2SchemaV2FieldListStruct)
		val := &val_
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *[]*S3TablesTableIcebergMetadataIcebergSchemaV2SchemaV2FieldListStruct; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2SchemaV2FieldListStructList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2SchemaV2FieldListStructList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_S3TablesTableIcebergMetadataIcebergSchemaV2SchemaV2FieldListStructList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewS3TablesTableIcebergMetadataIcebergSchemaV2SchemaV2FieldListStructListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if wrapsSet == nil {
		return fmt.Errorf("parameter wrapsSet is required, but nil was provided")
	}

	return nil
}

