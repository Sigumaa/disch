package disch

import (
	"errors"
	"fmt"
	"unicode"
)

// this is a sinmple library.

// Converts the given text to Discord Emojis.
// example: "Hello World" -> ":regional_indicator_h: :regional_indicator_e: :regional_indicator_l: :regional_indicator_l: :regional_indicator_o: :regional_indicator_w: :regional_indicator_o: :regional_indicator_r: :regional_indicator_l: :regional_indicator_d:"
// It is Case Insensitive.

// When Hiragana ,Katakana is entered, it is automatically interpreted as romaji and converted to emoji.
// example: "こんにちは" -> ":regional_indicator_k: :regional_indicator_o: :regional_indicator_n: :regional_indicator_n: :regional_indicator_i: :regional_indicator_c: :regional_indicator_h: :regional_indicator_i: :regional_indicator_w: :regional_indicator_a:"

func isHiragana(r rune) bool {
	return 0x3040 <= r && r <= 0x309f
}

func isSpace(r rune) bool {
	return r == ' '
}

func isKatakana(r rune) bool {
	return 0x30a0 <= r && r <= 0x30ff
}

func isAlphabet(r rune) bool {
	return ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z')
}

func isNumber(r rune) bool {
	return '0' <= r && r <= '9'
}

func hikaToEmoji(r rune) (string, error) {
	switch r {
	case 'あ', 'ア':
		return ":regional_indicator_a:", nil
	case 'い', 'イ':
		return ":regional_indicator_i:", nil
	case 'う', 'ウ':
		return ":regional_indicator_u:", nil
	case 'え', 'エ':
		return ":regional_indicator_e:", nil
	case 'お', 'オ':
		return ":regional_indicator_o:", nil
	case 'か', 'カ':
		return ":regional_indicator_k: :regional_indicator_a:", nil
	case 'き', 'キ':
		return ":regional_indicator_k: :regional_indicator_i:", nil
	case 'く', 'ク':
		return ":regional_indicator_k: :regional_indicator_u:", nil
	case 'け', 'ケ':
		return ":regional_indicator_k: :regional_indicator_e:", nil
	case 'こ', 'コ':
		return ":regional_indicator_k: :regional_indicator_o:", nil
	case 'さ', 'サ':
		return ":regional_indicator_s: :regional_indicator_a:", nil
	case 'し', 'シ':
		return ":regional_indicator_s: :regional_indicator_h: :regional_indicator_i:", nil
	case 'す', 'ス':
		return ":regional_indicator_s: :regional_indicator_u:", nil
	case 'せ', 'セ':
		return ":regional_indicator_s: :regional_indicator_e:", nil
	case 'そ', 'ソ':
		return ":regional_indicator_s: :regional_indicator_o:", nil
	case 'た', 'タ':
		return ":regional_indicator_t: :regional_indicator_a:", nil
	case 'ち', 'チ':
		return ":regional_indicator_c: :regional_indicator_h: :regional_indicator_i:", nil
	case 'つ', 'ツ':
		return ":regional_indicator_t: :regional_indicator_u:", nil
	case 'て', 'テ':
		return ":regional_indicator_t: :regional_indicator_e:", nil
	case 'と', 'ト':
		return ":regional_indicator_t: :regional_indicator_o:", nil
	case 'な', 'ナ':
		return ":regional_indicator_n: :regional_indicator_a:", nil
	case 'に', 'ニ':
		return ":regional_indicator_n: :regional_indicator_i:", nil
	case 'ぬ', 'ヌ':
		return ":regional_indicator_n: :regional_indicator_u:", nil
	case 'ね', 'ネ':
		return ":regional_indicator_n: :regional_indicator_e:", nil
	case 'の', 'ノ':
		return ":regional_indicator_n: :regional_indicator_o:", nil
	case 'は', 'ハ':
		return ":regional_indicator_h: :regional_indicator_a:", nil
	case 'ひ', 'ヒ':
		return ":regional_indicator_h: :regional_indicator_i:", nil
	case 'ふ', 'フ':
		return ":regional_indicator_h: :regional_indicator_u:", nil
	case 'へ', 'ヘ':
		return ":regional_indicator_h: :regional_indicator_e:", nil
	case 'ほ', 'ホ':
		return ":regional_indicator_h: :regional_indicator_o:", nil
	case 'ま', 'マ':
		return ":regional_indicator_m: :regional_indicator_a:", nil
	case 'み', 'ミ':
		return ":regional_indicator_m: :regional_indicator_i:", nil
	case 'む', 'ム':
		return ":regional_indicator_m: :regional_indicator_u:", nil
	case 'め', 'メ':
		return ":regional_indicator_m: :regional_indicator_e:", nil
	case 'も', 'モ':
		return ":regional_indicator_m: :regional_indicator_o:", nil
	case 'や', 'ヤ':
		return ":regional_indicator_y: :regional_indicator_a:", nil
	case 'ゆ', 'ユ':
		return ":regional_indicator_y: :regional_indicator_u:", nil
	case 'よ', 'ヨ':
		return ":regional_indicator_y: :regional_indicator_o:", nil
	case 'ら', 'ラ':
		return ":regional_indicator_r: :regional_indicator_a:", nil
	case 'り', 'リ':
		return ":regional_indicator_r: :regional_indicator_i:", nil
	case 'る', 'ル':
		return ":regional_indicator_r: :regional_indicator_u:", nil
	case 'れ', 'レ':
		return ":regional_indicator_r: :regional_indicator_e:", nil
	case 'ろ', 'ロ':
		return ":regional_indicator_r: :regional_indicator_o:", nil
	case 'わ', 'ワ':
		return ":regional_indicator_w: :regional_indicator_a:", nil
	case 'を', 'ヲ':
		return ":regional_indicator_w: :regional_indicator_o:", nil
	case 'ん', 'ン':
		return ":regional_indicator_n:", nil
	default:
		return "", errors.New("invalid character")

	}
}

func numToEmoji(r rune) (string, error) {
	switch r {
	case '0':
		return ":zero:", nil
	case '1':
		return ":one:", nil
	case '2':
		return ":two:", nil
	case '3':
		return ":three:", nil
	case '4':
		return ":four:", nil
	case '5':
		return ":five:", nil
	case '6':
		return ":six:", nil
	case '7':
		return ":seven:", nil
	case '8':
		return ":eight:", nil
	case '9':
		return ":nine:", nil
	default:
		return "", errors.New("not number")
	}
}

func Convert(text string) (res string, err error) {
	for _, r := range text {
		if isHiragana(r) || isKatakana(r) {
			emoji, err := hikaToEmoji(r)
			if err != nil {
				return "", err
			}
			res += emoji
		} else if isNumber(r) {
			emoji, err := numToEmoji(r)
			if err != nil {
				return "", err
			}
			res += emoji
		} else if isAlphabet(r) {
			r = unicode.ToLower(r)
			emoji := fmt.Sprintf(":regional_indicator_%c:", r)
			res += emoji
		} else if isSpace(r) {
			res += " "
		} else {
			return "", errors.New("invalid character")
		}
	}
	return res, nil
}
