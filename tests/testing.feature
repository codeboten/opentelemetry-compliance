Feature: Receive HTTP Server Span

  Scenario: Some scenario
    Given Receive a span
    When A span is received
    Then Span.name is "HTTP GET"
