Feature: get version
  In order to verify external API calls
  As a test creator
  I need to be able to request json from httpbin


  Scenario: should get json from httpbin
    When I send "GET" request to external "https://httpbin.org/json"
    Then the external response code should be 200
    And the external response should match json:
      """
        {
          "slideshow": {
            "author": "Yours Truly",
            "date": "date of publication",
            "slides": [
              {
                "title": "Wake up to WonderWidgets!",
                "type": "all"
              },
              {
                "items": [
                  "Why <em>WonderWidgets</em> are great",
                  "Who <em>buys</em> WonderWidgets"
                ],
                "title": "Overview",
                "type": "all"
              }
            ],
            "title": "Sample Slide Show"
          }
        }
      """