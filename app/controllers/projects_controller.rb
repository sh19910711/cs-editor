class ProjectsController < ActionController::Base
  protect_from_forgery :with => :exception
  layout 'application'

  def index
    @projects = Project.all
  end

  def show
    @project = Project.find(params[:id])
  end

  def create
    @project = Project.new(project_params)

    if @project.save
      @project.mkdir
      redirect_to @project, msg: 'created'
    else
      render :new
    end
  end

  def new
    @project = Project.new
  end

  def destroy
    @project = Project.find(params[:id])
    @project.destroy
    @project.rmdir

    redirect_to projects_path
  end

  private

    def project_params
      params.require(:project).permit(:name)
    end
end
