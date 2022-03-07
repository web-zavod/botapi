__version__ = '1.1.0'
from setuptools import setup, find_packages
import os
import sys

with open("README.md", "r") as file:
    long_description = file.read()

botapi_root = "./gen/py/botapi"

botapi_modules = [module for module in os.listdir(botapi_root) if module != "__init__.py"]

# For the import to work properly, we need to transform
# the directory into a python package
with open(f"{botapi_root}/__init__.py", "a") as file:
    [file.write(f"from botapi.{module}.v1 import *\n") for module in botapi_modules]

for module in botapi_modules:
    os.system(f"touch {botapi_root}/{module}/__init__.py")
    compiled_protos = [proto[:-3] for proto in os.listdir(f"{botapi_root}/{module}/v1")]
    with open(f"{botapi_root}/{module}/v1/__init__.py", "w") as file:
        file.write(f"__all__ = {compiled_protos}")

requires_python_old = [
        "grpcio==1.40.0",
        "grpcio-tools==1.40.0",
        "googleapis-common-protos==1.54.0",
        "grpc-google-longunning-v2==0.8.1",
    ]

requires_python_new = [
        "grpcio==1.43.0",
        "grpcio-tools==1.43.0",
        "googleapis-common-protos==1.54.0",
        "grpc-google-longrunning-v2==0.8.1",
        ]

setup(
    name="botapi",
    version=__version__,
    package_dir={
        "botapi": botapi_root,
        },
    packages=find_packages(where="./gen/py"),
    install_requires = requires_python_old if sys.version_info < (3,10) else requires_python_new
    )
