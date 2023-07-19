from flask import Flask, render_template
from flask_login import LoginManager
import requests
import os
from src.handlers.load_user import load_user
from src.blueprints.auth.auth import auth_bp
from src.blueprints.home.home import home_bp

app = Flask(__name__)
app.config["SECRET_KEY"] = "secret-key"
login_manager = LoginManager()
login_manager.init_app(app)
login_manager.user_loader(load_user)

@app.after_request
def add_cors_headers(response):
    response.headers["Access-Control-Allow-Origin"] = "*"
    response.headers["Access-Control-Allow-Headers"] = "Content-Type,Authorization"
    response.headers["Access-Control-Allow-Methods"] = "GET,PUT,POST,DELETE"
    response.headers["Access-Control-Allow-Credentials"] = "true"
    return response


app.register_blueprint(auth_bp)
app.register_blueprint(home_bp)


# Test route to check if frontend is able to communicate with backend
@app.route("/backend")
def api():
    return requests.get(os.environ.get("BACKEND_URL", None)).content


@app.errorhandler(404)
def page_not_found(e):
    return render_template("404.html"), 404
