import { useState } from 'react';
import { Flex } from '@chakra-ui/react';
import MapContainer from '../components/maps/MapContainer';
import { useSelector } from 'react-redux';
import { IGreenObject } from '../slices/greenObject.slice';
import { RootState } from '../slices/store';
import { IFeatureInfo, transCords } from '../components/maps/MapContainer/utils';
import { v4 as uuid } from 'uuid';
import GreenObjectForm from '../components/forms/GreenObjectForm';
import GreenScoreForm from '../components/forms/GreenScoreForm';


const GreenObjectsPage = () => {
  
    const selectedGreenObject = useSelector<RootState, IGreenObject>((state) => state.greenObjects.selectedGreenObject);
    const [greenObjects, setGreenObjects] = useState<IFeatureInfo[]>([
        {id: uuid(), coords: transCords([19.828483, 45.247179], true)},
        {id: uuid(), coords: transCords([19.847427, 45.251757], true)},
    ]);

    const addGreenObject = (id: string, coords: number[]) => {
        setGreenObjects([...greenObjects, {id, coords: coords}]);
    }

    return (
        <Flex direction={"column"}>
            <MapContainer edit={true} greenObjects={greenObjects}/>
            {selectedGreenObject.id && <form>
                <p>Selected id: {selectedGreenObject.id}</p>
                <p>Long, Lang: {JSON.stringify(transCords(selectedGreenObject.coords))}</p>
                <GreenScoreForm/>
            </form>}
            {!selectedGreenObject.id &&
                <div>
                    {/* add successCallback to GreenObject and get cords from redux */}
                <GreenObjectForm successCallback={addGreenObject}/>
                {/* <p>Long, Lang: {JSON.stringify(transCords(newCoords))}</p> */}
                {/* <Button onClick={addGreenObject}>Add Green Object</Button> */}
                </div>
            }
        </Flex>
    )

}

export default GreenObjectsPage;