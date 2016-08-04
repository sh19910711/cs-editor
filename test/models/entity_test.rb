require 'test_helper'

class EntityTest < ActiveSupport::TestCase
  test "#touchfile creates file if not exists" do
    p = Project.new(name: 'touchfile1')
    m = p.entities.new(path: 'path/to/file')

    m.mkdirp
    m.touchfile

    assert File.exists?('test/workspace/touchfile1/path/to/file')
  end

  test "#writefile writes content to real file" do
    p = Project.new(name: 'writefile1')
    m = p.entities.new(path: 'path/to/file')

    m.mkdirp
    m.touchfile
    m.writefile 'this is content'

    assert_equal 'this is content', File.read('test/workspace/writefile1/path/to/file')
  end
end
