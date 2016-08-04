class AddProjectName < ActiveRecord::Migration[5.0]
  def change
    add_column :projects, :name, :string, :unique => true
  end
end
