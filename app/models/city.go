package models

type City struct {
	OGCFID      NullInt64  `db:"ogc_fid" json:"ogc_fid"`
	WbkGeometry byte       `db:"wkb_geometry" json:"wkb_geometry"`
	Region      NullString `db:"region" json:"region"`
	Province    NullString `db:"province" json:"province"`
	Regency     NullString `db:"regency" json:"regency"`
	Name        NullString `db:"name" json:"name"`
	Type        NullString `db:"type" json:"type"`
	ProvCap     NullString `db:"prov_cap" json:"prov_cap"`
	RegCap      NullString `db:"reg_cap" json:"reg_cap"`
	Kota        NullString `db:"kota" json:"kota"`
	Longitude   NullString `db:"longitude" json:"longitude"`
	Latitude    NullString `db:"latitude" json:"latitude"`
	Source      NullString `db:"source" json:"source"`
	GeoCodDate  NullTime   `db:"geocoddate" json:"geocoddate"`
	GeoCodeBy   NullString `db:"geocodedby" json:"geocodedby"`
	Notes       NullString `db:"notes" json:"notes"`
}
