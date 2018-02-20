package status

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

import (
	"testing"
	"time"

	"github.com/apache/incubator-trafficcontrol/lib/go-tc"
	"github.com/apache/incubator-trafficcontrol/traffic_ops/traffic_ops_golang/auth"
	"github.com/apache/incubator-trafficcontrol/traffic_ops/traffic_ops_golang/test"
	"github.com/jmoiron/sqlx"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func getTestStatuses() []tc.Status {
	cdns := []tc.Status{}
	testStatus := tc.Status{
		Description: "description",
		ID:          1,
		Name:        "cdn1",
		LastUpdated: tc.Time{Time: time.Now()},
	}
	cdns = append(cdns, testStatus)

	testStatus2 := testStatus
	testStatus2.Name = "cdn2"
	testStatus2.Description = "description2"
	cdns = append(cdns, testStatus2)

	return cdns
}

func TestReadStatuses(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	db := sqlx.NewDb(mockDB, "sqlmock")
	defer db.Close()

	refType := GetRefType()

	testStatuses := getTestStatuses()
	cols := test.ColsFromStructByTag("db", tc.Status{})
	rows := sqlmock.NewRows(cols)

	for _, ts := range testStatuses {
		rows = rows.AddRow(
			ts.Description,
			ts.ID,
			ts.LastUpdated,
			ts.Name,
		)
	}
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	v := map[string]string{"dsId": "1"}

	servers, errs, _ := refType.Read(db, v, auth.CurrentUser{})
	if len(errs) > 0 {
		t.Errorf("cdn.Read expected: no errors, actual: %v", errs)
	}

	if len(servers) != 2 {
		t.Errorf("cdn.Read expected: len(servers) == 2, actual: %v", len(servers))
	}
}