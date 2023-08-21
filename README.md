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
        <details>
        <summary>response</summary>

        ```json
        {
          "data": {
            "userList": {
              "totalCount": 10,
              "users": [
                {
                  "id": 1,
                  "name": "A-san",
                  "email": "j@hoge.jp",
                  "birthday": "1997-01-01 00:00:00"
                },
                {
                  "id": 2,
                  "name": "B-san",
                  "email": "i@hoge.jp",
                  "birthday": "1998-01-02 00:00:00"
                },
                {
                  "id": 3,
                  "name": "C-san",
                  "email": "h@hoge.jp",
                  "birthday": "1999-01-03 00:00:00"
                },
                {
                  "id": 4,
                  "name": "D-san",
                  "email": "g@hoge.jp",
                  "birthday": "1991-01-04 00:00:00"
                },
                {
                  "id": 5,
                  "name": "E-san",
                  "email": "f@hoge.jp",
                  "birthday": "1995-01-05 00:00:00"
                },
                {
                  "id": 6,
                  "name": "F-san",
                  "email": "e@hoge.jp",
                  "birthday": "1996-01-06 00:00:00"
                },
                {
                  "id": 7,
                  "name": "G-san",
                  "email": "d@hoge.jp",
                  "birthday": "1992-01-07 00:00:00"
                },
                {
                  "id": 8,
                  "name": "H-san",
                  "email": "c@hoge.jp",
                  "birthday": "1993-01-08 00:00:00"
                },
                {
                  "id": 9,
                  "name": "I-san",
                  "email": "b@hoge.jp",
                  "birthday": "1994-01-09 00:00:00"
                },
                {
                  "id": 10,
                  "name": "J-san",
                  "email": "a@hoge.jp",
                  "birthday": "1990-01-10 00:00:00"
                }
              ]
            }
          }
        }
        ```
        </details>
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
        <details>
        <summary>response</summary>

        ```json
        {
          "data": {
            "userList": {
              "totalCount": 10,
              "users": [
                {
                  "id": 10,
                  "name": "J-san",
                  "email": "a@hoge.jp",
                  "birthday": "1990-01-10 00:00:00"
                },
                {
                  "id": 7,
                  "name": "G-san",
                  "email": "d@hoge.jp",
                  "birthday": "1992-01-07 00:00:00"
                },
                {
                  "id": 8,
                  "name": "H-san",
                  "email": "c@hoge.jp",
                  "birthday": "1993-01-08 00:00:00"
                },
                {
                  "id": 9,
                  "name": "I-san",
                  "email": "b@hoge.jp",
                  "birthday": "1994-01-09 00:00:00"
                },
                {
                  "id": 5,
                  "name": "E-san",
                  "email": "f@hoge.jp",
                  "birthday": "1995-01-05 00:00:00"
                }
              ]
            }
          }
        }
        ```
        </details>
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
        <details>
        <summary>response</summary>

        ```json
        {
          "data": {
            "userList": {
              "totalCount": 10,
              "users": [
                {
                  "id": 1,
                  "name": "A-san",
                  "email": "j@hoge.jp",
                  "birthday": "1997-01-01 00:00:00",
                  "friendsEach": [
                    {
                      "id": 2,
                      "name": "B-san",
                      "email": "i@hoge.jp",
                      "birthday": "1998-01-02 00:00:00"
                    },
                    {
                      "id": 3,
                      "name": "C-san",
                      "email": "h@hoge.jp",
                      "birthday": "1999-01-03 00:00:00"
                    },
                    {
                      "id": 4,
                      "name": "D-san",
                      "email": "g@hoge.jp",
                      "birthday": "1991-01-04 00:00:00"
                    },
                    {
                      "id": 5,
                      "name": "E-san",
                      "email": "f@hoge.jp",
                      "birthday": "1995-01-05 00:00:00"
                    },
                    {
                      "id": 10,
                      "name": "J-san",
                      "email": "a@hoge.jp",
                      "birthday": "1990-01-10 00:00:00"
                    }
                  ]
                },
                {
                  "id": 2,
                  "name": "B-san",
                  "email": "i@hoge.jp",
                  "birthday": "1998-01-02 00:00:00",
                  "friendsEach": [
                    {
                      "id": 1,
                      "name": "A-san",
                      "email": "j@hoge.jp",
                      "birthday": "1997-01-01 00:00:00"
                    }
                  ]
                },
                {
                  "id": 3,
                  "name": "C-san",
                  "email": "h@hoge.jp",
                  "birthday": "1999-01-03 00:00:00",
                  "friendsEach": [
                    {
                      "id": 1,
                      "name": "A-san",
                      "email": "j@hoge.jp",
                      "birthday": "1997-01-01 00:00:00"
                    }
                  ]
                },
                {
                  "id": 4,
                  "name": "D-san",
                  "email": "g@hoge.jp",
                  "birthday": "1991-01-04 00:00:00",
                  "friendsEach": [
                    {
                      "id": 1,
                      "name": "A-san",
                      "email": "j@hoge.jp",
                      "birthday": "1997-01-01 00:00:00"
                    }
                  ]
                },
                {
                  "id": 5,
                  "name": "E-san",
                  "email": "f@hoge.jp",
                  "birthday": "1995-01-05 00:00:00",
                  "friendsEach": [
                    {
                      "id": 1,
                      "name": "A-san",
                      "email": "j@hoge.jp",
                      "birthday": "1997-01-01 00:00:00"
                    }
                  ]
                },
                {
                  "id": 6,
                  "name": "F-san",
                  "email": "e@hoge.jp",
                  "birthday": "1996-01-06 00:00:00",
                  "friendsEach": [
                    {
                      "id": 7,
                      "name": "G-san",
                      "email": "d@hoge.jp",
                      "birthday": "1992-01-07 00:00:00"
                    },
                    {
                      "id": 8,
                      "name": "H-san",
                      "email": "c@hoge.jp",
                      "birthday": "1993-01-08 00:00:00"
                    }
                  ]
                },
                {
                  "id": 7,
                  "name": "G-san",
                  "email": "d@hoge.jp",
                  "birthday": "1992-01-07 00:00:00",
                  "friendsEach": [
                    {
                      "id": 6,
                      "name": "F-san",
                      "email": "e@hoge.jp",
                      "birthday": "1996-01-06 00:00:00"
                    }
                  ]
                },
                {
                  "id": 8,
                  "name": "H-san",
                  "email": "c@hoge.jp",
                  "birthday": "1993-01-08 00:00:00",
                  "friendsEach": [
                    {
                      "id": 6,
                      "name": "F-san",
                      "email": "e@hoge.jp",
                      "birthday": "1996-01-06 00:00:00"
                    }
                  ]
                },
                {
                  "id": 9,
                  "name": "I-san",
                  "email": "b@hoge.jp",
                  "birthday": "1994-01-09 00:00:00",
                  "friendsEach": []
                },
                {
                  "id": 10,
                  "name": "J-san",
                  "email": "a@hoge.jp",
                  "birthday": "1990-01-10 00:00:00",
                  "friendsEach": []
                }
              ]
            }
          }
        }
        ```
        </details>
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
        <details>
        <summary>response</summary>

        ```json
        {
          "data": {
            "userList": {
              "totalCount": 10,
              "users": [
                {
                  "id": 1,
                  "name": "A-san",
                  "email": "j@hoge.jp",
                  "birthday": "1997-01-01 00:00:00",
                  "friendsEach": [
                    {
                      "id": 2,
                      "name": "B-san",
                      "friendsEach": [
                        {
                          "id": 1,
                          "name": "A-san"
                        }
                      ]
                    },
                    {
                      "id": 3,
                      "name": "C-san",
                      "friendsEach": [
                        {
                          "id": 1,
                          "name": "A-san"
                        }
                      ]
                    },
                    {
                      "id": 4,
                      "name": "D-san",
                      "friendsEach": [
                        {
                          "id": 1,
                          "name": "A-san"
                        }
                      ]
                    },
                    {
                      "id": 5,
                      "name": "E-san",
                      "friendsEach": [
                        {
                          "id": 1,
                          "name": "A-san"
                        }
                      ]
                    },
                    {
                      "id": 10,
                      "name": "J-san",
                      "friendsEach": []
                    }
                  ]
                },
                {
                  "id": 2,
                  "name": "B-san",
                  "email": "i@hoge.jp",
                  "birthday": "1998-01-02 00:00:00",
                  "friendsEach": [
                    {
                      "id": 1,
                      "name": "A-san",
                      "friendsEach": [
                        {
                          "id": 2,
                          "name": "B-san"
                        },
                        {
                          "id": 3,
                          "name": "C-san"
                        },
                        {
                          "id": 4,
                          "name": "D-san"
                        },
                        {
                          "id": 5,
                          "name": "E-san"
                        },
                        {
                          "id": 10,
                          "name": "J-san"
                        }
                      ]
                    }
                  ]
                },
                {
                  "id": 3,
                  "name": "C-san",
                  "email": "h@hoge.jp",
                  "birthday": "1999-01-03 00:00:00",
                  "friendsEach": [
                    {
                      "id": 1,
                      "name": "A-san",
                      "friendsEach": [
                        {
                          "id": 2,
                          "name": "B-san"
                        },
                        {
                          "id": 3,
                          "name": "C-san"
                        },
                        {
                          "id": 4,
                          "name": "D-san"
                        },
                        {
                          "id": 5,
                          "name": "E-san"
                        },
                        {
                          "id": 10,
                          "name": "J-san"
                        }
                      ]
                    }
                  ]
                },
                {
                  "id": 4,
                  "name": "D-san",
                  "email": "g@hoge.jp",
                  "birthday": "1991-01-04 00:00:00",
                  "friendsEach": [
                    {
                      "id": 1,
                      "name": "A-san",
                      "friendsEach": [
                        {
                          "id": 2,
                          "name": "B-san"
                        },
                        {
                          "id": 3,
                          "name": "C-san"
                        },
                        {
                          "id": 4,
                          "name": "D-san"
                        },
                        {
                          "id": 5,
                          "name": "E-san"
                        },
                        {
                          "id": 10,
                          "name": "J-san"
                        }
                      ]
                    }
                  ]
                },
                {
                  "id": 5,
                  "name": "E-san",
                  "email": "f@hoge.jp",
                  "birthday": "1995-01-05 00:00:00",
                  "friendsEach": [
                    {
                      "id": 1,
                      "name": "A-san",
                      "friendsEach": [
                        {
                          "id": 2,
                          "name": "B-san"
                        },
                        {
                          "id": 3,
                          "name": "C-san"
                        },
                        {
                          "id": 4,
                          "name": "D-san"
                        },
                        {
                          "id": 5,
                          "name": "E-san"
                        },
                        {
                          "id": 10,
                          "name": "J-san"
                        }
                      ]
                    }
                  ]
                },
                {
                  "id": 6,
                  "name": "F-san",
                  "email": "e@hoge.jp",
                  "birthday": "1996-01-06 00:00:00",
                  "friendsEach": [
                    {
                      "id": 7,
                      "name": "G-san",
                      "friendsEach": [
                        {
                          "id": 6,
                          "name": "F-san"
                        }
                      ]
                    },
                    {
                      "id": 8,
                      "name": "H-san",
                      "friendsEach": [
                        {
                          "id": 6,
                          "name": "F-san"
                        }
                      ]
                    }
                  ]
                },
                {
                  "id": 7,
                  "name": "G-san",
                  "email": "d@hoge.jp",
                  "birthday": "1992-01-07 00:00:00",
                  "friendsEach": [
                    {
                      "id": 6,
                      "name": "F-san",
                      "friendsEach": [
                        {
                          "id": 7,
                          "name": "G-san"
                        },
                        {
                          "id": 8,
                          "name": "H-san"
                        }
                      ]
                    }
                  ]
                },
                {
                  "id": 8,
                  "name": "H-san",
                  "email": "c@hoge.jp",
                  "birthday": "1993-01-08 00:00:00",
                  "friendsEach": [
                    {
                      "id": 6,
                      "name": "F-san",
                      "friendsEach": [
                        {
                          "id": 7,
                          "name": "G-san"
                        },
                        {
                          "id": 8,
                          "name": "H-san"
                        }
                      ]
                    }
                  ]
                },
                {
                  "id": 9,
                  "name": "I-san",
                  "email": "b@hoge.jp",
                  "birthday": "1994-01-09 00:00:00",
                  "friendsEach": []
                },
                {
                  "id": 10,
                  "name": "J-san",
                  "email": "a@hoge.jp",
                  "birthday": "1990-01-10 00:00:00",
                  "friendsEach": []
                }
              ]
            }
          }
        }
        ```
        </details>
