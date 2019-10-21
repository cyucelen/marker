#!/usr/bin/env python3
import os
import subprocess
from subprocess import DEVNULL, PIPE
import tempfile
import argparse
import sys
import io

from PIL import Image


input_dir = '../examples'
output_dir = '../assets/png'

html_template = """
<!DOCTYPE html>
<html>
    <head>
        <style>
            {css}
        </style>
    </head>
    <body>
        <div class="term-container">
{term_body}
        </div>
    </body>
</html>
"""

def generate(input_dir, output_dir, css, width):
    with tempfile.TemporaryDirectory() as build_directory:
        for src_dir, name in get_examples(input_dir):
            image_name = name + '.png'
            builded_binary = os.path.join(build_directory, 'main')
            print("Generate image {} using go package: {}".format(image_name, src_dir))
            subprocess.run(['go', 'build', '-o', builded_binary, src_dir], check=True, stdout=DEVNULL)

            main = subprocess.Popen(['script', '-q', '-c', builded_binary], stdout=PIPE)
            tth = subprocess.Popen(['terminal-to-html'], stdin=main.stdout, stdout=PIPE)
            html = render_html(tth.stdout.read().decode(), css)

            image_raw = subprocess.run(['wkhtmltoimage', '--disable-smart-width', '--width', str(width), '-f', 'png', '-',  '-'], 
                                       check=True, input=html.encode(), stdout=PIPE, stderr=DEVNULL)
            image = Image.open(io.BytesIO(image_raw.stdout))
            final_image = os.path.join(output_dir, image_name)
            image.save(final_image)

            print("Image for example {} generated. Path: {}\n".format(src_dir, final_image))

def render_html(term_body, css):
    return html_template.format_map(dict(term_body=term_body, css=css))

def get_examples(input_dir):
    for entry in os.scandir(input_dir):
        if entry.is_dir():
            yield entry.path, entry.name

def check_dependencies():
    try:
        subprocess.run(['go', 'version'], stdout=DEVNULL)
    except FileNotFoundError as e:
        print("go binary not found. You should install Go package.")
        return False

    try:
        subprocess.run(['terminal-to-html', '--help'], stdout=DEVNULL)
    except FileNotFoundError as e:
        print('terminal-to-html is not found.')
        return False

    try:
        subprocess.run(['wkhtmltoimage', '-h'], stdout=DEVNULL)
    except FileNotFoundError as e:
        print('wkhtmltoimage is not found.')
        return False
    
    return True


def parse_args():
    parser = argparse.ArgumentParser(description='Generate example svg-image from Go packages')
    parser.add_argument('--input_dir', dest='input_dir', help='directory to Go packages', default=input_dir)
    parser.add_argument('--output_dir', dest='output_dir', help='directory where images will be saved', default=output_dir)
    parser.add_argument('--template', dest='template', help='path to css template that will be used during generation', default='./templates/terminal.css')
    parser.add_argument('--width', dest='width', help='width of image', type=int, default=1500)
    return parser.parse_args()

if __name__ == '__main__':
    if not check_dependencies():
        sys.exit(1)

    args = parse_args()
    with open(args.template) as f:
        css = f.read()

    generate(args.input_dir, args.output_dir, css, args.width)
