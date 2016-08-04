class Entity < ApplicationRecord
  belongs_to :project

  def save_and_touchfile
    if filepath.exist?
      logger.warn "Entity#save_and_touchfile: #{filepath} already exists"
      return false
    else
      mkdirp
      touchfile
      save
      return true
    end
  end

  # create parent directories
  def mkdirp
    logger.debug "Entity#mkdirp: #{filepath.dirname}"
    FileUtils.mkdir_p(filepath.dirname)
  end

  def touchfile
    logger.debug "Entity#touch: #{filepath}"
    FileUtils.touch(filepath)
  end

  def readfile
    logger.debug "Entity#readfile: #{filepath}"
    File.read(filepath)
  end

  def writefile(content)
    logger.debug "Entity#writefile: #{filepath}"
    File.write(filepath, content)
  end

  def to_param
    path
  end

  def readfile!
    assign_attributes :content => readfile
  end

  private

    def filepath
      Pathname(Editor::Application.config.workspace).join(project.name, path)
    end
end
