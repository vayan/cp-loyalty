# ITW CP Loyalty Go

# Pour run l'app

`go get github.com/onsi/ginkgo github.com/onsi/gomega github.com/gorilla/mux github.com/jinzhu/gorm github.com/jinzhu/gorm/dialects/sqlite`

`go build`

`./itw-cp-loyalty`

Et op! Dispo sur http://localhost:8080

# Description

Les routes sont dans `app.go` (`initializeRoutes`)

Les controllers sont dans `controllers.go`

Il y a 3 models

* `LoyaltyRank`
* `Ride`
* `User`

Dans `loyalty_rank.go`, `ride.go` et `user.go`

Tu peux run les tests avec `go test`
