require 'test_helper'

class RidesInteractionsTest < ActionDispatch::IntegrationTest
  test 'can create an user ride' do
    get new_user_ride_path(users(:first_user).id)
    assert_response :success

    post user_rides_path(users(:first_user).id), params: { ride: { price: 42 } }
    assert_response :redirect
    follow_redirect!
    assert_response :success
    assert_equal 42, json_response['rides'][0]['price']
    assert_equal 4, json_response['rides_left_before_next_rank']
  end

  test 'redirect on error when creating a ride' do
    get new_user_ride_path(users(:first_user).id)
    assert_response :success

    post user_rides_path(users(:first_user).id), params: { ride: { price: 0 } }
    assert_response :unprocessable_entity
  end

  test 'raise users loyalty rank when needed' do
    post user_rides_path(users(:big_user).id), params: { ride: { price: 42 } }
    assert_response :redirect
    follow_redirect!
    assert_response :success
    assert_equal 'silver', json_response['loyalty_rank']['name']
  end

  test 'raise users loyalty points' do
    post user_rides_path(users(:new_user).id), params: { ride: { price: 42 } }
    post user_rides_path(users(:new_user).id), params: { ride: { price: 42 } }
    assert_response :redirect
    follow_redirect!
    assert_response :success
    assert_equal 84, json_response['loyalty_points']
  end
end
