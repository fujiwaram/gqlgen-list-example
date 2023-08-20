# gqlgen-list-example
gqlgen(GraphQL) list example

## Usage

1. Build and start application
    ```
    make up
    ```
2. Open GraphiQL web page  
    http://localhost:8080/
3. Execute GraphQL command
   - Get all user list
        ```graphql
        query UserList{
            users {
                users {
                    id
                    name
                    email
                    birthday
                }
            }
        }
        ```
   - Get conditional user list
        ```graphql
        query UserListWithCondition{
            users(param: {
                filter: {
                    id: {
                        value: 5
                        condition: GTE
                    }
                    birthday: {
                        value: "2000-01-01 00:00:01"
                        condition: LTE
                    }
                }
                sort: {
                    birthday: DESC
                    id: ASC
                }
                limit: 5
                offset: 0
            }) {
                users {
                    id
                    name
                    email
                    birthday
                }
            }
        }
        ```