// clickhouse package define a client to send request to a ClickHouse server.
// It implements repository.Repository interface in order to run business
// logic related to our counter component.
//
// NOTE: ClickHouse does not fit as a counter store but the main purpose
// of this repository is to explore how to setup testcontainers
// to run integration tests with different kind of SGBD.
package clickhouse
