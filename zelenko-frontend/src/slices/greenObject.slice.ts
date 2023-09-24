import { createSlice } from '@reduxjs/toolkit'

export interface IGreenObject {
    id: string;
    coords: number[];
}

export interface IGreenObjectsState {
    selectedGreenObject: IGreenObject,
    newCords: number[];
}

const initialState: IGreenObjectsState = {
    selectedGreenObject: {
        id: '',
        coords: [0, 0],
    },
    newCords: [0, 0],
}

export const greenObjectsSlice = createSlice({
    name: 'greenObjects',
    initialState,
    reducers: {
        setNewCords: (state, action) => {
            state.newCords = action.payload;
        },
        setSelectedGreenObject: (state, action) => {
            state.selectedGreenObject = action.payload;
        }
    },
})

// Action creators are generated for each case reducer function
export const { setNewCords, setSelectedGreenObject } = greenObjectsSlice.actions
export default greenObjectsSlice.reducer