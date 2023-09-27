import { IGreenObject } from "../common/dtos";
import { fetchLike } from "./base";

export const addOneScore = async (body: IGreenObject): Promise<IGreenObject> => {
    const response = await fetchLike("POST", "addScore", body);
    return response;
}

export const subOneScore = async (body: IGreenObject): Promise<IGreenObject> => {
    const response = await fetchLike("POST", "subScore", body);
    return response;
}