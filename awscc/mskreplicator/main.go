// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskreplicator

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mskReplicator.MskReplicator",
		reflect.TypeOf((*MskReplicator)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "currentVersion", GoGetter: "CurrentVersion"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "kafkaClusters", GoGetter: "KafkaClusters"},
			_jsii_.MemberProperty{JsiiProperty: "kafkaClustersInput", GoGetter: "KafkaClustersInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logDeliveryInput", GoGetter: "LogDeliveryInput"},
			_jsii_.MemberMethod{JsiiMethod: "markWriteOnlyAttribute", GoMethod: "MarkWriteOnlyAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putKafkaClusters", GoMethod: "PutKafkaClusters"},
			_jsii_.MemberMethod{JsiiMethod: "putLogDelivery", GoMethod: "PutLogDelivery"},
			_jsii_.MemberMethod{JsiiMethod: "putReplicationInfoList", GoMethod: "PutReplicationInfoList"},
			_jsii_.MemberMethod{JsiiMethod: "putTags", GoMethod: "PutTags"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "registerProviderFeatureUsage", GoMethod: "RegisterProviderFeatureUsage"},
			_jsii_.MemberProperty{JsiiProperty: "replicationInfoList", GoGetter: "ReplicationInfoList"},
			_jsii_.MemberProperty{JsiiProperty: "replicationInfoListInput", GoGetter: "ReplicationInfoListInput"},
			_jsii_.MemberProperty{JsiiProperty: "replicatorArn", GoGetter: "ReplicatorArn"},
			_jsii_.MemberProperty{JsiiProperty: "replicatorName", GoGetter: "ReplicatorName"},
			_jsii_.MemberProperty{JsiiProperty: "replicatorNameInput", GoGetter: "ReplicatorNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogDelivery", GoMethod: "ResetLogDelivery"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberProperty{JsiiProperty: "serviceExecutionRoleArn", GoGetter: "ServiceExecutionRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "serviceExecutionRoleArnInput", GoGetter: "ServiceExecutionRoleArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_MskReplicator{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorConfig",
		reflect.TypeOf((*MskReplicatorConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorKafkaClusters",
		reflect.TypeOf((*MskReplicatorKafkaClusters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorKafkaClustersAmazonMskCluster",
		reflect.TypeOf((*MskReplicatorKafkaClustersAmazonMskCluster)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorKafkaClustersAmazonMskClusterOutputReference",
		reflect.TypeOf((*MskReplicatorKafkaClustersAmazonMskClusterOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "mskClusterArn", GoGetter: "MskClusterArn"},
			_jsii_.MemberProperty{JsiiProperty: "mskClusterArnInput", GoGetter: "MskClusterArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMskClusterArn", GoMethod: "ResetMskClusterArn"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MskReplicatorKafkaClustersAmazonMskClusterOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorKafkaClustersApacheKafkaCluster",
		reflect.TypeOf((*MskReplicatorKafkaClustersApacheKafkaCluster)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorKafkaClustersApacheKafkaClusterOutputReference",
		reflect.TypeOf((*MskReplicatorKafkaClustersApacheKafkaClusterOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apacheKafkaClusterId", GoGetter: "ApacheKafkaClusterId"},
			_jsii_.MemberProperty{JsiiProperty: "apacheKafkaClusterIdInput", GoGetter: "ApacheKafkaClusterIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "bootstrapBrokerString", GoGetter: "BootstrapBrokerString"},
			_jsii_.MemberProperty{JsiiProperty: "bootstrapBrokerStringInput", GoGetter: "BootstrapBrokerStringInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetApacheKafkaClusterId", GoMethod: "ResetApacheKafkaClusterId"},
			_jsii_.MemberMethod{JsiiMethod: "resetBootstrapBrokerString", GoMethod: "ResetBootstrapBrokerString"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MskReplicatorKafkaClustersApacheKafkaClusterOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorKafkaClustersClientAuthentication",
		reflect.TypeOf((*MskReplicatorKafkaClustersClientAuthentication)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorKafkaClustersClientAuthenticationMtls",
		reflect.TypeOf((*MskReplicatorKafkaClustersClientAuthenticationMtls)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorKafkaClustersClientAuthenticationMtlsOutputReference",
		reflect.TypeOf((*MskReplicatorKafkaClustersClientAuthenticationMtlsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecretArn", GoMethod: "ResetSecretArn"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "secretArn", GoGetter: "SecretArn"},
			_jsii_.MemberProperty{JsiiProperty: "secretArnInput", GoGetter: "SecretArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationMtlsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorKafkaClustersClientAuthenticationOutputReference",
		reflect.TypeOf((*MskReplicatorKafkaClustersClientAuthenticationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "mtls", GoGetter: "Mtls"},
			_jsii_.MemberProperty{JsiiProperty: "mtlsInput", GoGetter: "MtlsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putMtls", GoMethod: "PutMtls"},
			_jsii_.MemberMethod{JsiiMethod: "putSaslOAuthBearer", GoMethod: "PutSaslOAuthBearer"},
			_jsii_.MemberMethod{JsiiMethod: "putSaslScram", GoMethod: "PutSaslScram"},
			_jsii_.MemberMethod{JsiiMethod: "resetMtls", GoMethod: "ResetMtls"},
			_jsii_.MemberMethod{JsiiMethod: "resetSaslOAuthBearer", GoMethod: "ResetSaslOAuthBearer"},
			_jsii_.MemberMethod{JsiiMethod: "resetSaslScram", GoMethod: "ResetSaslScram"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "saslOAuthBearer", GoGetter: "SaslOAuthBearer"},
			_jsii_.MemberProperty{JsiiProperty: "saslOAuthBearerInput", GoGetter: "SaslOAuthBearerInput"},
			_jsii_.MemberProperty{JsiiProperty: "saslScram", GoGetter: "SaslScram"},
			_jsii_.MemberProperty{JsiiProperty: "saslScramInput", GoGetter: "SaslScramInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearer",
		reflect.TypeOf((*MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerClientCredentials",
		reflect.TypeOf((*MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerClientCredentials)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerClientCredentialsAssertion",
		reflect.TypeOf((*MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerClientCredentialsAssertion)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerClientCredentialsAssertionOutputReference",
		reflect.TypeOf((*MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerClientCredentialsAssertionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "audience", GoGetter: "Audience"},
			_jsii_.MemberProperty{JsiiProperty: "audienceInput", GoGetter: "AudienceInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetAudience", GoMethod: "ResetAudience"},
			_jsii_.MemberMethod{JsiiMethod: "resetSigningAlgorithm", GoMethod: "ResetSigningAlgorithm"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenRequestSecretArn", GoMethod: "ResetTokenRequestSecretArn"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "signingAlgorithm", GoGetter: "SigningAlgorithm"},
			_jsii_.MemberProperty{JsiiProperty: "signingAlgorithmInput", GoGetter: "SigningAlgorithmInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "tokenRequestSecretArn", GoGetter: "TokenRequestSecretArn"},
			_jsii_.MemberProperty{JsiiProperty: "tokenRequestSecretArnInput", GoGetter: "TokenRequestSecretArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerClientCredentialsAssertionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerClientCredentialsOutputReference",
		reflect.TypeOf((*MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerClientCredentialsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenRequestSecretArn", GoMethod: "ResetTokenRequestSecretArn"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "tokenRequestSecretArn", GoGetter: "TokenRequestSecretArn"},
			_jsii_.MemberProperty{JsiiProperty: "tokenRequestSecretArnInput", GoGetter: "TokenRequestSecretArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerClientCredentialsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerIamJwtBearer",
		reflect.TypeOf((*MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerIamJwtBearer)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerIamJwtBearerOutputReference",
		reflect.TypeOf((*MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerIamJwtBearerOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "audience", GoGetter: "Audience"},
			_jsii_.MemberProperty{JsiiProperty: "audienceInput", GoGetter: "AudienceInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetAudience", GoMethod: "ResetAudience"},
			_jsii_.MemberMethod{JsiiMethod: "resetSigningAlgorithm", GoMethod: "ResetSigningAlgorithm"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenRequestSecretArn", GoMethod: "ResetTokenRequestSecretArn"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "signingAlgorithm", GoGetter: "SigningAlgorithm"},
			_jsii_.MemberProperty{JsiiProperty: "signingAlgorithmInput", GoGetter: "SigningAlgorithmInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "tokenRequestSecretArn", GoGetter: "TokenRequestSecretArn"},
			_jsii_.MemberProperty{JsiiProperty: "tokenRequestSecretArnInput", GoGetter: "TokenRequestSecretArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerIamJwtBearerOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference",
		reflect.TypeOf((*MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "clientCredentials", GoGetter: "ClientCredentials"},
			_jsii_.MemberProperty{JsiiProperty: "clientCredentialsAssertion", GoGetter: "ClientCredentialsAssertion"},
			_jsii_.MemberProperty{JsiiProperty: "clientCredentialsAssertionInput", GoGetter: "ClientCredentialsAssertionInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientCredentialsInput", GoGetter: "ClientCredentialsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "iamJwtBearer", GoGetter: "IamJwtBearer"},
			_jsii_.MemberProperty{JsiiProperty: "iamJwtBearerInput", GoGetter: "IamJwtBearerInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putClientCredentials", GoMethod: "PutClientCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "putClientCredentialsAssertion", GoMethod: "PutClientCredentialsAssertion"},
			_jsii_.MemberMethod{JsiiMethod: "putIamJwtBearer", GoMethod: "PutIamJwtBearer"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientCredentials", GoMethod: "ResetClientCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientCredentialsAssertion", GoMethod: "ResetClientCredentialsAssertion"},
			_jsii_.MemberMethod{JsiiMethod: "resetIamJwtBearer", GoMethod: "ResetIamJwtBearer"},
			_jsii_.MemberMethod{JsiiMethod: "resetScope", GoMethod: "ResetScope"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenEndpointAuthenticationMethod", GoMethod: "ResetTokenEndpointAuthenticationMethod"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenEndpointTlsCertificateArn", GoMethod: "ResetTokenEndpointTlsCertificateArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenEndpointUrl", GoMethod: "ResetTokenEndpointUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scope", GoGetter: "Scope"},
			_jsii_.MemberProperty{JsiiProperty: "scopeInput", GoGetter: "ScopeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "tokenEndpointAuthenticationMethod", GoGetter: "TokenEndpointAuthenticationMethod"},
			_jsii_.MemberProperty{JsiiProperty: "tokenEndpointAuthenticationMethodInput", GoGetter: "TokenEndpointAuthenticationMethodInput"},
			_jsii_.MemberProperty{JsiiProperty: "tokenEndpointTlsCertificateArn", GoGetter: "TokenEndpointTlsCertificateArn"},
			_jsii_.MemberProperty{JsiiProperty: "tokenEndpointTlsCertificateArnInput", GoGetter: "TokenEndpointTlsCertificateArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "tokenEndpointUrl", GoGetter: "TokenEndpointUrl"},
			_jsii_.MemberProperty{JsiiProperty: "tokenEndpointUrlInput", GoGetter: "TokenEndpointUrlInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorKafkaClustersClientAuthenticationSaslScram",
		reflect.TypeOf((*MskReplicatorKafkaClustersClientAuthenticationSaslScram)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorKafkaClustersClientAuthenticationSaslScramOutputReference",
		reflect.TypeOf((*MskReplicatorKafkaClustersClientAuthenticationSaslScramOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "mechanism", GoGetter: "Mechanism"},
			_jsii_.MemberProperty{JsiiProperty: "mechanismInput", GoGetter: "MechanismInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMechanism", GoMethod: "ResetMechanism"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecretArn", GoMethod: "ResetSecretArn"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "secretArn", GoGetter: "SecretArn"},
			_jsii_.MemberProperty{JsiiProperty: "secretArnInput", GoGetter: "SecretArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslScramOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorKafkaClustersEncryptionInTransit",
		reflect.TypeOf((*MskReplicatorKafkaClustersEncryptionInTransit)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorKafkaClustersEncryptionInTransitOutputReference",
		reflect.TypeOf((*MskReplicatorKafkaClustersEncryptionInTransitOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionType", GoGetter: "EncryptionType"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionTypeInput", GoGetter: "EncryptionTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetEncryptionType", GoMethod: "ResetEncryptionType"},
			_jsii_.MemberMethod{JsiiMethod: "resetRootCaCertificate", GoMethod: "ResetRootCaCertificate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "rootCaCertificate", GoGetter: "RootCaCertificate"},
			_jsii_.MemberProperty{JsiiProperty: "rootCaCertificateInput", GoGetter: "RootCaCertificateInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MskReplicatorKafkaClustersEncryptionInTransitOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorKafkaClustersList",
		reflect.TypeOf((*MskReplicatorKafkaClustersList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_MskReplicatorKafkaClustersList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorKafkaClustersOutputReference",
		reflect.TypeOf((*MskReplicatorKafkaClustersOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "amazonMskCluster", GoGetter: "AmazonMskCluster"},
			_jsii_.MemberProperty{JsiiProperty: "amazonMskClusterInput", GoGetter: "AmazonMskClusterInput"},
			_jsii_.MemberProperty{JsiiProperty: "apacheKafkaCluster", GoGetter: "ApacheKafkaCluster"},
			_jsii_.MemberProperty{JsiiProperty: "apacheKafkaClusterInput", GoGetter: "ApacheKafkaClusterInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientAuthentication", GoGetter: "ClientAuthentication"},
			_jsii_.MemberProperty{JsiiProperty: "clientAuthenticationInput", GoGetter: "ClientAuthenticationInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionInTransit", GoGetter: "EncryptionInTransit"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionInTransitInput", GoGetter: "EncryptionInTransitInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putAmazonMskCluster", GoMethod: "PutAmazonMskCluster"},
			_jsii_.MemberMethod{JsiiMethod: "putApacheKafkaCluster", GoMethod: "PutApacheKafkaCluster"},
			_jsii_.MemberMethod{JsiiMethod: "putClientAuthentication", GoMethod: "PutClientAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "putEncryptionInTransit", GoMethod: "PutEncryptionInTransit"},
			_jsii_.MemberMethod{JsiiMethod: "putVpcConfig", GoMethod: "PutVpcConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetAmazonMskCluster", GoMethod: "ResetAmazonMskCluster"},
			_jsii_.MemberMethod{JsiiMethod: "resetApacheKafkaCluster", GoMethod: "ResetApacheKafkaCluster"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientAuthentication", GoMethod: "ResetClientAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "resetEncryptionInTransit", GoMethod: "ResetEncryptionInTransit"},
			_jsii_.MemberMethod{JsiiMethod: "resetVpcConfig", GoMethod: "ResetVpcConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpcConfig", GoGetter: "VpcConfig"},
			_jsii_.MemberProperty{JsiiProperty: "vpcConfigInput", GoGetter: "VpcConfigInput"},
		},
		func() interface{} {
			j := jsiiProxy_MskReplicatorKafkaClustersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorKafkaClustersVpcConfig",
		reflect.TypeOf((*MskReplicatorKafkaClustersVpcConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorKafkaClustersVpcConfigOutputReference",
		reflect.TypeOf((*MskReplicatorKafkaClustersVpcConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecurityGroupIds", GoMethod: "ResetSecurityGroupIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetSubnetIds", GoMethod: "ResetSubnetIds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIds", GoGetter: "SecurityGroupIds"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIdsInput", GoGetter: "SecurityGroupIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIds", GoGetter: "SubnetIds"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIdsInput", GoGetter: "SubnetIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MskReplicatorKafkaClustersVpcConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorLogDelivery",
		reflect.TypeOf((*MskReplicatorLogDelivery)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorLogDeliveryOutputReference",
		reflect.TypeOf((*MskReplicatorLogDeliveryOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putReplicatorLogDelivery", GoMethod: "PutReplicatorLogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "replicatorLogDelivery", GoGetter: "ReplicatorLogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "replicatorLogDeliveryInput", GoGetter: "ReplicatorLogDeliveryInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetReplicatorLogDelivery", GoMethod: "ResetReplicatorLogDelivery"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MskReplicatorLogDeliveryOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorLogDeliveryReplicatorLogDelivery",
		reflect.TypeOf((*MskReplicatorLogDeliveryReplicatorLogDelivery)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorLogDeliveryReplicatorLogDeliveryCloudwatchLogs",
		reflect.TypeOf((*MskReplicatorLogDeliveryReplicatorLogDeliveryCloudwatchLogs)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorLogDeliveryReplicatorLogDeliveryCloudwatchLogsOutputReference",
		reflect.TypeOf((*MskReplicatorLogDeliveryReplicatorLogDeliveryCloudwatchLogsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "logGroupInput", GoGetter: "LogGroupInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogGroup", GoMethod: "ResetLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MskReplicatorLogDeliveryReplicatorLogDeliveryCloudwatchLogsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorLogDeliveryReplicatorLogDeliveryFirehose",
		reflect.TypeOf((*MskReplicatorLogDeliveryReplicatorLogDeliveryFirehose)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorLogDeliveryReplicatorLogDeliveryFirehoseOutputReference",
		reflect.TypeOf((*MskReplicatorLogDeliveryReplicatorLogDeliveryFirehoseOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStream", GoGetter: "DeliveryStream"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStreamInput", GoGetter: "DeliveryStreamInput"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeliveryStream", GoMethod: "ResetDeliveryStream"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MskReplicatorLogDeliveryReplicatorLogDeliveryFirehoseOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorLogDeliveryReplicatorLogDeliveryOutputReference",
		reflect.TypeOf((*MskReplicatorLogDeliveryReplicatorLogDeliveryOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLogs", GoGetter: "CloudwatchLogs"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLogsInput", GoGetter: "CloudwatchLogsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "firehose", GoGetter: "Firehose"},
			_jsii_.MemberProperty{JsiiProperty: "firehoseInput", GoGetter: "FirehoseInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putCloudwatchLogs", GoMethod: "PutCloudwatchLogs"},
			_jsii_.MemberMethod{JsiiMethod: "putFirehose", GoMethod: "PutFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "putS3", GoMethod: "PutS3"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudwatchLogs", GoMethod: "ResetCloudwatchLogs"},
			_jsii_.MemberMethod{JsiiMethod: "resetFirehose", GoMethod: "ResetFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3", GoMethod: "ResetS3"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "s3", GoGetter: "S3"},
			_jsii_.MemberProperty{JsiiProperty: "s3Input", GoGetter: "S3Input"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MskReplicatorLogDeliveryReplicatorLogDeliveryOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorLogDeliveryReplicatorLogDeliveryS3",
		reflect.TypeOf((*MskReplicatorLogDeliveryReplicatorLogDeliveryS3)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorLogDeliveryReplicatorLogDeliveryS3OutputReference",
		reflect.TypeOf((*MskReplicatorLogDeliveryReplicatorLogDeliveryS3OutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "bucketInput", GoGetter: "BucketInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "prefix", GoGetter: "Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "prefixInput", GoGetter: "PrefixInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetBucket", GoMethod: "ResetBucket"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefix", GoMethod: "ResetPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MskReplicatorLogDeliveryReplicatorLogDeliveryS3OutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorReplicationInfoListConsumerGroupReplication",
		reflect.TypeOf((*MskReplicatorReplicationInfoListConsumerGroupReplication)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorReplicationInfoListConsumerGroupReplicationOutputReference",
		reflect.TypeOf((*MskReplicatorReplicationInfoListConsumerGroupReplicationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "consumerGroupOffsetSyncMode", GoGetter: "ConsumerGroupOffsetSyncMode"},
			_jsii_.MemberProperty{JsiiProperty: "consumerGroupOffsetSyncModeInput", GoGetter: "ConsumerGroupOffsetSyncModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "consumerGroupsToExclude", GoGetter: "ConsumerGroupsToExclude"},
			_jsii_.MemberProperty{JsiiProperty: "consumerGroupsToExcludeInput", GoGetter: "ConsumerGroupsToExcludeInput"},
			_jsii_.MemberProperty{JsiiProperty: "consumerGroupsToReplicate", GoGetter: "ConsumerGroupsToReplicate"},
			_jsii_.MemberProperty{JsiiProperty: "consumerGroupsToReplicateInput", GoGetter: "ConsumerGroupsToReplicateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "detectAndCopyNewConsumerGroups", GoGetter: "DetectAndCopyNewConsumerGroups"},
			_jsii_.MemberProperty{JsiiProperty: "detectAndCopyNewConsumerGroupsInput", GoGetter: "DetectAndCopyNewConsumerGroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetConsumerGroupOffsetSyncMode", GoMethod: "ResetConsumerGroupOffsetSyncMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetConsumerGroupsToExclude", GoMethod: "ResetConsumerGroupsToExclude"},
			_jsii_.MemberMethod{JsiiMethod: "resetDetectAndCopyNewConsumerGroups", GoMethod: "ResetDetectAndCopyNewConsumerGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetSynchroniseConsumerGroupOffsets", GoMethod: "ResetSynchroniseConsumerGroupOffsets"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "synchroniseConsumerGroupOffsets", GoGetter: "SynchroniseConsumerGroupOffsets"},
			_jsii_.MemberProperty{JsiiProperty: "synchroniseConsumerGroupOffsetsInput", GoGetter: "SynchroniseConsumerGroupOffsetsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MskReplicatorReplicationInfoListConsumerGroupReplicationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorReplicationInfoListStruct",
		reflect.TypeOf((*MskReplicatorReplicationInfoListStruct)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorReplicationInfoListStructList",
		reflect.TypeOf((*MskReplicatorReplicationInfoListStructList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_MskReplicatorReplicationInfoListStructList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorReplicationInfoListStructOutputReference",
		reflect.TypeOf((*MskReplicatorReplicationInfoListStructOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "consumerGroupReplication", GoGetter: "ConsumerGroupReplication"},
			_jsii_.MemberProperty{JsiiProperty: "consumerGroupReplicationInput", GoGetter: "ConsumerGroupReplicationInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putConsumerGroupReplication", GoMethod: "PutConsumerGroupReplication"},
			_jsii_.MemberMethod{JsiiMethod: "putTopicReplication", GoMethod: "PutTopicReplication"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourceKafkaClusterArn", GoMethod: "ResetSourceKafkaClusterArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourceKafkaClusterId", GoMethod: "ResetSourceKafkaClusterId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetKafkaClusterArn", GoMethod: "ResetTargetKafkaClusterArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetKafkaClusterId", GoMethod: "ResetTargetKafkaClusterId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sourceKafkaClusterArn", GoGetter: "SourceKafkaClusterArn"},
			_jsii_.MemberProperty{JsiiProperty: "sourceKafkaClusterArnInput", GoGetter: "SourceKafkaClusterArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "sourceKafkaClusterId", GoGetter: "SourceKafkaClusterId"},
			_jsii_.MemberProperty{JsiiProperty: "sourceKafkaClusterIdInput", GoGetter: "SourceKafkaClusterIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetCompressionType", GoGetter: "TargetCompressionType"},
			_jsii_.MemberProperty{JsiiProperty: "targetCompressionTypeInput", GoGetter: "TargetCompressionTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetKafkaClusterArn", GoGetter: "TargetKafkaClusterArn"},
			_jsii_.MemberProperty{JsiiProperty: "targetKafkaClusterArnInput", GoGetter: "TargetKafkaClusterArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetKafkaClusterId", GoGetter: "TargetKafkaClusterId"},
			_jsii_.MemberProperty{JsiiProperty: "targetKafkaClusterIdInput", GoGetter: "TargetKafkaClusterIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "topicReplication", GoGetter: "TopicReplication"},
			_jsii_.MemberProperty{JsiiProperty: "topicReplicationInput", GoGetter: "TopicReplicationInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorReplicationInfoListTopicReplication",
		reflect.TypeOf((*MskReplicatorReplicationInfoListTopicReplication)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorReplicationInfoListTopicReplicationOutputReference",
		reflect.TypeOf((*MskReplicatorReplicationInfoListTopicReplicationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "copyAccessControlListsForTopics", GoGetter: "CopyAccessControlListsForTopics"},
			_jsii_.MemberProperty{JsiiProperty: "copyAccessControlListsForTopicsInput", GoGetter: "CopyAccessControlListsForTopicsInput"},
			_jsii_.MemberProperty{JsiiProperty: "copyTopicConfigurations", GoGetter: "CopyTopicConfigurations"},
			_jsii_.MemberProperty{JsiiProperty: "copyTopicConfigurationsInput", GoGetter: "CopyTopicConfigurationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "detectAndCopyNewTopics", GoGetter: "DetectAndCopyNewTopics"},
			_jsii_.MemberProperty{JsiiProperty: "detectAndCopyNewTopicsInput", GoGetter: "DetectAndCopyNewTopicsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putStartingPosition", GoMethod: "PutStartingPosition"},
			_jsii_.MemberMethod{JsiiMethod: "putTopicNameConfiguration", GoMethod: "PutTopicNameConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetCopyAccessControlListsForTopics", GoMethod: "ResetCopyAccessControlListsForTopics"},
			_jsii_.MemberMethod{JsiiMethod: "resetCopyTopicConfigurations", GoMethod: "ResetCopyTopicConfigurations"},
			_jsii_.MemberMethod{JsiiMethod: "resetDetectAndCopyNewTopics", GoMethod: "ResetDetectAndCopyNewTopics"},
			_jsii_.MemberMethod{JsiiMethod: "resetStartingPosition", GoMethod: "ResetStartingPosition"},
			_jsii_.MemberMethod{JsiiMethod: "resetTopicNameConfiguration", GoMethod: "ResetTopicNameConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetTopicsToExclude", GoMethod: "ResetTopicsToExclude"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "startingPosition", GoGetter: "StartingPosition"},
			_jsii_.MemberProperty{JsiiProperty: "startingPositionInput", GoGetter: "StartingPositionInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "topicNameConfiguration", GoGetter: "TopicNameConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "topicNameConfigurationInput", GoGetter: "TopicNameConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "topicsToExclude", GoGetter: "TopicsToExclude"},
			_jsii_.MemberProperty{JsiiProperty: "topicsToExcludeInput", GoGetter: "TopicsToExcludeInput"},
			_jsii_.MemberProperty{JsiiProperty: "topicsToReplicate", GoGetter: "TopicsToReplicate"},
			_jsii_.MemberProperty{JsiiProperty: "topicsToReplicateInput", GoGetter: "TopicsToReplicateInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorReplicationInfoListTopicReplicationStartingPosition",
		reflect.TypeOf((*MskReplicatorReplicationInfoListTopicReplicationStartingPosition)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorReplicationInfoListTopicReplicationStartingPositionOutputReference",
		reflect.TypeOf((*MskReplicatorReplicationInfoListTopicReplicationStartingPositionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetType", GoMethod: "ResetType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationStartingPositionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorReplicationInfoListTopicReplicationTopicNameConfiguration",
		reflect.TypeOf((*MskReplicatorReplicationInfoListTopicReplicationTopicNameConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorReplicationInfoListTopicReplicationTopicNameConfigurationOutputReference",
		reflect.TypeOf((*MskReplicatorReplicationInfoListTopicReplicationTopicNameConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetType", GoMethod: "ResetType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationTopicNameConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorTags",
		reflect.TypeOf((*MskReplicatorTags)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorTagsList",
		reflect.TypeOf((*MskReplicatorTagsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_MskReplicatorTagsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorTagsOutputReference",
		reflect.TypeOf((*MskReplicatorTagsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "keyInput", GoGetter: "KeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetKey", GoMethod: "ResetKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetValue", GoMethod: "ResetValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_MskReplicatorTagsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
}
