/*
 * Copyright (c) 2024 flowerinsnow
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package service

import "github.com/flowerinsnowdh/weblogin/dao"

type Service dao.DAO

func NewService(d *dao.DAO) *Service {
	return (*Service)(d)
}
