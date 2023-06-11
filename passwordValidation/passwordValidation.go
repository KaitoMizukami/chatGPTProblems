/*

問題: "ユーザーが新しいアカウントを作成する際に、パスワードの強度を検証するサーバーサイドの機能を実装してください。
パスワードは、以下の条件を満たす必要があります。

少なくとも8文字以上であること
少なくとも1つの大文字と1つの小文字を含んでいること
少なくとも1つの数字を含んでいること
少なくとも1つの特殊文字（例: !@#$%^&*）を含んでいること
ユーザーが提供したパスワードが条件を満たしているかどうかを検証する関数 validatePassword(password string) bool を作成し、
適切なテストケースで検証してください。また、パスワードが条件を満たしていない場合には適切なエラーメッセージを返すようにしてください。"

*/

package passwordValidation

import (
	"regexp"
	"strings"
	"unicode"
)

func hasUpperAndLowerAndSpecial(password string) bool {
	hasUpper := false
	hasLower := false
	hasSpecial := false

	/* chatGPTの例 こっちの方が可読性高いかも

	switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case strings.ContainsAny(string(r), "!@#$%^&*"):
			hasSpecial = true
		}

	*/
	for _, r := range password {
		if unicode.IsUpper(r) {
			hasUpper = true
		} else if unicode.IsLower(r) {
			hasLower = true
		} else if strings.ContainsAny(string(r), "!@#$%^&*") {
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasSpecial
}

// hasNumber -> hasDigitの方が適切
func hasDigit(password string) bool {
	return regexp.MustCompile(`\d`).MatchString(password)
}

func validatePassword(password string) bool {
	if len(password) <= 8 {
		return false
	}
	if !hasUpperAndLowerAndSpecial(password) {
		return false
	}
	if !hasDigit(password) {
		return false
	}

	return true
}
