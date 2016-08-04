class Project < ApplicationRecord
  has_many :entities, :dependent => :destroy
  validates :name, uniqueness: true

  def mkdir
    FileUtils.mkdir_p dirpath
  end

  def rmdir
    FileUtils.rm_rf dirpath
  end

  private

    def dirpath
      Pathname(Editor::Application.config.workspace).join(name)
    end
end
