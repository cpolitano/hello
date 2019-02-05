const geoJson = {
    type: "Polygon",
    coordinates: [
        [[30, 10], [40, 40], [20, 40], [10, 20], [30, 10]]
    ]
}

console.log(JSON.stringify(geoJson));

const formattedGeoJson = '{"type": "Polygon", "coordinates": [[[30, 10], [40, 40], [20, 40], [10, 20], [30, 10]]]}'
