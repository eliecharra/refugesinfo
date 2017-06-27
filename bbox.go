package main

import (
	"io"
	"log"
	"net/http"
	"os"
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
