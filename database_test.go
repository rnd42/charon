/*
 *  Charon: A game authentication server
 *  Copyright (C) 2014-2016  Alex Mayfield <alexmax2742@gmail.com>
 *
 *  This program is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU Affero General Public License as published by
 *  the Free Software Foundation, either version 3 of the License, or
 *  (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU Affero General Public License for more details.
 *
 *  You should have received a copy of the GNU Affero General Public License
 *  along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package charon

import "testing"

func TestNewDatabase(t *testing.T) {
	_, err := NewDatabase(NewConfig(nil))
	if err != nil {
		t.Errorf("%s", err.Error())
	}
}

func TestAddUser(t *testing.T) {
	database, err := NewDatabase(NewConfig(nil))
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	err = database.AddUser("TestUser", "testuser@example.com", "VsGnJghDUW6C")
	if err != nil {
		t.Errorf("%s", err.Error())
	}
}

func TestAddUserErrors(t *testing.T) {
	database, err := NewDatabase(NewConfig(nil))
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	err = database.Import("fixture/user.sql")
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	err = database.AddUser("TestUser", "testuser@example.com", "VsGnJghDUW6C")
	if err == nil {
		t.Errorf("charon: user added despite uniqueness constraint")
	}
}

func TestFindUser(t *testing.T) {
	database, err := NewDatabase(NewConfig(nil))
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	err = database.Import("fixture/user.sql")
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	_, err = database.FindUser("TestUser")
	if err != nil {
		t.Errorf("%s", err.Error())
	}
}

func TestLoginUser(t *testing.T) {
	database, err := NewDatabase(NewConfig(nil))
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	database.Import("fixture/user.sql")
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	_, err = database.LoginUser("TestUser", "VsGnJghDUW6C")
	if err != nil {
		t.Errorf("%s", err.Error())
	}
}
