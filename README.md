# Milo

Milo is an open source system for managing iptables across multiple hosts.

Its goal is to provide light, modern and useful iptables management.

#  Create deb Package

To create deb package you can run `make package` , it will create directory **PACKAGES** in root of project

## dependency for ubuntu 

    sudo apt install gox ruby ruby-dev rubygems build-essential
    
   ## Installing [FPM](https://fpm.readthedocs.io/en/latest/installing.html)
   fpm is a command-line program designed to help you build packages.

    gem install --no-ri --no-rdoc fpm
# Build
Just run `make build`, it will create binary file in root of project

# Development
we will use **dep** for dependency manager

    dep ensure
# How to use

coming soon...


## Maintainer

 - Roman Kredentser
