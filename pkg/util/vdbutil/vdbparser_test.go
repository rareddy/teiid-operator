/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package vdbutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabase(t *testing.T) {
	var ddl = `CREATE DATABASE /* foo */ 
		customer OPTIONS (ANNOTATION 'Customer VDB');
	USE DATABASE customer;

	CREATE FOREIGN DATA WRAPPER postgresql;
	CREATE SERVER /* foo */
		sampledb TYPE "NONE" FOREIGN DATA WRAPPER postgresql;

	CREATE SERVER mongo TYPE 'NONE' FOREIGN DATA WRAPPER mongodb;

	CREATE SCHEMA accounts SERVER sampledb;
	CREATE VIRTUAL SCHEMA portfolio;`

	db, err := Parse(ddl)

	assert.Nil(t, err)
	assert.NotNil(t, db)

	assert.Equal(t, "customer", db.Name)
	assert.NotNil(t, db.Options)
	assert.Equal(t, "Customer VDB", db.Options["ANNOTATION"])

	assert.NotNil(t, db.Translators)
	assert.Equal(t, "postgresql", db.Translators[0].Name)

}
