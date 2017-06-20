# ITW CP Loyalty

J'ai fait chaque etape dans 1 commit.

Chaque commit a une description pour expliquer un peu.

J'ai prefere faire des formulaires pour les creations d'object c'est plus facile si tu veux tester l'app.

Mais les GET retourne du JSON.


# Pour run l'app

j'assume que tu as ruby deja install

`gem install bundler`

`bundle install`

`rake db:create`

`rake db:migrate`

`rake db:seed`

`rails s`

Et op! Dispo sur http://localhost:3000


Les routes : `config/routes.rb`

# Controllers

J'ai deux controllers :

* Pour les Rides : `app/controllers/rides_controller.rb`
* Pour les Users : `app/controllers/users_controller.rb`

## User

`/users/new`

Views basique pour creer un User pour tester l'app

`/users/:id`

Retourne les infos de l'user avec ces rides, loyalty rank et le nombre de rides a faire pour atteindre le prochain rank


## Ride

`/users/:user_id/rides/new`

View basique pour creer des rides pour un User pour tester l'app.

Les loyalty points et rank son update a la creation du Ride par deux
services :

`app/services/loyalty_point_updater.rb`

`app/services/loyalty_rank_updater.rb`

C'est pour eviter d'avoir trop de logique dans le controller.


# Model

Il y a 3 models

* `LoyaltyRank`
* `Ride`
* `User`

Dans `app/models`

Ils ont tous les validations necessaires.


# Tests

Tu peux run les tests facilement avec `rails test`

Le code des test est dans `test/integration/`

J'ai fait que des test d'integration pour valider le fonctionement de l'app
