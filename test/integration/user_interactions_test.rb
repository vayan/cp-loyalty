require 'test_helper'

class UserInteractionsTest < ActionDispatch::IntegrationTest
  test 'can create an user' do
    get '/users/new'
    assert_response :success

    post '/users', params: { user: { loyalty_points: 11 } }
    assert_response :redirect
    follow_redirect!
    assert_response :success
    assert_equal 11, json_response['loyalty_points']
  end

  test 'redirect on error when creation a user' do
    get '/users/new'
    assert_response :success

    post '/users', params: { user: { loyalty_points: -1 } }
    assert_response :unprocessable_entity
  end

  test 'shows users loyalty points' do
    get user_path(users(:first_user).id)
    assert_response :success
    assert_equal 242, json_response['loyalty_points']
  end
end
