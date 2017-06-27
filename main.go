package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"

	"gopkg.in/go-playground/validator.v9"
)

var version string
var validate *validator.Validate
var apiBaseURL = "https://www.refuges.info/api"
var output string

func main() {
	app := cli.NewApp()
	validate = validator.New()

	app.Version = version
	app.Usage = "Refuges.info API CLI tool"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "output",
			Usage:       "Give a file name to save output",
			Destination: &output,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "bbox",
			Usage: "Export des points contenus dans une bbox",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "format",
					Value: "geojson",
					Usage: "Le format de l'export.",
				},
				cli.StringFlag{
					Name:  "text-format",
					Value: "bbcode",
					Usage: "Le formatage du texte, que vous devrez retravailler par la suite.",
				},
				cli.IntFlag{
					Name:  "nb-coms",
					Value: 0,
					Usage: "Nombre de commentaires maximum à retourner. Retourne aucun commentaire pour 0. Retourne les n commentaires les plus récents.",
				},
				cli.IntFlag{
					Name:  "nb-points",
					Value: 121,
					Usage: "Nombre de points à exporter (le choix est fait par un algorithme interne avec prioritées élevées pour les abris et cabanes, et faibles pour refuges, sommets, cols...). 0 retournera tous les points de la zone, mais à utiliser avec précautions (lecture illisible et charge serveur importante).",
				},
				cli.StringFlag{
					Name:  "detail",
					Value: "simple",
					Usage: "Les détails du point, par défaut uniquement long, lat, altitude, nom, type, id et lien. complet est disponible uniquement lorsque format est geojson, xml. Aussi disponible en gpx (pour avoir un fichier léger) et en rss (complet conseillé pour afficher les remarques diverses).",
				},
				cli.StringFlag{
					Name:  "point-type",
					Value: "all",
					Usage: "Les types de point à exporter, parmis la liste suivante : cabane, refuge, gite, pt_eau, sommet, pt_passage, bivouac et lac",
				},
			},
			Action: func(c *cli.Context) error {
				if c.NArg() != 1 {
					return cli.NewExitError("Missing bbox argument", 1)
				}

				options := &Options{
					c.String("format"),
					c.String("text-format"),
					c.String("nb-coms"),
					c.String("nb-points"),
					c.String("detail"),
					c.String("point-type"),
				}

				err := validate.Struct(options)
				if err != nil {
					if _, ok := err.(*validator.InvalidValidationError); ok {
						fmt.Println(err)
					}
					for _, err := range err.(validator.ValidationErrors) {
						return cli.NewExitError("L'option '"+err.Field()+"' est invalide", 1)
					}
				}

				bbox(c.Args().First(), options)

				return nil
			},
		},
	}

	app.Run(os.Args)
}
