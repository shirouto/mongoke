package testutil

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	"github.com/graphql-go/graphql"
	goke "github.com/remorses/goke/src"
	"go.mongodb.org/mongo-driver/bson"
)

const MONGODB_URI = "mongodb://localhost/testdb"
const FIRESTORE_PROJECT_ID = "example"

func Bsonify(t *testing.T, x interface{}) map[string]interface{} {
	b, err := bson.Marshal(x)
	if err != nil {
		t.Error(err)
	}
	var d map[string]interface{}
	err = bson.Unmarshal(b, &d)
	if err != nil {
		t.Error(err)
	}
	return d
}

func Query(t *testing.T, schema graphql.Schema, query string, root goke.Map) interface{} {
	res := graphql.Do(graphql.Params{Schema: schema, RequestString: query, Context: context.Background(), RootObject: root})
	if res.Errors != nil && len(res.Errors) > 0 {
		t.Error(res.Errors[0])
		return nil
	}
	return res.Data
}

// func QuerySchema(t *testing.T, schema graphql.Schema, query string) interface{} {
// 	return Query(t, schema, query, nil)
// }

func QueryShouldFail(t *testing.T, schema graphql.Schema, query string, root goke.Map) error {
	res := graphql.Do(graphql.Params{Schema: schema, RequestString: query, RootObject: root})
	if res.Errors != nil && len(res.Errors) > 0 {
		return res.Errors[0]
	}
	t.Fatal("query should have failed")
	return nil
}

func PrettyPrint(x ...interface{}) {
	for _, x := range x {
		json, err := json.MarshalIndent(x, "", "   ")
		if err != nil {
			panic(err)
		}
		println(string(json))
	}
	println()
}

func Pretty(x ...interface{}) string {
	res := ""
	for _, x := range x {
		json, err := json.MarshalIndent(x, "", "   ")
		if err != nil {
			panic(err)
		}
		res += string(json)
		res += "\n"
	}
	return res
}

func ConvertToPlainMap(in interface{}) map[string]interface{} {
	var inInterface map[string]interface{}
	inrec, err := json.Marshal(in)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(inrec, &inInterface)
	return inInterface
}

func FormatJson(t *testing.T, s string) string {
	var expectedObj map[string]interface{}
	err := json.Unmarshal([]byte(s), &expectedObj)
	if err != nil {
		t.Error(err)
	}
	return Pretty(expectedObj)
}

var UserSchema = `
type User {
	name: String
	age: Int
}
`

var UserQuery1 = `
{
	findOneUser(where: {name: {eq: "dsf"}}) {
		name
		age
	}
}`

var UserQuery2 = `
{
	findManyUser(last: 1, where: {name: {eq: "sdfsdf"}}) {
	  nodes {
		name
		age
	  }
	  pageInfo {
		endCursor
		hasNextPage
		hasPreviousPage
		startCursor
	  }
	}
}`

var YamlConfig = `

jwt:
    type: H256

schema: |
    scalar Address
    scalar Url
    # scalar ObjectId
    type Task {
        _id: ObjectId
        address: Address
    }
    type WindowedEvent {
        value: Int
        timestamp: Int
    }
    type Guest {
        type: String
        _id: ObjectId
        name: String
    }
    enum Letter {
        a
        b
        c
    }
    type User {
        type: String
        _id: ObjectId
        name: String
        surname: String
        friends_ids: [ObjectId]
        url: Url
        letter: Letter
    }
    union Human = User | Guest

types:
    Task:
        collection: tasks
        exposed: false
    User:
        collection: users
        

relations:
    -   from: Task
        to: WindowedEvent
        type: to_many
        field: events
        where: {}

`

const IntrospectionQuery = `
  query IntrospectionQuery {
    __schema {
      queryType { name }
      mutationType { name }
      subscriptionType { name }
      types {
        ...FullType
      }
      directives {
        name
        description
        locations
        args {
          ...InputValue
        }
        # deprecated, but included for coverage till removed
        onOperation
        onFragment
        onField
      }
    }
  }
  fragment FullType on __Type {
    kind
    name
    description
    fields(includeDeprecated: true) {
      name
      description
      args {
        ...InputValue
      }
      type {
        ...TypeRef
      }
      isDeprecated
      deprecationReason
    }
    inputFields {
      ...InputValue
    }
    interfaces {
      ...TypeRef
    }
    enumValues(includeDeprecated: true) {
      name
      description
      isDeprecated
      deprecationReason
    }
    possibleTypes {
      ...TypeRef
    }
  }
  fragment InputValue on __InputValue {
    name
    description
    type { ...TypeRef }
    defaultValue
  }
  fragment TypeRef on __Type {
    kind
    name
    ofType {
      kind
      name
      ofType {
        kind
        name
        ofType {
          kind
          name
          ofType {
            kind
            name
            ofType {
              kind
              name
              ofType {
                kind
                name
                ofType {
                  kind
                  name
                }
              }
            }
          }
        }
      }
    }
  }
`
