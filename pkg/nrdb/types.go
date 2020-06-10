// Code generated by typegen; DO NOT EDIT.

package nrdb

import (
	"github.com/newrelic/newrelic-client-go/internal/serialization"
)

/* EpochMilliseconds - The `EpochMilliseconds` scalar represents the number of milliseconds since the Unix epoch */
type EpochMilliseconds serialization.EpochTime

/* EventAttributeDefinition - A human-readable definition of an NRDB Event Type Attribute */
type EventAttributeDefinition struct {
	/* Category - This attribute's category */
	Category string `json:"category"`
	/* Definition - A short description of this attribute */
	Definition string `json:"definition"`
	/* DocumentationUrl - The New Relic docs page for this attribute */
	DocumentationUrl string `json:"documentationUrl"`
	/* Label - The human-friendly formatted name of the attribute */
	Label string `json:"label"`
	/* Name - The name of the attribute */
	Name string `json:"name"`
}

/* EventDefinition - A human-readable definition of an NRDB Event Type */
type EventDefinition struct {
	/* Attributes - A list of attribute definitions for this event type */
	Attributes []EventAttributeDefinition `json:"attributes"`
	/* Definition - A short description of this event */
	Definition string `json:"definition"`
	/* Label - The human-friendly formatted name of the event */
	Label string `json:"label"`
	/* Name - The name of the event */
	Name string `json:"name"`
}

/* NrdbMetadata - An object containing metadata about the query and result. */
type NrdbMetadata struct {
	/* EventTypes - A list of the event types that were queried. */
	EventTypes []string `json:"eventTypes"`
	/* Facets - A list of facets that were queried. */
	Facets []string `json:"facets"`
	/* Messages - Messages from NRDB included with the result. */
	Messages []string `json:"messages"`
	/* TimeWindow - Details about the query time window. */
	TimeWindow NrdbMetadataTimeWindow `json:"timeWindow"`
}

/* NrdbMetadataTimeWindow - An object representing details about a query's time window. */
type NrdbMetadataTimeWindow struct {
	/* Begin - Timestamp marking the query begin time. */
	Begin EpochMilliseconds `json:"begin"`
	/* CompareWith - A clause representing the comparison time window. */
	CompareWith string `json:"compareWith"`
	/* End - Timestamp marking the query end time. */
	End EpochMilliseconds `json:"end"`
	/* Since - SINCE clause resulting from the query */
	Since string `json:"since"`
	/* Until - UNTIL clause resulting from the query */
	Until string `json:"until"`
}

/* NrdbResult - This scalar represents a NRDB Result. It is a `Map` of `String` keys to values.

The shape of these objects reflect the query used to generate them, the contents
of the objects is not part of the GraphQL schema. */
type NrdbResult map[string]interface{}

/* NrdbResultContainer - A data structure that contains the results of the NRDB query along
with other capabilities that enhance those results.

Direct query results are available through `results`, `totalResult` and
`otherResult`. The query you made is accessible through `nrql`, along with
`metadata` about the query itself. Enhanced capabilities include
`eventDefinitions`, `suggestedFacets` and more. */
type NrdbResultContainer struct {
	/* CurrentResults - In a `COMPARE WITH` query, the `currentResults` contain the results for the current comparison time window. */
	CurrentResults []NrdbResult `json:"currentResults"`
	/* EmbeddedChartUrl - Generate a publicly sharable Embedded Chart URL for the NRQL query.

	For more details, see [our docs](https://docs.newrelic.com/docs/apis/graphql-api/tutorials/query-nrql-through-new-relic-graphql-api#embeddable-charts). */
	EmbeddedChartUrl string `json:"embeddedChartUrl"`
	/* EventDefinitions - Retrieve a list of event type definitions, providing descriptions
	of the event types returned by this query, as well as details
	of their attributes. */
	EventDefinitions []EventDefinition `json:"eventDefinitions"`
	/* Metadata - Metadata about the query and result. */
	Metadata NrdbMetadata `json:"metadata"`
	/* Nrql - The [NRQL](https://docs.newrelic.com/docs/insights/nrql-new-relic-query-language/nrql-resources/nrql-syntax-components-functions) query that was executed to yield these results. */
	Nrql Nrql `json:"nrql"`
	/* OtherResult - In a `FACET` query, the `otherResult` contains the aggregates representing the events _not_
	contained in an individual `results` facet */
	OtherResult NrdbResult `json:"otherResult"`
	/* PreviousResults - In a `COMPARE WITH` query, the `previousResults` contain the results for the previous comparison time window. */
	PreviousResults []NrdbResult `json:"previousResults"`
	/* Results - The query results. This is a flat list of objects who's structure matches the query submitted. */
	Results []NrdbResult `json:"results"`
	/* StaticChartUrl - Generate a publicly sharable static chart URL for the NRQL query. */
	StaticChartUrl string `json:"staticChartUrl"`
	/* SuggestedFacets - Retrieve a list of suggested NRQL facets for this NRDB query, to be used with
	the `FACET` keyword in NRQL.

	Results are based on historical query behaviors.

	If the query already has a `FACET` clause, it will be ignored for the purposes
	of suggesting facets.

	For more details, see [our docs](https://docs.newrelic.com/docs/apis/graphql-api/tutorials/nerdgraph-graphiql-nrql-tutorial#suggest-facets). */
	SuggestedFacets []NrqlFacetSuggestion `json:"suggestedFacets"`
	/* SuggestedQueries - Suggested queries that could help explain an anomaly in your timeseries based on either statistical differences in the data or historical usage.

	If no `anomalyTimeWindow` is supplied, we will attempt to detect a spike in the NRQL results. If no spike is found, the suggested query results will be empty.

	Input NRQL must be a TIMESERIES query and must have exactly one result. */
	SuggestedQueries SuggestedNrqlQueryResponse `json:"suggestedQueries"`
	/* TotalResult - In a `FACET` query, the `totalResult` contains the aggregates representing _all_ the events,
	whether or not they are contained in an individual `results` facet */
	TotalResult NrdbResult `json:"totalResult"`
}

/* Nrql - This scalar represents a NRQL query string.

See the [NRQL Docs](https://docs.newrelic.com/docs/insights/nrql-new-relic-query-language/nrql-resources/nrql-syntax-components-functions) for more information about NRQL syntax. */
type Nrql string

/* NrqlFacetSuggestion - A suggested NRQL facet. Facet suggestions may be either a single attribute, or
a list of attributes in the case of multi-attribute facet suggestions. */
type NrqlFacetSuggestion struct {
	/* Attributes - A list of attribute names comprising the suggested facet.

	Raw attribute names will be returned here. Attribute names may need to be
	backtick-quoted before inclusion in a NRQL query. */
	Attributes []string `json:"attributes"`
	/* Nrql - A modified version of the input NRQL, with a `FACET ...` clause appended.
	If the original NRQL had a `FACET` clause already, it will be replaced. */
	Nrql Nrql `json:"nrql"`
}

/* SuggestedNrqlQuery - Interface type representing a query suggestion. */
type SuggestedNrqlQuery interface{}

/* SuggestedNrqlQueryResponse - result type encapsulating suggested queries */
type SuggestedNrqlQueryResponse struct {
	/* Suggestions - List of suggested queries. */
	Suggestions []SuggestedNrqlQuery `json:"suggestions"`
}