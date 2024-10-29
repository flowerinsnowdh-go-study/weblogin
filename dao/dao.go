/*
 * Copyright (c) 2024 flowerinsnow
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package dao

import (
	"database/sql"
)

type DAO sql.DB

func NewDao(db *sql.DB) *DAO {
	return (*DAO)(db)
}
