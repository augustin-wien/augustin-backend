package structs

import (
	"github.com/jackc/pgx/v5/pgtype"
)

// Attributes have to be uppercase to be exported
// TODO: I would avoid using different json names

type Settings struct {
	ID   int32
	Color string
	Logo  string
	Items  []Item
}

type Item struct {
	ID   int32
	Name  string
	Price float32
}

type Vendor struct {
	Credit   float64 `json:"credit"`
	QRcode   string  `json:"qrcode"`
	IDnumber string  `json:"idnumber"`
}

type Account struct {
	ID   pgtype.Int4
	Name pgtype.Text
}

type PaymentType struct {
	ID   pgtype.Int4
	Name pgtype.Text
}

type Payment struct {
	ID        pgtype.Int8
	Timestamp pgtype.Timestamp
	Sender    pgtype.Int4
	Receiver  pgtype.Int4
	Type      pgtype.Int4
	Amount    pgtype.Float4
}

type PaymentBatch struct {
	Payments []Payment
}
