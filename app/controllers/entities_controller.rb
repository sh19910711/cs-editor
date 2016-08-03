class EntitiesController < ActionController::Base
  protect_from_forgery :with => :exception

  before_filter do
    @project = Project.find(params[:project_id])
  end

  def index
  end

  def new
    @entity = @project.entities.new
  end

  def create
    @entity = @project.entities.new(entity_params)
    if @entity.save
      redirect_to project_entity_path([@project, @entity], :path => @entity.path)
    else
      redirect_to @project
    end
  end

  def show
    @path = params[:path]
  end

  private

    def entity_params
      params.require(:entity).permit(:path)
    end
end
