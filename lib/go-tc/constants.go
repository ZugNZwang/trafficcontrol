package tc

/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// ErrorConstant is used for error messages.
type ErrorConstant string

// Error converts ErrorConstants to a string.
func (e ErrorConstant) Error() string { return string(e) }

// DBError is an error message for database errors.
const DBError = ErrorConstant("database access error")

// NilTenantError can used when a Tenantable object finds that TentantID in the
// request is nil.
const NilTenantError = ErrorConstant("tenancy is enabled but request tenantID is nil")

// TenantUserNotAuthError is used when a user does not have access to a
// requested resource tenant.
const TenantUserNotAuthError = ErrorConstant("user not authorized for requested tenant")

// TenantDSUserNotAuthError is used when a user does not have access to a
// requested resource tenant for a delivery service.
const TenantDSUserNotAuthError = ErrorConstant("user not authorized for requested delivery service tenant")
const NeedsAtLeastOneIPError = ErrorConstant("both IP and IP6 addresses are empty")
const EmptyAddressCannotBeAServiceAddressError = ErrorConstant("an empty IP or IPv6 address cannot be marked as a service address")

// AlertLevel is used for specifying or comparing different levels of alerts.
type AlertLevel int

const (
	// SuccessLevel indicates that an action is successful.
	SuccessLevel AlertLevel = iota

	// InfoLevel indicates that the message is supplementary and is not directly
	// the result of the user's request.
	InfoLevel

	// WarnLevel indicates dangerous but non-failing conditions.
	WarnLevel

	// ErrorLevel indicates that the request failed.
	ErrorLevel
)

var alertLevels = [4]string{"success", "info", "warning", "error"}

// String returns the string representation of an AlertLevel.
func (a AlertLevel) String() string {
	return alertLevels[a]
}

// CachegroupCoordinateNamePrefix is a string that all cache group coordinate
// names are prefixed with.
const CachegroupCoordinateNamePrefix = "from_cachegroup_"
