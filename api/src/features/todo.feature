Feature: Todo List
    Scenario: Create new item in database
        When sending "POST" to "/item" with body:
            """
            {
                "title": "new item",
                "details":"new details"
            }
            """
        Then should get response status code of "201"
        And sending "GET" to "/items" should return items containing:
            """
            new item
            """
        And sending "GET" to "/items" should return items containing:
            """
            new details
            """
    Scenario: Update new item in database
        When sending "PUT" to "/item/1" with body:
            """
            {
                "title": "updated item",
                "details":"updated details"
            }
            """
        Then should get response status code of "204"
        And sending "GET" to "/item/1" should return items containing:
            """
            updated item
            """
        And sending "GET" to "/item/1" should return items containing:
            """
            updated details
            """