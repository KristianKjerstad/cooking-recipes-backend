# Get started
* Make sure you have installed [air](https://github.com/air-verse/air) and [Go](https://go.dev/)

To start the app:
1. Start the mongo db server using docker-compose: in the terminal run `docker-compose up --build` 
2. go to the terminal and enter `air`. 
(Alternatively the app can be started using `go run server.go`) 


A GraphQL playground will now start at `localhost:8060`.



## Example Graph QL queries 
```GraphQL
mutation AddDog {
  createDog(input: {name: "x", isGoodBoi: true}) {
    _id
    name
    isGoodBoi
  }
}

query AllDogs {
  dogs {
    name
    _id
    isGoodBoi
    
  }
}


query getDog {
  dog(_id:"6672e1fc003309c534c5ff34") {
    name
    _id
    isGoodBoi
  }
}
```


## Modify the code

The Go code related to Graph QL is autogenerated. After making changes to the GraphQL schema in graph/schema.graphqls, open the terminal and run the command `go run github.com/99designs/gqlgen generate`.

The setup is based on [this guide](https://www.howtographql.com/graphql-go/1-getting-started/).