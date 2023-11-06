package domain

import "image/color"

type ColorSettings struct {
	UserID  int64    `bson:"userId"`
	Pallets []Pallet `bson:"pallets"`
}

type Pallet struct {
	ID     int64           `bson:"id"`
	Colors []AltitudeColor `bson:"colors"`
}

type AltitudeColor struct {
	Altitude float64    `bson:"altitude"`
	Color    color.RGBA `bson:"inline"`
}
