package main

// Options passed to api
type Options struct {
	Format     string `validate:"eq=gpx|eq=geojson|eq=kmz|eq=kml|eq=gml|eq=csv|eq=xml|eq=rss"`
	TextFormat string `validate:"eq=bbcode|eq=texte|eq=markdown|eq=html"`
	NbComs     string
	NbPoints   string
	Detail     string `validate:"eq=simple|eq=complet"`
	PointType  string `validate:"eq=cabane|eq=refuge|eq=gite|eq=pt_eau|eq=sommet|eq=pt_passage|eq=bivouac|eq=lac|eq=all"`
}
