#!/usr/bin/env python3

"""License Checker

This script find all the projects inside certain folder and list their licenses

Todo:
    * support arguments
    * limit search level
    * resolve path
    * pretty print
    * more detail license information, i.e. GPL -> GPLv3
    * some project have vendor folder in scm, result in duplication, like testify
"""

import os
import sys
from collections import deque

__author__ = "at15"


def find_projects(vendor_folder):
    """Find all the projects in sub folders

    Any folder with a license file is treated as a project,
    because although all go packages should have host/owner/project style,
    gopkg.in results in gopkg.in/package.v1 style.
    Also there might be vendored bash scripts

    """

    projects = []
    folders_to_process = deque()
    folders_to_process.append(vendor_folder)
    while len(folders_to_process) > 0:
        current_folder = folders_to_process.popleft()
        for entry in os.scandir(current_folder):
            if entry.is_file() and entry.name.lower() == 'license':
                license_type = get_license(current_folder + "/" + entry.name)
                projects.append({'p': current_folder, 'l': license_type})
            if entry.is_dir():
                folders_to_process.append(current_folder + "/" + entry.name)
    return projects

def get_license(license_file):
    """Extract license type from license file

    Common license files can be found on:
    https://github.com/github/choosealicense.com/tree/gh-pages/_licenses

    """
    file_content = open(license_file, 'r').read()
    license_keywords = {'MIT': 'MIT', 'GNU': 'GNU', 'Apache': 'Apache',
                        'ISC': 'ISC', 'WTFPL': 'FUCK'}
    for license_type, keyword in license_keywords.items():
        if keyword in file_content:
            return license_type
    return 'UNKNOWN'


if __name__ == "__main__":
    if len(sys.argv) > 1:
        print(find_projects(sys.argv[1]))
    else:
        print("base folder not specified, using current folder")
        print(find_projects("."))
    print("finished")
