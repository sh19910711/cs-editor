class Entity < ApplicationRecord
  belongs_to :project

  def touch
    workspace = Pathname(Editor::Application.config.workspace)
    logger.debug "touch: #{workspace.join(project.name, path)}"
    FileUtils.touch(workspace.join(project.name, path))
  end

  def readfile
    workspace = Pathname(Editor::Application.config.workspace)
    logger.debug "read: #{workspace.join(project.name, path)}"
    File.read(workspace.join(project.name, path))
  end

  def to_param
    path
  end
end
