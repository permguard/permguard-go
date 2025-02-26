// Copyright 2024 Nitro Agility S.r.l.
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

package main

import (
	"fmt"

	"github.com/permguard/permguard-go"
	"github.com/permguard/permguard-go/az/azreq"
)

func checkAtomicEvaluation() {
	// Create a new Permguard client
	azClient := permguard.NewAZClient(
		permguard.WithEndpoint("localhost", 9094),
	)

	// Create the Principal
	principal := azreq.NewPrincipalBuilder("amy.smith@acmecorp.com").Build()

	// Create the entities
	entities := []map[string]any{
		{
			"uid": map[string]any{
				"type": "MagicFarmacia::Platform::BranchInfo",
				"id":   "subscription",
			},
			"attrs": map[string]any{
				"active": true,
			},
			"parents": []any{},
		},
	}

	req := azreq.NewAZAtomicRequestBuilder(204510383118, "181e252e247747338ad062abad0086a5",
		"amy.smith@acmecorp.com", "MagicFarmacia::Platform::Subscription", "MagicFarmacia::Platform::Action::view").
		// RequestID
		WithRequestID("1234").
		// Principal
		WithPrincipal(principal).
		// Entities
		WithEntitiesItems(azreq.CedarEntityKind, entities).
		// Subject
		WithSubjectKind("user").
		WithSubjectSource("keycloack").
		WithSubjectProperty("isSuperUser", true).
		// Resource
		WithResourceID("e3a786fd07e24bfa95ba4341d3695ae8").
		WithResourceProperty("isEnabled", true).
		// Action
		WithActionProperty("isEnabled", true).
		WithContextProperty("time", "2025-01-23T16:17:46+00:00").
		WithContextProperty("isSubscriptionActive", true).
		Build()

	// Check the authorization
	decsion, response, _ := azClient.Check(req)
	if decsion {
		fmt.Println("✅ Authorization Permitted")
	} else {
		fmt.Println("❌ Authorization Denied")
		if response.Context.ReasonAdmin != nil {
			fmt.Printf("-> Reason Admin: %s\n", response.Context.ReasonAdmin.Message)
		}
		if response.Context.ReasonUser != nil {
			fmt.Printf("-> Reason User: %s\n", response.Context.ReasonUser.Message)
		}
		for _, eval := range response.Evaluations {
			if eval.Context.ReasonUser != nil {
				fmt.Printf("-> Reason Admin: %s\n", eval.Context.ReasonAdmin.Message)
				fmt.Printf("-> Reason User: %s\n", eval.Context.ReasonUser.Message)
			}
		}
	}
}

// checkMultipleEvaluations checks multiple evaluations
func checkMultipleEvaluations() {
	// Create a new Permguard client
	azClient := permguard.NewAZClient(
		permguard.WithEndpoint("localhost", 9094),
	)

	// Create a new subject
	subject := azreq.NewSubjectBuilder("amy.smith@acmecorp.com").
		WithKind("user").
		WithSource("keycloack").
		WithProperty("isSuperUser", true).
		Build()

	// Create a new resource
	resource := azreq.NewResourceBuilder("MagicFarmacia::Platform::Subscription").
		WithID("e3a786fd07e24bfa95ba4341d3695ae8").
		WithProperty("isEnabled", true).
		Build()

	// Create ations
	actionView := azreq.NewActionBuilder("MagicFarmacia::Platform::Action::view").
		WithProperty("isEnabled", true).
		Build()

	actionCreate := azreq.NewActionBuilder("MagicFarmacia::Platform::Action::create").
		WithProperty("isEnabled", true).
		Build()

	// Create a new Context
	context := azreq.NewContextBuilder().
		WithProperty("time", "2025-01-23T16:17:46+00:00").
		WithProperty("isSubscriptionActive", true).
		Build()

	// Create evaluations
	evaluationView := azreq.NewEvaluationBuilder(subject, resource, actionView).
		WithRequestID("1234").
		WithContext(context).
		Build()

	evaluationCreate := azreq.NewEvaluationBuilder(subject, resource, actionCreate).
		WithRequestID("7890").
		WithContext(context).
		Build()

	// Create the Principal
	principal := azreq.NewPrincipalBuilder("amy.smith@acmecorp.com").Build()

	// Create the entities
	entities := []map[string]any{
		{
			"uid": map[string]any{
				"type": "MagicFarmacia::Platform::BranchInfo",
				"id":   "subscription",
			},
			"attrs": map[string]any{
				"active": true,
			},
			"parents": []any{},
		},
	}

	// Create a new request
	req := azreq.NewAZRequestBuilder(204510383118, "181e252e247747338ad062abad0086a5").
		WithPrincipal(principal).
		WithEntitiesItems(azreq.CedarEntityKind, entities).
		WithEvaluation(evaluationView).
		WithEvaluation(evaluationCreate).
		Build()

	// Check the authorization
	decsion, response, _ := azClient.Check(req)
	if decsion {
		fmt.Println("✅ Authorization Permitted")
	} else {
		fmt.Println("❌ Authorization Denied")
		if response.Context.ReasonAdmin != nil {
			fmt.Printf("-> Reason Admin: %s\n", response.Context.ReasonAdmin.Message)
		}
		if response.Context.ReasonUser != nil {
			fmt.Printf("-> Reason User: %s\n", response.Context.ReasonUser.Message)
		}
		for _, eval := range response.Evaluations {
			if eval.Context.ReasonUser != nil {
				fmt.Printf("-> Reason Admin: %s\n", eval.Context.ReasonAdmin.Message)
				fmt.Printf("-> Reason User: %s\n", eval.Context.ReasonUser.Message)
			}
		}
	}
}

func main() {
	checkAtomicEvaluation()
	//checkMultipleEvaluations()
}
