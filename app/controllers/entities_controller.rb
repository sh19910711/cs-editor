class EntitiesController < ActionController::Base
  protect_from_forgery :with => :exception
  layout 'application'

  before_filter do
    @project = Project.find(params[:project_id])
  end

  def index
    render :json => @project.entities.all
  end

  def new
    @entity = @project.entities.new
  end

  def create
    @entity = @project.entities.new(new_entity_params)

    if @entity.save_and_touchfile
      if c = entity_params[:content]
        @entity.writefile c
      end

      redirect_to project_entity_path(@project, @entity)
    else
      redirect_to @project, notice: "couldn't create a file"
    end
  end

  def show
    @path = params[:path]
    @entity = @project.entities.find_by(path: params[:path])
    @entity.readfile!
  end

  def update
    @entity = @project.entities.find_by(path: params[:path])

    if c = entity_params[:content]
      @entity.writefile c
      redirect_to project_entity_path(@project, @entity), flash: { :success => 'Updated :-)' }
    else
      redirect_to @project
    end
  end

  private

    def entity_params
      params.require(:entity).permit(:path, :content)
    end

    def new_entity_params
      entity_params.except(:content)
    end
end
