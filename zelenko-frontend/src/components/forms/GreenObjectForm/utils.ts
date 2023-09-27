import { ICreateGreenObject, IGreenObject } from "../../../common/dtos";

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

export const emptyGreenObjectForm = (coords: number[]): ICreateGreenObject => {
    return {
        LocationName: "",
        Latitude: coords[0],
        Longitude: coords[1],
        Street: "",
        City: "",
        Country: "",
        Shape: "can",
        TrashType: "all",
    }
}

export const transFromGreenObject = (greenObject: IGreenObject) : ICreateGreenObject => {
    return {
        LocationName: greenObject.LocationName,
        Latitude: greenObject.Location.Latitude,
        Longitude: greenObject.Location.Longitude,
        Street: greenObject.Location.Street,
        City: greenObject.Location.City,
        Country: greenObject.Location.Country,
        Shape: greenObject.Shape,
        TrashType: greenObject.TrashType,
    }
}