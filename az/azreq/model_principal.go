// Copyright 2025 Nitro Agility S.r.l.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package azreq

// Principal is the principal object.
type Principal struct {
	kind   string
	id     string
	source string
}

// GetKind returns the kind of the principal.
func (u *Principal) GetKind() string {
	return u.kind
}

// GetID returns the ID of the principal.
func (u *Principal) GetID() string {
	return u.id
}

// GetSource returns the source of the principal.
func (u *Principal) GetSource() string {
	return u.source
}

// PrincipalBuilder is the builder for the principal object.
type PrincipalBuilder struct {
	principal *Principal
}

// NewPrincipalBuilder creates a new principal builder.
func NewPrincipalBuilder(id string) *PrincipalBuilder {
	return &PrincipalBuilder{
		principal: &Principal{
			id: id,
		},
	}
}

// WithKind sets the kind of the principal.
func (b *PrincipalBuilder) WithKind(kind string) *PrincipalBuilder {
	b.principal.kind = kind
	return b
}

// WithProperty sets the property of the principal.
func (b *PrincipalBuilder) WithSource(source string) *PrincipalBuilder {
	b.principal.source = source
	return b
}

// Build builds the principal object.
func (b *PrincipalBuilder) Build() *Principal {
	return b.principal
}
