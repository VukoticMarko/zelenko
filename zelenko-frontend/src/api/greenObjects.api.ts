import { ICreateGreenObjectFormData } from "../components/forms/GreenObjectForm/utils";
import { fetchLike } from "./base";

export const createObjectApi = async (body: ICreateGreenObjectFormData) => {
    const response = await fetchLike("POST", "addObject", body);
    return response;
}