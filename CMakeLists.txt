cmake_minimum_required(VERSION 3.10)

# Require Clang as the compiler
find_program(CLANG_COMPILER NAMES clang++ REQUIRED)
set(CMAKE_CXX_COMPILER ${CLANG_COMPILER})

# Require Clang as the C compiler
find_program(CLANG_C_COMPILER NAMES clang REQUIRED)
set(CMAKE_C_COMPILER ${CLANG_C_COMPILER})

# Set the project name and version
project(contain-yourself VERSION 1.0)

# Specify the C++ standard
set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED True)


# Add the executable
add_executable(hello-world cpp/main.cpp)
