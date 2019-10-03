package main

var (
	jsonSample = `{
  "nhits": 1,
  "parameters": {
    "dataset": "velib-disponibilite-en-temps-reel",
    "timezone": "UTC",
    "rows": 10,
    "format": "json",
    "geofilter.distance": [
      "48.853169,2.402782,100"
    ],
    "facet": [
      "overflowactivation",
      "creditcard",
      "kioskstate",
      "station_state"
    ]
  },
  "records": [
    {
      "datasetid": "velib-disponibilite-en-temps-reel",
      "recordid": "de73b813e2f224f9f9bcf075da0004f820511263",
      "fields": {
        "station_state": "Operative",
        "maxbikeoverflow": 22,
        "densitylevel": "1",
        "nbbikeoverflow": 0,
        "dist": "82.0709088808",
        "nbedock": 0,
        "station_name": "Haies - Réunion",
        "kioskstate": "yes",
        "nbfreeedock": 0,
        "station_type": "no",
        "station_code": "20116",
        "creditcard": "no",
        "station": "{\"code\": \"20116\", \"name\": \"Haies - Réunion\", \"state\": \"Operative\", \"type\": \"no\", \"dueDate\": 1521154800, \"gps\": {\"latitude\": 48.85386896766491, \"longitude\": 2.402426162238911}}",
        "nbfreedock": 21,
        "duedate": "2018-03-15",
        "nbebikeoverflow": 0,
        "nbebike": 1,
        "overflow": "no",
        "nbdock": 22,
        "geo": [
          48.85386896766491,
          2.402426162238911
        ],
        "overflowactivation": "no",
        "nbbike": 0
      },
      "geometry": {
        "type": "Point",
        "coordinates": [
          2.402426162238911,
          48.85386896766491
        ]
      },
      "record_timestamp": "2019-10-02T19:11:01.032000+00:00"
    }
  ],
  "facet_groups": [
    {
      "facets": [
        {
          "count": 1,
          "path": "no",
          "state": "displayed",
          "name": "no"
        }
      ],
      "name": "overflowactivation"
    },
    {
      "facets": [
        {
          "count": 1,
          "path": "Operative",
          "state": "displayed",
          "name": "Operative"
        }
      ],
      "name": "station_state"
    },
    {
      "facets": [
        {
          "count": 1,
          "path": "yes",
          "state": "displayed",
          "name": "yes"
        }
      ],
      "name": "kioskstate"
    },
    {
      "facets": [
        {
          "count": 1,
          "path": "no",
          "state": "displayed",
          "name": "no"
        }
      ],
      "name": "creditcard"
    }
  ]
}`
)
