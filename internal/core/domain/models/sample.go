package models

import "time"

type Sample struct {
	// サンプルの一意識別子
	ID string
	// 文字列値
	StringVal string
	// 整数値
	IntVal int
	// 文字列配列
	ArrayVal []string
	// メールアドレス
	Email string
	// 作成日時
	CreatedAt time.Time
	// 更新日時
	UpdatedAt time.Time
}
