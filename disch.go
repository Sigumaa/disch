package disch

import (
	"errors"
)

// this is a sinmple library.

// Converts the given text to Discord Emojis.
// example: "Hello World" -> ":regional_indicator_h: :regional_indicator_e: :regional_indicator_l: :regional_indicator_l: :regional_indicator_o: :regional_indicator_w: :regional_indicator_o: :regional_indicator_r: :regional_indicator_l: :regional_indicator_d:"
// It is Case Insensitive.

// When Hiragana ,Katakana is entered, it is automatically interpreted as romaji and converted to emoji.
// example: "こんにちは" -> ":regional_indicator_k: :regional_indicator_o: :regional_indicator_n: :regional_indicator_n: :regional_indicator_i: :regional_indicator_c: :regional_indicator_h: :regional_indicator_i: :regional_indicator_w: :regional_indicator_a:"

var (
	ErrIC = errors.New("invalid character")
	ErrNN = errors.New("not number")
)

func isHiragana(r rune) bool {
	return 0x3040 <= r && r <= 0x309f
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

func isSpace(r rune) bool {
	return r == ' '
}

// ! ?
func isExclamation(r rune) bool {
	return r == '!' || r == '！'
}

func isQuestion(r rune) bool {
	return r == '?' || r == '？'
}

func hikaToEmoji(r rune) (string, error) {
	switch r {
	case 'あ', 'ア':
		return a, nil
	case 'い', 'イ':
		return i, nil
	case 'う', 'ウ':
		return u, nil
	case 'え', 'エ':
		return e, nil
	case 'お', 'オ':
		return o, nil
	case 'か', 'カ':
		return ka, nil
	case 'き', 'キ':
		return ki, nil
	case 'く', 'ク':
		return ku, nil
	case 'け', 'ケ':
		return ke, nil
	case 'こ', 'コ':
		return ko, nil
	case 'さ', 'サ':
		return sa, nil
	case 'し', 'シ':
		return shi, nil
	case 'す', 'ス':
		return su, nil
	case 'せ', 'セ':
		return se, nil
	case 'そ', 'ソ':
		return so, nil
	case 'た', 'タ':
		return ta, nil
	case 'ち', 'チ':
		return chi, nil
	case 'つ', 'ツ':
		return tsu, nil
	case 'て', 'テ':
		return te, nil
	case 'と', 'ト':
		return to, nil
	case 'な', 'ナ':
		return na, nil
	case 'に', 'ニ':
		return ni, nil
	case 'ぬ', 'ヌ':
		return nu, nil
	case 'ね', 'ネ':
		return ne, nil
	case 'の', 'ノ':
		return no, nil
	case 'は', 'ハ':
		return ha, nil
	case 'ひ', 'ヒ':
		return hi, nil
	case 'ふ', 'フ':
		return hu, nil
	case 'へ', 'ヘ':
		return he, nil
	case 'ほ', 'ホ':
		return ho, nil
	case 'ま', 'マ':
		return ma, nil
	case 'み', 'ミ':
		return mi, nil
	case 'む', 'ム':
		return mu, nil
	case 'め', 'メ':
		return me, nil
	case 'も', 'モ':
		return mo, nil
	case 'や', 'ヤ':
		return ya, nil
	case 'ゆ', 'ユ':
		return yu, nil
	case 'よ', 'ヨ':
		return yo, nil
	case 'ら', 'ラ':
		return ra, nil
	case 'り', 'リ':
		return ri, nil
	case 'る', 'ル':
		return ru, nil
	case 'れ', 'レ':
		return re, nil
	case 'ろ', 'ロ':
		return ro, nil
	case 'わ', 'ワ':
		return wa, nil
	case 'を', 'ヲ':
		return wo, nil
	case 'ん', 'ン':
		return n, nil
	case 'が', 'ガ':
		return ga, nil
	case 'ぎ', 'ギ':
		return gi, nil
	case 'ぐ', 'グ':
		return gu, nil
	case 'げ', 'ゲ':
		return ge, nil
	case 'ご', 'ゴ':
		return goo, nil
	case 'ざ', 'ザ':
		return za, nil
	case 'じ', 'ジ':
		return zi, nil
	case 'ず', 'ズ':
		return zu, nil
	case 'ぜ', 'ゼ':
		return ze, nil
	case 'ぞ', 'ゾ':
		return zo, nil
	case 'だ', 'ダ':
		return da, nil
	case 'ぢ', 'ヂ':
		return zi, nil
	case 'づ', 'ヅ':
		return zu, nil
	case 'で', 'デ':
		return de, nil
	case 'ど', 'ド':
		return do, nil
	case 'ば', 'バ':
		return ba, nil
	case 'び', 'ビ':
		return bi, nil
	case 'ぶ', 'ブ':
		return bu, nil
	case 'べ', 'ベ':
		return be, nil
	case 'ぼ', 'ボ':
		return bo, nil
	case 'ぱ', 'パ':
		return pa, nil
	case 'ぴ', 'ピ':
		return pi, nil
	case 'ぷ', 'プ':
		return pu, nil
	case 'ぺ', 'ペ':
		return pe, nil
	case 'ぽ', 'ポ':
		return po, nil
	default:
		return "", ErrIC

	}
}

func numToEmoji(r rune) (string, error) {
	switch r {
	case '0':
		return zero, nil
	case '1':
		return one, nil
	case '2':
		return two, nil
	case '3':
		return three, nil
	case '4':
		return four, nil
	case '5':
		return five, nil
	case '6':
		return six, nil
	case '7':
		return seven, nil
	case '8':
		return eight, nil
	case '9':
		return nine, nil
	default:
		return "", ErrNN
	}
}

func alToEmoji(rr rune) (string, error) {
	switch rr {
	case 'a', 'A':
		return a, nil
	case 'b', 'B':
		return b, nil
	case 'c', 'C':
		return c, nil
	case 'd', 'D':
		return d, nil
	case 'e', 'E':
		return e, nil
	case 'f', 'F':
		return f, nil
	case 'g', 'G':
		return g, nil
	case 'h', 'H':
		return h, nil
	case 'i', 'I':
		return i, nil
	case 'j', 'J':
		return j, nil
	case 'k', 'K':
		return k, nil
	case 'l', 'L':
		return l, nil
	case 'm', 'M':
		return m, nil
	case 'n', 'N':
		return n, nil
	case 'o', 'O':
		return o, nil
	case 'p', 'P':
		return p, nil
	case 'q', 'Q':
		return q, nil
	case 'r', 'R':
		return r, nil
	case 's', 'S':
		return s, nil
	case 't', 'T':
		return t, nil
	case 'u', 'U':
		return u, nil
	case 'v', 'V':
		return v, nil
	case 'w', 'W':
		return w, nil
	case 'x', 'X':
		return x, nil
	case 'y', 'Y':
		return y, nil
	case 'z', 'Z':
		return z, nil
	default:
		return "", ErrIC
	}
}

func Convert(text string) (res string, err error) {
	for _, r := range text {
		if isHiragana(r) || isKatakana(r) {
			emoji, err := hikaToEmoji(r)
			if err != nil {
				return "", err
			}
			res += emoji + " "
		} else if isNumber(r) {
			emoji, err := numToEmoji(r)
			if err != nil {
				return "", err
			}
			res += emoji + " "
		} else if isAlphabet(r) {
			emoji, err := alToEmoji(r)
			if err != nil {
				return "", err
			}
			res += emoji + " "
		} else if isSpace(r) {
			res += " "
		} else if isExclamation(r) {
			res += exclamation + " "
		} else if isQuestion(r) {
			res += question + " "
		} else {
			return "", ErrIC
		}
	}

	return res, nil
}
