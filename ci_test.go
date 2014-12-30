/*
 * Alexandria CMDB - Open source common.management database
 * Copyright (C) 2014  Ryan Armstrong <ryan@cavaliercoder.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 * package controllers
 */
package main

import (
	"fmt"
	"testing"
)

const (
	ciDB   = "TestCIDB"
	ciType = "TestCIType"
)

func TestCI(t *testing.T) {
	// Create a temporary cmdb
	uri := V1Uri("/cmdbs")
	body := fmt.Sprintf(`{"name":"%s"}`, ciDB)
	dburl := Post(t, uri, body)
	defer Delete(t, dburl)

	// Create temporary CI Type
	uri = V1Uri(fmt.Sprintf("/cmdbs/%s/citypes", ciDB))
	body = `{
		"name":"TestCIType",
		"description": "A test CI Type",
		"attributes": []
		}`
	typUrl := Post(t, uri, body)
	defer Delete(t, typUrl)

	// Test POST ../CI
	uri = V1Uri(fmt.Sprintf("/cmdbs/%s/%s", ciDB, ciType))
	body = `{
		"firstAttribute":"String value"
	}`
	Crud(t, uri, body, false)
}
