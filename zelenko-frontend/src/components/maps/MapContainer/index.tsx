import { useState, useRef, useEffect } from 'react';
import { useDispatch } from 'react-redux';

import { Feature, Map } from 'ol';
import VectorLayer from 'ol/layer/Vector';
import VectorSource from 'ol/source/Vector';
import { Geometry, Point } from 'ol/geom';
import { Flex } from '@chakra-ui/react';

import { createFeature as createMarker, initializeMap, transCords } from './utils';
import { setSelectedGreenObject } from '../../../slices/greenObject.slice';

import './index.css';
import 'ol/ol.css';
import { IFeatureInfo } from '../../../common/dtos';

interface IMapContainer {
  edit: boolean;
  features: IFeatureInfo[];
}

const MapContainer = ({ edit, features }: IMapContainer) => {

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

    const {initialMap, initialFeaturesLayer, markers} = initializeMap(features, edit, setSelectedObject);

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
    if (!features) return;
    const foundObject = features.find(go => go.id === selectedObjectId);
    dispatch(setSelectedGreenObject(foundObject ?? {id: "", coords: [0, 0]}));
  }, [selectedObjectId]);

  const handeLoad = () => {
    const newMarkers = features.map(feat => createMarker(feat));
    featuresLayer?.getSource()?.addFeatures(newMarkers);
    setMarkers([...newMarkers, ...markers]);
  }

  const hanldeEdit = () => {
    for (const feat of features) {
      const matcherdMarker = markers.find(mark => feat.id == mark.get('greenObjectId'));
      if (!matcherdMarker) continue;
      const markerCoords = matcherdMarker.getGeometry()?.getFlatCoordinates();
      if (!markerCoords) continue;
      const coordsAreSame = (markerCoords[0] === feat.coords[0]) && (markerCoords[1] === feat.coords[1]) 
      !coordsAreSame && matcherdMarker.getGeometry()?.setCoordinates([...feat.coords]);
    }
  }

  const handleAdd = () => {
    const newMarker = createMarker(features[features.length - 1]);
    featuresLayer?.getSource()?.addFeature(newMarker);
    setMarkers([...markers, newMarker]);
  }

  useEffect(() => {
    if (features.length === 0) return;
    if (markers.length === 1) {
      handeLoad();
      return;
    }
    if (markers.length - 1 === features.length) {
      hanldeEdit()
      return;
    }
    handleAdd();
  }, [JSON.stringify(features)]);
  
  return (
    <Flex direction={"column"}>
      <div ref={mapElement} className="map-container"></div>
    </Flex>
  )
}

export default MapContainer;