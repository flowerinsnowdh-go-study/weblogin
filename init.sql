/*
 * Copyright (c) 2024 flowerinsnow
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
DROP TABLE IF EXISTS `user`;

CREATE TABLE IF NOT EXISTS `user` (
    `id`              INTEGER PRIMARY KEY AUTOINCREMENT,
    `name`            VARCHAR(16) UNIQUE NOT NULL,
    `salt`            BINARY(32) NOT NULL,
    `password`        BINARY(32) NOT NULL,
    `access_token`    BINARY(32),
    `admin`           BOOLEAN NOT NULL,
    `create_time`     DATETIME NOT NULL,
    `last_login_time` DATETIME
);
