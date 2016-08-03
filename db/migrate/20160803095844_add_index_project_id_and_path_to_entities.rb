class AddIndexProjectIdAndPathToEntities < ActiveRecord::Migration[5.0]
  def change
    add_index :entities, [:project_id, :path], :unique => true, :name => 'project_entity_index'
  end
end
