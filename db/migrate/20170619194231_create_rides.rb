class CreateRides < ActiveRecord::Migration[5.0]
  def change
    create_table :rides do |t|
      t.integer :price, null: false
      t.references :user

      t.timestamps
    end
  end
end
