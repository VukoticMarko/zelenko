import { createSlice } from '@reduxjs/toolkit'
import { IFeatureInfo, IGreenObject } from '../common/dtos';

export interface IFeatureInfoWithObject {
    featureInfo: IFeatureInfo;
    greenObject: IGreenObject | undefined;
}

export interface IGreenObjectsState {
    selectedFeature: IFeatureInfoWithObject;
    greenObjects: IGreenObject[];
    newCords: number[];
    isEdit: boolean;
}

const initialState: IGreenObjectsState = {
    selectedFeature: {
        featureInfo: {
            id: '',
            coords: [0, 0],
        },
        greenObject: undefined,
    },
    newCords: [0, 0],
    greenObjects: [],
    isEdit: false,
}

export const greenObjectsSlice = createSlice({
    name: 'greenObjects',
    initialState,
    reducers: {
        setNewCords: (state, action) => {
            state.newCords = action.payload;
        },
        setSelectedGreenObject: (state, action) => {
            if (state.isEdit) {
                return;
            }
            const greenObject = state.greenObjects.find(go => go.ID == action.payload.id);
            state.selectedFeature = {
                featureInfo: {...action.payload},
                greenObject
            };
        },
        setGreenObjects: (state, action) => {
            state.greenObjects = action.payload;
        },
        addGreenObject: (state, action) => {
            state.greenObjects.push(action.payload);
        },
        updateGreenScore: (state, action) => {
            const index = state.greenObjects.findIndex(go => go.ID == action.payload.ID);
            if (index === -1) return;
            state.greenObjects[index] = action.payload;
            state.selectedFeature.greenObject = state.greenObjects[index];
        },
        switchEdit: (state, action) => {
            if (!state.selectedFeature) return;
            state.isEdit = action.payload;
        },
    },
})

// Action creators are generated for each case reducer function
export const { 
    setNewCords, 
    setSelectedGreenObject, 
    setGreenObjects, 
    addGreenObject, 
    updateGreenScore,
    switchEdit,
} = greenObjectsSlice.actions;
export default greenObjectsSlice.reducer