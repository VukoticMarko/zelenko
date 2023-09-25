export interface ICreateGreenObjectFormData {
	LocationName: string;
    Latitude: number;
	Longitude: number;
	Street: string;
	City: string;
	Country: string;
	Shape: string;
	TrashType: string;
}

export interface IGreenScoreForm {
    ID: string;
	Verification: number;
	Report: number;
	TrashRank: string;
}

export interface ILocationForm {
	Latitude: number;
	Longitude: number;
	Street: string;
	City: string;
	Country: string;
}

export const shapeOptions: string[] = [
    'can',
    'dumpster',
    'underground',
    'animal_garbage',
    'junkyard',
    'city_department',
    'waste_disposal',
    'waste_transfer',
    'waste_process_complex',
    'incineration_plant',
    'recycle_center',
    'landfill'
];

export const trashTypeOptions: string[] = [
    'all',
    'pet',
    'plastic',
    'paper',
    'metal',
    'weee',
    'wood',
    'glass',
    'biodegradable',
    'medical',
    'toxic',
    'nuclear',
    'organic',
    'radioactive',
    'construction',
    'battery',
    'liquid',
    'hazard'
];

export const emptyGreenObjectForm = (coords: number[]): ICreateGreenObjectFormData => {
    return {
        LocationName: "",
        Latitude: coords[0],
        Longitude: coords[1],
        Street: "",
        City: "",
        Country: "",
        Shape: "",
        TrashType: "",
    }
}