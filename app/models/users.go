package models

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// モデル
// 構造体定義、DBにアクセスする関数、レシーバメソッドを記載する。

// ユーザーリクエスト構造体
type UserReq struct {
	UserId	string	`json:"userId"`
	Password	string	`json:"password"`
}

// ユーザーレスポンス構造体
type UserRes struct {
	UserId	string	`json:"userId"`
	SessionId	string	`json:"sessionId"`
	CreatedAt	time.Time	`json:"createdAt"`
	UpdatedAt	time.Time	`json:"updatedAt"`
}

// ユーザーを取得する
func GetUser(sessionId string) (userRes UserRes, err error) {
	userRes.SessionId = sessionId

	// セッションIDからユーザーIDを取得する
	cmd := `
		select user_id
		from login_sessions
		where session_id = ?
		limit 1
	`

	err = Db.QueryRow(cmd, sessionId).Scan(
		&userRes.UserId,
	)

	if err != nil {
		fmt.Println(err)
		return userRes, err
	}

	// ユーザーIDからユーザーを取得する
	cmd = `
		select created_at, updated_at
		from users
		where user_id = ?
	`

	err = Db.QueryRow(cmd, userRes.UserId).Scan(
		&userRes.CreatedAt,
		&userRes.UpdatedAt,
	)

	return userRes, err
}

// ユーザーを登録する
func (u *UserReq) RegisterUser() (userRes UserRes, err error) {
	cmd := `
		insert into users (user_id, password)
		values (?, ?)
	`

	_, err = Db.Exec(cmd, u.UserId, encrypt(u.Password))
	if err != nil {
		fmt.Println(err)
	}

	// 登録した最新レコードを取得する
	cmd = `
		select user_id, created_at, updated_at
		from users
		order by updated_at desc
		limit 1
	`

	err = Db.QueryRow(cmd).Scan(
		&userRes.UserId,
		&userRes.CreatedAt,
		&userRes.UpdatedAt,
	)

	return userRes, err
}

// ログインする
func Login(userId string) (sessionId string, err error) {
	cmd := `
		insert into login_sessions (session_id, user_id)
		values (?, ?)
	`

	_, err = Db.Exec(cmd, createUuid(), userId)
	if err != nil {
		fmt.Println(err)
	}

	// 登録した最新レコードを取得する
	cmd = `
		select session_id
		from login_sessions
		order by created_at desc
		limit 1
	`

	err = Db.QueryRow(cmd).Scan(&sessionId)

	return sessionId, err
}

// ログインチェックをする
func (u *UserReq) CheckLogin() (isOk bool, userRes UserRes, err error) {
	cmd := `
		select user_id, created_at, updated_at
		from users
		where user_id = ? and password = ?
	`

	err = Db.QueryRow(cmd, u.UserId, encrypt(u.Password)).Scan(
		&userRes.UserId,
		&userRes.CreatedAt,
		&userRes.UpdatedAt,
	)

	if err != nil || userRes.UserId == "" {
		isOk = false
	} else {
		isOk = true
	}

	return isOk, userRes, err
}

// ログアウトする
func Logout(sessionId string) (err error) {
	cmd := `
		delete
		from login_sessions
		where session_id = ?
	`

	_, err = Db.Exec(cmd, sessionId)

	return err
}

// セッションIDからログインチェックをする
func CheckLoginByAccessToken(accessToken string) (isOk bool, err error) {
	var userId string

	// アクセストークンからユーザーIDを取得する
	cmd := `
		select user_id
		from logins
		where uuid = $1
		limit 1
	`

	err = Db.QueryRow(cmd, accessToken).Scan(&userId)
	if err != nil {
		fmt.Println(err)
	}

	if err != nil || userId == "" {
		isOk = false
	} else {
		isOk = true
	}

	return isOk, err
}

// 文字列を暗号化する
func encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}

// UUIDを生成する
func createUuid() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}