from flask import Flask, render_template
import requests

app = Flask(__name__)

@app.route('/')
def index():
    return render_template('index.html')

@app.route('/about')
def about():
    return 'About page'

@app.route('/backend')
def api():
    return requests.get('http://backend:5001').content