import { IsInt, Min, Max} from "class-validator";

export class Bounds {

  @IsInt()
  @Min(-90)
  @Max(90)
  private west : number;

  @IsInt()
  @Min(-180)
  @Max(180)
  private south : number;

  @IsInt()
  @Min(-90)
  @Max(90)
  private east : number;

  @IsInt()
  @Min(-180)
  @Max(180)
  private north : number;

  constructor(
    west : number,
    south : number,
    east : number,
    north : number
  )  {
      this.west = west;
      this.south = south;
      this.east = east;
      this.north = north;
    }

  public toString() : string {
    return this.west + ',' + this.south + ',' + this.east + ',' + this.north;
  }
}
