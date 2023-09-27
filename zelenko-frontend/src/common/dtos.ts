export interface ICreateGreenObject {
	LocationName: string;
    Latitude: number;
	Longitude: number;
	Street: string;
	City: string;
	Country: string;
	Shape: string;
	TrashType: string;
}

export interface IUpdateGreenObject extends ICreateGreenObject{
	ID: string;
}

export interface IGreenObject {
    ID: string;
	LocationName: string;
	Location: ILocation;
	Shape: string;
	TrashType: string;
	GreenScore: IGreenScore;
	Disabled: boolean;
}

export interface IGreenScore {
    ID: string;
	Verification: number;
	Report: number;
	TrashRank: string;
}

export interface ILocation {
    ID: string;
	Latitude: number;
	Longitude: number;
	Street: string;
	City: string;
	Country: string;
}

export interface IFeatureInfo {
    id: string;
    coords: number[];
}