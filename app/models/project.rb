class Project < ApplicationRecord
  has_many :entities, :dependent => :destroy
  validates :name, uniqueness: true

  def mkdir
    FileUtils.mkdir_p dirpath
  end

  def rmdir
    FileUtils.rm_rf dirpath
  end

  if Rails.env.development?
    require 'rubygems/package'
    require 'httpclient'

    def __send_request(request_method, request_url, request_params)
      c = ::HTTPClient.new
      c.send(http_client_method(request_method), request_url, { file: tar_io }.merge(request_params))
    end

    private

      def tar_io
        io = ::StringIO.new
        io.class.class_eval { attr_accessor :path }
        io.path = "file.tar"

        ::Gem::Package::TarWriter.new(io) do |w|
          entities.each do |e|
            w.add_file(e.path, e.stat.mode) { |f| f.write(e.readfile) }
          end
        end

        io.rewind
        io
      end

      def http_client_method(request_method)
        case request_method
        when 'POST'
          :post_content
        when 'PUT'
          :put_content
        else
          raise 'unknown method'
        end
      end
  end

  private

    def dirpath
      Pathname(Editor::Application.config.workspace).join(name)
    end
end
