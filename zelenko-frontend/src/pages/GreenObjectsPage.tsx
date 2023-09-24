import { useState } from 'react';
import { Button, Flex } from '@chakra-ui/react';
import MapContainer from '../components/maps/MapContainer';
import { useSelector } from 'react-redux';
import { IGreenObject } from '../slices/greenObject.slice';
import { RootState } from '../slices/store';
import { IFeatureInfo, transCords } from '../components/maps/MapContainer/utils';
import { v4 as uuid } from 'uuid';


const GreenObjectsPage = () => {
  
    const selectedGreenObject = useSelector<RootState, IGreenObject>((state) => state.greenObjects.selectedGreenObject);
    const newCoords = useSelector<RootState, number[]>((state) => state.greenObjects.newCords);
    const [greenObjects, setGreenObjects] = useState<IFeatureInfo[]>([
        {id: uuid(), coords: transCords([19.828483, 45.247179], true)},
        {id: uuid(), coords: transCords([19.847427, 45.251757], true)},
    ]);

    const addGreenObject = () => {
        setGreenObjects([...greenObjects, {id: uuid(), coords: newCoords}]);
    }

    return (
        <Flex direction={"column"}>
            <MapContainer edit={true} greenObjects={greenObjects}/>
            {selectedGreenObject.id && <form>
                <p>Selected id: {selectedGreenObject.id}</p>
                <p>Long, Lang: {JSON.stringify(transCords(selectedGreenObject.coords))}</p>
            </form>}
            {!selectedGreenObject.id && <form>
                <p>Selected id: Yet to be</p>
                <p>Long, Lang: {JSON.stringify(transCords(newCoords))}</p>
                <Button onClick={addGreenObject}>Add Green Object</Button>
            </form>}
        </Flex>
    )

}

export default GreenObjectsPage;