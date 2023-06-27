/*

以下の要件を満たすユーザー情報の管理プログラムを作成してください。

User構造体を定義し、以下のフィールドを持つようにしてください:

ID: ユーザーの一意の識別子 (文字列)
Name: ユーザー名 (文字列)
Email: メールアドレス (文字列)
UserManager構造体を定義し、以下のメソッドを持つようにしてください:

AddUser(user User): 引数として与えられたユーザーをユーザーリストに追加します。
RemoveUser(id string): 指定されたIDのユーザーをユーザーリストから削除します。
GetUser(id string) (User, error): 指定されたIDのユーザーを取得します。ユーザーが存在しない場合はエラーを返します。
ListUsers() []User: 現在のユーザーリストを返します。

*/

package main

import (
	"errors"
)

type User struct {
	ID    string
	Name  string
	Email string
}

type UserManager struct {
	users map[string]User
}

func (um *UserManager) AddUser(user User) error {
	// 重複するIDのユーザーが存在する場合はエラーを返す
	if _, ok := um.users[user.ID]; ok {
		return errors.New("user already exists")
	}

	um.users[user.ID] = user
	return nil
}

func (um *UserManager) RemoveUser(id string) error {
	// 指定されたIDのユーザーが存在しない場合はエラーを返す
	if _, ok := um.users[id]; !ok {
		return errors.New("user not found")
	}

	delete(um.users, id)
	return nil
}

func (um *UserManager) GetUser(id string) (User, error) {
	// 指定されたIDのユーザーが存在しない場合はエラーを返す
	if user, ok := um.users[id]; ok {
		return user, nil
	}

	return User{}, errors.New("user not found")
}

func (um *UserManager) ListUsers() []User {
	users := make([]User, 0, len(um.users))

	for _, user := range um.users {
		users = append(users, user)
	}

	return users
}
