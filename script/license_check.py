#!/usr/bin/env python3
import os

# This script find all the projects inside certain folder and check if the have license like GPL

def find_projects(vendor_folder):
    # although most path should be host/owner/project
    # there are projects using gopkg.in which result in path like gopkg.in
    
    # - TODO:limit level to 5
    # - TODO: some project have vendor folder in scm, result in duplication, like testify
    # - any folder with a license file is a project

    projects = []
    # use as a queue
    q = [vendor_folder]
    while len(q) > 0:
        d = q.pop()
        for entry in os.scandir(d):
            if entry.is_file() and entry.name.lower() == 'license':
                license = get_license(d + "/" + entry.name)
                projects.append({'p':d, 'l':license})
            if entry.is_dir():
                q.insert(0, d + "/" + entry.name)
    return projects

# Common license files  https://github.com/github/choosealicense.com/tree/gh-pages/_licenses
def get_license(license_file):
    s = open(license_file,'r').read()
    # TODO: switch case style?
    if 'MIT' in s:
        return 'MIT'
    if 'GNU' in s:
        # TODO: which GNU, GPLv?, LGPL
        return 'GNU'
    if 'Apache' in s:
        # TODO: which Apache
        return 'Apache'
    if 'Creative Commons' in s:
        # TODO: CC?
        return 'CC'
    if 'ISC' in s:
        return 'ISC'
    if 'FUCK' in s:
        return 'WTFPL'
    return 'UNKNOWN'
    

if __name__ == "__main__":
    print(find_projects("."))
    print("finished")
