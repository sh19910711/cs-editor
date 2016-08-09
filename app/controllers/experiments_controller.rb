# *** USE ONLY IN DEVELOPMENT :-( ***
class ExperimentsController < ActionController::Base
  protect_from_forgery :with => :exception
  layout 'application'

  # GET /sendto
  def sendto_show
    @project = Project.find(params[:project_id])
  end

  # POST /sendto
  def sendto
    @project = Project.find(params[:project_id])
    res = @project.__send_request(params[:request_method], params[:request_url], params[:request_params])
    render :json => { msg: 'ok', response: res }
  rescue Errno::ECONNREFUSED => e
    render :json => { msg: 'error', error: e }
  end
end
