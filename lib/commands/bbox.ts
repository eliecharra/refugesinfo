import { Detail, PointType, TextFormat, Format, Bounds } from '../types';
import { Validator } from "class-validator";
import { RefugesApi } from '../api';

const validator = new Validator();

exports.command = 'bbox [west] [south] [east] [north]'
exports.desc = ' Export des points contenus dans une bbox'
exports.builder = function (yargs: any) {
  return yargs.options({
    west: {
      number: 'number',
      demandOption: true
    },
    south: {
      type: 'number',
      demandOption: true
    },
    east: {
      type: 'number',
      demandOption: true
    },
    north: {
      type: 'number',
      demandOption: true
    },
    format: {
      default: Format.geojson,
      choices: ['geojson', 'kmz', 'kml', 'gml', 'gpx', 'csv', 'xml', 'rss'],
      description: 'Le format de l\'export.'
    },
    textFormat: {
      default: TextFormat.bbcode,
      choices: ['bbcode', 'texte', 'markdown', 'html'],
      description: 'Le formatage du texte, que vous devrez retravailler par la suite.'
    },
    nbComs: {
      default: 0,
      type: 'number',
      description: 'Nombre de commentaires maximum à retourner. Retourne aucun commentaire pour 0. Retourne les n commentaires les plus récents.'
    },
    nbPoints: {
      default: 121,
      type: 'number',
      description: 'Nombre de points à exporter (le choix est fait par un algorithme interne avec prioritées élevées pour les abris et cabanes, et faibles pour refuges, sommets, cols...). all retournera tous les points de la zone, mais à utiliser avec précautions (lecture illisible et charge serveur importante).'
    },
    detail: {
      default: Detail.simple,
      choices: ['simple', 'complet'],
      description: 'Les détails du point, par défaut uniquement long, lat, altitude, nom, type, id et lien. complet est disponible uniquement lorsque format est geojson, xml. Aussi disponible en gpx (pour avoir un fichier léger) et en rss (complet conseillé pour afficher les remarques diverses).'
    },
    pointTypes: {
      type: 'array',
      default: [PointType.all],
      description: 'Les types de point à exporter, parmis la liste suivante : cabane, refuge, gite, pt_eau, sommet, pt_passage, bivouac et lac ou leur équivament numérique: 7, 10, 9, 23, 6, 3, 19, 16. La valeur all sélectionne tous les types.'
    },
  }).coerce(['west', 'east'], function(arg: any){
    if (
      validator.isNumber(arg) &&
      validator.max(arg, 90) &&
      validator.min(arg, -90)) {
      return Number(arg);
    }

    throw new Error('Invalid coordinates');
  }).coerce(['south', 'north'], function(arg: any){
    if (
      validator.isNumber(arg) &&
      validator.max(arg, 180) &&
      validator.min(arg, -180)) {
      return Number(arg);
    }

    throw new Error('Invalid coordinates');
  });
}
exports.handler = function (argv: any) {

  let bounds = new Bounds(
    argv.west,
    argv.south,
    argv.east,
    argv.north
  );

  var api = new RefugesApi();

  api.bbox(
    bounds,
    argv.format,
    argv.textFormat,
    argv.nbComs,
    argv.nbPoints,
    argv.detail,
    argv.pointTypes
  );

}
