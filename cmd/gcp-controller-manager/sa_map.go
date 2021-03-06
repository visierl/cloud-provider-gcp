/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

// gsaEmail identifies a GCP service account in email format.
type gsaEmail string

// serviceAccount identifies a K8s service account object by its namespace and name.  Empty
// Namespace indicates the corresponding Kubernetes object was created in the "default" namespace.
type serviceAccount struct {
	Namespace, Name string
}

// MarshalText implements the encoding.TextMarshaler interface.
func (sa serviceAccount) MarshalText() ([]byte, error) {
	return []byte(sa.String()), nil
}

// String returns sa in a string as "<namespace>/<name>" or "default/<name>" if sa.Namespace is
// empty.
func (sa serviceAccount) String() string {
	if sa.Namespace == "" {
		return fmt.Sprintf("default/%s", sa.Name)
	}
	return fmt.Sprintf("%s/%s", sa.Namespace, sa.Name)
}

// saMap is a Mutax protected map of gsaEmail keyed by serviceAccount.
type saMap struct {
	sync.RWMutex
	ma map[serviceAccount]gsaEmail
}

func newSAMap() *saMap {
	t := make(map[serviceAccount]gsaEmail)
	return &saMap{
		ma: t,
	}
}

// Add stores the mapping from sa to gsa to m and returns the previous gsa if sa already existed.
func (m *saMap) add(sa serviceAccount, gsa gsaEmail) (gsaEmail, bool) {
	m.Lock()
	defer m.Unlock()
	lastGSA, found := m.ma[sa]
	m.ma[sa] = gsa
	return lastGSA, found
}

// Remove removes the entry keyed by sa in m and returns its gsa if sa existed.
func (m *saMap) remove(sa serviceAccount) (gsaEmail, bool) {
	m.Lock()
	defer m.Unlock()
	removedGSA, found := m.ma[sa]
	if found {
		delete(m.ma, sa)
	}
	return removedGSA, found
}

// Get looks up sa from m and returns its gsa if sa exists.
func (m *saMap) get(sa serviceAccount) (gsaEmail, bool) {
	m.RLock()
	defer m.RUnlock()
	gsa, ok := m.ma[sa]
	return gsa, ok
}

// Serialize returns m in its JSON encoded format or error if serialization had failed.
func (m *saMap) serialize() ([]byte, error) {
	m.RLock()
	defer m.RUnlock()
	return json.Marshal(m.ma)
}
