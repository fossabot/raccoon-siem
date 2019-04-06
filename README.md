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
  * [Console](#console)

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
* Executing user-defined actions, like active lists operations, notifications, calling external processes
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

Currently, Raccoon is using [NATS](https://nats.io/) to provide bus functionality. 

### Active lists storage
Active lists storage is an executable which is responsible for:
* Providing active lists operations
* Persisting and expiration of active list records

Currently, Raccoon is using [Redis](https://redis.io/) to provide active lists functionality..

### Event storage
Event storage is an executable which is responsible for:
* Storing normalized events (base and correlated)
* Providing search functionality for [Raccoon console](https://github.com/tephrocactus/raccoon-console)
via [Raccoon Core](#core)

Currently, Raccoon is using [Elasticsearch](https://www.elastic.co/products/elasticsearch) to provide event storage functionality.

### Console
Console is a GUI tool which allows:
* Compose and deploy configuration
* Search for events
* Manage users, assets, active lists and dictionaries
* Create reports

Currently, [console](https://github.com/tephrocactus/raccoon-console) is in early development stage.

## Architecture overview
Raccoon SIEM has quite flexible architecture which can be simplified or extended to meet your needs. For example, if you don't need correlation, you can skip [correlator](#correlator) deployment. Or, in case you need some extra processing, you can attach your services to [Bus](#bus) or configure [collector](#collector)/[correlator](#correlator) to output events to your service or just fetch normalized events from [event storage](#event-storage).

One common setup example might look like this:

![Architecture overview](https://github.com/tephrocactus/raccoon-siem/blob/master/docs/example_arch.png)

## Entities overview

### Connector
Connector's goal is to actively fetch or passively receive log records. It can be used whithin [collector](#collector) and [correlator](#correlator).

### Normalizer
Normalier's goal is to parse and convert raw log records to normalized event according to mapping rules provided by user. It can be used within [collector](#collector) only.

### Filter
Filter can be used whithin [collector](#collector) and [correlator](#correlator) to:
* Drop undesired (noisy) events
* Provide event selection mechanism for enrichment, aggregation and correlation rules.

### Mapping rule
Mapping rule is used to tie up raw log field or active list record field with normalized event field.

### Enrichment rule
Enrichment rule can be used whithin [collector](#collector) and [correlator](#correlator) to fill normalized events with additional data which can be taken from: constants, active list and dictionary records, event fields and third-party systems.

### Mutation rule
Mutation rule can be used within enrichment rules to alter values of event fields in a various ways.

### Aggregation rule
Aggregation rule can be used whithin [collector](#collector) to represent multiple identical or similar normalized events as single event with abillity to sum or concatenate values of original event fields. This technique may drammaticaly minimize collector output. 

### Correlation rule
Correlation rule can be used whithin [correlator](#correlator) to describe the signature of information security incidents.
That signature can be based on single or multiple normalized events.

### Action
Action allow user to react to correlation rule triggers: enrich correlated events, interact with active lists and third-party systems, send notifications, call executables and so on.

### Dictionary
Dictionary is a component's local static data source filled by user which can be used within enrichment rules to alter the values of event fields.

### Active list
Active list is a remote dynamic data source filled by user or correlation rule which can be used to represent some state (session tracking, host or user status and so on) or a as remote dictionary. Active list operations can take place within filters and actions.


### Destination
Destination's goal is to send normalized events to various endpoints. For example, to event storage or correlators.
It can be used whithin [collector](#collector) and [correlator](#correlator).
