Feature: HTTP Server Span

  Scenario: Server Span Name with available route
    Given There is an HTTP server with a low cardinality route available.
    And That low cardinality route available is "abc".
    And There is an OpenTelemetry HTTP instrumentation for that server.
    When The instrumentation creates a Server Span for a GET operation.
    Then The Server Span Name SHOULD be "HTTP GET abc".

  Scenario: Server Span Name without an available route
    Given There is an HTTP server without a low cardinality route available.
    And There is an OpenTelemetry HTTP instrumentation for that server.
    When The instrumentation creates a Server Span for a GET operation.
    Then The Server Span Name SHOULD be "HTTP GET".
