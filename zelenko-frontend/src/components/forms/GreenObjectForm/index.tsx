import { useEffect } from 'react';
import { useFormik } from 'formik';
import { ICreateGreenObjectFormData, emptyGreenObjectForm, shapeOptions, trashTypeOptions } from './utils';
import { Button, FormControl, FormLabel, HStack, Input, Select, VStack } from '@chakra-ui/react';
import { RootState } from '../../../slices/store';
import { useSelector } from 'react-redux';
import { transCords } from '../../maps/MapContainer/utils';
import { v4 as uuid } from 'uuid';
import { createObjectApi } from '../../../api/greenObjects.api';

interface IGreenObjectFormProps {
    successCallback?: any;
}

const GreenObjectForm = ({ successCallback }: IGreenObjectFormProps) => {

    const newCoords = useSelector<RootState, number[]>((state) => state.greenObjects.newCords);
    const postGreenObject = async (values: ICreateGreenObjectFormData) => {
        console.log("values => ", values);
        try {
            const response = await createObjectApi(values);
            if (successCallback) successCallback(uuid(), transCords([values.Latitude, values.Longitude], true));
            console.log("resp => ", response);
        } catch (err) {
            console.log(err);
        }
    }

    const formik = useFormik<ICreateGreenObjectFormData>({
        initialValues: emptyGreenObjectForm(transCords(newCoords)),
        enableReinitialize: true,
        onSubmit: postGreenObject
    });

    useEffect(() => {
        if (!newCoords) return;
        const coords = transCords(newCoords);
        const values = formik.values;
        values.Latitude = coords[0];
        values.Longitude = coords[1];
        formik.setValues({...values});
    }, [newCoords]);


    return (<>
        <form onSubmit={formik.handleSubmit}>
            <VStack spacing={4} align="flex-start" marginTop={4}>
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
                        <FormLabel htmlFor="Shape">Shape</FormLabel>
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
                <Button type='submit'>Add Green Object</Button>
            </VStack>
        </form>
    </>);
}

export default GreenObjectForm;