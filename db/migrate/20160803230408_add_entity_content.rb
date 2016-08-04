class AddEntityContent < ActiveRecord::Migration[5.0]
  def change
    add_column :entities, :content, :text
  end
end
