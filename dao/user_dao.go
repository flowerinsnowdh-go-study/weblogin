package dao

import (
	"database/sql"
	"errors"
	"github.com/flowerinsnowdh/weblogin/cerror"
	"github.com/flowerinsnowdh/weblogin/pojo"
)

// InsertUser 会插入一个指定用户名、盐和密码的用户数据库行
// cipherPassword 应该是
func (d *DAO) InsertUser(name string, salt []byte, cipherPassword []byte) (*pojo.User, error) {
	var db *sql.DB = (*sql.DB)(d)
	var tx *sql.Tx
	var err error
	if tx, err = db.Begin(); err != nil { // 开启事务
		return nil, err
	}

	var (
		insertErr,
		selectErr error
		id int
	)

	// 插入数据
	if _, insertErr = tx.Exec(
		"INSERT INTO `user` (`name`, `salt`, `password`, `admin`, `create_time`) VALUES (?, ?, ?, TRUE, DATETIME('NOW', 'LOCALTIME'))",
		name, salt, cipherPassword,
	); insertErr == nil {
		// 插入数据没问题，查询数据
		var rows *sql.Rows

		if rows, selectErr = tx.Query("SELECT LAST_INSERT_ROWID()"); selectErr == nil {
			if !rows.Next() {
				panic(errors.New("unexpected cerror")) // TODO DEBUG
			}
			if err := rows.Scan(&id); err != nil {
				panic(errors.New("unexpected cerror")) // TODO DEBUG
			}
		}
	}

	if insertErr != nil || selectErr != nil { // 插入或查询数据有问题
		if rbErr := tx.Rollback(); rbErr != nil { // 回滚
			return nil, rbErr // 返回回滚有问题
		} else if insertErr != nil { // 插入数据有问题
			return nil, insertErr // 返回插入数据有问题
		} else { // 查询数据有问题
			return nil, selectErr // 返回查询数据有问题
		}
	}

	if err = tx.Commit(); err != nil { // 提交数据
		return nil, err // 返回提交数据有问题
	}

	// 一切正常，返回用户
	return &pojo.User{
		Id:       id,
		Name:     name,
		Salt:     salt,
		Password: cipherPassword,
	}, nil
}

// IsUserIdExists 返回一个用户 Id 是否已存在
// 如果查询 SQL 时出现错误，返回 error
func (d *DAO) IsUserIdExists(id int) (bool, error) {
	var db *sql.DB = (*sql.DB)(d)

	if rows, err := db.Query(
		"SELECT COUNT(*) > 0 FROM `user` WHERE `id` = ?",
		id,
	); err != nil {
		return false, err
	} else {
		defer func(rows *sql.Rows) {
			_ = rows.Close()
		}(rows)
		if rows.Next() {
			var exists bool
			if err := rows.Scan(&exists); err != nil {
				return false, err
			} else {
				return exists, nil
			}
		} else {
			panic(cerror.NewUnexpectedError(""))
		}
	}
}

// IsUserExistsExists 返回一个用户 Id 是否已存在
// 如果查询 SQL 时出现错误，返回 error
func (d *DAO) IsUserExistsExists(name string) (bool, error) {
	var db *sql.DB = (*sql.DB)(d)

	if rows, err := db.Query(
		"SELECT COUNT(*) > 0 FROM `user` WHERE `name` = ?",
		name,
	); err != nil {
		return false, err
	} else {
		defer func(rows *sql.Rows) {
			_ = rows.Close()
		}(rows)
		if rows.Next() {
			var exists bool
			if err := rows.Scan(&exists); err != nil {
				return false, err
			} else {
				return exists, nil
			}
		} else { // SELECT 语句选取的是一个 BOOL 类型，不应该出现空结果的情况
			panic(cerror.NewUnexpectedError(""))
		}
	}
}

// SelectUserById 通过 Id 查找用户
// 如果查询 SQL 时出现错误，返回 error
func (d *DAO) SelectUserById(id int) (*pojo.User, error) {
	var db *sql.DB = (*sql.DB)(d)

	if rows, err := db.Query("SELECT `name`, `salt`, `password`, `access_token`, `admin`, `create_time`, `last_login_time` FROM `user` WHERE `id` = ?"); err != nil {
		return nil, err
	} else {
		defer func(rows *sql.Rows) {
			_ = rows.Close()
		}(rows)
		if rows.Next() {
			var user *pojo.User = &pojo.User{
				Id:            id,
				Name:          "",
				Salt:          nil,
				Password:      nil,
				AccessToken:   nil,
				Admin:         false,
				CreateTime:    nil,
				LastLoginTime: nil,
			}
			if err := rows.Scan(&user.Name, &user.Salt, &user.Password, &user.AccessToken, &user.Admin, &user.CreateTime, &user.LastLoginTime); err != nil {
				return nil, err
			} else {
				return user, nil
			}
		} else {
			return nil, nil
		}
	}
}

// SelectUserByName 通过用户名查找用户
// 如果查询 SQL 时出现错误，返回 error
func (d *DAO) SelectUserByName(name string) (*pojo.User, error) {
	var db *sql.DB = (*sql.DB)(d)

	if rows, err := db.Query("SELECT `id`, `salt`, `password`, `access_token`, `admin`, `create_time`, `last_login_time` FROM `user` WHERE `id` = ?"); err != nil {
		return nil, err
	} else {
		defer func(rows *sql.Rows) {
			_ = rows.Close()
		}(rows)
		if rows.Next() {
			var user *pojo.User = &pojo.User{
				Id:            -1,
				Name:          name,
				Salt:          nil,
				Password:      nil,
				AccessToken:   nil,
				Admin:         false,
				CreateTime:    nil,
				LastLoginTime: nil,
			}
			if err := rows.Scan(&user.Id, &user.Salt, &user.Password, &user.AccessToken, &user.Admin, &user.CreateTime, &user.LastLoginTime); err != nil {
				return nil, err
			} else {
				return user, nil
			}
		} else {
			return nil, nil
		}
	}
}
