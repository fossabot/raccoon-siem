# Raccoon SIEM
Raccoon is an open source SIEM designed for high traffic volume environments.

> The project is in active development stage and not yet ready for production usage. 
If you want to give it a try - use [releases](https://github.com/tephrocactus/raccoon-siem/releases)
since master branch is unstable.

![Raccoon logo](https://github.com/tephrocactus/raccoon-siem/blob/master/docs/logo_v2.png)

## Table of contents
* [Components overview](#components-overview)
  * [Collector](#collector)
  * [Correlator](#correlator)
  * [Core](#core)
  * [Console](#console)
  * [Bus](#bus)
  * [Active lists storage](#active-lists-storage)
  * [Event storage](#event-storage)
* [Entities overview](#entities-overview)
  * [Connector](#connector)
  * [Normalizer](#normalizer)
  * [Normalized event](#normalized-event)
  * [Filter](#filter)
  * [Mapping rule](#mapping-rule)
  * [Enrichment rule](#enrichment-rule)
  * [Mutation rule](#mutation-rule)
  * [Aggregation rule](#aggregation-rule)
  * [Correlation rule](#correlation-rule)
  * [Event selector](#event-selector)
  * [Action](#action)
  * [Dictionary](#dictionary)
  * [Active list](#active-list)
  * [Destination](#destination)
* [Architecture overview](#architecture-overview)
* [Collector processing flow](#collector-processing-flow)
* [Correlator processing flow](#correlator-processing-flow)
* [Configuration](#configuration)

## Components overview

### Collector
Collector is an executable which is responsible for:
* [Collecting logs](#connector) from various systems
* [Parsing and normalization](#normalizer)
* [Filtration](#filter)
* [Enrichment](#enrichment-rule) and [mutation](#mutation-rule) of [normalized event](#normalized-event) fields
* [Aggregation](#aggregation-rule)
* Sending [normalized events](#normalized-event) to various [destinations](#destination)

> Development status: beta.

### Correlator
Correlator is an executable which is responsible for:
* Real time correlation of [normalized events](#normalized-event) (security incidents detection)
* Executing user-defined [actions](#action): 
[active lists](#active-list) operations, 
notifications, third-party process calls
* Sending [correlated events](#normalized-event) to various [destinations](#destination)

> Development status: beta.

### Core
Core is an executable which is responsible for:
* Configuration management and deployment
* Asset management
* Event storage and active lists storage housekeeping
* Providing REST API for [console](#console) and your custom services

> Development status: alpha.

### Console
Console is a GUI tool which allows you to:
* Compose and deploy configuration
* Search for [events](#normalized-event)
* Manage searches, users, assets, [active lists](#active-list) and [dictionaries](#dictionary)
* Create reports

> Development status: alpha. [Github project](https://github.com/tephrocactus/raccoon-console).

### Bus
Bus is an executable which is responsible for:
* Passing [normalized events](#normalized-event) from [collectors](#collector) to [correlators](#correlator)
* Distributing configuration changes
* Other IPC tasks

Currently, Raccoon is using [NATS](https://nats.io/) to provide bus functionality.

> Development status: third-party product.

### Active lists storage
Active lists storage is an executable which is responsible for:
* Providing [active lists](#active-list) operations
* Persisting and expiration of [active list](#active-list) records

Currently, Raccoon is using [Redis](https://redis.io/) to provide active lists functionality.

> Development status: third-party product.

### Event storage
Event storage is an executable which is responsible for:
* Storing [normalized events](#normalized-event) (base and correlated)
* Providing search functionality for [console](#console)
via [core](#core)

Currently, Raccoon is using [Elasticsearch](https://www.elastic.co/products/elasticsearch) to provide event storage functionality.

> Development status: third-party product.

## Entities overview

### Connector
Connector's goal is to actively fetch or passively receive raw log records or [normalized events](#normalized-event). It can be used whithin [collector](#collector) and [correlator](#correlator).

> Currently implemented connectors: 
> * TCP/UDP listener
> * Netflow v9
> * [NATS](https://nats.io/) ([bus](#bus))
> * [Kafka](https://kafka.apache.org/)

> TODO:
> * Windows event log (currently, [Beats](https://www.elastic.co/products/beats) + [Kafka](https://kafka.apache.org/) can be used)
> * Text files
> * [MSSQL](https://www.microsoft.com/en-us/sql-server/)
> * [MySQL](https://www.mysql.com/)
> * [PostgreSQL](https://www.postgresql.org/)
> * [CockroachDB](https://www.cockroachlabs.com/)
> * [Oracle](https://www.oracle.com/database/)
> * [Elasticsearch](https://www.elastic.co/products/elasticsearch)
> * [MongoDB](https://www.mongodb.com/)
> * [Cassandra](http://cassandra.apache.org/)
> * [ScyllaDB](https://www.scylladb.com/)
> * [Redis](https://redis.io/)
> * [Tarantool](https://www.tarantool.io/ru/)
> * [RabbitMQ](https://www.rabbitmq.com/)

### Normalizer
Normalier's goal is to parse and convert raw log records to [normalized event](#normalized-event) according to [mapping rules](#mapping-rule) provided by user. It can be used within [collector](#collector) only.

> Currently implemented normalizers: 
> * JSON
> * CEF
> * CSV/TSV (with configurable delimiter)
> * Key/Value (with configurable delimiter)
> * Regexp
> * Syslog (RFC3164 & RFC5424)

> TODO:
> * XML

### Normalized event
Normalized event is a special structure wich defines a static [set of fields](https://github.com/tephrocactus/raccoon-siem/blob/master/sdk/normalization/event.go#L32) available for mapping, comparission, e.t.c. It is passed over the network (between Raccoon components) in JSON format. Normalized event can represent *base*, *aggregated* or *correlated* information security event.

### Filter
Filter can be used whithin [collector](#collector) and [correlator](#correlator) to:
* Drop undesired (noisy) events
* Provide event selection mechanism for 
[enrichment](#enrichment-rule),
[aggregation](#aggregation-rule) and 
[correlation](#correlation-rule) rules.

> Currently implemented comparission operators: 
> * **=** , **!=** , **>** , **>=**, **<** , **<=**
> * **inSubnet** , **!inSubnet**
> * **contains** , **!contains**

> Currently implemented comparission sources: 
> * [Active lists](#active-list)
> * [Normalization event's](#normalization-event) fields
> * Constant

### Mapping rule
Mapping rule is used to describe a relation between raw log field (or [active list](#active-list) record field) and [normalized event](#normalized-event) field.

### Enrichment rule
Enrichment rule can be used whithin [collector](#collector) and [correlator](#correlator) to fill [normalized events](#normalized-event) with additional data which can be taken from various sources.

> Currently implemented enrichment sources: 
> * [Dictionaries](#dictionary)
> * [Active lists](#active-list)
> * [Normalized event's](#normalized-event) fields
> * Constant

> TODO:
> * [Vulners](https://vulners.com/landing)
> * Custom HTTP/HTTPS API's

### Mutation rule
Mutation rule can be used within [enrichment rules](#enrichment-rule) to alter values of event fields in a various ways.

> Currently implemented mutators:
> * Regexp
> * Substring
> * Lowercase

### Aggregation rule
Aggregation rule can be used whithin [collector](#collector) to represent multiple identical or similar [normalized events](#normalized-event) as single event with abillity to sum or concatenate values of original event fields. This technique may drammaticaly minimize [collector](#collector) output. 

### Correlation rule
Correlation rule can be used whithin [correlator](#correlator) to describe the signature of information security incidents.
That signature can be based on single or multiple [normalized events](#normalized-event).

### Event selector
Event selectors are used to describe [normalized events](#normalized-event) expected by [correlation rules](#correlation-rule).

### Action
Action allow user to define reactions to [correlation rule](#correlation-rule) triggers.

> Currently implemented actions:
> * [Correlated event](#normalization-event) [enrichment](#enrichment-rule)
> * Set/Del operations with [active lists](#active list)

> TODO:
> * Notifications
> * External process calls
> * Custom HTTP/HTTPS API's interaction 

### Dictionary
Dictionary is a component's local static data source filled by user which can be used within [enrichment rules](#enrichment-rule) to alter the values of [normalized event](#normalized-event) fields.

### Active list
Active list is a remote dynamic data source filled by user or [correlation rule](#correlation-rule) which can be used to represent some state (session tracking, host or user status, e.t.c) or a as remote [dictionary](#dictionary). Active list operations can take place within [filters](#filter) and [actions](#action).

### Destination
Destination's goal is to send [normalized events](#normalized-event) to various endpoints. 
For example, to [event storage](#event-storage) or [correlators](#correlator).
It can be used whithin [collector](#collector) and [correlator](#correlator).

> Currently implemented destinations:
> * [Elasticsearch](https://www.elastic.co/products/elasticsearch) ([event storage](#event-storage))
> * [NATS](https://nats.io) ([bus](#bus))

> TODO:
> * [Kafka](https://kafka.apache.org/)
> * [RabbitMQ](https://www.rabbitmq.com/)
> * TCP/UDP socket

## Architecture overview
Raccoon SIEM has quite flexible architecture which can be simplified or extended to meet your needs. For example, if you don't need correlation, you can skip [correlator](#correlator) deployment. Or, in case you need some extra processing, you can attach your services to [bus](#bus) or configure [collector](#collector)/[correlator](#correlator) to output events to your service or just fetch [normalized events](#normalized-event) from [event storage](#event-storage).

One common setup example might look like this:

![Architecture overview](https://github.com/tephrocactus/raccoon-siem/blob/master/docs/architecture_overview.png)

## Collector processing flow
1. First of all, [collector](#collector) needs to actively fetch or passively receive raw logs. For that purpose various [connectors](#connector) are used. 
1. After log record was grabbed by a connector, it is passed to one of the processing workers in a round-robin fashion. 
1. Inside processing worker a log record is passed to a [normalizer](#normalizer), which parses and converts it to a [normalized event](#normalized-event) according to the user-defined [mapping rules](#mapping-rule). 
1. Next, a [normalized event](#normalized-event) is passed through the [filters](#filter), if any defined. If an [event](#normalization-event) did not pass all the [filters](#filter), it is dropped. 
1. An [event](#normalization-event) is passed through the [enrichment rules](#enrichment-rule), if any defined. 
Every [enrichment rule](#enrichment-rule) can have a [filter](#filter) in front of it. 
1. An [event](#normalization-event) is passed through the [aggregation rules](#aggregation-rule), if any defined.
Every [aggregation rule](#aggregation-rule) can have a [filter](#filter) in front of it. 
1. Finally, if an [event](#normalization-event) wasn't aggregated, it is sent to each [destination](#destination) defined by user immediately. 
Otherwise it will be sent after the aggregation timeout or threshold exceeded. 

![Collector processing flow](https://github.com/tephrocactus/raccoon-siem/blob/master/docs/collector_processing_flow.png)

## Correlator processing flow
1. First of all, [correlator](#correlator) needs to receive [normalized events](#normalized-event) from [bus](#bus) via [connector](#connector).
1. After [normalized event](#normalized-event) was received, it is passed to one of the processing workers in a round-robin fashion.
1. Inside processing worker [normalized event](#normalized-event) is passed through defined [correlation rules](#correlation-rule). Every [correlation rule](#correlation-rule) can have single or multiple [event selectors](#event-selector) with [filter](#filter) on-board.
1. If [normalized event](#normalized-event) wasn't selected by any of [correlation rules](#correlation-rule), it is dropped. We assume that [collector](#collector) has already put [base events](#normalized-event) to the [event storage](#event-storage).
1. If [correlated event](#normalized-event) was created by any of [correlation rules](#correlation-rule), it is linked with [base events](#normalized-event) and sent to each [destination](#destination) defined by user immediately. [Base events](#normalized-event) itself are dropped.

![Correlator processing flow](https://github.com/tephrocactus/raccoon-siem/blob/master/docs/correlator_processing_flow.png)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Ftephrocactus%2Fraccoon-siem.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Ftephrocactus%2Fraccoon-siem?ref=badge_shield)

## Configuration
> Comming soon...


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Ftephrocactus%2Fraccoon-siem.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Ftephrocactus%2Fraccoon-siem?ref=badge_large)