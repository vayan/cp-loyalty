Rails.application.routes.draw do
  resources :users, only: [:new, :create, :show] do
    resources :rides, only: [:new, :create]
  end
end
