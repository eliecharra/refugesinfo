import request = require('request');
import config = require('config');

import { Bounds, Format, TextFormat, Detail, PointType } from "./types";

export class RefugesApi {

  private baseUrl : any;

  constructor(){
    this.baseUrl = config.get('api.baseUrl');
  }

  public bbox (
      bbox : Bounds,
      format: Format = Format.geojson,
      textFormat: TextFormat = TextFormat.bbcode,
      nbComs: number = 0,
      nbPoints: string|number = 121,
      detail: Detail = Detail.simple,
      pointTypes : PointType[] = [PointType.all]
    ) : void {
      let params = {
        bbox: bbox.toString(),
        format,
        format_texte: textFormat,
        nb_coms: nbComs,
        nb_points: nbPoints,
        detail,
        type_points: pointTypes.join()
      }

      request.get({url: this.baseUrl + 'bbox', qs: params},
        function(error, response, body){
          console.log(body);
        }
      );
  }
}
