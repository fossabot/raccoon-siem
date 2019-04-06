# Raccoon SIEM
Raccoon is an open source SIEM designed for high traffic volume environments.

The project is in active development stage. If you want to give it a try - do not use master branch since it is unstable. 
Use [releases](https://github.com/tephrocactus/raccoon-siem/releases) instead.

## Table of contents
* [Components overview](#components-overview)
  * [Collector](#collector)
  * [Correlator](#correlator)
  * [Core](#core)
  * [Bus](#bus)
  * [Active lists storage](#active-lists-storage)
  * [Event storage](#event-storage)

## Components overview

### Collector
Collector is an executable which is responsible for:
* Collecting logs from various systems
* Parsing
* Normalization
* Filtration
* Enrichment (from constants, active lists, dictionaries, event fields)
* Mutation of normalized event fields
* Aggregation
* Sending normalized events to various destinations, like storage and correlators

### Correlator
Correlator is an executable which is responsible for:
* Real time correlation of normalized events (security incidents detection)
* Executing various user-defined actions, like active lists interactions, notifications, external processes
* Sending correlated events to various destinations, like storage and correlators

### Core
Core is an executable which is responsible for:
* Configuration management and deployment
* Event and active lists storage housekeeping
* Providing REST API for [Raccoon console](https://github.com/tephrocactus/raccoon-console)

### Bus
Bus is an executable which is responsible for:
* Passing normalized events from collectors to correlators
* Distributing configuration changes
* Other IPC tasks

Currently, we are using [NATS](https://nats.io/) to provide bus functionality. 

### Active lists storage
Active lists storage is an executable which is responsible for:
* Providing active lists operations
* Persisting and expiration of active list records

Currently, we are using [Redis](https://redis.io/) to provide active lists functionality..

### Event storage
Event storage is an executable which is responsible for:
* Storing normalized events (base and correlated)
* Providing search functionality for [Raccoon console](https://github.com/tephrocactus/raccoon-console)

Currently, we are using [Elasticsearch](https://www.elastic.co/products/elasticsearch) to provide event storage functionality.

## Modules overview

### Connector
Connector's goal is to actively fetch or passively receive log records. It can be used whithin [collector](#collector) and [correlator](#correlator).

### Normalizer
Normalier's goal is to parse and convert raw log records to normalized event according to mapping rules provided by user. It can be used within [collector](#collector) only.

### Destination
Destination's goal is to send normalized events to various endpoints. For example, to event storage or correlators.
It can be used whithin [collector](#collector) and [correlator](#correlator).

### 
