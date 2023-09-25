import { useState, useRef, useEffect } from 'react';
import { useDispatch } from 'react-redux';

import { Feature, Map } from 'ol';
import VectorLayer from 'ol/layer/Vector';
import VectorSource from 'ol/source/Vector';
import { Geometry, Point } from 'ol/geom';
import { Flex } from '@chakra-ui/react';

import { IFeatureInfo, createFeature, initializeMap } from './utils';
import { setSelectedGreenObject } from '../../../slices/greenObject.slice';

import './index.css';
import 'ol/ol.css';

interface IMapContainer {
  edit: boolean;
  greenObjects: IFeatureInfo[];
}

const MapContainer = ({ edit, greenObjects }: IMapContainer) => {

  const [map, setMap] = useState<Map | undefined>();
  const [featuresLayer, setFeaturesLayer] = useState<VectorLayer<VectorSource<Geometry>> | undefined>();
  const [markers, setMarkers] = useState<Feature<Point>[]>([]);
  const [selectedObjectId, setSelectedObject] = useState<string | undefined>(undefined);

  const dispatch = useDispatch();
  const mapElement = useRef<HTMLDivElement | null>(null);

  useEffect(() => {
    if (mapElement.current === null) return;
    if (map) return;
    if (featuresLayer) return;

    const {initialMap, initialFeaturesLayer, markers} = initializeMap(greenObjects, edit, setSelectedObject);

    initialMap.setTarget(mapElement.current);
    setMap(initialMap);
    setFeaturesLayer(initialFeaturesLayer);
    setMarkers(markers);

    // on dismount
    return () => initialMap.setTarget("");
  }, []);

  useEffect(() => {
    if (!map) return;
    map.updateSize();
  }, [map]);

  useEffect(() => {
    const foundObject = greenObjects.find(go => go.id === selectedObjectId);
    dispatch(setSelectedGreenObject(foundObject ?? {id: "", coords: [0, 0]}));
  }, [selectedObjectId]);

  useEffect(() => {
    const newFeature = createFeature(greenObjects[greenObjects.length - 1]);
    featuresLayer?.getSource()?.addFeature(newFeature);
  }, [JSON.stringify(greenObjects)]);
  
  return (
    <Flex direction={"column"}>
      <div ref={mapElement} className="map-container"></div>
    </Flex>
  )
}

export default MapContainer;