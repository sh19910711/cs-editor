class ProjectsController < ActionController::Base
  protect_from_forgery :with => :exception

  def index
    render json: Project.all
  end

  def show
    render json: Project.find(params[:id])
  end

  def create
    @project = Project.new

    if @project.save
      redirect_to @project, msg: 'created'
    else
      render :new
    end
  end

  def new
    @project = Project.new
  end
end
