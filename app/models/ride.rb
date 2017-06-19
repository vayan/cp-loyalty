class Ride < ApplicationRecord
  belongs_to :user

  validates :price, presence: true, numericality: { greater_than: 0 }
  validates :user, presence: true
end
