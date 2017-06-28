package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli"
	validator "gopkg.in/go-playground/validator.v9"
)

func bbox(bbox string, options *Options) {

	req, err := http.NewRequest("GET", apiBaseURL+"/bbox", nil)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("bbox", bbox)
	q.Add("format", options.Format)
	q.Add("format_texte", options.TextFormat)
	q.Add("nb_coms", options.NbComs)
	q.Add("nb_points", options.NbPoints)
	q.Add("detail", options.Detail)
	q.Add("type_points", options.PointType)
	req.URL.RawQuery = q.Encode()

	resp, err := http.Get(req.URL.String())
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	stream := os.Stdout
	if output != "" {
		stream, err = os.Create(output)
	}

	io.Copy(stream, resp.Body)
}

var bboxCmd = cli.Command{
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
			cli.ShowSubcommandHelp(c)
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
}
