Feature: Todo List
    Scenario:
        When sending "POST" to "/item" with body:
            """
            {
                "title": "new item"
            }
            """
        Then should get response status code of "200"
