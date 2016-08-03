class Project < ApplicationRecord
  has_many :entities, :dependent => :destroy
  validates :name, uniqueness: true

  def mkdir
    workspace = Pathname(Editor::Application.config.workspace)
    FileUtils.mkdir_p workspace.join(name) # TODO: abstract
  end
end
