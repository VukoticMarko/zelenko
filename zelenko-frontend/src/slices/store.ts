import { configureStore } from '@reduxjs/toolkit'
import greenObjectSlice from './greenObject.slice'

export const store = configureStore({
  reducer: {
    greenObjects: greenObjectSlice,
  },
})

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch