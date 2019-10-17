#!/usr/bin/env python3
import os
import shutil
import subprocess
from subprocess import DEVNULL, PIPE
import tempfile
import argparse
import sys

from bs4 import BeautifulSoup

input_dir = '../examples'
output_dir = '../assets/svg'

def generate(input_dir, output_dir, template, cols):
    with tempfile.TemporaryDirectory() as build_directory:
        for src_dir, name in get_examples(input_dir):
            image_name = name + '.svg'
            builded_binary = os.path.join(build_directory, 'main')
            print("Generate image {} using go package: {}".format(image_name, src_dir))
            subprocess.run(['go', 'build', '-o', builded_binary, src_dir], check=True, stdout=DEVNULL)

            print("Binary builded, running binary to guess geometry...")
            min_cols, min_rows = get_minimal_geometry_command(builded_binary)
            if min_cols > cols:
                print('WARNING: output have string with length={}, but you specified cols={}, so image will be cropped'.format(min_cols, cols))
            svg_picture = os.path.join(build_directory, 'output.svg')

            print("Minimal geometry {}x{}, building svg image...".format(min_cols, min_rows))
            geometry = '{}x{}'.format(cols, min_rows + 1)
            subprocess.run(['termtosvg', svg_picture, '-c', builded_binary, '-g', geometry, '-t', template], check=True, stdout=DEVNULL, stderr=DEVNULL)

            print("Cropping image by one line...")
            crop_svg(svg_picture, min_rows + 1)

            print("Copy image to output dir...")
            final_image = os.path.join(output_dir, image_name)
            shutil.copyfile(svg_picture, final_image)

            print("Image for example {} generated. Path: {}\n\n".format(src_dir, final_image))

def get_examples(input_dir):
    for entry in os.scandir(input_dir):
        if entry.is_dir():
            yield entry.path, entry.name

def get_minimal_geometry(output):
    rows, cols = 0, 0
    for line in output.splitlines():
        len_line = len(line)
        if len_line > cols:
            cols = len_line
        rows += 1
    return cols, rows

def get_minimal_geometry_command(command):
    output = subprocess.run([command], stdout=PIPE, stderr=DEVNULL)
    return get_minimal_geometry(output.stdout.decode())

def crop_svg(svg_filename, lines):
    with open(svg_filename) as f:
        soup = BeautifulSoup(f, 'lxml-xml')

    screen = soup.find('svg', id='screen')
    height = int(screen['height'])
    screen['height'] = str((height // lines) * (lines - 1))

    with open(svg_filename, 'w') as f:
        f.write(str(soup))


def check_dependencies():
    try:
        subprocess.run(['go', 'version'], stdout=DEVNULL)
    except FileNotFoundError as e:
        print("go binary not found. You should install Go package.")
        return False

    try:
        subprocess.run(['termtosvg', '--help'], stdout=DEVNULL)
    except FileNotFoundError as e:
        print('termtosvg command not found. Install it with "pip3 install -r requirements.txt".')
        return False
    return True


def parse_args():
    parser = argparse.ArgumentParser(description='Generate example svg-image from Go packages')
    parser.add_argument('--input_dir', dest='input_dir', help='directory to Go packages', default=input_dir)
    parser.add_argument('--output_dir', dest='output_dir', help='directory where svg images will be saved', default=output_dir)
    parser.add_argument('--template', dest='template', help='template that will be used during generation', default='gjm8')
    parser.add_argument('--cols', dest='cols', help='width of svg image in chars', type=int, default=120)
    return parser.parse_args()

if __name__ == '__main__':
    if not check_dependencies():
        sys.exit(1)

    args = parse_args()
    generate(args.input_dir, args.output_dir, args.template, args.cols)
