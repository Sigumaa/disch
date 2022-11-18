package disch

import "testing"

func TestConvert(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
		wantErr bool
	}{
		{
			name:    "いろいろまじってるやつ",
			args:    args{text: "あア1"},
			wantRes: ":regional_indicator_a: :regional_indicator_a: :one: ",
			wantErr: false,
		},
		{
			name:    "空白があると",
			args:    args{text: "あ ア 1"},
			wantRes: ":regional_indicator_a:  :regional_indicator_a:  :one: ",
			wantErr: false,
		},
		{
			name:    "ちょんちょん",
			args:    args{text: "ばか"},
			wantRes: ":regional_indicator_b: :regional_indicator_a: :regional_indicator_k: :regional_indicator_a: ",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := Convert(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRes != tt.wantRes {
				t.Errorf("Convert() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_hikaToEmoji(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "ひらがな",
			args:    args{r: 'あ'},
			want:    ":regional_indicator_a:",
			wantErr: false,
		},
		{
			name:    "カタカナ",
			args:    args{r: 'ア'},
			want:    ":regional_indicator_a:",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := hikaToEmoji(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("hikaToEmoji() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("hikaToEmoji() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isAlphabet(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "アルファベットカナ！？",
			args: args{r: 'a'},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAlphabet(tt.args.r); got != tt.want {
				t.Errorf("isAlphabet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isHiragana(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ひらがな、なの、カナ！？",
			args: args{r: 'あ'},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHiragana(tt.args.r); got != tt.want {
				t.Errorf("isHiragana() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isKatakana(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "カタカナって...コト！？",
			args: args{r: 'ア'},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isKatakana(tt.args.r); got != tt.want {
				t.Errorf("isKatakana() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isNumber(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "数字って...コト！？",
			args: args{r: '1'},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber(tt.args.r); got != tt.want {
				t.Errorf("isNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSpace(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "スペースって...コト！？",
			args: args{r: ' '},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSpace(tt.args.r); got != tt.want {
				t.Errorf("isSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

// !
func Test_isQuestion(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "？って...コト！？",
			args: args{r: '?'},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isQuestion(tt.args.r); got != tt.want {
				t.Errorf("isQuestion() = %v, want %v", got, tt.want)
			}
		})
	}

}

func Test_isExclamation(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "！って...コト！？",
			args: args{r: '!'},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isExclamation(tt.args.r); got != tt.want {
				t.Errorf("isExclamation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numToEmoji(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "数字をかえちゃうよっ！",
			args:    args{r: '1'},
			want:    ":one:",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := numToEmoji(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("numToEmoji() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("numToEmoji() got = %v, want %v", got, tt.want)
			}
		})
	}
}
