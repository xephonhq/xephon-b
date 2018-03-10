#!/usr/bin/env python3

import os
from shutil import copytree, ignore_patterns, rmtree


# hack use latest libtsdb-go

def copy_to_vendor(src_repo, dst_repo):
    """
    src_repo: github.com/dyweb/gommon
    dst_repo: github.com/benchhub/benchboard/agent
    """
    gopath = os.getenv('GOPATH')
    print('gopath is', gopath)
    repos = os.path.join(gopath, 'src')
    src = os.path.join(repos, src_repo)
    dst = os.path.join(repos, dst_repo, 'vendor', src_repo)
    if os.path.isdir(dst):
        print('remove existing directory', dst)
        rmtree(dst)
    print('copy from', src, 'to', dst)
    copytree(src, dst, ignore=ignore_patterns(
        '.git', 'vendor', '.idea', 'node_modules'))


if __name__ == '__main__':
    copy_to_vendor('github.com/libtsdb/libtsdb-go',
                   'github.com/xephonhq/xephon-b')
