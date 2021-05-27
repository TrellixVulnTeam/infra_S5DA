// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package frontend

import (
	"testing"

	"go.chromium.org/luci/appengine/gaetesting"
	"go.chromium.org/luci/gae/service/datastore"
)

// TestReadActionEntityFromEmptyDatastore check that a read from a consistent datastore with
// nothing in it consistently produces no action entities.
func TestReadActionEntityFromEmptyDatastore(t *testing.T) {
	t.Parallel()
	ctx := gaetesting.TestingContext()
	datastore.GetTestable(ctx).Consistent(true)
	q := MakeAllActionEntitiesQuery("")
	es, err := q.Next(ctx, 100)
	token := q.Token
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if len(es) != 0 {
		t.Errorf("unexpected entities: %v", es)
	}
	if token != "" {
		t.Errorf("expected tok to be empty not %s", token)
	}
}

// TestReadObservationEntityFromEmptyDatastore check that a read from a consistent datastore with
// nothing in it consistently produces no observation entities.
func TestReadObservationEntityFromEmptyDatastore(t *testing.T) {
	t.Parallel()
	ctx := gaetesting.TestingContext()
	datastore.GetTestable(ctx).Consistent(true)
	q := MakeAllObservationEntitiesQuery("")
	es, err := q.Next(ctx, 100)
	token := q.Token
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if len(es) != 0 {
		t.Errorf("unexpected entities: %v", es)
	}
	if token != "" {
		t.Errorf("expected tok to be empty not %s", token)
	}
}

// TestReadSingleActionEntityFromDatastore tests putting a single action entity into datastore
// and then reading it back out.
func TestReadSingleActionEntityFromDatastore(t *testing.T) {
	t.Parallel()
	ctx := gaetesting.TestingContext()
	datastore.GetTestable(ctx).Consistent(true)
	if err := PutActionEntities(ctx, &ActionEntity{
		ID: "hi",
	}); err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	q := MakeAllActionEntitiesQuery("")
	es, err := q.Next(ctx, 100)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if len(es) != 1 {
		t.Errorf("unexpected entities: %v", es)
	}
}

// TestReadSingleObservationEntityFromDatastore tests putting a single observation entity into datastore
// and then reading it back out.
func TestReadSingleObservationEntityFromDatastore(t *testing.T) {
	t.Parallel()
	ctx := gaetesting.TestingContext()
	datastore.GetTestable(ctx).Consistent(true)
	if err := PutObservationEntities(ctx, &ObservationEntity{
		ID: "hi",
	}); err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	q := MakeAllObservationEntitiesQuery("")
	es, err := q.Next(ctx, 100)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if len(es) != 1 {
		t.Errorf("unexpected entities: %v", es)
	}
}

// TestReadActionEntitiesFromDatastoreOneAtAtime reads several action entities out of datastore one at a time
// using pagination.
func TestReadActionEntitiesFromDatastoreOneAtAtime(t *testing.T) {
	t.Parallel()
	ctx := gaetesting.TestingContext()
	datastore.GetTestable(ctx).Consistent(true)
	if err := PutActionEntities(
		ctx,
		&ActionEntity{
			ID: "entity1",
		},
		&ActionEntity{
			ID: "entity2",
		},
	); err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	// Test that extracting the first ActionEntity produces some action entity.
	var resultTok string
	{
		q := MakeAllActionEntitiesQuery("")
		es, err := q.Next(ctx, 1)
		resultTok = q.Token
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
		if len(es) != 1 {
			t.Errorf("unexpected entities: %v", es)
		}
		if q.Token == "" {
			t.Errorf("token should not be empty!")
		}
	}
	// Test that extracting the second ActionEntity produces some action entity.
	{
		q := MakeAllActionEntitiesQuery(resultTok)
		es, err := q.Next(ctx, 10)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
		if len(es) != 1 {
			t.Errorf("unexpected entities: %v", es)
		}
		// Even though there are no further entities to extract,
		// we haven't observed the end of the query yet and therefore
		// should not see an end token.
		if q.Token != "" {
			t.Errorf("unexpected token: %q", q.Token)
		}
	}
}