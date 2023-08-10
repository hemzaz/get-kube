#!/usr/bin/env python3

import os
import yaml

# Load the directory structure and files from .project.yaml
with open('.project.yaml', 'r') as file:
    dirs = yaml.safe_load(file)

# Create directories and files
for dir, files in dirs.items():
    os.makedirs(dir, exist_ok=True)
    for file in files:
        with open(os.path.join(dir, file), 'w') as f:
            pass

print("Directory structure and files created successfully!")
