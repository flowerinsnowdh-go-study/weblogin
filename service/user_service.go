/*
 * Copyright (c) 2024 flowerinsnow
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package service

import (
	"crypto/rand"
	"github.com/flowerinsnowdh/weblogin/dao"
	"github.com/flowerinsnowdh/weblogin/pojo"
	"github.com/tjfoc/gmsm/sm3"
	"hash"
)

func (s *Service) InsertDefaultUser() (*pojo.User, error) {
	const admin = "admin"

	var (
		d             *dao.DAO  = (*dao.DAO)(s)
		salt          []byte    = make([]byte, 32) // 盐
		plainPassword []byte    = []byte(admin)    // 默认密码
		h             hash.Hash = sm3.New()        // 哈希函数
	)

	// 随机生成盐
	_, _ = rand.Read(salt)

	// 哈希盐+密码
	h.Write(salt)
	h.Write(plainPassword)
	var password []byte = h.Sum(make([]byte, 0))

	return d.InsertUser(admin, salt, password)
}
