import { ICreateGreenObject, IGreenObject, IUpdateGreenObject } from "../common/dtos";
import { fetchLike } from "./base";

export const createObjectApi = async (body: ICreateGreenObject): Promise<IGreenObject> => {
    const response = await fetchLike("POST", "addObject", body);
    return response;
}

export const updateObjectApi = async (body: IUpdateGreenObject): Promise<IGreenObject> => {
    const response = await fetchLike("POST", "updateObject", body);
    return response;
}

export const getAllGreenObjectsApi = async () => {
    const response = await fetchLike("GET", "getAll");
    if (!response) return [];
    return response;
}