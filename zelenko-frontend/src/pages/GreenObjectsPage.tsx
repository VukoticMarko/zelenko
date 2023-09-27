import { useEffect, useState } from 'react';
import { Flex } from '@chakra-ui/react';
import MapContainer from '../components/maps/MapContainer';
import { useDispatch, useSelector } from 'react-redux';
import { RootState } from '../slices/store';
import { transCords } from '../components/maps/MapContainer/utils';
import GreenObjectForm from '../components/forms/GreenObjectForm';
import GreenScoreForm from '../components/forms/GreenScoreForm';
import { getAllGreenObjectsApi } from '../api/greenObjects.api';
import { IFeatureInfo, IGreenObject } from '../common/dtos';
import { IFeatureInfoWithObject, setGreenObjects } from '../slices/greenObject.slice';


const GreenObjectsPage = () => {
  
    const dispatch = useDispatch();
    const selectedGreenObject = useSelector<RootState, IFeatureInfoWithObject>((state) => state.greenObjects.selectedFeature);
    const greenObjects = useSelector<RootState, IGreenObject[]>((state) => state.greenObjects.greenObjects);
    const isEdit = useSelector<RootState, boolean>((state) => state.greenObjects.isEdit);
    const [features, setFeatures] = useState<IFeatureInfo[]>([]);

    useEffect(() => {
        if (greenObjects.length !== 0) return;
        getAllGreenObjectsApi()
            .then((resp) => dispatch(setGreenObjects(resp)))
            .catch(e => console.log(e)); 
    }, []);

    useEffect(() => {
        if (greenObjects.length === 0) return;
        const mappedFeatureInfo = greenObjects.map((go) => ({
            id: go.ID, 
            coords: transCords([go.Location.Latitude, go.Location.Longitude], true)}
        ));
        setFeatures([...mappedFeatureInfo]);
    }, [greenObjects])

    return (
        <Flex direction={"column"}>
            <MapContainer edit={true} features={features}/>
            {selectedGreenObject.greenObject && !isEdit && <div>
                <GreenScoreForm greenObject={selectedGreenObject.greenObject}/>
            </div>}
            {!selectedGreenObject.featureInfo.id && <GreenObjectForm/>}
            {selectedGreenObject.featureInfo.id && isEdit && <GreenObjectForm isEdit={true} greenObject={selectedGreenObject.greenObject}/>
            }
        </Flex>
    )

}

export default GreenObjectsPage;