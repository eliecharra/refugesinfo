refugesinfo
===========

Refuges.info CLI api client

Usage
-----

```
$ ./refugesinfo

NAME:
   refugesinfo - Refuges.info API CLI tool

USAGE:
   refugesinfo [global options] command [command options] [arguments...]

VERSION:
   e9a8a7c

COMMANDS:
     bbox     Export des points contenus dans une bbox
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --output value  Give a file name to save output
   --help, -h      show help
   --version, -v   print the version
```

Bbox
----

```
➜ ./refugesinfo bbox --help  
NAME:
   refugesinfo bbox - Export des points contenus dans une bbox

USAGE:
   refugesinfo bbox [command options] [arguments...]

OPTIONS:
   --format value       Le format de l'export. (default: "geojson")
   --text-format value  Le formatage du texte, que vous devrez retravailler par la suite. (default: "bbcode")
   --nb-coms value      Nombre de commentaires maximum à retourner. Retourne aucun commentaire pour 0. Retourne les n commentaires les plus récents. (default: 0)
   --nb-points value    Nombre de points à exporter (le choix est fait par un algorithme interne avec prioritées élevées pour les abris et cabanes, et faibles pour refuges, sommets, cols...). 0 retournera tous les points de la zone, mais à utiliser avec précautions (lecture illisible et charge serveur importante). (default: 121)
   --detail value       Les détails du point, par défaut uniquement long, lat, altitude, nom, type, id et lien. complet est disponible uniquement lorsque format est geojson, xml. Aussi disponible en gpx (pour avoir un fichier léger) et en rss (complet conseillé pour afficher les remarques diverses). (default: "simple")
   --point-type value   Les types de point à exporter, parmis la liste suivante : cabane, refuge, gite, pt_eau, sommet, pt_passage, bivouac et lac (default: "all")
```
