package dto

import "time"

type Sample struct {
	// サンプルの一意識別子
	ID string `json:"id"`
	// 文字列値
	StringVal string `json:"stringVal"`
	// 整数値
	IntVal int `json:"intVal"`
	// 文字列配列
	ArrayVal []string `json:"arrayVal"`
	// メールアドレス
	Email string `json:"email"`
	// 作成日時
	CreatedAt time.Time `json:"createdAt"`
	// 更新日時
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateSampleInput struct {
	// 文字列値
	StringVal string `json:"stringVal"`
	// 整数値
	IntVal int `json:"intVal"`
	// 文字列配列
	ArrayVal []string `json:"arrayVal"`
	// メールアドレス
	Email string `json:"email"`
}

type CreateSamplePayload struct {
	// 作成されたサンプル
	Sample *Sample `json:"sample"`
}
