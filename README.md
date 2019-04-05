# Raccoon SIEM
Raccoon is an open source SIEM designed for high traffic volume environments.

The project is in active development stage. If you want to give it a try - do not use master branch since it is unstable. 
Use [releases](https://github.com/tephrocactus/raccoon-siem/releases) instead.

### Terminology

#### Collector
Collector is an executable which is responsibe for:
* Collecting logs from various systems
* Parsing
* Normalization
* Filtration
* Enrichment
* Aggregation
* Sending logs to various destinations, like storage and correlators
