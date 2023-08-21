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
            userList {
                totalCount
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
            userList(param: {
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
                    birthday: ASC
                }
                limit: 5
                offset: 0
            }) {
                totalCount
                users {
                    id
                    name
                    email
                    birthday
                }
            }
        }
        ```
   - Get all user and friends
        ```graphql
        query UserListAndFriends{
            userList {
                totalCount
                users {
                    id
                    name
                    email
                    birthday
                    friendsEach {
                        id
                        name
                        email
                        birthday  
                    }
                }
            }
        }
        ```
   - Get all user and friends more
        ```graphql
        query UserListAndFriendsMore{
            userList {
                totalCount
                users {
                    id
                    name
                    email
                    birthday
                    friendsEach {
                        id, name
                        friendsEach {
                            id, name
                        }
                    }
                }
            }
        }
        ```
