package random

import (
	"math/rand"
	"strconv"
	"time"
)

// CardType holds the card type information.
type CardType struct {
	Name   string
	Length []int
	Prefix []string
}

// SupportedCards holds all supported card list.
var SupportedCards = []CardType{
	{
		Name:   "AmericanExpress",
		Length: []int{15},
		Prefix: []string{
			"34", "37",
		},
	},
	{
		Name:   "DinersClub",
		Length: []int{14},
		Prefix: []string{
			"36",
			"300", "301", "302", "303", "304", "305",
		},
	},
	{
		Name:   "DinersClubUS",
		Length: []int{16},
		Prefix: []string{
			"54", "55",
		},
	},
	{
		Name:   "Discover",
		Length: []int{16, 19},
		Prefix: []string{
			"65",
			"644", "645", "646", "647", "648", "649",
			"6011", "6222", "6223", "6224", "6225", "6226", "6227", "6228",
			"62213", "62214", "62215", "62216", "62217", "62218", "62219",
			"62290", "62291",
			"622126", "622127", "622128", "622129", "622920", "622921",
			"622922", "622923", "622924", "622925",
		},
	},
	{
		Name:   "JCB",
		Length: []int{15, 16},
		Prefix: []string{
			"353", "354", "355", "356", "357", "358", "1800", "2131", "3528",
			"3529",
		},
	},
	{
		Name:   "Laser",
		Length: []int{16, 17, 18, 19},
		Prefix: []string{
			"6304", "6706", "6709", "6771",
		},
	},
	{
		Name:   "Maestro",
		Length: []int{12, 13, 14, 15, 16, 17, 18, 19},
		Prefix: []string{
			"5018", "5020", "5038", "6304", "6759", "6761", "6762", "6763",
			"6764", "6765", "6766", "6772",
		},
	},
	{
		Name:   "Mastercard",
		Length: []int{16},
		Prefix: []string{
			"23", "24", "25", "26", "51", "52", "53", "54", "55",
			"223", "224", "225", "226", "227", "228", "229", "271",
			"2221", "2222", "2223", "2224", "2225", "2226", "2227", "2228",
			"2229", "2720",
		},
	},
	{
		Name:   "Solo",
		Length: []int{16, 18, 19},
		Prefix: []string{
			"6334", "6767",
		},
	},
	{
		Name:   "Unionpay",
		Length: []int{16, 17, 18, 19},
		Prefix: []string{
			"6222", "6223", "6224", "6225", "6226", "6227", "6228",
			"62213", "62214", "62215", "62216", "62217", "62218", "62219",
			"62290", "62291",
			"622126", "622127", "622128", "622129", "622920", "622921",
			"622922", "622923", "622924", "622925",
		},
	},
	{
		Name:   "Visa",
		Length: []int{13, 16, 19},
		Prefix: []string{
			"4",
		},
	},
	{
		Name:   "Mir",
		Length: []int{13, 16},
		Prefix: []string{
			"2200", "2201", "2202", "2203", "2204",
		},
	},
}

// GetRandomCardType return a random card type.
func GetRandomCardType() CardType {
	rand.Seed(time.Now().Unix())
	randType := rand.Intn(len(SupportedCards))

	return SupportedCards[randType]
}

// GetCardType returns card type named or random card type.
func GetCardType(name string) CardType {
	for _, v := range SupportedCards {
		if v.Name == name {
			return v
		}
	}

	return GetRandomCardType()
}

// Generate returns card information according to Card Type.
func (c CardType) Generate() string {
	rand.Seed(time.Now().Unix())

	length := c.Length[rand.Intn(len(c.Length))]
	prefix := c.Prefix[rand.Intn(len(c.Prefix))]

	card := prefix
	randomNumberLength := length - (len(prefix) + 1)
	for i := 0; i < randomNumberLength; i++ {
		rand := rand.Intn(10)
		card += strconv.Itoa(rand)
	}

	// Generates last digit according to Luhn algorithm
	sum := 0
	for i := 0; i < len(card); i++ {
		digit, _ := strconv.Atoi((string(card[i])))

		if i%2 == 0 {
			digit = digit * 2

			if digit > 9 {
				digit = (digit / 10) + (digit % 10)
			}
		}

		sum += digit
	}

	mod := sum % 10
	if mod == 0 {
		return card + strconv.Itoa(0)
	}

	return card + strconv.Itoa(10-mod)
}
