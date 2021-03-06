World Locations I/O
===

| branch   | status |
| -------- |:------:|
| master   | [![CircleCI](https://circleci.com/gh/SamuelTissot/worldlocations/tree/master.svg?style=svg)](https://circleci.com/gh/SamuelTissot/worldlocations/tree/master) |


API Documentation
----
The API documentation is hosted with postman and can be viewed here

[**WORLDLOCATIONS DOCUMENTATION**](https://documenter.getpostman.com/view/6999284/S17qTpUm)


About
----
Worldlocations I/O is a personal project that I did in order to explore [Buffalo](https://github.com/gobuffalo/buffalo) 
capabilities. Keeping this in mind I do not guaranty the full accuracy of the data returned by the API. However, if any 
discrepancies or omission are notice, I will be more than happy to correct then. Please submit an issue on the github 
repo.

Thanks.
 

Goal of the Project
----
It is a project that aims at mapping the relationships between different geo/political entities.

`countries -> political subdivision -> cities`

for example 
it is possible to query all the cities, or subdivision within a country.
it is also possible to query cities of subdivisions.

```
+--------------------------------------+
|                                      |
|               Country                |
|                  +                   |
|        +--------------------+        |
|        |                    |        |
|        v                    v        |
|   Subdivision             cities     |
|        |                             |
|        |                             |
|        v                             |
|     cities                           |
|                                      |
|                                      |
+--------------------------------------+

```

Worldlocations in Numbers
----


| section | records |
|:------|------:|
| countries | 273* |
| subdivisions | 5194 |
| languages | 104 |
| cities | 186455 |

* includes none official zones 

Authentication
---
The service is currently free with a rate limit of 15 req/mins. 

Error Codes
---
200 OK or 404 Not found

Rate limit
---
15 req/mins

If you require more bandwidth
please contact me at worldlocations.io [AT] gmail [DOT] com

----

This project was build with Makes Bates [Buffalo](https://github.com/gobuffalo/buffalo) ecosystem. Thank You!