// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rdsdbproxytargetgroup


type RdsDbProxyTargetGroupConnectionPoolConfigurationInfo struct {
	// The number of seconds for a proxy to wait for a connection to become available in the connection pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_db_proxy_target_group#connection_borrow_timeout RdsDbProxyTargetGroup#connection_borrow_timeout}
	ConnectionBorrowTimeout *float64 `field:"optional" json:"connectionBorrowTimeout" yaml:"connectionBorrowTimeout"`
	// One or more SQL statements for the proxy to run when opening each new database connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_db_proxy_target_group#init_query RdsDbProxyTargetGroup#init_query}
	InitQuery *string `field:"optional" json:"initQuery" yaml:"initQuery"`
	// The maximum size of the connection pool for each target in a target group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_db_proxy_target_group#max_connections_percent RdsDbProxyTargetGroup#max_connections_percent}
	MaxConnectionsPercent *float64 `field:"optional" json:"maxConnectionsPercent" yaml:"maxConnectionsPercent"`
	// Controls how actively the proxy closes idle database connections in the connection pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_db_proxy_target_group#max_idle_connections_percent RdsDbProxyTargetGroup#max_idle_connections_percent}
	MaxIdleConnectionsPercent *float64 `field:"optional" json:"maxIdleConnectionsPercent" yaml:"maxIdleConnectionsPercent"`
	// Each item in the list represents a class of SQL operations that normally cause all later statements in a session using a proxy to be pinned to the same underlying database connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_db_proxy_target_group#session_pinning_filters RdsDbProxyTargetGroup#session_pinning_filters}
	SessionPinningFilters *[]*string `field:"optional" json:"sessionPinningFilters" yaml:"sessionPinningFilters"`
}

