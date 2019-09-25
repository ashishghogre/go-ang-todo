Feature: Todo List
    Scenario:
        When sending "POST" to "/item" with body:
            """
            {
                "title": "new item"
            }
            """
        Then should get response status code of "201"
        And sending "GET" to "/items" should return items containing:
            """
            new item
            """