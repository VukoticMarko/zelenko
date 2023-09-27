import { Map, View, Feature, MapBrowserEvent } from 'ol';
import VectorLayer from 'ol/layer/Vector';
import VectorSource from 'ol/source/Vector';
import TileLayer from 'ol/layer/Tile';
import OSM from 'ol/source/OSM'
import { fromLonLat as projectLongLat, transform as transformCoords } from 'ol/proj';
import { Geometry, Point } from 'ol/geom';
import { Style, Icon } from 'ol/style';
import { defaults as getDefaultMapControls} from 'ol/control'
import { v4 as uuid } from 'uuid';

import {store} from '../../../slices/store';
import { setNewCords } from '../../../slices/greenObject.slice';
import { IFeatureInfo } from '../../../common/dtos';


const CENTER_CORDS = [19.833549, 45.267136];
const TRANS_CENTER_CORDS = transformCoords(CENTER_CORDS, "EPSG:4326", "EPSG:3857");
const TRASH_CAN = "https://cdn-icons-png.flaticon.com/512/542/542724.png"
const MOOVABLE_ICON = "https://play-lh.googleusercontent.com/5WifOWRs00-sCNxCvFNJ22d4xg_NQkAODjmOKuCQqe57SjmDw8S6VOSLkqo6fs4zqis"


export const transCords = (cords: number[], fromHuman: boolean = false) => {
    if (fromHuman) {
        return transformCoords(cords, "EPSG:4326", "EPSG:3857");
    }
    return transformCoords(cords, "EPSG:3857", "EPSG:4326");
}

export interface IInitMapResult {
    initialMap: Map,
    initialFeaturesLayer: VectorLayer<VectorSource<Geometry>>;
    markers: Feature<Point>[];
    moovable?: Feature<Point>;
}

export const createFeature = (featureInfo: IFeatureInfo, icon?: string, iconScale?: number): Feature<Point> => {
    const markerFeature = new Feature({
        geometry: new Point(featureInfo.coords),
    });

    markerFeature.setStyle(new Style({
        image: new Icon({
        scale: iconScale ?? 0.05,
        src: icon ?? TRASH_CAN,
        })
    }));

    markerFeature.setProperties({
        ...markerFeature.getProperties(),
        "greenObjectId": featureInfo.id,
    });

    return markerFeature;
}

export const addFeature = (features: Feature[], featureInfo: IFeatureInfo) => {
    const markerFeature = createFeature(featureInfo);
    features.push(markerFeature);
}

export const createMoovableFeature = () => {
    const markerFeature = createFeature({id: uuid(), coords: TRANS_CENTER_CORDS}, MOOVABLE_ICON, 0.1);
    return markerFeature;
}

const moveFeatureCallback = (markerFeature: Feature<Point>) => {
    return (event: MapBrowserEvent<any>) => {
        markerFeature.getGeometry()?.setCoordinates(event.coordinate);
        store.dispatch(setNewCords(event.coordinate));
    } 
}

export const initializeMap = (features: IFeatureInfo[], edit?: boolean, setSelected?: any): IInitMapResult => {
    const markerFeatures = features.map(feat => createFeature(feat));
    const moovableFeature = edit ? createMoovableFeature() : undefined; 
    edit && markerFeatures.push(moovableFeature!);

    const initalFeaturesLayer = new VectorLayer({
        source: new VectorSource({
            features: markerFeatures,
            wrapX: true,
        }),
    });

    const initialMap = new Map({
        layers: [
            new TileLayer({
            source: new OSM(),
            }),
            initalFeaturesLayer
        ],
        view: new View({
            projection: 'EPSG:3857',
            center: transformCoords(CENTER_CORDS, "EPSG:4326", "EPSG:3857"),
            zoom: 12,
            minZoom: 5,
            maxZoom: 20,
        }),
        controls: getDefaultMapControls()
    });

    edit && initialMap.on("singleclick", moveFeatureCallback(moovableFeature!));
    initialMap.on('click', (event) => {
        let clickedFeature = initialMap.forEachFeatureAtPixel(event.pixel, (feature) => {
            return feature;
        });
        edit && setSelected(clickedFeature?.get('greenObjectId'));
    });

    const view = initialMap.getView();
    view.setCenter(projectLongLat(CENTER_CORDS, view.getProjection()));

    return {
        initialMap: initialMap,
        initialFeaturesLayer: initalFeaturesLayer,
        markers: markerFeatures,
        moovable: moovableFeature,
    }
}
