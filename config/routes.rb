Rails.application.routes.draw do
  # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
  root 'projects#index'
  resources :projects do
    resources :entities, :param => :path
    get '/entities/*path', :to => 'entities#show', :as => 'Entity', :format => false
  end
end
