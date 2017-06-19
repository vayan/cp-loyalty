class UsersController < ApplicationController
  before_action :set_user, only: [:show]

  def show
    render json: @user
  end

  def new
    @user ||= User.new
  end

  def create
    @user = User.create(user_params)
    if @user.persisted?
      redirect_to user_path(@user)
    else
      render :new, status: :unprocessable_entity
    end
  end

private
  def user_params
    params.require(:user).permit(:loyalty_points)
  end

  def set_user
    @user = User.find(params[:id])
  end
end
