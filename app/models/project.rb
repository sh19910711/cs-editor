class Project < ApplicationRecord
  has_many :entities, :dependent => :destroy
end
