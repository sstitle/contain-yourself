cmake_minimum_required(VERSION 3.10)

# Set the project name and version
project(contain-yourself VERSION 1.0)

# Specify the C++ standard
set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED True)

# Add the executable
add_executable(hello-world cpp/main.cpp)
