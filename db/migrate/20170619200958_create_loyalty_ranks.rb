class CreateLoyaltyRanks < ActiveRecord::Migration[5.0]
  def change
    create_table :loyalty_ranks do |t|
      t.string :name, null: false
      t.integer :required_rides_count, null: false, default: 0

      t.timestamps
    end
  end
end
