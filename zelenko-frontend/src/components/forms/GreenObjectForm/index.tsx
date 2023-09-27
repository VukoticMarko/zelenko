import { useEffect } from 'react';
import { useFormik } from 'formik';
import { emptyGreenObjectForm, shapeOptions, transFromGreenObject, trashTypeOptions } from './utils';
import { Button, Card, FormControl, FormLabel, HStack, Heading, Input, Select, VStack } from '@chakra-ui/react';
import { RootState } from '../../../slices/store';
import { useDispatch, useSelector } from 'react-redux';
import { transCords } from '../../maps/MapContainer/utils';
import { createObjectApi, updateObjectApi } from '../../../api/greenObjects.api';
import { ICreateGreenObject, IGreenObject } from '../../../common/dtos';
import { addGreenObject, switchEdit, updateGreenScore } from '../../../slices/greenObject.slice';

interface IGreenObjectFormProps {
    successCallback?: any;
    isEdit?: boolean;
    greenObject?: IGreenObject;
}

const GreenObjectForm = ({ isEdit, greenObject }: IGreenObjectFormProps) => {

    const dispatch = useDispatch();
    const newCoords = useSelector<RootState, number[]>((state) => state.greenObjects.newCords);
    const postGreenObject = async (values: ICreateGreenObject) => {
        try {
            const newGreenObject = await createObjectApi({...values, Latitude: +values.Latitude, Longitude: +values.Longitude});
            dispatch(addGreenObject(newGreenObject));
            formik.resetForm();
        } catch (err) {
            console.log(err);
        }
    }

    const patchGreenObject = async (values: ICreateGreenObject) => {
        try {
            if (!greenObject) return;
            const patchedGreenObject = await updateObjectApi(
                {...values, ID: greenObject.ID, Latitude: +values.Latitude, Longitude: +values.Longitude}
            );
            dispatch(updateGreenScore(patchedGreenObject));
            dispatch(switchEdit(false));
            formik.resetForm();
        } catch (err) {
            console.log(err);
        }
    }

    const formik = useFormik<ICreateGreenObject>({
        initialValues: greenObject ? transFromGreenObject(greenObject) : emptyGreenObjectForm(transCords(newCoords)),
        // enableReinitialize: true,
        onSubmit: isEdit ? patchGreenObject : postGreenObject
    });

    useEffect(() => {
        if (!newCoords) return;
        const coords = transCords(newCoords);
        const values = formik.values;
        values.Latitude = coords[0];
        values.Longitude = coords[1];
        formik.setValues({...values});
    }, [newCoords]);


    return (<Card borderWidth={'thick'} padding={'10px'} marginTop={4} alignItems={'flex-start'}>
        <Heading>{isEdit ? "Edit Green Object" : "Create Green Object"}</Heading>
        <form onSubmit={formik.handleSubmit} style={{width: '100%'}}>
            <VStack spacing={4} align="flex-start">
                <FormControl>
                    <FormLabel htmlFor="LocationName">Location Name</FormLabel>
                    <Input
                        id="LocationName"
                        name="LocationName"
                        type="LocationName"
                        variant="filled"
                        onChange={formik.handleChange}
                        value={formik.values.LocationName}
                    />
                </FormControl>
                <HStack gap={'10px'} width={"100%"}>
                    <FormControl>
                        <FormLabel htmlFor="Latitude">Latitude</FormLabel>
                        <Input
                            id="Latitude"
                            name="Latitude"
                            type="Latitude"
                            variant="filled"
                            onChange={formik.handleChange}
                            value={formik.values.Latitude}
                        />
                    </FormControl>
                    <FormControl>
                        <FormLabel htmlFor="Longitude">Longitude</FormLabel>
                        <Input
                            id="Longitude"
                            name="Longitude"
                            type="Longitude"
                            variant="filled"
                            onChange={formik.handleChange}
                            value={formik.values.Longitude}
                        />
                    </FormControl>
                </HStack>
                <HStack width={"100%"}>
                    <FormControl>
                        <FormLabel htmlFor="Street">Street</FormLabel>
                        <Input
                            id="Street"
                            name="Street"
                            type="Street"
                            variant="filled"
                            onChange={formik.handleChange}
                            value={formik.values.Street}
                        />
                    </FormControl>
                    <FormControl>
                        <FormLabel htmlFor="Longitude">City</FormLabel>
                        <Input
                            id="City"
                            name="City"
                            type="City"
                            variant="filled"
                            onChange={formik.handleChange}
                            value={formik.values.City}
                        />
                    </FormControl>
                    <FormControl>
                        <FormLabel htmlFor="Country">Country</FormLabel>
                        <Input
                            id="Country"
                            name="Country"
                            type="Country"
                            variant="filled"
                            onChange={formik.handleChange}
                            value={formik.values.Country}
                        />
                    </FormControl>
                </HStack>
                <HStack gap={'10px'} width={"100%"}>
                    <FormControl>
                        <FormLabel htmlFor="Shape">Shape</FormLabel>
                        <Select
                            id="Shape"
                            name="Shape"
                            variant="filled"
                            onChange={formik.handleChange}
                            value={formik.values.Shape}
                        >
                            {shapeOptions.map(shape => (
                                <option key={shape} value={shape}>{shape}</option>
                            ))};
                        </Select>
                    </FormControl>
                    <FormControl>
                        <FormLabel htmlFor="Type">Type</FormLabel>
                        <Select
                            id="TrashType"
                            name="TrashType"
                            variant="filled"
                            onChange={formik.handleChange}
                            value={formik.values.TrashType}
                        >
                            {trashTypeOptions.map(trashType => (
                                <option key={trashType} value={trashType}>{trashType}</option>
                            ))};
                        </Select>
                    </FormControl>
                </HStack>
                <HStack>
                    <Button type='submit'>{isEdit ? "Edit Green Object" : "Add Green Object"}</Button>
                    {isEdit && <Button onClick={() => dispatch(switchEdit(false))}>Cancel</Button>}
                </HStack>
            </VStack>
        </form>
    </Card>);
}

export default GreenObjectForm;