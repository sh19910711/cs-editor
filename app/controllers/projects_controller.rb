class ProjectsController < ActionController::Base
  protect_from_forgery :with => :exception

  def index
    @projects = Project.all
  end

  def show
    @project = Project.find(params[:id])
  end

  def create
    @project = Project.new(project_params)

    if @project.save
      redirect_to @project, msg: 'created'
    else
      render :new
    end
  end

  def new
    @project = Project.new
  end

  private

    def project_params
      params.require(:project).permit(:name)
    end
end
