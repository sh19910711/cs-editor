class CreateEntities < ActiveRecord::Migration[5.0]
  def change
    create_table :entities do |t|
      t.string :path
      t.integer :project_id

      t.timestamps
    end
  end
end
