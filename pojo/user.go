/*
 * Copyright (c) 2024 flowerinsnow
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package pojo

import "time"

type User struct {
	Id            int
	Name          string
	Salt          []byte
	Password      []byte
	AccessToken   []byte
	Admin         bool
	CreateTime    *time.Time
	LastLoginTime *time.Time
}
