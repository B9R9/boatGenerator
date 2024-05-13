import React, { useEffect, useRef, useState } from 'react';
import L from 'leaflet';
import 'leaflet/dist/leaflet.css';

import DisplayBoatsDetails from './DisplayBoatsDetails';

const isInValidRange = (latitude, longitude) => {
    // Vérifier si la latitude est dans la plage [-5000, 5000]
    if (latitude < -5000 || latitude > 5000) {
        console.log("Orig Value error")
        return true;
    }
    
    // Vérifier si la longitude est dans la plage [-5000, 5000]
    if (longitude < -5000 || longitude > 5000) {
        console.log("Orig Value error")
        return true;
    }
    
    // Si les coordonnées sont dans les plages valides, retourner true
    return false;
};

const DisplayBoats = ({ boats }) => {
    const mapRef = useRef(null);
    const [map, setMap] = useState(null);
    const [selectedBoat, setSelectedBoat] = useState(null);
    const [boatDetails, setBoatDetails] = useState(null);
    const [boatHistory, setBoatHistory] = useState(false);
    const [index, setIndex] = useState(0)
    const mapSize = 10;

    const handleClick = (event) => {
        setSelectedBoat(null)
    }

    const handleIndexListBoat = (event) => {
        let newIndex = index + 1
        if (newIndex > boats.length) {
            newIndex = 0
        }
        setIndex(newIndex)
    }

    useEffect(() => {
        console.log ("SelectedBoat: ", selectedBoat)
        if (selectedBoat && boats.find(boat => boat.id === selectedBoat.id)) {
            const foundBoat = boats.find(boat => boat.id === selectedBoat.id)
            console.log("FoundBoat: ", foundBoat)
            setBoatDetails(foundBoat);
        } else {
            console.log("FoundBoat: NULL ")
            setSelectedBoat(null)
            setBoatDetails(null)
        }
    }, [boats, selectedBoat])

    
    useEffect(() => {
        if (mapRef.current && !map) {
            if (!mapRef.current._leaflet_id) {
                if (mapRef.current && mapRef.current._leaflet_id) {
                    mapRef.current._leaflet_id = null;
                }
                const initialMap = L.map(mapRef.current,{
                    crs: L.CRS.Simple,
                    minZoom: -5,
                    maxBounds: [[-5000, -5000], [5000, 5000]],
                })// Définir le point de vue initial
    
                // Définir les limites de la carte
                const bounds = [[-5000, -5000], [5000, 5000]];
    
                // Ajouter l'image en tant que calque
                // const image = L.imageOverlay('uqm_map_full.png', bounds).addTo(initialMap);
                initialMap.getContainer().style.background = 'blue';
                // Adapter la carte pour afficher l'image entière
                initialMap.fitBounds(bounds);
                      // Ajouter les lignes représentant les axes x et y
                const xAxis = L.polyline([[-5000, 0], [5000, 0]], { color: 'black' }).addTo(initialMap);
                const yAxis = L.polyline([[0, -5000], [0, 5000]], { color: 'black' }).addTo(initialMap);
                       // Ajouter des marques de mesure sur les bords de la carte
                const measureOptions = { color: 'black', weight: 1, opacity: 0.5 };
                const measureLeft = L.polyline([[-5000, -5000], [-5000, 5000]], measureOptions).addTo(initialMap);
                const measureRight = L.polyline([[5000, -5000], [5000, 5000]], measureOptions).addTo(initialMap);
                const measureTop = L.polyline([[-5000, -5000], [5000, -5000]], measureOptions).addTo(initialMap);
                const measureBottom = L.polyline([[-5000, 5000], [5000, 5000]], measureOptions).addTo(initialMap);
                
                initialMap.setView([0, 0], 0)
                setMap(initialMap);
            }
        }
    }, [mapRef, map]);
    
    useEffect(() => {
        if (map) {
            // Supprimer tous les marqueurs existants
            map.eachLayer(layer => {
                if (layer instanceof L.CircleMarker) {
                    map.removeLayer(layer);
                }
            });
    
            boats.forEach((boat, index) => {
                console.log("Adding marker for boat:", boat);
                if (isInValidRange(boat.Latitude, boat.Longitude)) {
                    console.log("EROOORRRRRRRR------>")
                }
                let markerColor = 'green'; // Couleur par défaut pour les marqueurs de bateaux non sélectionnés
                let markerSize = 4
                if (selectedBoat && selectedBoat.id === boat.id) {
                    markerColor = 'red'; // Si le bateau est sélectionné, changer la couleur du marqueur en rouge
                    markerSize = 8
                }
                const marker = L.circleMarker([boat.Latitude, boat.Longitude], { radius: markerSize, color: markerColor, boatId: boat.id })
                marker.bindTooltip(`<b>Informations du bateau :</b><br>ID:${boat.id}<br>Cap: ${boat.Cap}<br>LastDirection: ${boat.LastDirection}<br>Latitude: ${boat.Latitude}<br>Longitude: ${boat.Longitude}<br>Speed: ${boat.Speed}`, { permanent: false, direction: 'top' });
                marker.on('click', () => {
                    if (selectedBoat) {
                        setSelectedBoat(null)
                    } else {
                    setSelectedBoat(boat)
                    }
                });
                console.log("Adding marker to map:", marker);
                marker.addTo(map);
            });
        }
    }, [map, boats, selectedBoat]);

    useEffect(() => {
        return () => {
            if (map) {
                map.remove();
                setMap(null);
            }
        };
    }, [map]);

    return (
    <>    
    <div id="mapid" style={{ height: '60vh', width: '60vw', margin: 'auto', display: 'block'}}  ref={mapRef}></div>
    {boatDetails &&  <DisplayBoatsDetails boat={boatDetails} handleClick={handleClick} handleIndexListBoat={handleIndexListBoat}/> }
    </>

    )

};

export default  DisplayBoats