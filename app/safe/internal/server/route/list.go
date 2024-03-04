/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package route

import (
	"net/http"
)

// List returns all registered workloads to the system with some metadata
// that is secure to share. For example, it returns secret names but not values.
//
// - cid: A string representing the client identifier.
// - w: An http.ResponseWriter used to write the HTTP response.
// - r: A pointer to an http.Request representing the received HTTP request.
// - spiffeid: spiffe id of the caller.
func List(cid string, w http.ResponseWriter, r *http.Request, spiffeid string) {
	doList(cid, w, r, spiffeid, false)
}

// ListEncrypted returns all registered workloads to the system. Similar to `List`
// it return meta information; however, it also returns encrypted secret values
// where an operator can decrypt if they know the VSecM root key.
//
// - cid: A string representing the client identifier.
// - w: An http.ResponseWriter used to write the HTTP response.
// - r: A pointer to an http.Request representing the received HTTP request.
// - spiffeid: spiffe id of the caller.
func ListEncrypted(cid string, w http.ResponseWriter, r *http.Request, spiffeid string) {
	doList(cid, w, r, spiffeid, true)
}