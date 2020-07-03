package v1_test

import (
	"reflect"
	"testing"

	random "github.com/erdaltsksn/random/v1"
)

func TestGetRandomCardType(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Card Type"},
	}
	for _, tt := range tests {
		for i := 0; i < 1000; i++ {
			got := random.GetRandomCardType()
			t.Run(tt.name, func(t *testing.T) {
				if !inSupportedCards(got) {
					t.Error("Got:", got, "Want:", random.SupportedCards)
				}
			})
		}
	}
}

func inSupportedCards(c random.CardType) bool {
	for _, v := range random.SupportedCards {
		if reflect.DeepEqual(v, c) {
			return true
		}
	}
	return false
}

func TestGetCardType(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want random.CardType
	}{
		{"AmericanExpress",
			args{name: "AmericanExpress"},
			random.CardType{
				Name:   "AmericanExpress",
				Length: []int{15},
				Prefix: []string{"34", "37"},
			},
		},
		{"Laser",
			args{name: "Laser"},
			random.CardType{
				Name:   "Laser",
				Length: []int{16, 17, 18, 19},
				Prefix: []string{
					"6304", "6706", "6709", "6771",
				},
			},
		},
		{"Mir",
			args{name: "Mir"},
			random.CardType{

				Name:   "Mir",
				Length: []int{13, 16},
				Prefix: []string{
					"2200", "2201", "2202", "2203", "2204",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := random.GetCardType(tt.args.name)
			if !reflect.DeepEqual(got, tt.want) {
				t.Error("Got:", got, "Want:", random.SupportedCards)
			}
		})
	}
}

func TestCardType_Generate(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name       string
		args       args
		wantPrefix []string
		wantLength []int
	}{
		{"AmericanExpress",
			args{name: "AmericanExpress"},
			[]string{
				"34", "37",
			},
			[]int{15},
		},
		{"JCB",
			args{name: "JCB"},
			[]string{
				"353", "354", "355", "356", "357", "358", "1800", "2131", "3528",
				"3529",
			},
			[]int{15, 16},
		},
		{"Visa",
			args{name: "Visa"}, []string{"4"}, []int{13, 16, 19},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 1000; i++ {
				got := random.GetCardType(tt.args.name).Generate()
				if !hasCCPrefix(got, tt.wantPrefix) || !hasCCLength(got, tt.wantLength) {
					t.Error("Got:", got, "Want:", "prefix:", tt.wantPrefix, "length:", tt.wantLength)
				}
			}
		})
	}
}

func hasCCPrefix(text string, prefix []string) bool {
	for _, v := range prefix {
		if len(text) >= len(v) && text[0:len(v)] == v {
			return true
		}
	}

	return false
}

func hasCCLength(text string, length []int) bool {
	for _, v := range length {
		if len(text) == v {
			return true
		}
	}

	return false
}
