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

var database *Database = newDatabase()

// Database is the type of the parser result
type Database struct {
	Name        string
	Version     string
	Options     map[string]string
	Translators []Translator
	Servers     []Server
}

// Translator --
type Translator struct {
	Name     string
	BaseType string
	Options  map[string]string
}

// Server --
type Server struct {
	Name       string
	Translator *Translator
	Options    map[string]string
}

// Statement --
type Statement struct {
}

func newDatabase() *Database {
	return &Database{
		Translators: make([]Translator, 0),
		Servers:     make([]Server, 0),
	}
}

func setDatabase(name string, version string, options map[string]string) {
	database.Name = name
	database.Version = version
	database.Options = options
}

func addComment(str string) {
	// nothing
}

func addDataWrapper(name string, typeName string, options map[string]string) {
	translator := Translator{
		Name:     name,
		BaseType: typeName,
		Options:  options,
	}
	database.Translators = append(database.Translators, translator)
}

func addServer(name string, translatorName string, options map[string]string) {
	var translator *Translator
	for _, t := range database.Translators {
		if t.Name == translatorName {
			translator = &t
			break
		}
	}

	if translator == nil {
		log.Info("No Translator definition found in the DDL with name ", translatorName, " adding as known type")
		t := Translator{
			Name: translatorName,
		}
		database.Translators = append(database.Translators, t)
		translator = &t
	}

	server := Server{
		Name:       name,
		Translator: translator,
		Options:    options,
	}
	database.Servers = append(database.Servers, server)
}
