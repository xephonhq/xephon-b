# check license
import os

def find_projects(vendor_folder):
    # although most path should be host/owner/project
    # there are projects using gopkg.in which result in path like gopkg.in
    # - limit level to 5
    # - any folder with a license file is a project

    # use as a queue
    q = [vendor_folder]
